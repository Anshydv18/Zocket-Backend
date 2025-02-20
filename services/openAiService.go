package services

import (
	"backend/env"
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func GenerateTaskName(ctx *context.Context, task string) {
	client := openai.NewClient(env.Get("OPEN_AI_API_KEY"))

	res, err := client.CreateChatCompletion(*ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "You are an assistant that generates smart task suggestions for a task management app.",
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("Suggest three tasks for building a real-time task related to %s.", task),
			},
		},
		MaxTokens: 150,
	})

	fmt.Println(res)
	fmt.Println(err)
}
