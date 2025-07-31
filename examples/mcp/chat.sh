#!/bin/bash
. "${HOME}/osprey.sh"

: <<'COMMENT'
✋ if you are running this script in a Docker container, 
you need to export the MODEL_RUNNER_BASE_URL environment variable to point to the model runner service.
export MODEL_RUNNER_BASE_URL=http://model-runner.docker.internal/engines/llama.cpp/v1

✋ if you are working with devcontainer, it's already set.
COMMENT

DMR_BASE_URL=${MODEL_RUNNER_BASE_URL:-http://localhost:12434/engines/llama.cpp/v1}
MODEL=${MODEL_RUNNER_TOOL_MODEL:-"ai/qwen2.5:latest"}
TEMPERATURE=${MODEL_RUNNER_TEMPERATURE:-"0.0"}

DEBUG_MODE=false
if [ "$1" = "-d" ] || [ "$1" = "--debug" ]; then
  DEBUG_MODE=true
fi

clear

echo " ▄▄       ▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄       ▄▄       ▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄ ";
echo "▐░░▌     ▐░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌     ▐░░▌     ▐░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌";
echo "▐░▌░▌   ▐░▐░▌ ▀▀▀▀█░█▀▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌     ▐░▌░▌   ▐░▐░▌▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌";
echo "▐░▌▐░▌ ▐░▌▐░▌     ▐░▌     ▐░▌       ▐░▌     ▐░▌▐░▌ ▐░▌▐░▌▐░▌          ▐░▌       ▐░▌";
echo "▐░▌ ▐░▐░▌ ▐░▌     ▐░▌     ▐░█▄▄▄▄▄▄▄█░▌     ▐░▌ ▐░▐░▌ ▐░▌▐░▌          ▐░█▄▄▄▄▄▄▄█░▌";
echo "▐░▌  ▐░▌  ▐░▌     ▐░▌     ▐░░░░░░░░░░░▌     ▐░▌  ▐░▌  ▐░▌▐░▌          ▐░░░░░░░░░░░▌";
echo "▐░▌   ▀   ▐░▌     ▐░▌     ▐░█▀▀▀▀▀▀▀▀▀      ▐░▌   ▀   ▐░▌▐░▌          ▐░█▀▀▀▀▀▀▀▀▀ ";
echo "▐░▌       ▐░▌     ▐░▌     ▐░▌               ▐░▌       ▐░▌▐░▌          ▐░▌          ";
echo "▐░▌       ▐░▌ ▄▄▄▄█░█▄▄▄▄ ▐░▌               ▐░▌       ▐░▌▐░█▄▄▄▄▄▄▄▄▄ ▐░▌          ";
echo "▐░▌       ▐░▌▐░░░░░░░░░░░▌▐░▌               ▐░▌       ▐░▌▐░░░░░░░░░░░▌▐░▌          ";
echo " ▀         ▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀                 ▀         ▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀           ";
echo "                                                                                   ";

echo "🧠 using model ${MODEL}"
docker model pull ${MODEL}

echo ""

MCP_SERVER=${MCP_SERVER:-"http://localhost:9090"}

read -r -d '' SYSTEM_INSTRUCTION <<- EOM
You are a MiP self-balancing robot.
You have tools that actually call a physical robot that you are connected to via Bluetooth.
Use your tools to respond to human requests.
Keep your responses short and to the point.
EOM

MCP_TOOLS=$(get_mcp_http_tools "$MCP_SERVER")
TOOLS=$(transform_to_openai_format "$MCP_TOOLS")

if [[ "$DEBUG_MODE" == "true" ]]; then
  echo "---------------------------------------------------------"
  echo "Available tools:"
  echo "${TOOLS}" 
  echo "---------------------------------------------------------"
fi

: <<'COMMENT'
Example of commands to send to your MiP robot:
Turn the chest LED purple
COMMENT

# needed to handle the ASSISTANT_RESPONSE being created from a subprocess.
shopt -s lastpipe

# Initialize conversation history array
CONVERSATION_HISTORY=()
ASSISTANT_RESPONSE=""

function callback() {
  ASSISTANT_RESPONSE+="$1"
}

while true; do
  STOPPED="false"
  USER_CONTENT=$(gum write --placeholder "🎤 What can I do for you?")
  
  if [[ "$USER_CONTENT" == "/bye" ]]; then
    echo "Goodbye!"
    break
  fi

  echo "🎤 ${USER_CONTENT}"

  # Add user message to conversation history
  add_user_message CONVERSATION_HISTORY "${USER_CONTENT}"

  while [ "$STOPPED" != "true" ]; do
    # Build messages array conversation history
    MESSAGES=$(build_messages_array CONVERSATION_HISTORY)

    read -r -d '' DATA <<- EOM
{
  "model": "${MODEL}",
  "options": {
    "temperature": ${TEMPERATURE}
  },
  "messages": [${MESSAGES}],
  "tools": ${TOOLS},
  "parallel_tool_calls": true,
  "tool_choice": "auto"
}
EOM

    RESULT=$(osprey_tool_calls ${DMR_BASE_URL} "${DATA}")

    if [[ "$DEBUG_MODE" == "true" ]]; then
      echo "📝 raw JSON response:"
      print_raw_response "${RESULT}"
    fi

    FINISH_REASON=$(get_finish_reason "${RESULT}")
    case $FINISH_REASON in
      tool_calls)
        # Get tool calls for further processing
        TOOL_CALLS=$(get_tool_calls "${RESULT}")

        if [[ -n "$TOOL_CALLS" ]]; then
            echo "⏳ making robot request..."
            add_tool_calls_message CONVERSATION_HISTORY "${TOOL_CALLS}"

            for tool_call in $TOOL_CALLS; do
                FUNCTION_NAME=$(get_function_name "$tool_call")
                FUNCTION_ARGS=$(get_function_args "$tool_call")
                CALL_ID=$(get_call_id "$tool_call")

                echo "🛠️ calling robot function: $FUNCTION_NAME with args: $FUNCTION_ARGS"

                # Execute function via MCP
                MCP_RESPONSE=$(call_mcp_http_tool "$MCP_SERVER" "$FUNCTION_NAME" "$FUNCTION_ARGS")
                RESULT_CONTENT=$(get_tool_content_http "$MCP_RESPONSE")

                echo "✅ robot result: $RESULT_CONTENT"
                echo ""

                TOOL_RESULT=$(echo "${RESULT_CONTENT}" | jq -r '.content')
                add_tool_message CONVERSATION_HISTORY "${CALL_ID}" "${TOOL_RESULT}"
            done
        else
          if [[ "$DEBUG_MODE" == "true" ]]; then
            echo "🔵 no tool calls found in response"
          fi
        fi
        ;;

      stop)
        STOPPED="true"
        ASSISTANT_MESSAGE=$(echo "${RESULT}" | jq -r '.choices[0].message.content')
        echo "🤖 ${ASSISTANT_MESSAGE}"

        # Add assistant response to conversation history (from callback)
        add_assistant_message CONVERSATION_HISTORY "${ASSISTANT_MESSAGE}"
        ;;

      *)
        echo "🔵 unexpected finish reason"
        ;;
    esac

  done
  echo ""
done
