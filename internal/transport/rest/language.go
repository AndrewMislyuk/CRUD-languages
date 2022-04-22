package rest

import (
	"github.com/AndrewMislyuk/CRUD-languages/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get Language By ID
// @Security ApiKeyAuth
// @Tags language
// @Description get language by id
// @ID get-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} domain.Language
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /language/{id} [get]
func (h *Handler) GetLanguageById(c *gin.Context) {
	id := c.Param("id")

	language, err := h.languageService.GetByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, language)
}

// @Summary Update Language
// @Security ApiKeyAuth
// @Tags language
// @Description update language by id
// @ID update-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param input body domain.UpdateLanguageInput true "language info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /language/{id} [put]
func (h *Handler) UpdateLanguage(c *gin.Context) {
	id := c.Param("id")

	var input domain.UpdateLanguageInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.languageService.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete Language
// @Security ApiKeyAuth
// @Tags language
// @Description delete language by id
// @ID delete-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /language/{id} [delete]
func (h *Handler) DeleteLanguage(c *gin.Context) {
	id := c.Param("id")

	err := h.languageService.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Create Language
// @Security ApiKeyAuth
// @Tags language
// @Description create language
// @ID create-language
// @Accept  json
// @Produce  json
// @Param input body domain.Language true "language info"
// @Success 200 {object} getCreationId
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /language/ [post]
func (h *Handler) CreateLanguage(c *gin.Context) {
	var language domain.Language
	if err := c.BindJSON(&language); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.languageService.Create(language)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getCreationId{
		Id: id,
	})
}

// @Summary Get Languages List
// @Security ApiKeyAuth
// @Tags language
// @Description get languages list
// @ID get-languages
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllLanguagesResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /language/ [get]
func (h *Handler) GetLanguageList(c *gin.Context) {
	language, err := h.languageService.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllLanguagesResponse{
		Data: language,
	})
}
