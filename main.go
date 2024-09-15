package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/linchengweiii/splitter/pkg/expense"
	"github.com/linchengweiii/splitter/pkg/group"
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

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		g, err := json.Marshal(defaultGroup)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(g)
	}).Methods(http.MethodGet)

	router.HandleFunc("/expense", func(w http.ResponseWriter, r *http.Request) {
		var expenseInput struct {
			Description string             `json:"description"`
			Paid		map[string]float64 `json:"paid"`
			Owed		map[string]float64 `json:"owed"`
		}
		err := json.NewDecoder(r.Body).Decode(&expenseInput)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		expenseOutput, err := expenseService.Create(
			expenseInput.Description,
			expenseInput.Paid,
			expenseInput.Owed,
		)
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

		defaultGroup.Expenses = append(defaultGroup.Expenses, expense)
		groupService.Update(*defaultGroup)

		e, err := json.Marshal(expenseOutput)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(e)
	}).Methods(http.MethodPost)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
