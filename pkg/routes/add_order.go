package routes

import (
	"time"

	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/msg"

	"github.com/labstack/echo/v4"
)

type (
	AddOrderController struct {
		controller.Controller
		Client *ent.Client
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

func (c *AddOrderController) Get(ctx echo.Context) error {
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

func (c *AddOrderController) Post(ctx echo.Context) error {
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
		SetStatus(form.Status).
		SetBalanceDue(form.BalanceDue).
		Save(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "unable to create order")
	}

	ctx.Logger().Infof("order created: %s", o.ID)
	msg.Success(ctx, "The order was successfully added.")

	return c.Redirect(ctx, "orders")
}
