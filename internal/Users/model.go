package users

import "fmt"

type User struct {
	UUID  string
	Name  string
	Email string
	Phone string
	CV    string
}

func (u *User) CreateNewUsers(uuid, name, email, phone, cv_path string) error {
	if uuid == u.UUID || email == u.Email {
		fmt.Errorf("Such a user already exists")
	}
	u.UUID = uuid
	u.Name = name
	u.Email = email
	u.Phone = phone
	u.CV = cv_path
	return nil
}
