package handler

import (
	"encoding/json"
	"github.com/GlushenkovGleb/mentor-service/pkg/model"
	"github.com/GlushenkovGleb/mentor-service/pkg/service"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type Handler struct {
	serv *service.Service
}

func NewHandler(serve *service.Service) *Handler {
	return &Handler{serve}
}

func (h *Handler) SaveMentor(w http.ResponseWriter, r *http.Request) {
	var mentor model.CreateMentorRequest
	err := json.NewDecoder(r.Body).Decode(&mentor)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	err = h.serv.SaveMentor(mentor)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateMentorStatus(w http.ResponseWriter, r *http.Request) {
	var req model.UpdateStatusMentorRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	mentorIDRaw := chi.URLParam(r, "id")
	mentorID, err := strconv.Atoi(mentorIDRaw)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	err = h.serv.UpdateMentorStatus(mentorID, req.Status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetMentors(w http.ResponseWriter, r *http.Request) {
	res, err := h.serv.GetMentors()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}
