package rest

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/AndrewMislyuk/CRUD-languages/internal/domain"
)

type Language interface {
	Create(ctx context.Context, language domain.Language) error
	GetByID(ctx context.Context, id string) (domain.Language, error)
	GetAll(ctx context.Context) ([]domain.Language, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, inp domain.UpdateLanguageInput) error
}

type Handler struct {
	languageService Language
}

func NewHandler(lang Language) *Handler {
	return &Handler{
		languageService: lang,
	}
}

func (h *Handler) InitRouter() {
	http.HandleFunc("/language/", loggingMiddleware(h.handleLanguage))
}

func (h *Handler) handleLanguage(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/language/")

	switch r.Method {
	case http.MethodGet:
		if id != "" {
			h.GetLanguageById(w, r, id)
		} else {
			h.GetLanguageList(w, r)
		}
	case http.MethodPost:
		h.CreateLanguage(w, r)
	case http.MethodDelete:
		if id != "" {
			h.DeleteLanguage(w, r, id)
		}
	case http.MethodPut:
		if id != "" {
			h.UpdateLanguage(w, r, id)
		}
	default:
		w.WriteHeader(http.StatusBadGateway)
	}
}

func (h *Handler) GetLanguageById(w http.ResponseWriter, r *http.Request, id string) {
	language, err := h.languageService.GetByID(context.TODO(), id)
	if err != nil {
		log.Println("GetLanguageById() error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(language)
	if err != nil {
		log.Println("GetLanguageById() error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) UpdateLanguage(w http.ResponseWriter, r *http.Request, id string) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp domain.UpdateLanguageInput
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.languageService.Update(context.TODO(), id, inp)
	if err != nil {
		log.Println("UpdateLanguage() error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Language Has Been Updated"))
}

func (h *Handler) DeleteLanguage(w http.ResponseWriter, r *http.Request, id string) {
	err := h.languageService.Delete(context.TODO(), id)
	if err != nil {
		log.Println("DeleteLanguage() error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Language Has Been Deleted"))
}

func (h *Handler) CreateLanguage(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var language domain.Language
	if err = json.Unmarshal(reqBytes, &language); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.languageService.Create(context.TODO(), language)
	if err != nil {
		log.Println("CreateLanguage() error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Language Has Been Created"))
}

func (h *Handler) GetLanguageList(w http.ResponseWriter, r *http.Request) {
	language, err := h.languageService.GetAll(context.TODO())
	if err != nil {
		log.Println("GetLanguageList() error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(language)
	if err != nil {
		log.Println("GetLanguageList() error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}
