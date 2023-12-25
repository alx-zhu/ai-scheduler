package server

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	openai "github.com/sashabaranov/go-openai"
	"google.golang.org/api/calendar/v3"
)

const (
	asstInstructions = "I will be passing you events in my day with the format ${event-name, start-time, end-time}$. You will help to schedule new events and return events in the format ${event-name, start-time, end-time}$.\n	You will be a helpful assistant for scheduling events for my users, and will have a friendly conversation. A user of my application will ask you to help them find time in my day to schedule new events. Please start by asking the user what events they'd like to schedule.\n When responding, please describe the event by giving me the day of the week, time, and date using a normal time format and ask me for confirmation\n	After confirmation, send me the event formatted as ${event-name, start-time, end-time}$"
)

type Assistant struct {
	srv           *calendar.Service
	gpt           *openai.Client
	systemMessage string
}

type GPTMsg struct {
	Id      int
	Message string
	Events  []*EventInfo
}

func CreateAssistant() *Assistant {
	srv, err := createServiceFromCredentials()
	if err != nil {
		return nil
	}

	fmt.Printf("Key: %s\n", os.Getenv("OPENAI_API_KEY"))

	gpt := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	return &Assistant{
		srv:           srv,
		gpt:           gpt,
		systemMessage: "I will be passing you events in my day with the format [event-name, start-time, end-time]. You will help to schedule new events and return events in the format [event-name, start-time, end-time].",
	}
}

func (asst *Assistant) StartAssistant() error {
	// gpt := asst.gpt
	events := getEventsToday(asst.srv)

	// resp, err := gpt.CreateChatCompletion(
	// 	context.Background(),
	// 	openai.ChatCompletionRequest{
	// 		Model: openai.GPT3Dot5Turbo,
	//		ResponseFormat: ChatCompletionResponseFormatTypeJSONObject,
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
	re := regexp.MustCompile("\\${}\\$")
	for i, event := range events {
		bytes, _ := json.Marshal(event)
		fmt.Printf("%d. %v\n", i, bytes)
	}
	print("These are my events today: " + eventsToString(events))
	re.FindAllString("", -1)

	// fmt.Print(resp.Choices[0].Message.Content)
	return nil
}

// If the response does not contain any events, you MUST start it with the characters ***
