package main

import "fmt"
import (

	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Program struct{
	Name string
	Data string
}
   
func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("falta la correcta conección con la base de datos"))
	})

	r.Get("/getPrograms", getPrograms)

	r.Post("/postProgram", postProgram)

	r.Delete("/deleteProgram", deleteProgram)

	http.ListenAndServe(":3000", r)
}

func postProgram (w http.ResponseWriter, r *http.Request){
	request := map[string]string{}
	json.NewDecoder(r.Body).Decode(&request)
	program := Program{
		Name : request["Name"],
		Data : request["Data"],
	}
	json.NewEncoder(w).Encode(program)
}

func getPrograms (w http.ResponseWriter, r *http.Request){
	prog := Program{
		Name : "programaUno",
		Data : "dataProgramaUno",
	}

	progjosn, _ := json.Marshal(prog)
	progstring := string(progjosn)
	fmt.Println(progstring)
	json.NewEncoder(w).Encode(prog)
}

func deleteProgram (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("se eliminó "))
}