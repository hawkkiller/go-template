package template

import "context"

type Storage interface {
	Create(ctx context.Context, user *User) error
}
