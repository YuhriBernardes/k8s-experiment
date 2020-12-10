package server

import (
	"encoding/json"
	"mykube/database"
	"mykube/model"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	db     *database.Db
	router *mux.Router
}

func NewServer(db *database.Db) Server {
	return Server{
		db:     db,
		router: mux.NewRouter(),
	}
}

func (srvr Server) CreateUser() {
	srvr.router.Methods(http.MethodPost).Path("/users").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := new(CreateUserReq)

		if err := json.NewDecoder(r.Body).Decode(body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newUser := &model.User{
			Name: body.Name,
			Age:  body.Age,
		}

		if err := srvr.db.Create(newUser); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{"id": newUser.ID})
	})
}

func (srvr Server) GetAll() {
	srvr.router.Methods(http.MethodGet).Path("/users").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := srvr.db.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(res)

	})
}

func (srvr Server) Start() {
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodOptions})
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"*"})
	creds := handlers.AllowCredentials()
	http.ListenAndServe(":3030", handlers.CORS(creds, methods, origins, headers)(srvr.router))

}
