package rest

import (
	_ "github.com/AndrewMislyuk/CRUD-languages/docs"
	"github.com/AndrewMislyuk/CRUD-languages/internal/domain"
	"github.com/AndrewMislyuk/CRUD-languages/internal/service"
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

type User interface {
	CreateUser(user domain.SignUpInput) (string, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (string, error)
}

type Handler struct {
	languageService Language
	userService     User
}

type getAllLanguagesResponse struct {
	Data []domain.Language `json:"data"`
}

type getCreationId struct {
	Id string `json:"id"`
}

type getCreationToken struct {
	AccessToken string `json:"token"`
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		languageService: service.Language,
		userService:     service.User,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	auth := router.Group("/auth")
	{
		auth.Use(h.loggingMiddleware)

		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	language := router.Group("/language")
	{
		language.Use(h.userIdentify)
		language.Use(h.loggingMiddleware)

		language.GET("/", h.GetLanguageList)
		language.GET("/:id", h.GetLanguageById)
		language.POST("/", h.CreateLanguage)
		language.PUT("/:id", h.UpdateLanguage)
		language.DELETE("/:id", h.DeleteLanguage)
	}

	return router
}
