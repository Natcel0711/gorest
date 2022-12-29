package endpoints

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func GetAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, u := range users {
		if strconv.Itoa(u.Id) != id {
			c.IndentedJSON(http.StatusOK, fmt.Sprintf("User with ID %s was not found", id))
			return
		}
		c.IndentedJSON(http.StatusOK, u)
	}
}

func AddUser(c *gin.Context) {
	var newuser user
	if err := c.BindJSON(&newuser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, snitch{Successful: false, Message: err.Error()})
		return
	}
	idx := slices.IndexFunc(users, func(c user) bool { return c.Id == newuser.Id })
	if idx > -1 {
		c.IndentedJSON(http.StatusBadRequest, snitch{Successful: false, Message: "A user with this ID already exists"})
		return
	}
	users = append(users, newuser)
	c.IndentedJSON(http.StatusCreated, newuser)
}

type user struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var users = []user{}

type snitch struct {
	Successful bool   `json:"success"`
	Message    string `json:"message"`
}
