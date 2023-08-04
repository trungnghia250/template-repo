package delivery

import (
	"github.com/gin-gonic/gin"
)

func (u *UserHandler) UserAPIRoute(router *gin.RouterGroup) {
	router.GET("/user/list", u.GetAllUser())
	router.GET("/user", u.GetUserByID())
	router.POST("/user", u.CreateUser())
	router.PUT("/user", u.UpdateUser())
	router.DELETE("/user", u.DeleteUser())
}
