package expense

import "github.com/google/uuid"

type Repository interface {
	Create(expense IdentifiableExpense) error
	Read(id string) (IdentifiableExpense, error)
	Update(expense IdentifiableExpense) error
	Delete(id string) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository}
}

func (s *Service) Create(
	description string,
	paid map[string]float64,
	owed map[string]float64,
) (*IdentifiableExpense, error) {
	id := uuid.New()
	expense := &IdentifiableExpense{
		Id:   id.String(),
		Description: description,
		Paid: paid,
		Owed: owed,
	}
	err := s.repository.Create(*expense)
	if err != nil {
		return nil, err
	}

	return expense, nil
}

func (s *Service) Read(id string) (IdentifiableExpense, error) {
	return s.repository.Read(id)
}

func (s *Service) Update(expense IdentifiableExpense) error {
	return s.repository.Update(expense)
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}
