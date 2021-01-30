package model

type (
	User struct {
		UserName string
		Roles    []string
	}

	Organization struct {
		Organization string
		Users        []User
	}
)

func NewOrganization(organization string) Organization {
	return Organization{
		organization,
		make([]User, 0),
	}
}
func NewUser(userName string) User {
	return User{
		userName,
		make([]string, 0),
	}
}
