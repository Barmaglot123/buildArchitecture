package web

import (
	"github.com/gin-gonic/gin"
	"build_architecture/datasource"
	"build_architecture/model"
	"net/http"
)

func RunServer() {
	r := gin.Default()
	r.GET("/users", usersList)
	r.GET("/users/:id", userWithID)
	r.DELETE("/users/:id", deleteUserWithID)
	r.POST("users/:name/:surname/:phnumber", createUser)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	r.Run(":3000")
}

func usersList(c *gin.Context) {
	users := []model.User{}
	datasource.DataSource.Find(&users)
	c.JSON(http.StatusOK, users)
}
func userWithID (c *gin.Context) {
	user := []model.User{}
	datasource.DataSource.First(&user, c.Param("id"))
	c.JSON(http.StatusOK, user)

}

func deleteUserWithID (c *gin.Context){
	user  := []model.User{}
	datasource.DataSource.First(&user, c.Param("id"))
	datasource.DataSource.Delete(&user)
	c.JSON(http.StatusOK, user)
}

func createUser (c *gin.Context){
	user := model.User{     Name:c.PostForm("name"),
		Surname:c.PostForm("surname"),
		Phone:c.PostForm("phnumber")}
	datasource.DataSource.NewRecord(user)
	datasource.DataSource.Create(&user)
	c.JSON(http.StatusOK, "Done!")


}