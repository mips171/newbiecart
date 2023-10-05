package routes

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/productcategory"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/msg"
)

type (
	// Interfaces for ProductCategory
	ProductCategoryGetter interface {
		GetAll(ctx echo.Context) error
		GetByID(ctx echo.Context) error
	}

	ProductCategoryEditor interface {
		EditByID(ctx echo.Context) error
	}

	ProductCategoryAdder interface {
		Add(ctx echo.Context) error
	}

	// ProductCategoryController with integrated methods
	ProductCategoryController struct {
		controller.Controller
		Client *ent.Client
	}

	ProductCategory struct {
		Name        string
		Description string
		ID          int
	}

	ProductCategoryForm struct {
		ID          *int   `form:"id"`
		Name        string `form:"name" validate:"required"`
		Description string `form:"description"`
		Submission  controller.FormSubmission
	}
)

// ... Implement the methods (similar to ProductController) for ProductCategoryController...
func (c *ProductCategoryController) getProductCategoryIDFromContext(ctx echo.Context) (int, error) {
	categoryID := ctx.Param("id")
	return strconv.Atoi(categoryID)
}

func (c *ProductCategoryController) GetAll(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "product_categories/view_all"
	page.Metatags.Description = "Browse our product categories."
	page.Pager = controller.NewPager(ctx, 10)
	page.Data = c.fetchProductCategories(ctx, &page.Pager)
	return c.RenderPage(ctx, page)
}

func (c *ProductCategoryController) fetchProductCategories(ctx echo.Context, pager *controller.Pager) []ProductCategory {
	categories, _ := c.Container.ORM.ProductCategory.Query().
		Offset(pager.GetOffset()).
		Limit(pager.ItemsPerPage).
		All(ctx.Request().Context())

	pager.SetItems(len(categories))

	categoryList := make([]ProductCategory, len(categories))
	for i, category := range categories {
		categoryList[i] = ProductCategory{
			Name:        category.Name,
			ID:          category.ID,
			Description: category.Description,
		}
	}
	return categoryList
}

func (c *ProductCategoryController) GetByID(ctx echo.Context) error {
	categoryID, err := c.getProductCategoryIDFromContext(ctx)
	if err != nil {
		return c.handleError(err, "Invalid product category ID")
	}

	category, err := c.Container.ORM.ProductCategory.
		Query().
		Where(productcategory.IDEQ(categoryID)).
		Only(ctx.Request().Context())

	if err != nil {
		return c.handleError(err, "Unable to fetch product category details")
	}

	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "product_categories/view"
	page.Title = category.Name
	page.Data = category

	return c.RenderPage(ctx, page)
}

func (c *ProductCategoryController) Add(ctx echo.Context) error {
	switch ctx.Request().Method {
	case echo.GET:
		return c.handleAddGet(ctx)
	case echo.POST:
		return c.handleAddPost(ctx)
	default:
		return echo.ErrMethodNotAllowed
	}
}

func (c *ProductCategoryController) handleAddGet(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "product_categories/add"
	page.Title = "Add New Product Category"
	page.Form = ProductCategoryForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*ProductCategoryForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *ProductCategoryController) handleAddPost(ctx echo.Context) error {
	var form ProductCategoryForm
	ctx.Set(context.FormKey, &form)

	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse product category form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.handleAddGet(ctx)
	}

	category, err := c.Container.ORM.ProductCategory.
		Create().
		SetName(form.Name).
		SetDescription(form.Description).
		Save(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "unable to create product category")
	}

	ctx.Logger().Infof("Created category: %s", category.Name)
	msg.Success(ctx, "Category added successfully: "+category.Name)

	return c.Redirect(ctx, "product_categories.view_all")
}

func (c *ProductCategoryController) EditByID(ctx echo.Context) error {
	switch ctx.Request().Method {
	case echo.GET:
		return c.handleEditByIdGet(ctx)
	case echo.POST:
		return c.handleEditByIdPost(ctx)
	default:
		return echo.ErrMethodNotAllowed
	}
}

func (c *ProductCategoryController) handleEditByIdGet(ctx echo.Context) error {
	categoryID, err := c.getProductCategoryIDFromContext(ctx)
	if err != nil {
		return c.handleError(err, "Invalid product category ID")
	}

	category, err := c.Container.ORM.ProductCategory.Get(ctx.Request().Context(), categoryID)
	if err != nil {
		return c.Fail(err, "Unable to fetch product category")
	}

	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "product_categories/edit"
	page.Title = "Edit Product Category"
	page.Form = ProductCategoryForm{
		ID:   &category.ID,
		Name: category.Name,
		// ... populate other fields ...
	}

	return c.RenderPage(ctx, page)
}

func (c *ProductCategoryController) handleEditByIdPost(ctx echo.Context) error {
	var form ProductCategoryForm
	ctx.Set(context.FormKey, &form)

	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse product category form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.handleEditByIdGet(ctx)
	}

	ctx.Logger().Infof("Attempting to update product category with ID: %d", *form.ID)

	category, err := c.Container.ORM.ProductCategory.UpdateOneID(*form.ID).
		SetName(form.Name).
		SetDescription(form.Description).
		// ... set other fields ...
		Save(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "Unable to update product category")
	}

	ctx.Logger().Infof("Product category updated: %s", category.Name)
	msg.Success(ctx, "The product category was successfully updated.")

	return c.Redirect(ctx, "product_categories.view_all")
}

func (c *ProductCategoryController) handleError(err error, msg string) error {
	// Handle errors, e.g., logging, setting HTTP status
	return c.Fail(err, msg)
}
