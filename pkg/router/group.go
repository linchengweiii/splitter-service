package router

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/linchengweiii/splitter/pkg/group"
    "github.com/linchengweiii/splitter/pkg/expense"
)

type GroupRouter interface {
    GetGroup(w http.ResponseWriter, r *http.Request)

    PostExpense(w http.ResponseWriter, r *http.Request)
    PatchExpense(w http.ResponseWriter, r *http.Request)
    DeleteExpense(w http.ResponseWriter, r *http.Request)
}

type GroupRouterImpl struct {
    groupId         string
    groupService    *group.Service
    expenseService  *expense.Service
}

func (router *GroupRouterImpl) GetGroup(w http.ResponseWriter, r *http.Request) {
    group, err := router.groupService.Read(router.groupId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    g, err := json.Marshal(group)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(g)
}

func (router *GroupRouterImpl) PostExpense(w http.ResponseWriter, r *http.Request) {
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

func (router *GroupRouterImpl) PatchExpense(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    expenseId := vars["expenseId"]

    var expenseInput expense.Expense
    err := json.NewDecoder(r.Body).Decode(&expenseInput)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    expenseOutput, err := router.expenseService.Update(expenseId, expenseInput)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    g, err := router.groupService.Read(router.groupId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    for i, expense := range g.Expenses {
        if expense.Id == expenseId {
            g.Expenses[i] = group.Expense{
                Id: expenseOutput.Id,
                Description: expenseOutput.Description,
                Paid: expenseOutput.Paid,
                Owed: expenseOutput.Owed,
            }
            break
        }
    }

    e, err := json.Marshal(expenseOutput)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(e)
}

func (router *GroupRouterImpl) DeleteExpense(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    expenseId := vars["expenseId"]

    err := router.expenseService.Delete(expenseId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    group, err := router.groupService.Read(router.groupId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    for i, expense := range group.Expenses {
        if expense.Id == expenseId {
            group.Expenses = append(group.Expenses[:i], group.Expenses[i+1:]...)
            break
        }
    }

    router.groupService.Update(group)

    w.WriteHeader(http.StatusNoContent)
}

func NewGroupRouter(
    groupId string,
    groupService *group.Service,
    expenseService *expense.Service,
) GroupRouter {
    return &GroupRouterImpl{
        groupId,
        groupService,
        expenseService,
    }
}
