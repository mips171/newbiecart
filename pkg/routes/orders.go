package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/order"
	"github.com/mikestefanello/pagoda/pkg/controller"
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

	orderForm struct {
		ID          int     `form:"id" validate:"required"`
		Status      string  `form:"status" validate:"required"`
		PlacedAt    string  `form:"placed_at" validate:"required"`
		BalanceDue  float64 `form:"balance_due" validate:"required"`
		Customer    string  `form:"customer" validate:"required"`
		ProcessedBy float64 `form:"processed_by" validate:"required,gte=0"`
		Submission  controller.FormSubmission
	}
)

func NewOrdersController(client *ent.Client) *ordersController {
	return &ordersController{Client: client}
}

func (c *ordersController) Get(ctx echo.Context) error {
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
			Status:      o.Status,
			PlacedAt:    o.PlacedAt.Format("2006-01-02"),
			BalanceDue:  o.BalanceDue,
			Customer:    cust,
			ProcessedBy: processedBy,
			ID:          o.ID,
		}
	}

	return orders
}
