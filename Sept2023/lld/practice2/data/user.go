package data

type User struct {
	id           string
	name         string
	email        string
	mobileNumber string
}

func CreateUser(id string, name string, email string, mobileNumber string) *User {
	return &User{
		id:           id,
		name:         name,
		email:        email,
		mobileNumber: mobileNumber,
	}
}

func (u *User) Id() string {
	return u.id
}
