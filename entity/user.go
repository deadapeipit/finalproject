package entity

import "time"

//User same struct as table
type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type UserUpdate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserGetComment struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUpdateOutput struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) ToUserUpdateOutput() *UserUpdateOutput {
	out := &UserUpdateOutput{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Age:       u.Age,
		UpdatedAt: u.UpdatedAt,
	}
	return out
}
