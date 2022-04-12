package main

import (
	"dev11/config"
	"dev11/event"
	"log"
	"net/http"
	"strconv"
	"time"
)

var events = make(map[int]event.Event)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", createEvent)
	mux.HandleFunc("/update_event", updateEvent)
	mux.HandleFunc("/delete_event", deleteEvent)
	mux.HandleFunc("/event_for_day", eventForDay)
	mux.HandleFunc("/event_for_week", eventForWeek)
	mux.HandleFunc("/event_for_month", eventForMonth)

	log.Println("Запуск сервера")
	err := http.ListenAndServe(config.GetConfig().Port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var e event.Event
	var err error

	e.Id, err = strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		errorBadRequest(err, w)
		return
	}
	e.Login = r.URL.Query().Get("login")
	e.Date, err = time.Parse("02-01-2006", r.URL.Query().Get("date"))
	if err != nil {
		errorBadRequest(err, w)
		return
	}
	e.Text = r.URL.Query().Get("text")

	if _, ok := events[e.Id]; ok {
		errorServiceUnavailable(w, "Событие с таким id уже существует")
		return
	}

	events[e.Id] = event.NewEvent(e.Id, e.Login, e.Date, e.Text)
	j, err := events[e.Id].MarshalJson()
	if err != nil {
		delete(events, e.Id)
		errorInternalServerError(err, w)
		return
	}

	log.Println("Событие записано")
	w.Write(returnResult(string(j)))
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	var e event.Event
	var reservCopy event.Event
	var err error

	e.Id, err = strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		errorBadRequest(err, w)
		return
	}
	e.Login = r.URL.Query().Get("login")
	e.Date, err = time.Parse("02-01-2006", r.URL.Query().Get("date"))
	if err != nil {
		errorBadRequest(err, w)
		return
	}
	e.Text = r.URL.Query().Get("text")

	if _, ok := events[e.Id]; !ok {
		errorServiceUnavailable(w, "События с таким id не существует")
		return
	}

	reservCopy = events[e.Id]
	delete(events, e.Id)
	events[e.Id] = event.NewEvent(e.Id, e.Login, e.Date, e.Text)
	j, err := events[e.Id].MarshalJson()
	if err != nil {
		delete(events, e.Id)
		events[reservCopy.Id] = reservCopy
		errorInternalServerError(err, w)
		return
	}

	log.Println("Событие изменено")
	w.Write(returnResult(string(j)))
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	delID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		errorBadRequest(err, w)
		return
	}

	if _, ok := events[delID]; !ok {
		errorServiceUnavailable(w, "События с таким id не существует")
		return
	}

	delete(events, delID)
	log.Println("Событие удалено")
	w.Write(returnResult("Событие удалено"))
}

func eventForDay(w http.ResponseWriter, r *http.Request) {
	str := ""
	for _, ev := range events {
		if ev.Date.Sub(time.Now()).Hours() < 24 {
			str += strconv.Itoa(ev.Id) + " " + ev.Login + " " + ev.Date.String() + " " + ev.Text + "\n"
		}
	}

	log.Println("Выведены события за день")
	w.Write(returnResult(str))
}

func eventForWeek(w http.ResponseWriter, r *http.Request) {
	str := ""
	for _, ev := range events {
		if ev.Date.Sub(time.Now()).Hours() < 168 {
			str += strconv.Itoa(ev.Id) + " " + ev.Login + " " + ev.Date.String() + " " + ev.Text + "\n"
		}
	}

	log.Println("Выведены события за неделю")
	w.Write(returnResult(str))
}

func eventForMonth(w http.ResponseWriter, r *http.Request) {
	str := ""
	for _, ev := range events {
		if ev.Date.Sub(time.Now()).Hours() < 720 {
			str += strconv.Itoa(ev.Id) + " " + ev.Login + " " + ev.Date.String() + " " + ev.Text + "\n"
		}
	}

	log.Println("Выведены события за месяц")
	w.Write(returnResult(str))
}

func errorBadRequest(err error, w http.ResponseWriter) {
	log.Println(err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(returnError(err.Error()))
}

func errorServiceUnavailable(w http.ResponseWriter, str string) {
	log.Println(str)
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write(returnError(str))
}

func errorInternalServerError(err error, w http.ResponseWriter) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(returnError(err.Error()))
}

func returnError(err string) []byte {
	return []byte("{\"error\": " + err + "}")
}

func returnResult(res string) []byte {
	return []byte("{\"result\": " + res + "}")
}
