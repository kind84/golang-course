package controllers

import (
	"net/http"
	"time"

	"github.com/kind84/golang-course/mongo/ex3/models"
	"github.com/kind84/golang-course/mongo/ex3/sessions"
	"github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) models.User {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	c.MaxAge = sessions.SessionLength
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u models.User
	if s, ok := sessions.DbSessions[c.Value]; ok {
		s.LastActivity = time.Now()
		sessions.DbSessions[c.Value] = s
		u = sessions.DbUsers[s.Un]
	}
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := sessions.DbSessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		sessions.DbSessions[c.Value] = s
	}
	_, ok = sessions.DbUsers[s.Un]
	// refresh session
	c.MaxAge = sessions.SessionLength
	http.SetCookie(w, c)
	return ok
}
