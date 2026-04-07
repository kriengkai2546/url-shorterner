package url

import (
	"fmt"
	"math/rand"
	"os"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateURL(userID int, longURL string) (*CreateURLResponse, error) {
	shortCode := generateCode(6)

	u, err := s.repo.Save(userID, longURL, shortCode)
	if err != nil {
		return nil, err
	}

	baseURL := os.Getenv("BASE_URL")
	return &CreateURLResponse{
		ShortCode: u.ShortCode,
		ShortURL:  fmt.Sprintf("%s/%s", baseURL, u.ShortCode),
		LongURL:   u.LongURL,
	}, nil
}

func (s *Service) GetOriginalURL(code string) (string, error) {
	u, err := s.repo.FindByShortCode(code)
	if err != nil {
		return "", err
	}

	return u.LongURL, nil
}

func (s *Service) GetUserURLs(userID int) ([]URL, error) {
	return s.repo.FindByUserID(userID)
}

func (s *Service) DeleteURL(id, userID int) error {
	return s.repo.Delete(id, userID)
}

func generateCode(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}