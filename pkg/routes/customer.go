package routes

import (
	"fmt"
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

	editCustomerForm struct {
		ID         int    `form:"id" validate:"required"`
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
	page.Name = "customers/list"
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

func (c *customersController) ViewCustomer(ctx echo.Context) error {
	customerID := ctx.Param("id")

	// make customerID an int
	customerIDInt, err := strconv.Atoi(customerID)
	if err != nil {
		return c.Fail(err, "invalid customer ID")
	}

	// Fetch the customer from the database using its ID
	customer, err := c.Container.ORM.Customer.
		Query().
		Where(customer.IDEQ(customerIDInt)).
		Only(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "unable to fetch customer details")
	}

	page := controller.NewPage(ctx)
	page.Layout = "edit_entity"
	page.Name = "customers.view"
	page.Title = fmt.Sprint(customer.ID)
	page.Data = customer

	return c.RenderPage(ctx, page)
}

// Display the form to add a customer
func (c *customersController) ShowCreateCustomerForm(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "edit_entity"
	page.Name = "customers/add"
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
	page.Layout = "edit_entity"
	page.Name = "customers/add"
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

func (c *customersController) EditCustomer(ctx echo.Context) error {
	customerID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return c.Fail(err, "Invalid customer ID")
	}

	customer, err := c.Container.ORM.Customer.Get(ctx.Request().Context(), customerID)
	if err != nil {
		return c.Fail(err, "Unable to fetch customer")
	}

	page := controller.NewPage(ctx)
	page.Layout = "edit_entity"
	page.Name = "customers/edit"
	page.Title = "Edit customer"
	page.Form = editCustomerForm{
		ID:         customer.ID,
		Name:       customer.Name,
		Email:      customer.Email,
		Password:   customer.Password,
		Address:    customer.Address,
		Status:     string(customer.Status),
		Submission: controller.FormSubmission{},
	}

	return c.RenderPage(ctx, page)
}

// EditCustomer processes the customer editing form
func (c *customersController) EditCustomerPost(ctx echo.Context) error {
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
