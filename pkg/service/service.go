package service

import (
	"github.com/GlushenkovGleb/mentor-service/pkg/model"
	"github.com/GlushenkovGleb/mentor-service/pkg/store"
)

type Service struct {
	st *store.Store
}

func NewService(st *store.Store) *Service {
	return &Service{st: st}
}

func (s *Service) SaveMentor(mentor model.CreateMentorRequest) error {
	mentor.Status = "IN_REVIEW"
	return s.st.SaveMentor(mentor)
}

func (s *Service) UpdateMentorStatus(mentorID int, status string) error {
	return s.st.UpdateMentorStatus(mentorID, status)
}

func (s *Service) GetMentors() ([]model.Mentor, error) {
	return s.st.GetMentors()
}
