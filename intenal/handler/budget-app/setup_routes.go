package budget_app

func (h handler) SetupRoutes() {
	h.app.Get("/", h.Greet)

	usersGroup := h.app.Group("users")
	usersGroup.Get("/me", h.Me)
	usersGroup.Get("/", h.GetUsers)
	usersGroup.Get("/:username", h.GetUser)
	usersGroup.Post("/", h.CreateUser)
	// TODO update password using JWT token usersGroup.Put("/", h.UpdateUser)
	// TODO delete using JWT token usersGroup.Delete("/", h.DeleteUser)

	accessGroup := h.app.Group("access")
	accessGroup.Post("/login", h.LogIn)

	categoriesGroup := h.app.Group("categories")
	categoriesGroup.Get("/", h.GetCategories)
	categoriesGroup.Get("/:id", h.GetCategory)
	categoriesGroup.Post("/", h.CreateCategory)
	categoriesGroup.Put("/:id", h.UpdateCategory)
}
