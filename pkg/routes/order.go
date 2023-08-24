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
	ordersController struct {
		controller.Controller
		Client *ent.Client
	}
	Order struct {
		Status      string
		PlacedAt    string // Use time.Time if you want to work with time objects directly
		BalanceDue  float64
		Customer    ent.Customer
		ProcessedBy ent.StaffMember // Placeholder if you need to display the staff member's name
		ID          int
	}
	addOrderForm struct {
		Status      string  `form:"status" validate:"required"`
		PlacedAt    string  `form:"placed_at" validate:"required"`
		BalanceDue  float64 `form:"balance_due" validate:"required,gte=0"`
		Customer    string  `form:"customer" validate:"required"`
		ProcessedBy float64 `form:"processed_by" validate:"required,gte=0"`
		Submission  controller.FormSubmission
	}
)

func (c *ordersController) Get(ctx echo.Context) error {
	orderID := ctx.Param("id")

	// make orderID an int
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		return c.Fail(err, "invalid order ID")
	}

	// Fetch the order from the database using its ID
	order, err := c.Container.ORM.Order.
		Query().
		Where(order.IDEQ(orderIDInt)).
		Only(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "unable to fetch order details")
	}

	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "order"
	page.Title = fmt.Sprint(order.ID)
	page.Data = order

	return c.RenderPage(ctx, page)
}

func (c *ordersController) GetOrders(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "orders"
	page.Metatags.Description = "All Orders"
	page.Pager = controller.NewPager(ctx, 30)
	page.Data = c.fetchOrders(ctx, &page.Pager)

	return c.RenderPage(ctx, page)
}

func (c *ordersController) fetchOrders(ctx echo.Context, pager *controller.Pager) []Order {
	ords, _ := c.Container.ORM.Order.Query(). // always use the container's ORM
							WithCustomer().
							WithProcessedBy().
							Offset(pager.GetOffset()).
							Limit(pager.ItemsPerPage).
							Order(ent.Asc(order.FieldPlacedAt)).
							All(ctx.Request().Context())

	pager.SetItems(len(ords))

	orders := make([]Order, len(ords))
	for i, o := range ords {
		var cust ent.Customer
		if len(o.Edges.Customer) > 0 {
			cust = *o.Edges.Customer[0]
		}

		var processedBy ent.StaffMember
		if len(o.Edges.ProcessedBy) > 0 {
			processedBy = *o.Edges.ProcessedBy[0]
		}

		orders[i] = Order{
			Status:      string(o.Status),
			PlacedAt:    o.PlacedAt.Format("2006-01-02"),
			BalanceDue:  o.BalanceDue,
			Customer:    cust,
			ProcessedBy: processedBy,
			ID:          o.ID,
		}
	}

	return orders
}

func (c *ordersController) ShowCreateOrderForm(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "add_order"
	page.Title = "New Order"
	page.Form = addOrderForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*addOrderForm)
	}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*addOrderForm)
	} else {
		today := time.Now().Format("2006-01-02")
		page.Form = addOrderForm{
			PlacedAt: today,
		}
	}

	return c.RenderPage(ctx, page)
}

func (c *ordersController) CreateOrder(ctx echo.Context) error {
	var form addOrderForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse order form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.Get(ctx)
	}

	// Attempt creating the order
	o, err := c.Container.ORM.Order.
		Create().
		SetStatus(order.Status(form.Status)).
		SetBalanceDue(form.BalanceDue).
		Save(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "unable to create order")
	}

	ctx.Logger().Infof("order created: %s", o.ID)
	msg.Success(ctx, "The order was successfully added.")

	return c.Redirect(ctx, "orders")
}
