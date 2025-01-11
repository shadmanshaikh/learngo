package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

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
	r.GET("api/v1/home" , func(c *gin.Context){
		c.JSON(http.StatusOK , homeData)
	})

	//reading params in go
	r.GET("api/v1/user/:id" , func(c *gin.Context){
		userID := c.Param("id")
		res := gin.H{
			"userid" : userID,
		}
		c.JSON(http.StatusOK , res)
	})

	r.Run()
}