package users

import "github.com/luis16121013/libreria/storage"

type ServiceUser interface {
	SaveUser(username, pwd, role, email string) (*User, error)
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}

type ModelUser struct{}

//method return New Model User
func NewUser() *ModelUser {
	return &ModelUser{}
}

type Usrs []User

//method for saved users
func (u *ModelUser) SaveUser(username, pwd, role, email string) (*User, error) {
	stmt := "INSERT INTO userTable(username,pwd,role,email) VALUES(?,?,?,?)"
	db, err := storage.NewMyslRepository()

	if err != nil {
		return nil, err
	}

	r, err := db.Exec(stmt, username, pwd, role, email)
	if err != nil {
		return nil, err
	}

	id, _ := r.LastInsertId()
	user := &User{
		Id:       id,
		Username: username,
		Password: pwd,
		Role:     role,
		Email:    email,
	}
	return user, nil
}

func (u *ModelUser) GetUsers() ([]User, error) {
	stmt := "SELECT * FROM userTable"
	listUsers := Usrs{}

	db, err := storage.NewMyslRepository()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := User{}
		rows.Scan(&u.Id, &u.Username, &u.Password, &u.Role, &u.Email)
		listUsers = append(listUsers, u)
	}
	return listUsers, nil
}
