package routes

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/product"
	"github.com/mikestefanello/pagoda/pkg/controller"
)

type (

	// Interfaces to make actions explicit
	ProductGetter interface {
		GetAll(ctx echo.Context) error
		GetByID(ctx echo.Context) error
	}

	ProductEditor interface {
		EditByID(ctx echo.Context) error
	}

	ProductAdder interface {
		Add(ctx echo.Context) error
	}

	// ProductController with integrated methods
	ProductController struct {
		controller.Controller
		Client *ent.Client
	}

	Product struct {
		Name        string
		Description string
		Price       float64
		ID          int
	}

	ProductForm struct {
		ID          *int    `form:"id"` // Note use of pointer to allow for nil values, and no "required" validation
		Name        string  `form:"name" validate:"required"`
		Sku         string  `form:"sku" validate:"required"`
		Description string  `form:"description" validate:"required"`
		Price       float64 `form:"price" validate:"required,gte=0"`
		Quantity    int     `form:"quantity" validate:"required,gte=0"`
		Submission  controller.FormSubmission
	}
)

var _ ProductGetter = &ProductController{}
var _ ProductEditor = &ProductController{}
var _ ProductAdder = &ProductController{}

func (c *ProductController) GetAll(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "products/view_all" // Located at templates/pages/products/
	page.Metatags.Description = "Browse our products."
	page.Metatags.Keywords = []string{"Shopping", "Products", "Buy"}
	page.Pager = controller.NewPager(ctx, 10)
	page.Data = c.fetchProducts(ctx, &page.Pager)

	return c.RenderPage(ctx, page)
}

func (c *ProductController) GetByID(ctx echo.Context) error {
	productID, err := c.getProductIDFromContext(ctx)
	if err != nil {
		return c.handleError(err, "Invalid product ID")
	}

	product, err := c.Container.ORM.Product.
		Query().
		Where(product.IDEQ(productID)).
		Only(ctx.Request().Context())

	if err != nil {
		return c.handleError(err, "Unable to fetch product details")
	}

	page := controller.NewPage(ctx)
	page.Layout = "product_layout"
	page.Name = "products/view"
	page.Title = product.Name
	page.Data = product

	return c.RenderPage(ctx, page)
}

func (c *ProductController) EditByID(ctx echo.Context) error {
	switch ctx.Request().Method {
	case echo.GET:
		return c.handleEditByIdGet(ctx)
	case echo.POST:
		return c.handleEditByIdPost(ctx)
	default:
		return echo.ErrMethodNotAllowed
	}
}

func (c *ProductController) Add(ctx echo.Context) error {
	switch ctx.Request().Method {
	case echo.GET:
		return c.handleAddGet(ctx)
	case echo.POST:
		return c.handleAddPost(ctx)
	default:
		return echo.ErrMethodNotAllowed
	}
}

// Helper methods
func (c *ProductController) getProductIDFromContext(ctx echo.Context) (int, error) {
	productID := ctx.Param("id")
	return strconv.Atoi(productID)
}

func (c *ProductController) handleError(err error, msg string) error {
	// Handle errors, e.g., logging, setting HTTP status
	return c.Fail(err, msg)
}

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
			ID:          p.ID,
		}
	}
	return products
}
