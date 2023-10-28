package main

func startAssistant() {
	srv, err := createServiceFromCredentials()
	if err != nil {
		return
	}

	events := getEventsToday(srv)
	print(eventsToString(events))
}
