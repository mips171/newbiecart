package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/product"
	"github.com/mikestefanello/pagoda/pkg/controller"
)

type (
	ProductController struct {
		controller.Controller
		Client *ent.Client
	}

	Product struct {
		Name        string
		Description string
		Price       float64
	}
)

func (c *ProductController) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "products"
	page.Metatags.Description = "Browse our products."
	page.Metatags.Keywords = []string{"Shopping", "Products", "Buy"}
	page.Pager = controller.NewPager(ctx, 10)
	page.Data = c.fetchProducts(ctx, &page.Pager)

	return c.RenderPage(ctx, page)
}

// fetchProducts fetches products from the database and maps them to the local Product struct
func (c *ProductController) fetchProducts(ctx echo.Context, pager *controller.Pager) []Product {
	prods, _ := c.Container.ORM.Product.Query().
		Offset(pager.GetOffset()).
		Limit(pager.ItemsPerPage).
		Order(ent.Asc(product.FieldName)).
		All(ctx.Request().Context())

	pager.SetItems(len(prods))

	products := make([]Product, len(prods))
	for i, p := range prods {
		products[i] = Product{
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		}
	}
	return products
}
