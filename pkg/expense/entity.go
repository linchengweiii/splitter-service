package expense

type IdentifiableExpense struct {
    Id              string              `json:"id"`
    Description     string              `json:"description"`
    Paid            map[string]float64  `json:"paid"`
    Owed            map[string]float64  `json:"owed"`
}
