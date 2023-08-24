package routes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/customer"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
)

type (
	customersController struct {
		controller.Controller
		Client *ent.Client
	}

	Customer struct {
		ID       int
		Name     string
		Email    string
		Password string
		Address  string
		Status   string
	}

	addCustomerForm struct {
		Name       string `form:"name" validate:"required"`
		Email      string `form:"email" validate:"required"`
		Password   string `form:"password" validate:"required"`
		Address    string `form:"address" validate:"required"`
		Status     string `form:"status" validate:"required"`
		Submission controller.FormSubmission
	}
)

func NewCustomersController(client *ent.Client) *customersController {
	return &customersController{Client: client}
}

// GetCustomers displays a list of all customers
func (c *customersController) GetCustomers(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "customers"
	page.Metatags.Description = "All Customers"
	page.Pager = controller.NewPager(ctx, 30)
	page.Data = c.fetchCustomers(ctx, &page.Pager)

	return c.RenderPage(ctx, page)
}

// fetchCustomers fetches all the customers
func (c *customersController) fetchCustomers(ctx echo.Context, pager *controller.Pager) []Customer {
	custs, _ := c.Container.ORM.Customer.Query().
		Offset(pager.GetOffset()).
		Limit(pager.ItemsPerPage).
		Order(ent.Asc(customer.FieldName)).
		All(ctx.Request().Context())

	pager.SetItems(len(custs))

	customers := make([]Customer, len(custs))
	for i, cust := range custs {
		customers[i] = Customer{
			ID:       cust.ID,
			Name:     cust.Name,
			Email:    cust.Email,
			Password: cust.Password,
			Address:  cust.Address,
			Status:   string(cust.Status),
		}
	}

	return customers
}

// Display the form to add a customer
func (c *customersController) ShowCreateCustomerForm(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "add_customer"
	page.Title = "New Customer"

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*addCustomerForm)
	} else {
		page.Form = addCustomerForm{}
	}

	return c.RenderPage(ctx, page)
}

// CreateCustomer processes the customer creation form
func (c *customersController) CreateCustomer(ctx echo.Context) error {

	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "add_customer"
	page.Title = "New Customer"
	page.Form = addCustomerForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*addCustomerForm)
	} else {
		page.Form = addCustomerForm{}
	}

	// For now, just a basic error handling and insertion. This should be enhanced.
	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	address := ctx.FormValue("address")
	status := ctx.FormValue("status")

	_, err := c.Client.Customer.Create().
		SetName(name).
		SetEmail(email).
		SetPassword(password).
		SetAddress(address).
		SetStatus(customer.Status(status)).
		Save(ctx.Request().Context())

	if err != nil {
		// Handle the error more gracefully, possibly return an error page or JSON
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to create customer")
	}

	return ctx.Redirect(http.StatusSeeOther, "/customers")
}

// EditCustomer processes the customer editing form
func (c *customersController) EditCustomer(ctx echo.Context) error {
	id := ctx.Param("id") // Assuming id is an integer parameter in the route

	// Fetch customer based on ID
	// parse id to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		// Handle the error more gracefully, possibly return an error page or JSON
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to fetch customer")
	}

	cust, err := c.Client.Customer.Get(ctx.Request().Context(), idInt)
	if err != nil {
		// Handle the error more gracefully, possibly return an error page or JSON
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to fetch customer")
	}

	// For now, just a basic error handling and update. This should be enhanced.
	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	address := ctx.FormValue("address")
	status := ctx.FormValue("status")

	_, err = cust.Update().
		SetName(name).
		SetEmail(email).
		SetPassword(password).
		SetAddress(address).
		SetStatus(customer.Status(status)).
		Save(ctx.Request().Context())

	if err != nil {
		// Handle the error more gracefully, possibly return an error page or JSON
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to edit customer")
	}

	return ctx.Redirect(http.StatusSeeOther, "/customers")
}
