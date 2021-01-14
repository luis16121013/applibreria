package tests

import "testing"
import "github.com/luis16121013/libreria/users"

var (
	username = "carlos"
	password = "secretpwd"
	role     = "admin"
	email    = "email@test.com"
)

func TestCreateUser(t *testing.T) {
	UserModel := users.NewUser()
	u, _ := UserModel.SaveUser(username, password, role, email)
	if u.Username != username || u.Password != password || u.Role != role || u.Email != email {
		t.Fatal("user diferent")
	}
}
