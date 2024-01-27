package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func usersHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case http.MethodGet:
			message := ResponseMessage{Message: "Respuesta desde verbo GET, sigueme para aprender mas..."}

			jsonData, err := json.Marshal(message)

			if err != nil {
				http.Error(w, "Error al codificar como JSON", http.StatusInternalServerError)
				return		
			}

			w.Header().Set("Content-Type","application/json")
			w.Write(jsonData)

		case http.MethodPost:

			users := [] User {
				{ID: 1, Username: "charlesRootsDev", Email: "some@gmail.com"},
				{ID: 2, Username: "charlesRoots", Email: "som2@gmail.com"},
			}
		
			jsonData, err := json.Marshal(users)
		
			if err != nil {
				http.Error(w, "Error al codificar como JSON", http.StatusInternalServerError)
				return		
			}
		
			w.Header().Set("Content-Type","application/json")
			w.Write(jsonData)
		
		default:
			// RESPUESTA PARA OTROS VERBOS (PUT, DELETE, PATCH, ETC...)
			w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main(){
	http.HandleFunc("/users",usersHandler)
	log.Println("Server listen on port 9000...")
	log.Fatal(http.ListenAndServe(":9000",nil))
}