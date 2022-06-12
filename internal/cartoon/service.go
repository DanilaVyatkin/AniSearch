package cartoon

import (
	"context"
	"fmt"
)

type Service struct {
	repository Storage
}

func NewService(repository Storage) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetAll(ctx context.Context) (c []Cartoon, err error) {
	all, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all cartoon dto to error: %v", err)
	}

	return all, nil
}
