package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type metaData struct {
	Title string
	Header string
	Search string
}


func main(){
	//gin router
	r := gin.Default()

	//rendering html files
	r.LoadHTMLGlob("templates/*")

	// making a realestate landing page apis
	homeData :=[]gin.H{
		 {
			"name" : "2 BHK Villa",
			"area" : "Marina , Dubai",
			"price" : "1 million AED",
		},
		{
			"name" : "3 BHK Villa",
			"area" : "Downtown , Dubai",
			"price" : "3 million AED",
		},
	
	} 

	// routing using v1
	r.GET("api/v1/home" , func(c *gin.Context){
		c.JSON(http.StatusOK , homeData)
	})

	//reading path params in go
	r.GET("api/v1/user/:id" , func(c *gin.Context){
		userID := c.Param("id")
		res := gin.H{
			"userid" : userID,
		}
		c.JSON(http.StatusOK , res)
	})

	//contact us page
	r.GET("/contact-us" , func(c *gin.Context){
		c.HTML(http.StatusOK , "form.html" , nil)
	})

	r.POST("/submit" , func(c *gin.Context){
		res := gin.H{
			"message" : c.PostForm("message"),
			"name" : c.PostForm("username"),
			"email" : c.PostForm("email"),
		}
		c.JSON(http.StatusOK , res)
	})

	//homepage for real estate 
	r.GET("/" , func(c *gin.Context){
		c.JSON(http.StatusOK , gin.H{
			"Lets get started with : " : "GO!",
		})
	})

	//properties page
	r.GET("/properties" , func(c *gin.Context){
		data := metaData{
			Title:  "Real Estate Club",
			Header: "The Real Estate",
			Search: "Something constant",
		}
		c.HTML(http.StatusOK , "properties.html" , data)
	})


	r.POST("/getprops" , func(c *gin.Context){
		res := gin.H{
			"you search keyword" : c.PostForm("search"),
		}
		c.JSON( http.StatusOK,res)
	})

	// use of query params in go
	r.GET("/info" , func(c *gin.Context) {
		name := c.DefaultQuery("name" , "guest")
		age := c.Query("age")
		res := gin.H{
			"message" : "Hello there",
			"name" : name,
			"age" : age,
		} 
		c.JSON(http.StatusOK , res)
	}) 


	
	r.Run()
}