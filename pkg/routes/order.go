package routes

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/order"
	"github.com/mikestefanello/pagoda/pkg/controller"
)

type (
	orderDetail struct {
		controller.Controller
	}

	Order struct {
		Status      string
		PlacedAt    string // Use time.Time if you want to work with time objects directly
		BalanceDue  float64
		Customer    ent.Customer
		ProcessedBy ent.StaffMember // Placeholder if you need to display the staff member's name
		ID          int
	}
)

func (c *orderDetail) Get(ctx echo.Context) error {
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
