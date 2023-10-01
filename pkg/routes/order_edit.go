package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/msg"
)

type (
	editOrderForm struct {
		ID          int     `form:"id" validate:"required"`
		BalanceDue       float64 `form:"balance_due" validate:"required"`
		Submission  controller.FormSubmission
	}
)

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
	page.Title = "Edit Order"
	page.Form = editOrderForm{
		ID:          order.ID,

		// ... populate other fields ...
	}

	// ... rest of your code ...
	return c.RenderPage(ctx, page)
}

func (c *OrderController) handleEditByIdPost(ctx echo.Context) error {
	var form editOrderForm
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

	// Fetch the existing order and update its fields
	order, err := c.Container.ORM.Order.UpdateOneID(form.ID). // Assuming form.ID exists
		SetBalanceDue(form.BalanceDue).

		// ... set other fields ...

		Save(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "Unable to update order")
	}

	ctx.Logger().Infof("Order updated:", order)
	msg.Success(ctx, "The order was successfully updated.")

	return c.Redirect(ctx, "orders")
}
