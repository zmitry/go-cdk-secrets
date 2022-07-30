package hello

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"log"
)

type EchoRepository struct {
	sql *sql.DB
}

func NewEchoRepository(db *sql.DB) *EchoRepository {
	return &EchoRepository{
		sql: db,
	}
}

type Result struct {
	Data string
}

func (s *EchoRepository) Generate() (*Result, error) {
	sessionBytes := make([]byte, 16)
	if _, err := rand.Read(sessionBytes); err != nil {
		return nil, err
	}
	session := hex.EncodeToString(sessionBytes)
	return &Result{Data: session}, nil
}

func (s *EchoRepository) SaveMessage(ctx context.Context, ip string) (int, error) {
	_, err := s.sql.Exec("INSERT INTO messages (message) VALUES ($1)", ip)
	if err != nil {
		return 0, err
	}
	rows, err := s.sql.Query("SELECT count(*) FROM messages")
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var count int
	if rows.Next() {
		rows.Scan(&count)
	}
	log.Printf("%d messages", count)
	return count, nil
}

type EchoService struct {
	sessions map[string]bool
	repo     *EchoRepository
}

func NewEchoService(repo *EchoRepository) *EchoService {
	return &EchoService{
		sessions: make(map[string]bool),
		repo:     repo,
	}
}

type HelloRequest struct {
	Key string `query:"key"`
	IP  string `header:"IP"`
}

type HelloResponse struct {
	Session      string `header:"session"`
	MessageCount int    `json:"message_count"`
}

func (s *EchoService) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	session, err := s.repo.Generate()
	if err != nil {
		return nil, err
	}
	count, err := s.repo.SaveMessage(ctx, req.IP)
	if err != nil {
		return nil, err
	}
	s.sessions[session.Data] = true
	return &HelloResponse{
		Session:      session.Data,
		MessageCount: count,
	}, nil
}
