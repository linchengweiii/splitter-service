package group

type Expense struct {
    Id          string              `json:"id"`
    Description string              `json:"description"`
    Paid        map[string]float64  `json:"paid"`
    Owed        map[string]float64  `json:"owed"`
}

type Group struct {
    Id          string      `json:"id"`
    Name        string      `json:"name"`
    Expenses    []Expense   `json:"expenses"`
}
