package server

import "time"

func getBeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func findDateOfDay(dow time.Weekday, weeksAhead int) time.Time {
	// Get the current date and time
	currentTime := time.Now()

	// Calculate the difference between the target day and the current day
	daysUntilTargetDay := int((dow - currentTime.Weekday() + 7) % 7)

	// Calculate the date of the target day
	targetDate := currentTime.Add(time.Duration(daysUntilTargetDay+7*weeksAhead) * 24 * time.Hour)

	return getBeginningOfDay(targetDate)
}
