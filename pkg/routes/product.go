package routes

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/product"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/types"
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
		Price       types.Currency
		ID          int
	}

	ProductForm struct {
		ID          *int   `form:"id"` // Note use of pointer to allow for nil values, and no "required" validation
		Name        string `form:"name" validate:"required"`
		Sku         string `form:"sku" validate:"required"`
		Description string `form:"description" validate:"required"`
		Price       string `form:"price" validate:"required"`
		Quantity    int    `form:"quantity" validate:"required,gte=0"`
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

func (c *ProductController) handleEditByIdGet(ctx echo.Context) error {
	productID, err := c.getProductIDFromContext(ctx)
	if err != nil {
		return c.handleError(err, "Invalid product ID")
	}

	product, err := c.Container.ORM.Product.Get(ctx.Request().Context(), productID)
	if err != nil {
		return c.Fail(err, "Unable to fetch product")
	}

	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "products/edit"
	page.Title = "Edit Product"
	page.Form = ProductForm{
		ID:          &product.ID, // Make sure to pass the address of product.ID
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

func (c *ProductController) handleEditByIdPost(ctx echo.Context) error {
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
		return c.handleEditByIdGet(ctx)
	}

	ctx.Logger().Infof("Attempting to update product with ID: %d", form.ID)

	price, err := types.CurrencyFromString(form.Price)
	if err != nil {
		return c.Fail(err, "unable to process price")
	}

	// Fetch the existing product and update its fields
	product, err := c.Container.ORM.Product.UpdateOneID(*form.ID). // Assuming form.ID exists
									SetName(form.Name).
									SetSku(form.Sku).
									SetDescription(form.Description).
									SetPrice(price.String()).
									SetStockCount(form.Quantity).

		// ... set other fields ...
		Save(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "Unable to update product")
	}

	ctx.Logger().Infof("Product updated: %s", product.Name)
	msg.Success(ctx, "The product was successfully updated.")

	return c.Redirect(ctx, "products.view_all")
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
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "unable to create product")
	}

	ctx.Logger().Infof("created: %s", p.Name)
	msg.Success(ctx, "Added successfully: "+p.Name)

	return c.Redirect(ctx, "products.view_all")
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
		price, err := types.CurrencyFromString(p.Price)
		if err != nil {
			ctx.Logger().Errorf("could not convert price: %s", p.Name)
			continue
		}
		products[i] = Product{
			Name:        p.Name,
			Description: p.Description,
			Price:       price,
			ID:          p.ID,
		}
	}
	return products
}
