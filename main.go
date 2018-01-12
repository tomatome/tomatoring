package main

import (
	"fmt"
	"io"
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

	router.LoadHTMLGlob("pages/*.html")
	router.MaxMultipartMemory = 8 << 20
	router.GET("/", Handler)
	router.GET("/server/get", GetHandler)
	router.POST("/server/post", PostHandler)
	router.POST("/server/upload", UploadHandler)
	//router.PUT("/upload", UploadHandler)
	//router.DELETE("/server/delete", DeleteHandler)
	//监听端口
	os.Setenv("PORT", "12345")
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
func Handler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "文件上传下载页面"})

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
func UploadHandler(c *gin.Context) {
	log.Println("UploadHandler")
	name := c.PostForm("name")
	log.Println(name)
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	filename := header.Filename

	fmt.Println(file, err, filename)

	out, err := os.Create("pages/static/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusCreated, "upload successful")
	//c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully\n", file.Filename))
	return
}
func DeleteHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
	return
}
