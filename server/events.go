package server

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

type EventInfo struct {
	Summary string `json:"summary"`
	Start   string `json:"start"`
	End     string `json:"end"`
}

// ========== Get Event ==========

func getEventInfo(event *calendar.Event) *EventInfo {
	return &EventInfo{
		Summary: event.Summary,
		Start:   event.Start.DateTime,
		End:     event.End.DateTime,
	}
}

func getEvents(srv *calendar.Service, date time.Time) ([]*EventInfo, error) {
	var info *EventInfo
	result := make([]*EventInfo, 0)
	today := getBeginningOfDay(date).Format(time.RFC3339)
	tmrw := getBeginningOfDay(date).AddDate(0, 0, 1).Format(time.RFC3339)
	events, err := srv.Events.List("primary").TimeMin(today).TimeMax(tmrw).
		ShowDeleted(false).SingleEvents(true).Do()

	if err != nil {
		log.Fatalf("Unable to retrieve events: %v", err)
		return nil, err
	}

	for _, item := range events.Items {
		info = getEventInfo(item)
		result = append(result, info)
	}

	return result, err
}

func getEventsToday(srv *calendar.Service) []*EventInfo {
	result, err := getEvents(srv, time.Now())
	if err != nil {
		return nil
	}
	return result
}

func getEventsOnDow(srv *calendar.Service, dow time.Weekday, weeksAhead int) []*EventInfo {
	result, err := getEvents(srv, findDateOfDay(dow, weeksAhead))
	if err != nil {
		return nil
	}
	return result
}

// ========== Add Event ==========

func (asst *Assistant) insertEvent(event *EventInfo) (success bool) {
	new := &calendar.Event{
		Summary: event.Summary,
		Start:   &calendar.EventDateTime{DateTime: event.Start},
		End:     &calendar.EventDateTime{DateTime: event.End},
	}

	res, err := asst.srv.Events.Insert("primary", new).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
		return false
	}

	fmt.Printf("Event created: %s\n", res.HtmlLink)
	return true
}

func (asst *Assistant) insertEvents(events []*EventInfo) {
	var ok bool
	for _, event := range events {
		ok = asst.insertEvent(event)
		if !ok {
			fmt.Printf("Failed to create event %v\n", event)
		}
	}
}

// ========== Helpers ==========

func eventsToString(events []*EventInfo) string {
	result := ""
	for _, item := range events {
		bytes, _ := json.Marshal(item)
		result += string(bytes) + "\n"
	}
	return result
}
