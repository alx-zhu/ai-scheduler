package server

import openai "github.com/sashabaranov/go-openai"

var initialConversation = []openai.ChatCompletionMessage{
	{
		Role: openai.ChatMessageRoleUser,
		Content: `Here are my events for today:
			{"summary": "Project X meeting", "start": "2023-12-22T14:30:00-05:00", "end": "2023-10-28T15:30:00-05:00"} 
			{"summary": "Meeting with Chase", "start":"2023-12-22T16:15:00-05:00", "end": "2023-10-28T18:15:00-05:00"}
			{"summary": "Christmas Party", "start": "2023-12-22T20:00:00-05:00", "end": "2023-10-28T23:00:00-05:00"}`,
	},
	{
		Role: openai.ChatMessageRoleAssistant,
		Content: `{
				"id": 1,
				"message": "Great! It looks like you have some events scheduled for today: \n 1. Project X meeting today from 2:30 pm to 3:30 pm \n 2. Meeting with Chase from 4:15 pm to 6:15 pm \n 3. Christmas Party from 8:00 pm to 11:00 pm\n I'd be happy to help you schedule some events. What would you like to schedule?",
				"events": []
			}`,
	},
	{
		Role:    openai.ChatMessageRoleUser,
		Content: "I want to go to the gym this afternoon for 1 hour",
	},

	{
		Role: openai.ChatMessageRoleAssistant,
		Content: `{
				"id": 2,
				"message": "Sure, I can help you find a suitable time for your gym session. Based on your current schedule, it seems you have a couple of options:\n\n 1. Before Project X Meeting \n  Time: Anytime Before 2:30 PM \n  Date: Thursday, December 22, 2023\n\n 2. Before Meeting with Chase:\n  Time: Between 3:30 PM and 4:15 PM \n  Date: Thursday, December 22, 2023\n\n 3. After Meeting with Chase:\n  Time: Between 6:15 PM and 8:00 PM \n  Date: Thursday, December 22, 2023 \n\n Please let me know which option works better for you, or if you have a specific time in mind. Once you confirm, I'll schedule the gym session accordingly.",
				"events": []
			}`,
	},
	{
		Role:    openai.ChatMessageRoleUser,
		Content: "Before the Project X Meeting would be great.",
	},
	{
		Role: openai.ChatMessageRoleAssistant,
		Content: `{
			"id": 3,
			"message": "Great choice! Let's schedule your gym session before your 'Project X Meeting' event. How about we allocate 1 hour for your gym session?\n\n New Event: Gym Session \n  Start Time: 12:00 PM \n  End Time: 1:00 PM\n  Date: Thursday, December 22, 2023\n\n Does this time work for you? If it does, I'll go ahead and add it to your schedule.",
			"events": []
			}`,
	},
	{
		Role:    openai.ChatMessageRoleUser,
		Content: "No, I would prefer a little later, after lunch ends at 1:00pm",
	},
	{
		Role: openai.ChatMessageRoleAssistant,
		Content: `{
				"id": 4,
				"message": "No problem! How about we move your gym session to 1:00pm after your lunch?\n\n New Event: Gym Session \n  Start Time: 1:00 PM\n  End Time: 2:00 PM\n  Date: Thursday, December 22, 2023\n\nDoes this time work for you? If it does, I'll go ahead and add it to your schedule.",
				"events": []
			}`,
	},
	{
		Role:    openai.ChatMessageRoleUser,
		Content: "Yes, that works! Please add it to my schedule.",
	},
	{
		Role: openai.ChatMessageRoleAssistant,
		Content: `{
				"id": 5,
				"message": "Fantastic! Your gym session is now scheduled.",
				"events": [{"summary": "Gym session", "start": "2023-12-22T13:00:00-05:00", "end": "2023-10-28T14:00:00-05:00"}]
			}`,
	},
	{
		Role: openai.ChatMessageRoleSystem,
		Content: `I will be passing you events in my day in JSON format. You will help to schedule new events and respond in JSON format.
		
		You will be a helpful assistant for scheduling events for users, and will have a friendly conversation. A user of my application will ask you to help them find time in the day to schedule new events. Please start by asking the user what events they would like to schedule.
		
		When responding, please describe the event by giving the day of the week, time, and date using a normal time format and ask for confirmation.
		
		Please restart the conversation now.`,
	},
}
