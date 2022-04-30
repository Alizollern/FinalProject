package admins

type Admin struct {
	AUID         string
	Login        string
	Password     string
	Privilege_ID int
	IsActive     bool
}
