package sessions

import (
	"fmt"
	"time"

	"github.com/kind84/golang-course/mongo/ex3/models"
)

var DbUsers = map[string]models.User{}       // user ID, user
var DbSessions = map[string]models.Session{} // session ID, session
var DbSessionsCleaned time.Time

const SessionLength int = 30

func CleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	ShowSessions()              // for demonstration purposes
	for k, v := range DbSessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(DbSessions, k)
		}
	}
	DbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	ShowSessions()             // for demonstration purposes
}

// for demonstration purposes
func ShowSessions() {
	fmt.Println("********")
	for k, v := range DbSessions {
		fmt.Println(k, v.Un)
	}
	fmt.Println("")
}
