package server

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"google.golang.org/api/calendar/v3"
)

type Assistant struct {
	srv *calendar.Service
	gpt *openai.Client
	conversation []openai.ChatCompletionMessage
}

type GptMsg struct {
	Id      int          `json:"id"`
	Message string       `json:"message"`
	Events  []*EventInfo `json:"events"`
}

func CreateAssistant() *Assistant {
	srv, err := createServiceFromCredentials()
	if err != nil {
		return nil
	}

	fmt.Printf("Key: %s\n", os.Getenv("OPENAI_API_KEY"))

	gpt := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	return &Assistant{
		srv: srv,
		gpt: gpt,
		conversation: initialConversation,
	}
}

func (asst *Assistant) StartAssistant() error {
	err := asst.sendGptMessage("Here are my events for today:" + eventsToString(getEventsToday(asst.srv)))

	if err != nil {
		fmt.Printf("Error communicating with OpenAI: %v\n", err)
		return err
	}

	go asst.runAssistant()
	
	return nil
}

func (asst *Assistant) runAssistant() {
	// Running the assistant
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter your response: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
		}

		err = asst.sendGptMessage("User: " + input)

		if err != nil {
			fmt.Printf("Error communicating with OpenAI: %v\n", err)
			break
		}
	}
}

func (asst *Assistant) sendGptMessage (msg string) error {
	var gptMsg GptMsg

	// ==================== Sending message ====================
	myEventsMsg := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: msg,
	}
	asst.conversation = append(asst.conversation, myEventsMsg)

	// Send message
	resp, err := asst.gpt.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONObject,
			},
			Messages: asst.conversation,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return errors.New("Error setting up GPT client")
	}

	// ==================== Handle response ====================
	respMsg := resp.Choices[0].Message.Content
	err = json.Unmarshal([]byte(respMsg), &gptMsg)

	if err != nil {
		fmt.Printf("Unmarshalling error: %v\n", err)
		return err
	}

	fmt.Printf("Assistant:\n %s\n", gptMsg.Message)
	return nil
}


// Testing
func (asst *Assistant) sendTestMessage() error {
	var gptMsg GptMsg

	// ==================== Sending message ====================

	// Send message
	resp, err := asst.gpt.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONObject,
			},
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `I will be passing you events in my day in JSON format. You will help to schedule new events and respond in JSON format.
					
					You will be a helpful assistant for scheduling events for users, and will have a friendly conversation. A user of my application will ask you to help them find time in the day to schedule new events. Please start by asking the user what events they would like to schedule.
					
					When responding, please describe the event by giving the day of the week, time, and date using a normal time format and ask for confirmation.`,
				},
				{
					Role: openai.ChatMessageRoleUser,
					Content: `Here are my events for today:
						{"summary": "Project X meeting", "start": "2023-12-22T14:30:00-05:00", "end": "2023-10-28T15:30:00-05:00"} 
						{"summary": "Meeting with Chase", "start":"2023-12-22T16:15:00-05:00", "end": "2023-10-28T18:15:00-05:00"}
						{"summary": "Christmas Party", "start": "2023-12-22T20:00:00-05:00", "end": "2023-10-28T23:00:00-05:00"}`,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return errors.New("Error setting up GPT client")
	}

	// ==================== Handle response ====================
	respMsg := resp.Choices[0].Message.Content
	err = json.Unmarshal([]byte(respMsg), &gptMsg)

	if err != nil {
		fmt.Printf("Unmarshalling error: %v\n", err)
		return err
	}

	fmt.Printf("Assistant:\n %s\n", gptMsg.Message)
	return nil
}