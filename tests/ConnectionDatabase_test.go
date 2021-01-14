package tests

import "testing"
import "github.com/luis16121013/libreria/storage"

func TestConnectionDB(t *testing.T) {
	repository, _ := storage.NewMyslRepository()
	err := repository.Repo.Ping()
	if err != nil {
		t.Fatal("no connected DB")
	}
}
