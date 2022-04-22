package rest

import (
	"github.com/AndrewMislyuk/CRUD-languages/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary SignUp
// @Tags Auth
// @Description user sign-up
// @ID sign-up
// @Accept  json
// @Produce  json
// @Param input body domain.SignUpInput true "User info"
// @Success 200 {object} getCreationId
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) SignUp(c *gin.Context) {
	var input domain.SignUpInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.userService.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getCreationId{
		Id: id,
	})
}

// @Summary SignIn
// @Tags Auth
// @Description user sign-in
// @ID sign-in
// @Accept  json
// @Produce  json
// @Param input body domain.SignInInput true "User info"
// @Success 200 {object} getCreationToken
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) SignIn(c *gin.Context) {
	var input domain.SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.userService.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getCreationToken{
		AccessToken: token,
	})
}
