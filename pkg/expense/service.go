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
	expense Expense,
) (IdentifiableExpense, error) {
	id := uuid.New()
	expenseWithId := &IdentifiableExpense{
		Id:   id.String(),
		Description: expense.Description,
		Paid: expense.Paid,
		Owed: expense.Owed,
	}
	err := s.repository.Create(*expenseWithId)
	if err != nil {
		return IdentifiableExpense{}, err
	}

	return *expenseWithId, nil
}

func (s *Service) Read(id string) (IdentifiableExpense, error) {
	return s.repository.Read(id)
}

func (s *Service) Update(id string, expense Expense) (IdentifiableExpense, error) {
	expenseWithId := IdentifiableExpense{
		Id: id,
		Description: expense.Description,
		Paid: expense.Paid,
		Owed: expense.Owed,
	}
	err := s.repository.Update(expenseWithId)
	if err != nil {
		return IdentifiableExpense{}, err
	}
	return expenseWithId, nil

}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}
