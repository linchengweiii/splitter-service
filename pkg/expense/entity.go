package expense

type Expense struct {
    Description     string              `json:"description"`
    Paid            map[string]float64  `json:"paid"`
    Owed            map[string]float64  `json:"owed"`
}

type IdentifiableExpense struct {
    Id              string              `json:"id"`
    Description     string              `json:"description"`
    Paid            map[string]float64  `json:"paid"`
    Owed            map[string]float64  `json:"owed"`
}
