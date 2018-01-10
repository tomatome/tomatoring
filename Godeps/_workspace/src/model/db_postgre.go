package model

import (
	"database/sql"
	"log"
	"time"
)

//import (
//	"database/sql"
//	"fmt"
//	"log"

//	_ "github.com/lib/pq"
//)

type Users struct {
	uid      int
	userName string
	nickName string
	userPic  string
	openId   string
	authData string
	created  time.Time
	updated  time.Time
}

func (u *Users) TestSample() {
	u.uid = 2
	u.nickName = "john"
	u.userName = "hello"
}

func (u *Users) dbInsert(db *sql.DB) error {
	//插入数据
	stmt, err := db.Prepare("INSERT INTO \"WX_USERS\"(uid,\"userName\",\"nickName\",\"userPic\",\"openId\",\"authData\",created,updated) VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING uid")
	checkErr(err)

	now := time.Now()

	res, err := stmt.Exec(u.uid, u.userName, u.nickName, u.userPic, u.openId, u.authData, now, now)
	//这里的三个参数就是对应上面的$1,$2,$3了

	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	log.Println("dbInsert rows affect:", affect)
	return err
}
func (u *Users) dbDelete(db *sql.DB) error {
	//删除数据
	stmt, err := db.Prepare("delete from \"WX_USERS\" where uid=$1")
	checkErr(err)

	res, err := stmt.Exec(u.uid)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	log.Println("dbDelete rows affect:", affect)
	return err
}

func (u *Users) dbSelect(db *sql.DB) error {
	//查询数据
	rows, err := db.Query("SELECT uid,\"userName\",\"nickName\",\"userPic\",\"openId\",\"authData\", created, updated FROM \"WX_USERS\" where uid=2")
	checkErr(err)

	println("-----------")
	for rows.Next() {
		var u Users
		err = rows.Scan(&u.uid, &u.userName, &u.nickName, &u.userPic, &u.openId, &u.authData, &u.created, &u.updated)
		checkErr(err)
		log.Println("uid = ", u.uid, "\nname = ", u.userName, "\ndep = ", u.nickName, "\ncreated = ", u.created, "\n-----------")
	}
	return err
}
func (u *Users) dbUpdate(db *sql.DB) error {
	//更新数据
	stmt, err := db.Prepare("update \"WX_USERS\" set userName=$1 where uid=$2")
	checkErr(err)

	res, err := stmt.Exec("ficow", u.uid)

	affect, err := res.RowsAffected()
	checkErr(err)
	log.Println("dbUpdate rows affect:", affect)

	return err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
