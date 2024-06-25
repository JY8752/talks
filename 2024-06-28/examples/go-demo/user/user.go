package user

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) User {
	return User{name: name, age: age}
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetAge() int {
	return u.age
}

func UpdateName(name string, user *User) {
	user.name = name
}
