package main

import (
	"context"
	"fmt"
	"os"

	swarmgo "github.com/allurisravanth/swarmgo"
	"github.com/allurisravanth/swarmgo/llm"
	dotenv "github.com/joho/godotenv"
)

func main() {
	dotenv.Load()

	aiProvider := llm.NewOpenAILLM(os.Getenv("OPENAI_API_KEY"))

	client := swarmgo.NewSwarmWithCustomProvider(aiProvider, &swarmgo.Config{})

	agent := &swarmgo.Agent{
		Name:         "Agent",
		Instructions: "You are a helpful agent.",
		Model:        "gpt-3.5-turbo",
	}

	messages := []llm.Message{
		{Role: llm.RoleUser, Content: "Hi!"},
	}

	ctx := context.Background()
	response, err := client.Run(ctx, agent, messages, nil, "", false, false, 5, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Messages[len(response.Messages)-1].Content)
}
