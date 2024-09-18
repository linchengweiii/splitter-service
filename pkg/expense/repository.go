package expense

import "errors"

type InMemoryRepository struct {
    expenses []IdentifiableExpense
}

func NewInMemoryRepository() *InMemoryRepository {
    return &InMemoryRepository{
        make([]IdentifiableExpense, 0),
    }
}

func (r *InMemoryRepository) Create(expense IdentifiableExpense) error {
    for _, e := range r.expenses {
        if e.Id == expense.Id {
            return errors.New("Expense already exists")
        }
    }
    r.expenses = append(r.expenses, expense)
    return nil
}

func (r *InMemoryRepository) Read(id string) (IdentifiableExpense, error) {
    for _, expense := range r.expenses {
        if expense.Id == id {
            return expense, nil
        }
    }
    return IdentifiableExpense{}, errors.New("Expense not found")
}

func (r *InMemoryRepository) Update(expense IdentifiableExpense) error {
    for i, e := range r.expenses {
        if e.Id == expense.Id {
            r.expenses[i] = expense
            return nil
        }
    }
    return errors.New("Expense not found")
}

func (r *InMemoryRepository) Delete(id string) error {
    for i, expense := range r.expenses {
        if expense.Id == id {
            r.expenses = append(r.expenses[:i], r.expenses[i+1:]...)
            return nil
        }
    }
    return errors.New("Expense not found")
}
