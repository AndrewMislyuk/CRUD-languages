package rest

import (
	"net/http"

	_ "github.com/AndrewMislyuk/CRUD-languages/docs"
	"github.com/AndrewMislyuk/CRUD-languages/internal/domain"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Language interface {
	Create(language domain.Language) (string, error)
	GetByID(id string) (domain.Language, error)
	GetAll() ([]domain.Language, error)
	Delete(id string) error
	Update(id string, inp domain.UpdateLanguageInput) error
}

type Handler struct {
	languageService Language
}

type getAllLanguagesResponse struct {
	Data []domain.Language `json:"data"`
}

type getCreationId struct {
	Id string `json:"id"`
}

func NewHandler(lang Language) *Handler {
	return &Handler{
		languageService: lang,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	language := router.Group("/language", h.loggingMiddleware)
	{
		language.GET("/", h.GetLanguageList)
		language.GET("/:id", h.GetLanguageById)
		language.POST("/", h.CreateLanguage)
		language.PUT("/:id", h.UpdateLanguage)
		language.DELETE("/:id", h.DeleteLanguage)
	}

	return router
}

// @Summary Get Language By ID
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
// @Router /{id} [get]
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
// @Router /{id} [put]
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
// @Router /{id} [delete]
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
// @Router / [post]
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
// @Tags language
// @Description get languages list
// @ID get-languages
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllLanguagesResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router / [get]
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
