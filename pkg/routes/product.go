package routes

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent/product"
	"github.com/mikestefanello/pagoda/pkg/controller"
)

type (
	productDetail struct {
		controller.Controller
	}
)

func (c *productDetail) Get(ctx echo.Context) error {
	productID := ctx.Param("id")

	// make productID an int
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		return c.Fail(err, "invalid product ID")
	}

	// Fetch the product from the database using its ID
	product, err := c.Container.ORM.Product.
		Query().
		Where(product.IDEQ(productIDInt)).
		Only(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "unable to fetch product details")
	}

	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "product"
	page.Title = product.Name
	page.Data = product

	return c.RenderPage(ctx, page)

}
