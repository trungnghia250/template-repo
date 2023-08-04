package delivery

import (
	"FirstAssignment/service/domain/user/usecase"
	"FirstAssignment/service/model/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	user usecase.UserUseCase
}

func NewUserHandler(user usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		user: user,
	}
}

func (u *UserHandler) GetAllUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := u.user.GetAllUser(ctx)
		if err != nil {
			ctx.PureJSON(http.StatusBadRequest, err)
		}

		ctx.JSON(http.StatusOK, users)
	}
}

func (u *UserHandler) GetUserByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &dto.GetUserByIDRequest{}
		if err := ctx.ShouldBindQuery(req); err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
		}

		users, err := u.user.GetUserByID(ctx, *req)
		if err != nil {
			ctx.PureJSON(http.StatusBadRequest, err)
		}

		ctx.JSON(http.StatusOK, users)
	}
}

func (u *UserHandler) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &dto.CreateUserRequest{}
		if err := ctx.ShouldBind(req); err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
		}

		err := u.user.CreateUser(ctx, *req)
		if err != nil {
			ctx.PureJSON(http.StatusBadRequest, err)
		}

		ctx.JSON(http.StatusOK, nil)
	}
}

func (u *UserHandler) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &dto.UpdateUserRequest{}
		if err := ctx.ShouldBind(req); err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
		}

		err := u.user.UpdateUser(ctx, *req)
		if err != nil {
			ctx.PureJSON(http.StatusBadRequest, err)
		}

		ctx.JSON(http.StatusOK, nil)
	}
}

func (u *UserHandler) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &dto.DeleteUserRequest{}
		if err := ctx.ShouldBindQuery(req); err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
		}

		err := u.user.DeleteUser(ctx, *req)
		if err != nil {
			ctx.PureJSON(http.StatusBadRequest, err)
		}

		ctx.JSON(http.StatusOK, nil)
	}
}
