package routes

import (
	"fmt"

	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/msg"

	"github.com/labstack/echo/v4"
)


func (c *ProductController) handleAddGet(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "products/add"
	page.Title = "Add New Product"
	page.Form = ProductForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*ProductForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *ProductController) handleAddPost(ctx echo.Context) error {
	var form ProductForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse product form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.handleAddGet(ctx)
	}

	// Attempt creating the product
	p, err := c.Container.ORM.Product.
		Create().
		SetName(form.Name).
		SetSku(form.Sku).
		SetDescription(form.Description).
		SetPrice(form.Price).
		SetStockCount(form.Quantity).
		Save(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "unable to create product")
	}

	ctx.Logger().Infof("created: %s",p.Name)
	msg.Success(ctx, fmt.Sprintf("Added successfully: %s", p.Name))

	return c.Redirect(ctx, "products")
}
