package domain

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	City      string `json:"city"`
}

func NewUser(first_name, last_name, city string) *User {
	return &User{
		FirstName: first_name,
		LastName:  last_name,
		City:      city,
	}
}
