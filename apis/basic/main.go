package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.GET("/check" , func(c *gin.Context){
		res := gin.H{
			"chadmax" :"hello , this is chadmax",
			"gigachad" :"hi , what is that disturbed me",
		}
		c.JSON(http.StatusOK , res)
	})

	r.Run()
}