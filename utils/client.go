package utils

import (
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"google.golang.org/api/calendar/v3"
)

type Assistant struct {
	calendar *calendar.Service
	gpt *openai.Client
	systemMessage string
}

func CreateAssistant() *Assistant {
	calendar, err := createServiceFromCredentials()
	if err != nil {
		return nil
	}

	fmt.Printf("Key: %s\n", os.Getenv("OPENAI_API_KEY"))

	gpt := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	return &Assistant{
		calendar: calendar,
		gpt: gpt,
		systemMessage: "I will be passing you events in my day with the format [event-name, start-time, end-time]. You will help to schedule new events and return events in the format [event-name, start-time, end-time].",
	}
}

func (asst *Assistant) StartAssistant() error {
	// gpt := asst.gpt
	events := getEventsToday(asst.calendar)

	// resp, err := gpt.CreateChatCompletion(
	// 	context.Background(),
	// 	openai.ChatCompletionRequest{
	// 		Model: openai.GPT3Dot5Turbo,
	// 		Messages: []openai.ChatCompletionMessage{
	// 			{
	// 				Role:    openai.ChatMessageRoleSystem,
	// 				Content: asst.systemMessage,
	// 			},
	// 			{
	// 				Role:    openai.ChatMessageRoleUser,
	// 				Content: "These are my events today: " + eventsToString(events),
	// 			},
	// 		},
	// 	},
	// )

	// if err != nil {
	// 	fmt.Printf("ChatCompletion error: %v\n", err)
	// 	return errors.New("Error setting up GPT client")
	// }

	print("These are my events today: " + eventsToString(events))

	// fmt.Print(resp.Choices[0].Message.Content)
	return nil
}
