package router_dir

import (
	"RMS/handler_dir"
	"RMS/middleware_dir"
	"github.com/go-chi/chi"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/rms", func(rms chi.Router) {
		rms.Post("/login", handler_dir.Login)
		rms.Get("/logout", handler_dir.Logout)
		rms.Post("/newUser", handler_dir.CreateByUser)
		rms.Route("/home", func(home chi.Router) {
			home.Use(middleware_dir.AuthMiddleware)
			home.Get("/restaurants", handler_dir.GetRestaurant)
			home.Get("/dishes", handler_dir.GetDishes)
			home.Get("/distance", handler_dir.GetDistance)

			home.Route("/admin", func(admin chi.Router) {
				admin.Use(middleware_dir.AdminMiddleware)
				admin.Get("/users", handler_dir.GetAllUsers)
				admin.Post("/newUser", handler_dir.CreateByAdmin)
				admin.Get("/getSubAdmin", handler_dir.GetSubAdmin)
				admin.Post("/newDish", handler_dir.CreateDish)
				admin.Post("/newDishBulk", handler_dir.CreateBulkDish)
				admin.Post("/csvDish", handler_dir.CreateCsvDish)
				admin.Post("/newRestaurant", handler_dir.CreateRestaurant)
				admin.Post("/newRestaurantBulk", handler_dir.CreateBulkRestaurant)

			})
			home.Route("/subAdmin", func(subAdmin chi.Router) {
				subAdmin.Use(middleware_dir.SubAdminMiddleware)
				subAdmin.Get("/users", handler_dir.GetAllUsersBySub)
				subAdmin.Get("/restaurants", handler_dir.GetRestaurantBySub)
				subAdmin.Post("/newUser", handler_dir.CreateBySubAdmin)
				subAdmin.Post("/newDish", handler_dir.CreateDish)
				subAdmin.Post("/newDishBulk", handler_dir.CreateBulkDish)
				subAdmin.Post("/csvDish", handler_dir.CreateCsvDish)
				subAdmin.Post("/newRestaurant", handler_dir.CreateRestaurant)
				subAdmin.Post("/newRestaurantBulk", handler_dir.CreateBulkRestaurant)
			})

		})

	})

	return r
}
