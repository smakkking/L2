package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func (s *Server) EventByName(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	eventName, ok := m["event_name"]
	if !ok {
		makeJsonRespond(w, 400, jsonError("invalid request"))
		return

	}
	s.Mu.RLock()
	event, ok := s.Cache[eventName[0]]
	s.Mu.RUnlock()
	if !ok {
		makeJsonRespond(w, 500, jsonError("no event for this name"))
		return
	}
	data, err := json.Marshal(event)
	if err != nil {
		log.Println(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	makeJsonRespond(w, 200, jsonResult(string(data)))
}

func (s *Server) EventsForDay(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	day, ok := m["day"]
	if !ok {
		makeJsonRespond(w, 400, jsonError("invalid request"))
		return
	}
	handeledTime, err := strconv.Atoi(day[0])
	if err != nil {
		log.Println(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	result := make([]Event, 0)
	timeFrom := time.Unix(0, 0).Add(time.Duration(handeledTime) * 24 * time.Hour)
	timeTo := timeFrom.AddDate(0, 0, 1)
	s.Mu.RLock()
	for _, event := range s.Cache {
		if inTimeSpan(timeFrom, timeTo, event.Time) {
			result = append(result, event)
		}
	}
	s.Mu.RUnlock()
	data, err := json.Marshal(result)
	if err != nil {
		log.Panicln(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	makeJsonRespond(w, 200, jsonResult(string(data)))
}

func (s *Server) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	week, ok := m["week"]
	if !ok {
		makeJsonRespond(w, 400, jsonError("invalid request"))
		return
	}
	handeledTime, err := strconv.Atoi(week[0])
	if err != nil {
		log.Println(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	result := make([]Event, 0)
	timeFrom := time.Unix(0, 0).Add(time.Duration(handeledTime) * 24 * time.Hour * 7)
	timeTo := timeFrom.AddDate(0, 0, 7)
	s.Mu.RLock()
	for _, event := range s.Cache {
		if inTimeSpan(timeFrom, timeTo, event.Time) {
			result = append(result, event)
		}
	}
	s.Mu.RUnlock()
	data, err := json.Marshal(result)
	if err != nil {
		log.Panicln(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	makeJsonRespond(w, 200, jsonResult(string(data)))
}

func (s *Server) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	month, ok := m["month"]
	if !ok {
		makeJsonRespond(w, 400, jsonError("invalid request"))
		return
	}
	handeledTime, err := strconv.Atoi(month[0])
	if err != nil {
		log.Println(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	result := make([]Event, 0)
	s.Mu.RLock()
	for _, event := range s.Cache {
		if event.Time.Month() == time.Month(handeledTime) {
			result = append(result, event)
		}
	}
	s.Mu.RUnlock()
	data, err := json.Marshal(result)
	if err != nil {
		log.Panicln(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return
	}
	makeJsonRespond(w, 200, jsonResult(string(data)))
}

const (
	permissionError     int = 2
	valid               int = 1
	invalidData         int = 0
	internalServerError int = -1
)

func getDataFromRequest(r *http.Request) (Event, error) {
	event := Event{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return event, err
	}
	err = json.Unmarshal(data, &event)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (s *Server) validatePost(w http.ResponseWriter, event *Event, actionType string) int {
	s.Mu.RLock()
	data, ok := s.Cache[event.EventName]
	s.Mu.RUnlock()
	result := invalidData

	date := isValidDate(event)
	switch actionType {
	case "create":
		if !ok && date && event.EventName != "" {
			result = valid
		}
	case "update":
		if ok && date {
			result = valid
		}
	case "delete":
		if ok {
			if data.UserId == event.UserId {
				result = valid
			} else {
				result = permissionError
			}
		}
	default:
		result = internalServerError
	}
	return result
}

func (s *Server) validateAndRespond(w http.ResponseWriter, code int) bool {
	if code == valid {
		return true
	}
	switch code {
	case internalServerError:
		makeJsonRespond(w, 503, jsonError("internal server error"))
	case invalidData:
		makeJsonRespond(w, 400, jsonError("invalid data"))
	case permissionError:
		makeJsonRespond(w, 500, jsonError("permisson error"))
	}
	return false
}

func (s *Server) postRequestCheck(w http.ResponseWriter, r *http.Request, request string) (Event, error) {
	event := Event{}
	if r.Method != http.MethodPost {
		errorString := "method not allowed"
		makeJsonRespond(w, 500, jsonError(errorString))
		return event, fmt.Errorf(errorString)
	}
	event, err := getDataFromRequest(r)
	if err != nil {
		log.Println(err)
		makeJsonRespond(w, 503, jsonError("internal server error"))
		return event, err
	}
	validate := s.validatePost(w, &event, request)
	if !s.validateAndRespond(w, validate) {
		return event, fmt.Errorf("something being wrong")
	}
	return event, nil
}

func (s *Server) createAndUpdate(w http.ResponseWriter, r *http.Request, request string) {
	event, err := s.postRequestCheck(w, r, request)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.setEventCache(event)
	makeJsonRespond(w, 200, jsonResult("ok"))
}

func (s *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {
	s.createAndUpdate(w, r, "create")
}

func (s *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	s.createAndUpdate(w, r, "update")
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	event, err := s.postRequestCheck(w, r, "delete")
	if err != nil {
		return
	}
	s.deleteEventCache(event.EventName)
	makeJsonRespond(w, 200, jsonResult("ok"))
}
