package utils

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

type EventInfo struct {
	start   string
	end     string
	summary string
}

func getEventInfo(event *calendar.Event) *EventInfo {
	return &EventInfo{
		start:   event.Start.DateTime,
		end:     event.End.DateTime,
		summary: event.Summary,
	}
}

func getEventsToday(srv *calendar.Service) []*EventInfo {
	var info *EventInfo
	result := make([]*EventInfo, 0)

	today := getBeginningOfDay(time.Now()).Format(time.RFC3339)
	tmrw := getBeginningOfDay(time.Now().AddDate(0, 0, 1)).Format(time.RFC3339)
	events, err := srv.Events.List("primary").TimeMin(today).TimeMax(tmrw).
		ShowDeleted(false).SingleEvents(true).Do()

	if err != nil {
		log.Fatalf("Unable to retrieve events: %v", err)
	}

	for _, item := range events.Items {
		info = getEventInfo(item)
		result = append(result, info)
	}

	return result
}

func eventsToString(events []*EventInfo) string {
	var event string
	result := ""
	for _, item := range events {
		event = fmt.Sprintf("[%s, %s, %s]\n", 
			item.summary, item.start, item.end)
		result += event
	}

	return result
}