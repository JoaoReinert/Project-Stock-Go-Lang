package entities

type User struct {
	ID int64 `json:"id,omitempty"`

	Name string `json:"name"`

	Email string `json:"email"`

	Password string `json:"password"`

	Salt string
}

type UserLogin struct {
	Email string `json:"email"`

	Password string `json:"password"`
}

type UserToken struct {
	Token string `json:"token"`
}

type EmployeeSubject struct {
	ID *int64
}
