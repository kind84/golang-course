package controllers

import (
	"html/template"
	"net/http"
	"time"

	"github.com/kind84/golang-course/mongo/ex3/models"
	"github.com/kind84/golang-course/mongo/ex3/sessions"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type GeneralController struct {
	tpl *template.Template
}

func NewGeneralController(tpl *template.Template) *GeneralController {
	return &GeneralController{tpl}
}

func (gc GeneralController) Index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	sessions.ShowSessions() // for demonstration purposes
	gc.tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func (gc GeneralController) Bar(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := getUser(w, req)
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	sessions.ShowSessions() // for demonstration purposes
	gc.tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func (gc GeneralController) Signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		// username taken?
		if _, ok := sessions.DbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessions.SessionLength
		http.SetCookie(w, c)
		sessions.DbSessions[c.Value] = models.Session{un, time.Now()}
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = models.User{un, bs, f, l, r}
		sessions.DbUsers[un] = u
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	sessions.ShowSessions() // for demonstration purposes
	gc.tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func (gc GeneralController) Login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := sessions.DbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessions.SessionLength
		http.SetCookie(w, c)
		sessions.DbSessions[c.Value] = models.Session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	sessions.ShowSessions() // for demonstration purposes
	gc.tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func (gc GeneralController) Logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(sessions.DbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// clean up dbSessions
	if time.Now().Sub(sessions.DbSessionsCleaned) > (time.Second * 30) {
		go sessions.CleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
