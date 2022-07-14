package main

//importación múltiple
import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"contact.com/src/routes"
)
   
func main() {
	r := chi.NewRouter() //instancia de enrutador chi
	r.Use(middleware.Logger)

	routes.RoutesProgram(r) //llamado a las rutas de nuestro proyecto

	http.ListenAndServe(":3000", r) //puerto escuchado por nustro proyecto
}