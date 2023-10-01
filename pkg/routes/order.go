package routes

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/order"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/msg"
)

type (
	// Interfaces to make actions explicit
	OrderGetter interface {
		GetAll(ctx echo.Context) error
		GetByID(ctx echo.Context) error
	}

	OrderEditor interface {
		EditByID(ctx echo.Context) error
	}

	OrderAdder interface {
		Add(ctx echo.Context) error
	}

	// OrderController with integrated methods
	OrderController struct {
		controller.Controller
		Client *ent.Client
	}

	Order struct {
		Status      string
		PlacedAt    string // Use time.Time if you want to work with time objects directly
		BalanceDue  float64
		// Customer    ent.Customer
		// ProcessedBy ent.StaffMember // Placeholder if you need to display the staff member's name
		ID          int
	}

	OrderForm struct {
		ID          *int    `form:"id"` // Note use of pointer to allow for nil values, and no "required" validation
		Status      string  `form:"status" validate:"required"`
		PlacedAt    string  `form:"placed_at" validate:"required"`
		BalanceDue  float64 `form:"balance_due" validate:"required,gte=0"`
		// Customer    string  `form:"customer" validate:"required"`
		// ProcessedBy float64 `form:"processed_by" validate:"required,gte=0"`
		Submission  controller.FormSubmission
	}
)

var _ OrderGetter = &OrderController{}
var _ OrderEditor = &OrderController{}
var _ OrderAdder = &OrderController{}

func (c *OrderController) GetAll(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "orders/view_all" // Located at templates/pages/orders/
	page.Title = "Orders"
	page.Metatags.Description = "All orders"
	page.Metatags.Keywords = []string{"Shopping", "Orders", "Buy"}
	page.Pager = controller.NewPager(ctx, 10)
	page.Data = c.fetchOrders(ctx, &page.Pager)

	return c.RenderPage(ctx, page)
}

func (c *OrderController) GetByID(ctx echo.Context) error {
	orderID, err := c.getOrderIDFromContext(ctx)
	if err != nil {
		return c.handleError(err, "Invalid order ID")
	}

	order, err := c.Container.ORM.Order.
		Query().
		Where(order.IDEQ(orderID)).
		Only(ctx.Request().Context())

	if err != nil {
		return c.handleError(err, "Unable to fetch order details")
	}

	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "orders/view"
	page.Title = "Viewing Order"
	page.Data = order

	return c.RenderPage(ctx, page)
}

func (c *OrderController) EditByID(ctx echo.Context) error {
	switch ctx.Request().Method {
	case echo.GET:
		return c.handleEditByIdGet(ctx)
	case echo.POST:
		return c.handleEditByIdPost(ctx)
	default:
		return echo.ErrMethodNotAllowed
	}
}

func (c *OrderController) handleEditByIdGet(ctx echo.Context) error {
	orderID, err := c.getOrderIDFromContext(ctx)
	if err != nil {
		return c.handleError(err, "Invalid order ID")
	}

	order, err := c.Container.ORM.Order.Get(ctx.Request().Context(), orderID)
	if err != nil {
		return c.Fail(err, "Unable to fetch order")
	}

	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "orders/edit"
	page.Title = "Editing Order"
	page.Form = OrderForm{
		ID: &order.ID, // Make sure to pass the address of order.ID
		PlacedAt: order.PlacedAt.Format("2006-01-02T15:04"),
		BalanceDue: order.BalanceDue,
		Status: string(order.Status),
		// ... populate other fields ...
	}
	// ... rest of your code ...
	return c.RenderPage(ctx, page)
}

func (c *OrderController) handleEditByIdPost(ctx echo.Context) error {
	var form OrderForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse order form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.handleEditByIdGet(ctx)
	}

	ctx.Logger().Infof("Attempting to update order with ID: %d", form.ID)

	fmt.Println("got status:", form.Status)

	// Fetch the existing order and update its fields
	order, err := c.Container.ORM.Order.UpdateOneID(*form.ID). // Assuming form.ID exists
		SetStatus(order.Status(form.Status)).
		// SetBalanceDue(form.BalanceDue).
		// ... set other fields ...
		Save(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "Unable to update order")
	}

	ctx.Logger().Infof("Order updated: %s", order.ID)
	msg.Success(ctx, "The order was successfully updated.")

	return c.Redirect(ctx, "orders")
}

func (c *OrderController) Add(ctx echo.Context) error {
	switch ctx.Request().Method {
	case echo.GET:
		return c.handleAddGet(ctx)
	case echo.POST:
		return c.handleAddPost(ctx)
	default:
		return echo.ErrMethodNotAllowed
	}
}

func (c *OrderController) handleAddGet(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "orders/add"
	page.Title = "Add New Order"
	page.Form = OrderForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*OrderForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *OrderController) handleAddPost(ctx echo.Context) error {
	var form OrderForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse order form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.handleAddGet(ctx)
	}

	// Assuming you have an Order ORM similar to Product ORM, you can create the order as follows:
	// (I am assuming some fields like 'Status', 'PlacedAt', and 'BalanceDue' based on the provided schema)

	// turn form.PlacedAt into appropriate format
	placedAt, err := time.Parse("2006-01-02T15:04", form.PlacedAt)
	if err != nil {
		return c.Fail(err, "unable to parse form.PlacedAt")
	}

	o, err := c.Container.ORM.Order.
		Create().
		SetStatus(order.Status(form.Status)).
		SetPlacedAt(placedAt).
		SetBalanceDue(form.BalanceDue).
		// ... Any other fields based on your form and ORM
		Save(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "unable to create order")
	}

	ctx.Logger().Infof("created: Order ID %d with status %s", o.ID, o.Status)
	msg.Success(ctx, fmt.Sprintf("Order successfully added with ID: %d", o.ID))

	return c.Redirect(ctx, "orders")
}

// Helper methods
func (c *OrderController) getOrderIDFromContext(ctx echo.Context) (int, error) {
	orderID := ctx.Param("id")
	return strconv.Atoi(orderID)
}

func (c *OrderController) handleError(err error, msg string) error {
	// Handle errors, e.g., logging, setting HTTP status
	return c.Fail(err, msg)
}

func (c *OrderController) fetchOrders(ctx echo.Context, pager *controller.Pager) []Order {
	prods, _ := c.Container.ORM.Order.Query().
		Offset(pager.GetOffset()).
		Limit(pager.ItemsPerPage).
		Order(ent.Asc(order.FieldID)).
		All(ctx.Request().Context())

	pager.SetItems(len(prods))

	orders := make([]Order, len(prods))
	for i, p := range prods {
		orders[i] = Order{
			ID:         p.ID,
			PlacedAt: 	p.PlacedAt.Format("2006-01-02T15:04"),
			BalanceDue: p.BalanceDue,
		}
	}
	return orders
}
