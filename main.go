package main

import (
	"github.com/labstack/echo/v4"
	"database/sql"
	"github.com/mattn/go-sqlite3"
)

type Event struct {
	Name string
	Local string
}

var events []Event

func createEvent() {
	events = append(events, Event{Name: "NRDOPL", Local: "Praça"})
	events = append(events, Event{Name: "Observação de Chuva de Meteoros", Local: "Cedai"})
}


func main(){
	createEvent()
	e := echo.New()
	e.GET("/eventos", getEvents)
	e.POST("/eventos", newEvents)
	e.Logger.Fatal(e.Start(":8080"))
}

func getEvents(c echo.Context) error {
	return c.JSON(200, events)
}

func newEvents(c echo.Context) error {
	event := new(Event)
	if err := c.Bind(event); err != nil {
		return err
	}
	events = append(events, *event)
	saveEvent(*event)
	return c.JSON(200, events)
}

func saveEvent(event Event) error {
	db, err := sql.Open("sqlite3", "events.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO events (name, local) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(event.Name, event.Local)
	return nil
}