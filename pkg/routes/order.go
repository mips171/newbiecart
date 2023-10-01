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
	OrderController struct {
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

)

func (c *OrderController) GetByID(ctx echo.Context) error {
	orderID, err := c.getOrderIDFromContext(ctx)
	if err != nil {
		return c.handleError(err, "Invalid order ID")
	}

	// Fetch the order from the database using its ID
	order, err := c.Container.ORM.Order.
		Query().
		Where(order.IDEQ(orderID)).
		Only(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "unable to fetch order details")
	}

	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "orders/view"
	page.Title = fmt.Sprint(order.ID)
	page.Data = order

	return c.RenderPage(ctx, page)
}

func (c *OrderController) GetAll(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "orders/view_all"
	page.Metatags.Description = "All Orders"
	page.Pager = controller.NewPager(ctx, 30)
	page.Data = c.fetchOrders(ctx, &page.Pager)

	return c.RenderPage(ctx, page)
}

func (c *OrderController) fetchOrders(ctx echo.Context, pager *controller.Pager) []Order {
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

func (c *OrderController) AddOrder(ctx echo.Context) error {
	switch ctx.Request().Method {
	case echo.GET:
		return c.handleAddOrderGet(ctx)
	case echo.POST:
		return c.handleAddOrderPost(ctx)
	default:
		return echo.ErrMethodNotAllowed
	}
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

// Helper methods
func (c *OrderController) getOrderIDFromContext(ctx echo.Context) (int, error) {
	orderId := ctx.Param("id")
	return strconv.Atoi(orderId)
}

func (c *OrderController) handleError(err error, msg string) error {
	// Handle errors, e.g., logging, setting HTTP status
	return c.Fail(err, msg)
}