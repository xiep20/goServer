/*
 * @Author: xieping
 * @Date: 2023-09-08 14:18:56
 * @LastEditors: xieping
 * @LastEditTime: 2023-09-08 18:26:19
 * @Description:
 */
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//欢迎接口组
	v1 := router.Group("/come")
	{
		v1.GET("/", helloWorld)
		v1.GET("/:name/*action", nameIsAction)
		v1.GET("/welcome", welcome)
	}

	// 表单接口
	v2 := router.Group("/form")
	{
		v2.POST("/formSubmit", formSubmit)
		v2.POST("/upload", upload)
		v2.POST("/multiUpload", multiUpload)
	}

	//html页
	//加载模板
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	v3 := router.Group("/views/")
	{
		v3.GET("html/:htmlName", getHtml)
		v3.GET("tmpl/:htmlName", getTmpl)
	}

	// 静态文件服务
	// 显示当前文件夹下的所有文件/或者指定文件
	router.StaticFS("/showDir", http.Dir("."))
	router.StaticFS("/files", http.Dir("/upload"))

	//Static提供给定文件系统根目录中的文件。
	//router.Static("/files", "/bin")
	router.StaticFile("/image", "./upload/aa.svg")
	//静态资源
	router.Static("/static", "./static")

	//重定向
	v4 := router.Group("/redirect/")
	{
		v4.GET("bd", redirectBD)
	}

	router.Run(":8000")
}

// 默认
func helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

// 他是谁
func nameIsAction(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

// 欢迎
func welcome(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest") //可设置默认值
	//nickname := c.Query("nickname") // 是 c.Request.URL.Query().Get("nickname") 的简写
	c.String(http.StatusOK, fmt.Sprintf("Hello %s ", name))
}

// 表单提交
func formSubmit(c *gin.Context) {
	type1 := c.DefaultPostForm("type", "alert") //可设置默认值
	username := c.PostForm("username")
	password := c.PostForm("password")

	//hobbys := c.PostFormMap("hobby")
	//hobbys := c.QueryArray("hobby")
	hobbys := c.PostFormArray("hobby")

	c.String(http.StatusOK, fmt.Sprintf("type is %s, username is %s, password is %s,hobby is %v", type1, username, password, hobbys))

}

// 文件上传
func upload(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, "upload/"+file.Filename)

	/*
	   也可以直接使用io操作，拷贝文件数据。
	   out, err := os.Create(filename)
	   defer out.Close()
	   _, err = io.Copy(out, file)
	*/

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

// router.MaxMultipartMemory = 8 << 20 // 8 MiB
// 多文件上传
func multiUpload(c *gin.Context) {

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["files"]

	for _, file := range files {
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}

	c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files ", len(files)))
}

// 请求html
func getTmpl(c *gin.Context) {
	htmlName := c.Param("htmlName")
	//根据完整文件名渲染模板，并传递参数
	c.HTML(http.StatusOK, htmlName+".tmpl", gin.H{
		"title": htmlName,
	})
}

func getHtml(c *gin.Context) {
	htmlName := c.Param("htmlName")
	//返回html
	c.HTML(http.StatusOK, htmlName+".html", gin.H{"title": htmlName})
}

func redirectBD(c *gin.Context) {
	//支持内部和外部的重定向
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
}
