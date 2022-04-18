package template

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *CreateUserHashedDTO) NewUser() *User {
	return &User{
		Username: u.Username,
		Email:    u.Email,
		Password: string(u.Password),
	}
}

func (u *CreateUserDTO) Hashed(hashedPassword []byte) *CreateUserHashedDTO {
	return &CreateUserHashedDTO{
		Username: u.Username,
		Email:    u.Email,
		Password: hashedPassword,
	}
}

type CreateUserDTO struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateUserHashedDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}
