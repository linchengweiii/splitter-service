package router

import (
    "encoding/json"
    "net/http"

    "github.com/linchengweiii/splitter/pkg/group"
    "github.com/linchengweiii/splitter/pkg/expense"
)

type ExpenseRouter interface {
    PostExpense(w http.ResponseWriter, r *http.Request)
}

type ExpenseRouterImpl struct {
    groupId string
    groupService *group.Service
    expenseService *expense.Service
}

func(router *ExpenseRouterImpl) PostExpense(w http.ResponseWriter, r *http.Request) {
    var expenseInput expense.Expense
    err := json.NewDecoder(r.Body).Decode(&expenseInput)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return;
    }

    expenseOutput, err := router.expenseService.Create(expenseInput)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    expense := group.Expense{
        Id: expenseOutput.Id,
        Description: expenseOutput.Description,
        Paid: expenseOutput.Paid,
        Owed: expenseOutput.Owed,
    }

    group, err := router.groupService.Read(router.groupId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    group.Expenses = append(group.Expenses, expense)
    router.groupService.Update(group)

    e, err := json.Marshal(expenseOutput)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(e)
}

func NewExpenseRouter(groupId string, groupService *group.Service, expenseService *expense.Service) ExpenseRouter {
    return &ExpenseRouterImpl{
        groupId,
        groupService,
        expenseService,
    }
}
