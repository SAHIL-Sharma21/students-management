package types

// carting student struc
type Student struct {
	Id    int
	Name  string `validate:"required" json:"name"`
	Email string `validate:"required" json:"email"`
	Age   int    `validate:"required" json:"age"`
}
