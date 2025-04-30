# Timer MCP Server

A MCP server that displays the current time.

## Requirements

- Go

## How to Use

1. Clone this repository

```sh
git clone git@github.com:hond0413/Timer-MCP.git
```

2. Build this tool

```sh
go build .
```

3. Add MCP settings to your MCP client

Cursor case

```json
{
  "mcpServers": {
    "timer": {
      "command": "~/go/src/github.com/hond0413/time-mcp/time-mcp"
    }
  }
}
```

4. Call this MCP server
