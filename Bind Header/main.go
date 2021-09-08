package main
import (
	"githu.com/gin-gonic/gin"
)
type myForm struct {
    Colors []string `form:"colors[]"`
}
func formHandler(c *gin.Context) {
    var fakeForm myForm
    c.ShouldBind(&fakeForm)
    c.JSON(200, gin.H{"color": fakeForm.Colors})
}
func main()  {
	router:=gin.Default()
	router.GET("/colors",formHandler)
	router.Run()
}