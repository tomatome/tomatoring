package main

import (
	"fmt"
	"log"
	"model"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type DataCenter struct {
	bc *model.DBCenter
}

var dc *DataCenter

func main() {
	dc = new(DataCenter)
	dc.bc = model.InitDB()
	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router := gin.Default()
	router.GET("/server/get", GetHandler)
	router.POST("/server/post", PostHandler)
	router.PUT("/server/put", PutHandler)
	router.DELETE("/server/delete", DeleteHandler)
	//监听端口
	//os.Setenv("PORT", "12345")
	log.Println("listening...", os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
	if err != nil {
		panic(err)
	}
}
func GetHandler(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "the key is not exist!"
	}
	u := new(model.Users)
	u.TestSample()
	dc.bc.SetModelDao(u)
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
	return
}
func PostHandler(c *gin.Context) {
	type JsonHolder struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	holder := JsonHolder{Id: 1, Name: "my name"}
	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, holder)
	return
}
func PutHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("put success!\n"))
	return
}
func DeleteHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
	return
}
