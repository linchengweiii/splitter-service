package router

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/linchengweiii/splitter/pkg/expense"
)

type ExpenseRouter interface {
    GetExpense(w http.ResponseWriter, r *http.Request)
}

type ExpenseRouterImpl struct {
    expenseService *expense.Service
}

func (router *ExpenseRouterImpl) GetExpense(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    expenseId := vars["expenseId"]

    expense, err := router.expenseService.Read(expenseId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    e, err := json.Marshal(expense)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(e)
}

func NewExpenseRouter(expenseService *expense.Service) ExpenseRouter {
    return &ExpenseRouterImpl{expenseService}
}
