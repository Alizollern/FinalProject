package users

import "context"

type Service interface {
	CreateNewUsers(ctx context.Context, uuid, name, email, phone, cv_path string)
	DeleteUsers(ctx context.Context, uuid string)
	UpdateUsers(ctx context.Context, uuid, data string)
	GetUser(ctx context.Context, uuid string)
}

type service struct {
	// storage Store
}
