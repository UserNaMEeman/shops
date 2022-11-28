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
	// router.Use(h.userIdentity)
	// router.Use(middleware.Logger)
	router.Post("/api/user/register", h.signUp) //registration
	router.Post("/api/user/login", h.signIn)    //authentification

	router.Route("/api/user", func(router chi.Router) {
		// router.Use(middleware.Logger)
		router.Use(h.userIdentity)
		router.Use(h.GzipHandle)
		router.Post("/orders", h.uploadOrder)        //загрузка пользователем номера заказа для расчёта
		router.Get("/orders", h.getOrders)           //получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях
		router.Get("/balance", h.getBalance)         //получение текущего баланса счёта баллов лояльности пользователя
		router.Post("/balance/withdraw", h.withdraw) //запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа
		router.Get("/withdrawals", h.withdrawals)    //получение информации о выводе средств с накопительного счёта пользователем
		//
	})
	// router.Get("/api/user/balance/withdraw", signUp)  //получение информации о выводе средств с накопительного счёта пользователем
	// router.Get("/test", h.IsLoggedIn)
	return router
}
