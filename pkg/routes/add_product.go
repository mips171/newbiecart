package routes

import (
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/msg"

	"github.com/labstack/echo/v4"
)

type (
	AddProductController struct {
		controller.Controller
		Client *ent.Client
	}

	productForm struct {
		Name        string  `form:"name" validate:"required"`
		Sku         string  `form:"sku" validate:"required"`
		Description string  `form:"description" validate:"required"`
		Price       float64 `form:"price" validate:"required,gte=0"`
		Quantity    int     `form:"quantity" validate:"required,gte=0"`
		Submission  controller.FormSubmission
	}
)

func (c *AddProductController) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "add_product"
	page.Title = "Add New Product"
	page.Form = productForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*productForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *AddProductController) Post(ctx echo.Context) error {
	var form productForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse product form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.Get(ctx)
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

	ctx.Logger().Infof("product created: %s", p.Name)
	msg.Success(ctx, "The product was successfully added.")

	return c.Redirect(ctx, "products")
}
