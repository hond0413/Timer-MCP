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

	nowTool := mcp.NewTool("now",
		mcp.WithDescription("現在の日付と時刻を表示します（Shows the current date and time）"),
		mcp.WithString(timeZone,
			mcp.Description("時刻を取得するタイムゾーン（Timezone to get the time in）"),
		),
	)

	s.AddTool(nowTool, nowHandler)

	timeTool := mcp.NewTool("time",
		mcp.WithDescription("現在の時刻のみを表示します（Shows only the current time）"),
		mcp.WithString(timeZone,
			mcp.Description("時刻を取得するタイムゾーン（Timezone to get the time in）"),
		),
	)

	s.AddTool(timeTool, timeHandler)

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func nowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// タイムゾーンを考慮した時刻を取得
	now, err := getTimeWithZone(request)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(now.Format(time.RFC3339)), nil
}

// 共通の時刻取得ロジックを抽出した関数
func getTimeWithZone(request mcp.CallToolRequest) (time.Time, error) {
	var now time.Time
	tz, ok := request.Params.Arguments[timeZone].(string)
	if !ok {
		now = time.Now()
	} else {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			return time.Time{}, err
		}
		now = time.Now().In(loc)
	}

	return now, nil
}

func timeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// nowHandlerと共通のロジックを使用
	result, err := getTimeWithZone(request)
	if err != nil {
		return nil, err
	}

	// 時間のみのフォーマットで結果を返す
	return mcp.NewToolResultText(result.Format("15:04:05")), nil
}
