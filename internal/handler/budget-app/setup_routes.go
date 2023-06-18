package budget_app

func (h handler) SetupRoutes() {
	h.app.Get("/", h.Greet)

	v1Group := h.app.Group("v1")

	usersGroup := v1Group.Group("users")
	usersGroup.Get("/me", h.Me)
	usersGroup.Get("/", h.GetUsers)
	usersGroup.Get("/:username", h.GetUser)
	usersGroup.Post("/", h.CreateUser)
	// TODO update password using JWT token usersGroup.Put("/", h.UpdateUser)
	// TODO delete using JWT token usersGroup.Delete("/", h.DeleteUser)

	accessGroup := v1Group.Group("access")
	accessGroup.Post("/login", h.LogIn)

	categoriesGroup := v1Group.Group("categories")
	categoriesGroup.Get("/:type", h.GetCategories)
	categoriesGroup.Get("/by_id/:id", h.GetCategory)
	categoriesGroup.Post("/", h.CreateCategory)
	categoriesGroup.Put("/", h.UpdateCategory)
	categoriesGroup.Delete("/:id", h.DeleteCategory)

	entriesGroup := v1Group.Group("entries")
	entriesGroup.Get("/:type", h.GetEntries)
	entriesGroup.Get("by_id/:id", h.GetEntryByID)
}
