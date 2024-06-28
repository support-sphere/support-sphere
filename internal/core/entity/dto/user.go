package dto

type RegisterUser struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
