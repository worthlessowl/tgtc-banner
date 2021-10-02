package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/worthlessowl/tgtc-banner/backend/dictionary"
	"github.com/worthlessowl/tgtc-banner/backend/service"
)

func AddUser(w http.ResponseWriter, r *http.Request) {

	var p dictionary.User

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	p = User.AddUser(context.Background(), p)
	fmt.Fprintf(w, fmt.Sprint("success, id User : ", p.ID))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	idstring := r.URL.Query().Get("id")

	idInt64, err := strconv.ParseInt(idstring, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	p, err := User.GetUser(context.Background(), idInt64)
	if err != nil {
		// log.Fatal(err)
		fmt.Fprintf(w, err.Error())
	}

	val, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(val))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var p dictionary.User

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	User.DeleteUser(context.Background(), p.ID)
	fmt.Fprintf(w, "success")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var p dictionary.User

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	User.UpdateUser(context.Background(), p)

	fmt.Fprintf(w, "success")
}
}
