package handlers

import (
	"fmt"
	"time"
)

func (app *App) ClearSession() {
	Sessions, err := app.sessionService.GetAllSessionsTime()
	// fmt.Println("SESSIONS: ", Sessions)
	if err != nil {
		fmt.Println("error of getting all sesions time", err.Error())
	}

	for {
		time.Sleep(time.Second)
		for i, v := range Sessions {
			if v.Expiry.Before(time.Now()) {
				err := app.sessionService.DeleteSession(v.Token)
				if i == len(Sessions)-1 {
					Sessions = Sessions[:i]
				} else {
					Sessions = append(Sessions[:i], Sessions[i+1:]...)
				}
				if err != nil {
					fmt.Println("Session delete was failed", err.Error())
				} else {
					fmt.Printf("session for %s was deleted", v.Username)
				}
			}
		}
	}
}
