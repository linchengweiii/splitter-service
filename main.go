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

	groupRouter := router.NewGroupRouter(defaultGroup.Id, groupService)
	expenseRouter := router.NewExpenseRouter(defaultGroup.Id, groupService, expenseService)

	router := mux.NewRouter()

	router.HandleFunc(
		"/",
		groupRouter.GetGroup,
	).Methods(http.MethodGet)

	router.HandleFunc(
		"/expense",
		expenseRouter.PostExpense,
	).Methods(http.MethodPost)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
