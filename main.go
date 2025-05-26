package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const timeZone = "time-zone"

func main() {
	s := server.NewMCPServer(
		"Timer",
		"1.0.0",
	)

	tool := mcp.NewTool("now",
		mcp.WithDescription("現在の時刻を表示します（Shows the current time）"),
		mcp.WithString(timeZone,
			mcp.Description("時刻を取得するタイムゾーン（Timezone to get the time in）"),
		),
	)

	s.AddTool(tool, nowHandler)

	tool2 := mcp.NewTool("time",
		mcp.WithDescription("現在の時刻のみを表示します（Shows the current time only）"),
		mcp.WithString(timeZone,
			mcp.Description("時刻を取得するタイムゾーン（Timezone to get the time in）"),
		),
	)

	s.AddTool(tool2, nowHandler) // TODO: Change nowHandler to timeHandler

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func nowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var now time.Time
	tz, ok := request.Params.Arguments[timeZone].(string)
	if !ok {
		now = time.Now()
	} else {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			return nil, err
		}
		now = time.Now().In(loc)
	}

	return mcp.NewToolResultText(now.Format(time.RFC3339)), nil
}

func timeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var now time.Time
	tz, ok := request.Params.Arguments[timeZone].(string)
	if !ok {
		now = time.Now()
	} else {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			return nil, err
		}
		now = time.Now().In(loc)
	}

	return mcp.NewToolResultText(now.Format("15:04:05")), nil
}
