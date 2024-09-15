package group

import (
	"github.com/google/uuid"
)

type Repository interface {
    Create(group Group) error
    Read(id string) (Group, error)
    Update(group Group) error
    Delete(id string) error
}

type Service struct {
    repository Repository
}

func NewService(repository Repository) *Service {
    return &Service{repository}
}

func (s *Service) Create(name string) (*Group, error) {
    id := uuid.New()
    group := &Group{
        Id: id.String(),
        Name: name,
        Expenses: []Expense{},
    }
    err := s.repository.Create(*group)
    if err != nil {
        return nil, err
    }

    return group, nil
}

func (s *Service) Read(id string) (Group, error) {
    return s.repository.Read(id)
}

func (s *Service) Update(group Group) error {
    return s.repository.Update(group)
}

func (s *Service) Delete(id string) error {
    return s.repository.Delete(id)
}
