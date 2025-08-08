# MiP MCP Server

This program creates an [MCP server](https://modelcontextprotocol.io/overview) that provides tools for controlling the MiP robot from any model that has tool calling support.

## Install

```shell
go install github.com/hybridgroup/tinygo-mip/cmd/mcp-mip@latest
```

## Running

```shell
mcp-mip [MAC address or Bluetooth ID]
```

You can also use the `-port` flag to set a specific port for the MCP server.

```shell
$ mcp-mip D0:39:72:A2:4E:55
enabling...
start scan...
device: D0:39:72:A2:4E:55 -60 WowWee-MiP-8429
connected to D0:39:72:A2:4E:55
2025/08/02 12:43:00 MCP server listening on http :9090
```

Once it is running, you can call it from whatever MCP host/client that you wish.
