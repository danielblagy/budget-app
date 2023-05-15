package budget_app

func (h handler) SetupRoutes() {
	h.app.Get("/", h.Greet)

	usersGroup := h.app.Group("users")
	usersGroup.Get("/", h.GetUsers)
	usersGroup.Get("/:user_id", h.GetUser)
	usersGroup.Post("/", h.CreateUser)
}
