package rest

import (
	"net/http"

	"github.com/AndrewMislyuk/CRUD-languages/internal/domain"
	"github.com/gin-gonic/gin"
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

func NewHandler(lang Language) *Handler {
	return &Handler{
		languageService: lang,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

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

func (h *Handler) GetLanguageById(c *gin.Context) {
	id := c.Param("id")

	language, err := h.languageService.GetByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, language)
}

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

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

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
