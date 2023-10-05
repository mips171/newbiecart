package routes

import (
	"net/http"

	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/middleware"
	"github.com/mikestefanello/pagoda/pkg/services"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

// BuildRouter builds the router
func BuildRouter(c *services.Container) {
	// Static files with proper cache control
	// funcmap.File() should be used in templates to append a cache key to the URL in order to break cache
	// after each server restart
	c.Web.Group("", middleware.CacheControl(c.Config.Cache.Expiration.StaticFile)).
		Static(config.StaticPrefix, config.StaticDir)

	// Non static file route group
	g := c.Web.Group("")

	// Force HTTPS, if enabled
	if c.Config.HTTP.TLS.Enabled {
		g.Use(echomw.HTTPSRedirect())
	}

	g.Use(
		echomw.RemoveTrailingSlashWithConfig(echomw.TrailingSlashConfig{
			RedirectCode: http.StatusMovedPermanently,
		}),
		echomw.Recover(),
		echomw.Secure(),
		echomw.RequestID(),
		echomw.Gzip(),
		echomw.Logger(),
		middleware.LogRequestID(),
		echomw.TimeoutWithConfig(echomw.TimeoutConfig{
			Timeout: c.Config.App.Timeout,
		}),
		session.Middleware(sessions.NewCookieStore([]byte(c.Config.App.EncryptionKey))),
		middleware.LoadAuthenticatedUser(c.Auth),
		middleware.ServeCachedPage(c.Cache),
		echomw.CSRFWithConfig(echomw.CSRFConfig{
			TokenLookup: "form:csrf",
		}),
	)

	// Base controller
	ctr := controller.NewController(c)

	// Error handler
	err := errorHandler{Controller: ctr}
	c.Web.HTTPErrorHandler = err.Get

	// Example routes
	navRoutes(c, g, ctr)
	userRoutes(c, g, ctr)
	staffMemberRoutes(c, g, ctr)
}

func navRoutes(c *services.Container, g *echo.Group, ctr controller.Controller) {
	home := home{Controller: ctr}
	g.GET("/", home.Get).Name = "home"

	search := search{Controller: ctr}
	g.GET("/search", search.Get).Name = "search"

	about := about{Controller: ctr}
	g.GET("/about", about.Get).Name = "about"

	contact := contact{Controller: ctr}
	g.GET("/contact", contact.Get).Name = "contact"
	g.POST("/contact", contact.Post).Name = "contact.post"

	customers := customersController{Controller: ctr}
	// List of Customers
	g.GET("/customers", customers.GetCustomers).Name = "customers.list"
	// View a Specific Customer
	g.GET("/customers/:id", customers.ViewCustomer).Name = "customers.view"
	// Edit a Specific Customer
	g.GET("/customers/:id/edit", customers.EditCustomer).Name = "customers.edit"
	g.POST("/customers/:id/edit", customers.EditCustomerPost).Name = "customers.edit.submit"
	// Create a New Customer
	g.GET("/customers/add", customers.ShowCreateCustomerForm).Name = "customers.add"
	g.POST("/customers/add", customers.CreateCustomer).Name = "customers.add.submit"

	categories := ProductCategoryController{Controller: ctr}
	g.GET("/categories", categories.GetAll).Name = "product_categories.view_all"
	g.GET("/categories/:id", categories.GetByID).Name = "product_categories.view"
	g.GET("/categories/add", categories.Add).Name = "product_categories.add"
	g.POST("/categories/add", categories.Add).Name = "product_categories.add.post"
	g.GET("/categories/:id/edit", categories.EditByID).Name = "product_categories.edit"       // display the edit form
	g.POST("/categories/:id/edit", categories.EditByID).Name = "product_categories.edit.post" // submit the edit form

	products := ProductController{Controller: ctr}
	g.GET("/products", products.GetAll).Name = "products.view_all"
	g.GET("/products/:id", products.GetByID).Name = "products.view"
	g.GET("/products/add", products.Add).Name = "products.add"
	g.POST("/products/add", products.Add).Name = "products.add.post"
	g.GET("/products/:id/edit", products.EditByID).Name = "products.edit"       // display the edit form
	g.POST("/products/:id/edit", products.EditByID).Name = "products.edit.post" // submit the edit form

	orders := OrderController{Controller: ctr}
	g.GET("/orders", orders.GetAll).Name = "orders.view_all"
	g.GET("/orders/:id", orders.GetByID).Name = "orders.view"
	g.GET("/orders/add", orders.Add).Name = "orders.add"
	g.POST("/orders/add", orders.Add).Name = "orders.add.post"
	g.GET("/orders/:id/edit", orders.EditByID).Name = "orders.edit"       // display the edit form
	g.POST("/orders/:id/edit", orders.EditByID).Name = "orders.edit.post" // submit the edit form
}

func userRoutes(c *services.Container, g *echo.Group, ctr controller.Controller) {
	logout := logout{Controller: ctr}
	g.GET("/logout", logout.Get, middleware.RequireAuthentication()).Name = "logout"

	verifyEmail := verifyEmail{Controller: ctr}
	g.GET("/email/verify/:token", verifyEmail.Get).Name = "verify_email"

	noAuth := g.Group("/user", middleware.RequireNoAuthentication())
	login := login{Controller: ctr}
	noAuth.GET("/login", login.Get).Name = "login"
	noAuth.POST("/login", login.Post).Name = "login.post"

	register := register{Controller: ctr}
	noAuth.GET("/register", register.Get).Name = "register"
	noAuth.POST("/register", register.Post).Name = "register.post"

	forgot := forgotPassword{Controller: ctr}
	noAuth.GET("/password", forgot.Get).Name = "forgot_password"
	noAuth.POST("/password", forgot.Post).Name = "forgot_password.post"

	resetGroup := noAuth.Group("/password/reset",
		middleware.LoadUser(c.ORM),
		middleware.LoadValidPasswordToken(c.Auth),
	)
	reset := resetPassword{Controller: ctr}
	resetGroup.GET("/token/:user/:password_token/:token", reset.Get).Name = "reset_password"
	resetGroup.POST("/token/:user/:password_token/:token", reset.Post).Name = "reset_password.post"
}

func staffMemberRoutes(c *services.Container, g *echo.Group, ctr controller.Controller) {
	logout := logout{Controller: ctr}
	g.GET("/logout", logout.Get, middleware.RequireAuthentication()).Name = "logout"

	verifyEmail := verifyEmail{Controller: ctr}
	g.GET("/email/verify/:token", verifyEmail.Get).Name = "verify_email"

}
