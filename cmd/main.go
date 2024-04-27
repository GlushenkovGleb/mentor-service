package main

import (
	"github.com/GlushenkovGleb/mentor-service/pkg/handler"
	"github.com/GlushenkovGleb/mentor-service/pkg/service"
	"github.com/GlushenkovGleb/mentor-service/pkg/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	st := store.NewStore()

	serv := service.NewService(st)

	h := handler.NewHandler(serv)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Post("/mentors", h.SaveMentor)
	r.Patch("/mentors/{id}", h.UpdateMentorStatus)
	r.Get("/mentors", h.GetMentors)
	http.ListenAndServe(":3000", r)
}
