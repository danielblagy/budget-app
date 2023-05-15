package budget_app

func (h handler) SetupRoutes() {
	h.app.Get("/", h.Greet)

	usersGroup := h.app.Group("users")
	usersGroup.Get("/", h.GetUsers)
	usersGroup.Get("/:username", h.GetUser)
	usersGroup.Post("/", h.CreateUser)
	// TODO update password using JWT token usersGroup.Put("/", h.UpdateUser)
	// TODO delete using JWT token usersGroup.Delete("/", h.DeleteUser)
}
