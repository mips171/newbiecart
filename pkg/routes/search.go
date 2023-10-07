package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent/product"
	"github.com/mikestefanello/pagoda/pkg/controller"
)

type (
	search struct {
		controller.Controller
	}

	searchResult struct {
		Title string
		URL   string
	}
)

func (c *search) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "search"

	var results []searchResult

	search := ctx.QueryParam("query")
	if search != "" {
		// Search the database using ent ORM
		products, err := c.Container.ORM.Product.
			Query().
			Where(product.NameContainsFold(search)). // Using a predicate to find products containing the search query in their names
			All(ctx.Request().Context())             // Pass context for handling request lifecycle

		if err != nil {
			// Handle error. You might want to return or log this error.
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch data"})
		}

		// Convert the ORM results into your search results.
		for _, product := range products {
			results = append(results, searchResult{
				Title: product.Name,
				URL:   fmt.Sprintf("/products/%d", product.ID),
			})
		}

		// Search other entities here.
	}
	page.Data = results

	return c.RenderPage(ctx, page)
}

func (c *search) SearchProducts(ctx echo.Context) error {
	query := ctx.QueryParam("q")
	products, err := c.Container.ORM.Product.Query().
		Where(product.NameContains(query)).
		Limit(10).
		All(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch products"})
	}

	var response string
	for _, product := range products {
		response += fmt.Sprintf(
			`<li>
				<a href="#" hx-get="/products/%d/row" hx-target="#selected-products" hx-swap="beforeend">%s</a>
			</li>`,
			product.ID, product.Name)
	}

	return ctx.HTML(http.StatusOK, response)
}