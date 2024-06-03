package tui

import (
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/kmuju/TuiCalendar/cmd/model"
)

func TestNewCalendar(t *testing.T) {
	events := []model.Event{
		{Name: "Event 1", Start: time.Now(), End: time.Now().Add(time.Hour)},
		{Name: "Event 2", Start: time.Now().Add(2 * time.Hour), End: time.Now().Add(3 * time.Hour)},
	}
	height := 30
	width := 30
	listWidth := 20
	renderFrom := 0
	renderAmount := 3

	cal := NewCalendar(events, height, width, listWidth, renderFrom, renderAmount)

	if len(cal.events) != len(events) {
		t.Fatalf("Wrong length of events")
	}
	if cal.height != height {
		t.Fatalf("Wrong length of events")
	}
	if cal.width != width {
		t.Fatalf("Wrong length of events")
	}
	if cal.listWidth != listWidth {
		t.Fatalf("Wrong length of events")
	}
	if cal.renderFrom != renderFrom {
		t.Fatalf("Wrong length of events")
	}
	if cal.renderAmount != len(events) {
		t.Fatalf("Wrong length of events")
	}
}

func TestEmptyRender(t *testing.T) {
	cal := NewCalendar([]model.Event{}, 0, 12, 0, 0, 0)
	if cal.Render() != "Ingen Events" {
		t.Fatalf("Did not detect empty events slice\n")
	}
}

func TestDown(t *testing.T) {
	events := []model.Event{
		{Name: "A", Start: time.Now(), End: time.Now().Add(time.Hour)},
		{Name: "B", Start: time.Now().Add(2 * time.Hour), End: time.Now().Add(3 * time.Hour)},
		{Name: "C", Start: time.Now(), End: time.Now().Add(time.Hour)},
		{Name: "D", Start: time.Now(), End: time.Now().Add(time.Hour)},
	}
	cal := NewCalendar(events, 0, 0, 0, 0, 3)
	if cal.renderFrom != 0 || cal.renderAmount != 3 || cal.selected != 0 {
		t.Fatalf("wrong render from(%d), selected(%d) or amount(%d)\n", cal.renderFrom, cal.selected, cal.renderAmount)
	}
	// | -> Events rendered
	/*
		A | S
		B |
		C |
		D
	*/
	{
		/*
			A |
			B | S
			C |
			D
		*/
		cal.Down()
		if cal.selected != 1 {
			t.Fatalf("1:Did not go down: %d\n", cal.selected)
		}
		if cal.renderFrom != 0 {
			t.Fatalf("1:Changed render from: %d\n", cal.renderFrom)
		}
	}
	{
		/*
			A |
			B |
			C | S
			D
		*/
		cal.Down()
		if cal.selected != 2 {
			t.Fatalf("2:Did not go down: %d\n", cal.selected)
		}
		if cal.renderFrom != 0 {
			t.Fatalf("2:Changed render from: %d\n", cal.renderFrom)
		}
	}
	{
		/*
			A
			B |
			C |
			D | S
		*/
		cal.Down()
		if cal.selected != 3 {
			t.Fatalf("Did not go down: %d\n", cal.selected)
		}
		if cal.renderFrom != 1 {
			t.Fatalf("3: Did not change renderfrom: %d\n", cal.renderFrom)
		}
	}
	{
		/*
			A
			B |
			C |
			D | S
		*/
		cal.Down()
		if cal.selected != 3 {
			t.Fatalf("Went down %d\n", cal.selected)
		}
		if cal.renderFrom != 1 {
			t.Fatalf("4: Changed renderfrom %d\n", cal.renderFrom)
		}
	}

}

func TestUp(t *testing.T) {
	events := []model.Event{
		{Name: "A", Start: time.Now(), End: time.Now().Add(time.Hour)},
		{Name: "B", Start: time.Now().Add(2 * time.Hour), End: time.Now().Add(3 * time.Hour)},
		{Name: "C", Start: time.Now(), End: time.Now().Add(time.Hour)},
		{Name: "D", Start: time.Now(), End: time.Now().Add(time.Hour)},
	}
	cal := NewCalendar(events, 0, 0, 0, 0, 3)
	cal.renderFrom = 1
	cal.selected = 3
	if cal.renderFrom != 1 || cal.renderAmount != 3 || cal.selected != 3 {
		t.Fatalf("wrong render from(%d), selected(%d) or amount(%d)\n", cal.renderFrom, cal.selected, cal.renderAmount)
	}
	/*
		A
		B |
		C |
		D | S
	*/
	{
		/*
			A
			B |
			C | S
			D |
		*/
		cal.Up()
		if cal.renderFrom != 1 {
			t.Fatalf("1: Changed render from %d\n", cal.renderFrom)
		}
		if cal.selected != 2 {
			t.Fatalf("1: Wrong value for selected %d\n", cal.selected)
		}
	}
	{
		/*
			A
			B | S
			C |
			D |
		*/
		cal.Up()
		if cal.renderFrom != 1 {
			t.Fatalf("2: Changed render from %d\n", cal.renderFrom)
		}
		if cal.selected != 1 {
			t.Fatalf("2: Wrong value for selected %d\n", cal.selected)
		}
	}
	{
		/*
			A | S
			B |
			C |
			D
		*/
		cal.Up()
		if cal.renderFrom != 0 {
			t.Fatalf("3: Did not change render from %d\n", cal.renderFrom)
		}
		if cal.selected != 0 {
			t.Fatalf("3: Wrong value for selected %d\n", cal.selected)
		}
	}
	{
		/*
			A | S
			B |
			C |
			D
		*/
		cal.Up()
		if cal.renderFrom != 0 {
			t.Fatalf("3: Changed render from when zero %d\n", cal.renderFrom)
		}
		if cal.selected != 0 {
			t.Fatalf("3: Changed selected at top of list %d\n", cal.selected)
		}
	}
}

func TestRenderString(t *testing.T) {
	events := []model.Event{
		{Name: "Event 1", Start: time.Now(), End: time.Now().Add(time.Hour)},
		{Name: "Event 2", Start: time.Now().Add(2 * time.Hour), End: time.Now().Add(3 * time.Hour)},
	}
	height := 30
	width := 30
	listWidth := 20
	renderFrom := 0
	renderAmount := 3

	cal := NewCalendar(events, height, width, listWidth, renderFrom, renderAmount)
	output := cal.Render()
	if output != cal.String() {
		t.Fatalf("Render should be the same as String\n%s\n", output)
	}
}

func TestRenderEvent(t *testing.T) {
	width := 10
	event := model.Event{
		Name:        "Event 2",
		Start:       time.Now().Add(2 * time.Hour),
		End:         time.Now().Add(3 * time.Hour),
		Description: "this is the desc for the test",
	}

	str := renderEvent(event, width, false)
	output := strings.Split(str, "\n")
	if len(output) == 0 {
		t.Fatalf("No lines in renderEvent (%+v)\n", event)
	}
	if utf8.RuneCountInString(output[0]) != width {
		t.Fatalf("Wrong width: Expected(%d), Got(%d)\n'%s'\n%s\n",
			width,
			utf8.RuneCountInString(output[0]),
			output[0],
			str)
	}

}

func TestRenderDesc(t *testing.T) {
	width := 10
	event := model.Event{
		Name:        "Event 2",
		Start:       time.Now().Add(2 * time.Hour),
		End:         time.Now().Add(3 * time.Hour),
		Description: "this is the desc for the test",
	}

	str := renderDescription(event, width)
	output := strings.Split(str, "\n")
	if len(output) == 0 {
		t.Fatalf("No lines in renderEvent (%+v)\n", event)
	}
	if utf8.RuneCountInString(output[0]) != width {
		t.Fatalf("Wrong width: Expected(%d), Got(%d)\n'%s'\n%s\n", width, len(output[0]), output[0], str)
	}

}

func TestRenderDay(t *testing.T) {
	width := 30
	t_time, err := time.Parse(time.RFC3339, "2024-05-31T14:20:52Z")
	if err != nil {
		t.Fatalf("Wrong time format")
	}
	_, month, date := t_time.Date()
	output := renderDay(t_time, width, date, month)

	if len(output) == 0 {
		t.Fatalf("Did not render day: %s, Month: %d, Date %d\n", t_time.Format(time.RFC3339), month, date)
	}
	split := strings.Split(output, "\n")
	if len(split) != 1 {
		t.Fatalf("Wrong height for render(%d): %+v\n", len(split), split)
	}
	if utf8.RuneCountInString(output) != width {
		t.Fatalf("Wrong width for the render(%d): '%s'\n", utf8.RuneCountInString(output), output)
	}
	if !strings.Contains(output, "Lørdag 31. mai") {
		t.Fatalf("Wrong format for day, date and month: %s\n", output)
	}
}
