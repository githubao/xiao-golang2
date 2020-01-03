// gin框架
// author: baoqiang
// time: 2018/12/25 下午1:42
package lang

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"log"
)

// GET POST PUT PATCH DELETE OPTIONS
// 文件上传 分组路由 jsonP，参数处理

func RunGin() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := gin.Default()

	// set mode
	gin.SetMode(gin.ReleaseMode)

	// 使用中间件
	r.Use(func(ctx *gin.Context) {
		log.Printf("get request: %v", time.Now().Format(Layout))
		//fmt.Printf("get request: %v", time.Now())
		ctx.Next()
	})

	// 返回json
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 参数路由
	r.GET("/path/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.String(200, fmt.Sprintf("你的id是: %s", id))
	})

	// get参数
	r.GET("/get", func(ctx *gin.Context) {
		name := ctx.DefaultQuery("name", "xiaobao")
		ctx.Data(200, "text/plain; charset=utf-8",
			[]byte(fmt.Sprintf("你的名字叫: %s", name)))
	})

	// 分组
	g := r.Group("group")
	g.GET("/morning", func(ctx *gin.Context) {
		ctx.String(200, "morning from group")
	})
	g.GET("/afternoon", func(ctx *gin.Context) {
		ctx.String(200, "afternoon from group")
	})

	// 默认路由
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"msg": "page not found",
		})
	})
	//r.NoMethod(func(ctx *gin.Context) {
	//	ctx.JSON(405, gin.H{
	//		"msg": "no such method",
	//	})
	//})

	// post参数
	r.POST("/post", func(ctx *gin.Context) {
		name := ctx.DefaultPostForm("name", "xiaobao222")
		ctx.Data(200, "text/plain; charset=utf-8",
			[]byte(fmt.Sprintf("你提交的名字叫: %s", name)))
	})

	// post json
	r.POST("/json", func(ctx *gin.Context) {
		var jdata map[string]interface{}
		ctx.BindJSON(&jdata)
		ctx.JSON(200, jdata)
	})

	// 上传文件
	r.POST("upload", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")
		fmt.Printf("file name: %s", file.Filename)

		//Save(file,dst)
		ctx.String(200, fmt.Sprintf("file: %s uploaded!", file.Filename))
	})

	r.Run(":8080")
}
