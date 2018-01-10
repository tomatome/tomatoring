// db_test.go
package model

import (
	"fmt"
	"testing"
)

func TestInitDB(t *testing.T) {
	bc := InitDB()
	fmt.Println(bc.c)

	u := new(Users)
	u.uid = 2
	u.nickName = "john"
	u.userName = "hello"
	bc.SetModelDao(u)
}
