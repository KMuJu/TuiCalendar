package state

import (
	"testing"
	"time"

	"github.com/kmuju/TuiCalendar/cmd/model"

	"github.com/stretchr/testify/assert"
)

var (
	t1, _  = time.Parse(time.RFC3339, "2024-05-05T08:15:00Z")
	t2, _  = time.Parse(time.RFC3339, "2024-05-05T12:30:00Z")
	t3, _  = time.Parse(time.RFC3339, "2024-05-05T16:45:00Z")
	t4, _  = time.Parse(time.RFC3339, "2024-05-15T09:00:00Z")
	t5, _  = time.Parse(time.RFC3339, "2024-05-15T13:15:00Z")
	t6, _  = time.Parse(time.RFC3339, "2024-05-15T17:30:00Z")
	t7, _  = time.Parse(time.RFC3339, "2024-05-20T10:00:00Z")
	t8, _  = time.Parse(time.RFC3339, "2024-05-20T14:15:00Z")
	t9, _  = time.Parse(time.RFC3339, "2024-05-20T18:30:00Z")
	t10, _ = time.Parse(time.RFC3339, "2024-05-25T11:00:00Z")
)

func TestStartDayFilter(t *testing.T) {
	controller := EventController{Events: createEvents()}

	output := controller.GetEvents(StartDayFilter(2024, 5, 15))
	expected := []model.Event{
		{Start: t4, End: t4.Add(time.Hour)},
		{Start: t5, End: t5.Add(time.Hour)},
		{Start: t6, End: t6.Add(time.Hour)},
	}
	assert.Equal(t, len(expected), len(output))
	for i := 0; i < len(expected); i++ {
		assert.Equal(t, expected[i], output[i])
	}
}

func TestAfterDayFilter(t *testing.T) {
	controller := EventController{Events: createEvents()}

	output := controller.GetEvents(AfterDayFilter(2024, 5, 20))
	expected := []model.Event{
		{Start: t7, End: t7.Add(time.Hour)},
		{Start: t8, End: t8.Add(time.Hour)},
		{Start: t9, End: t9.Add(time.Hour)},
		{Start: t10, End: t10.Add(time.Hour)},
	}
	assert.Equal(t, len(expected), len(output))
	for i := 0; i < len(expected); i++ {
		assert.Equal(t, expected[i], output[i])
	}
}

func TestBetweenTimesFilter(t *testing.T) {
	controller := EventController{Events: createEvents()}

	output := controller.GetEvents(BetweenTimesFilter(t2, t6))
	expected := []model.Event{
		{Start: t3, End: t3.Add(time.Hour)},
		{Start: t4, End: t4.Add(time.Hour)},
		{Start: t5, End: t5.Add(time.Hour)},
	}
	assert.Equal(t, len(expected), len(output))
	for i := 0; i < len(expected); i++ {
		assert.Equal(t, expected[i], output[i])
	}
}

func createEvents() []model.Event {
	return []model.Event{
		{Start: t1, End: t1.Add(time.Hour)},
		{Start: t2, End: t2.Add(time.Hour)},
		{Start: t3, End: t3.Add(time.Hour)},
		{Start: t4, End: t4.Add(time.Hour)},
		{Start: t5, End: t5.Add(time.Hour)},
		{Start: t6, End: t6.Add(time.Hour)},
		{Start: t7, End: t7.Add(time.Hour)},
		{Start: t8, End: t8.Add(time.Hour)},
		{Start: t9, End: t9.Add(time.Hour)},
		{Start: t10, End: t10.Add(time.Hour)},
	}

}
