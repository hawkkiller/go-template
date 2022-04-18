package template

import "golang.org/x/crypto/bcrypt"

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

func (u *UpdateUserDTO) NewUser(oldUser *User) (*User, error) {
	ep, err := bcrypt.GenerateFromPassword([]byte(u.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return &User{}, nil
	}
	user := &User{
		Username: u.Username,
		Email:    u.Email,
		Password: string(ep),
		ID:       u.ID,
	}
	if user.Username == "" {
		user.Username = oldUser.Username
	}
	if user.Email == "" {
		user.Email = oldUser.Email
	}
	if user.Password == "" {
		user.Password = oldUser.Password
	}
	return user, nil
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

type GetUserByEmailAndPasswordDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	ID          string `json:"id" validate:"required"`
	Username    string `json:"username,omitempty" `
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	OldPassword string `json:"old_password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
}
