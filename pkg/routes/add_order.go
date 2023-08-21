package routes

import (
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
)

func (c *AddOrderController) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "add_order"
	page.Title = "New Order"
	page.Form = orderForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*orderForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *AddOrderController) Post(ctx echo.Context) error {
	var form orderForm
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
		Save(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "unable to create order")
	}

	ctx.Logger().Infof("order created: %s", o.ID)
	msg.Success(ctx, "The order was successfully added.")

	return c.Redirect(ctx, "orders")
}
