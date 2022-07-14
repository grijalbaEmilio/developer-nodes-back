package routes

import (
	"contact.com/src/controllers"
	"github.com/go-chi/chi/v5"
)

func RoutesProgram(r *chi.Mux){

	r.Get("/getPrograms", controllers.GetPrograms)

	r.Post("/postProgram", controllers.PostProgram)

	r.Delete("/deleteProgram", controllers.DeleteProgram)
}