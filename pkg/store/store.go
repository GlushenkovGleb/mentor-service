package store

import (
	"database/sql"
	"fmt"
	"github.com/GlushenkovGleb/mentor-service/pkg/model"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

const initSt = `CREATE TABLE IF NOT EXISTS mentors(
	   id serial PRIMARY KEY,
	   name varchar(255) NOT NULL,
	   family varchar(255) NOT NULL,
	   description varchar(255) NOT NULL,
	   course varchar(255) NOT NULL,
	   status varchar(255) NOT NULL,
	   photo_url varchar(255) NOT NULL
	)`

type Store struct {
	db *sql.DB
}

func NewStore() *Store {
	db := mustConnect()

	fmt.Println("Successfully connected!")

	return &Store{db: db}
}

func mustConnect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(initSt)
	if err != nil {
		panic(err)
	}

	return db
}

func (s *Store) SaveMentor(m model.CreateMentorRequest) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	query := `INSERT INTO "mentors"("name", "family", "course", "description", "status", "photo_url") values ($1, $2, $3, $4, $5, $6) returning id;`
	_, err = tx.Exec(query, m.Name, m.Family, m.Course, m.Description, m.Status, m.PhotoURL)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateMentorStatus(mentorID int, status string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	query := `update mentors set status = $1 where id = $2`
	_, err = tx.Exec(query, status, mentorID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetMentors() ([]model.Mentor, error) {
	rows, err := s.db.Query(`SELECT * FROM mentors WHERE status = 'APPROVED'`)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	mentors := []model.Mentor{}
	for rows.Next() {
		var m model.Mentor
		err := rows.Scan(&m.ID, &m.Name, &m.Family, &m.Course, &m.Description, &m.Status, &m.PhotoURL)
		if err != nil {
			return nil, err
		}
		mentors = append(mentors, m)
	}

	return mentors, nil
}
