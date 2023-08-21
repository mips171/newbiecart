package routes

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/msg"
)

type EditProductController struct {
	controller.Controller
	Client *ent.Client
}

func (c *EditProductController) Get(ctx echo.Context) error {
	productID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return c.Fail(err, "Invalid product ID")
	}

	product, err := c.Container.ORM.Product.Get(ctx.Request().Context(), productID)
	if err != nil {
		return c.Fail(err, "Unable to fetch product")
	}

	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "edit_product"
	page.Title = "Edit Product"
	page.Form = productForm{
		ID:          product.ID,
		Name:        product.Name,
		Sku:         product.Sku,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.StockCount,
		// ... populate other fields ...
	}

	// ... rest of your code ...
	return c.RenderPage(ctx, page)
}

func (c *EditProductController) Post(ctx echo.Context) error {
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

	ctx.Logger().Infof("Attempting to update product with ID: %d", form.ID)

	// Fetch the existing product and update its fields
	product, err := c.Container.ORM.Product.UpdateOneID(form.ID). // Assuming form.ID exists
									SetName(form.Name).
									SetSku(form.Sku).
									SetDescription(form.Description).
									SetPrice(form.Price).
									SetStockCount(form.Quantity).
		// ... set other fields ...
		Save(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "Unable to update product")
	}

	ctx.Logger().Infof("Product updated: %s", product.Name)
	msg.Success(ctx, "The product was successfully updated.")

	return c.Redirect(ctx, "products")
}