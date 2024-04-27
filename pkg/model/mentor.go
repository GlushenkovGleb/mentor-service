package model

type CreateMentorRequest struct {
	Name        string `json:"name"`
	Family      string `json:"family"`
	Description string `json:"description"`
	Course      string `json:"course"`
	Status      string `json:"status"`
	PhotoURL    string `json:"photoURL"`
}

type UpdateStatusMentorRequest struct {
	Status string `json:"status"`
}

type Mentor struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Family      string `json:"family"`
	Description string `json:"description"`
	Course      string `json:"course"`
	Status      string `json:"status"`
	PhotoURL    string `json:"photoURL"`
}
