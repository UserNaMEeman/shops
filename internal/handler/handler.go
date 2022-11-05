package handler

import (
	"github.com/UserNaMEeman/shops/internal/service"
	"github.com/go-chi/chi"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/api/user/register", h.signUp) //registration
	// router.Post("/api/user/login", signIn)            //authentification
	// router.Post("/api/user/orders", signUp)           //загрузка пользователем номера заказа для расчёта
	// router.Get("/api/user/orders", signUp)            //получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях
	// router.Get("/api/user/balance", signUp)           //получение текущего баланса счёта баллов лояльности пользователя
	// router.Post("/api/user/balance/withdraw", signUp) //запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа
	// router.Get("/api/user/balance/withdraw", signUp)  //получение информации о выводе средств с накопительного счёта пользователем
	return router
}
