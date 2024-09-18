package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/linchengweiii/splitter/pkg/expense"
	"github.com/linchengweiii/splitter/pkg/group"
	"github.com/linchengweiii/splitter/pkg/router"
)

func main() {
	groupRepository := group.NewInMemoryRepository()
	groupService := group.NewService(groupRepository)
	defaultGroup, err := groupService.Create("default")
	if err != nil {
		log.Fatal(err)
	}

	expenseRepository := expense.NewInMemoryRepository()
	expenseService := expense.NewService(expenseRepository)

	groupRouter := router.NewGroupRouter(defaultGroup.Id, groupService, expenseService)
	expenseRouter := router.NewExpenseRouter(expenseService)

	router := mux.NewRouter()

	router.HandleFunc(
		"/group",
		groupRouter.GetGroup,
	).Methods(http.MethodGet)

	router.HandleFunc(
		"/group/expense",
		groupRouter.PostExpense,
	).Methods(http.MethodPost)

	router.HandleFunc(
		"/group/expense/{expenseId}",
		groupRouter.PatchExpense,
	).Methods(http.MethodPatch)

	router.HandleFunc(
		"/group/expense/{expenseId}",
		groupRouter.DeleteExpense,
	).Methods(http.MethodDelete)

	router.HandleFunc(
		"/expense/{expenseId}",
		expenseRouter.GetExpense,
	).Methods(http.MethodGet)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
