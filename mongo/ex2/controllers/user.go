package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/kind84/golang-course/mongo/ex2/models"
	"github.com/satori/go.uuid"

	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	storedUsers map[string]models.User
}

func NewUserController() *UserController {
	bs, err := ioutil.ReadFile("../stored-users.txt")
	if err != nil {
		return &UserController{map[string]models.User{}}
	}

	var su map[string]models.User
	err = json.Unmarshal(bs, &su)
	if err != nil {
		return &UserController{map[string]models.User{}}
	}

	return &UserController{su}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Fetch user
	u, ok := uc.storedUsers[id]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	id, _ := uuid.NewV4()
	u.ID = id.String()

	// store the user in map
	uc.storedUsers[u.ID] = u

	uj, _ := json.Marshal(u)
	fmt.Println(string(uj))

	suj, _ := json.Marshal(uc.storedUsers)

	if err := storeUsers(suj); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	delete(uc.storedUsers, id)

	suj, _ := json.Marshal(uc.storedUsers)

	err := storeUsers(suj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")
}

func storeUsers(users []byte) error {
	// store users to file
	store, err := os.Create("../stored-users.txt")
	if err != nil {
		return err
	}
	defer store.Close()

	_, err = store.Write(users)
	if err != nil {
		return err
	}

	return nil
}
