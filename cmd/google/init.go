package google

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kmuju/TuiCalendar/cmd/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func GetService() (*calendar.Service, error) {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		return &calendar.Service{}, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
		return &calendar.Service{}, err
	}
	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	return srv, err
}

func GetInfo(srv *calendar.Service) ([]*calendar.Event, error) {
	// srv, err := GetService()
	// if err != nil {
	// 	log.Printf("Unable to retrieve Calendar client: %v", err)
	// 	return []*calendar.Event{}, err
	// }

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
		return []*calendar.Event{}, err
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		return []*calendar.Event{}, nil
	}

	return events.Items, nil
}

func createEvents(e []*calendar.Event) []model.Event {
	events := make([]model.Event, len(e))
	for i, calendarEvent := range e {
		starttime := calendarEvent.Start.DateTime
		if starttime == "" {
			starttime = calendarEvent.Start.Date
		}
		start, err := time.Parse(time.RFC3339, starttime)
		if err != nil {
			continue
		}
		endtime := calendarEvent.End.DateTime
		if endtime == "" {
			endtime = calendarEvent.End.Date
		}
		end, err := time.Parse(time.RFC3339, endtime)
		if err != nil {
			continue
		}
		event := model.Event{
			Id:          calendarEvent.Id,
			Name:        calendarEvent.Summary,
			Description: calendarEvent.Description,
			Status:      calendarEvent.Status,
			Start:       start,
			End:         end,
		}
		events[i] = event
	}
	return events
}
