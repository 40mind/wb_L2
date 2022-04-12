package event

import (
	"encoding/json"
	"time"
)

type Event struct {
	Id    int       `json:"id"`
	Login string    `json:"login"`
	Date  time.Time `json:"date"`
	Text  string    `json:"text"`
}

func NewEvent(id int, login string, date time.Time, text string) Event {
	return Event{
		Id:    id,
		Login: login,
		Date:  date,
		Text:  text,
	}
}

func (e Event) MarshalJson() ([]byte, error) {
	return json.Marshal(e)
}
