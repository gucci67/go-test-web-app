package controller

import (
	"go-test-web-app/api/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userService = service.UserService{}

func (u *UserController) GetUserList(c *gin.Context) {

	users, err := userService.GetUserList()
	if err != nil {
		log.Fatalln(err)
	}

	c.HTML(http.StatusFound, "user-list.html", gin.H{
		"list": users,
	})
}

func (u *UserController) DisplayEditUser(c *gin.Context) {

	userId, _ := strconv.Atoi(c.Param("id"))
	user, err := userService.GetUser(userId)
	if err != nil {
		log.Fatalln(err)
	}

	c.HTML(http.StatusFound, "user-edit.html", gin.H{
		"user": user,
	})
}

func (u *UserController) CreateUser(c *gin.Context) {

	name := c.PostForm("name")
	address := c.PostForm("address")

	if err := userService.CreateUser(name, address); err != nil {
		log.Fatalln(err)
	}

	c.Redirect(http.StatusSeeOther, "/user")
}

func (u *UserController) UpdateUser(c *gin.Context) {

	userId, _ := strconv.Atoi(c.Param("id"))
	name := c.PostForm("name")
	address := c.PostForm("address")

	if err := userService.UpdateUser(userId, name, address); err != nil {
		log.Fatalln(err)
	}

	c.Redirect(http.StatusSeeOther, "/user")
}
