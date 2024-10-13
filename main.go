package main

//import gin
import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// import "github.com/thinkerou/favicon"
func main() {
	// create a gin server
	ginServer := gin.Default()
	// ginServer.Use(favicon.New("./favicon.ico"))
	
	// load the html file
	ginServer.LoadHTMLGlob("templates/*")

	// load the static file
	ginServer.Static("/static", "./static")

	// access the address,deal with the request and response(Request Response)
	
	// Gin Restful API
	// ginServer.GET("/hello", func(context *gin.Context) {
	// 	context.JSON(200, gin.H{
	// 		"message": "hello gin",})
	// })
	// ginServer.POST("/user", func(context *gin.Context) {
	// 	context.JSON(200, gin.H{"message": "post user"})
	// })
	// ginServer.PUT("/user")
	// ginServer.DELETE("/user") 
	
	// response the html file
	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{
			"msg": "Hello from the backend!",
		})
	})

	// response the json data
	// usl?userid=1&username=brian
	ginServer.GET("/user/info", func(context *gin.Context) {
		userid := context.Query("userid")
		username := context.Query("username")
		context.JSON(200, gin.H{
			"userid": userid,
			"username": username,
		})
	})

	// /user/info/1/brian
	ginServer.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(200, gin.H{
			"userid": userid,
			"username": username,
		})
	})

	// response the json data
	ginServer.POST("/json", func(context *gin.Context) {
		data,_ := context.GetRawData()

		var m map[string]interface{}
		_ = json.Unmarshal(data, &m)
		context.JSON(200, m)
	})

	ginServer.POST("/user/add", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(200, gin.H{
			"msg":"ok",
			"username": username,
			"password": password,
		})
	})
	// router
	ginServer.GET("/test", func(context *gin.Context) {
		// redirect 
		context.JSON(301, gin.H{
			"msg": "redirect",
		})
	})

	// 404 
	ginServer.NoRoute(func(context *gin.Context) {
		context.JSON(404, "404.html")
	})

	userGroup := ginServer.Group("/user")
	{
		userGroup.GET("/add")
		userGroup.GET("/login")
		userGroup.GET("/logout")
	}
	orderGroup := ginServer.Group("/order")
	{
		orderGroup.GET("/add")
		orderGroup.GET("/list")
		orderGroup.GET("/delete")
	}

	// server port
	ginServer.Run(":8080")
	
}
