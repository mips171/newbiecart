package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/user"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
)

type UserPage struct {
    Controller controller.Controller
}

func (u *UserPage) Details(ctx echo.Context) error {
    // Get the authenticated user from the context
    authUser := ctx.Get(context.AuthenticatedUserKey).(*ent.User)

    // Query the user details from the database
    usr, err := u.Controller.Container.ORM.User.
        Query().
        Where(user.ID(authUser.ID)).
        Only(ctx.Request().Context())

    if err != nil {
        return u.Controller.Fail(err, "unable to fetch user details")
    }

    // Render the user details page
    page := controller.NewPage(ctx)
    page.Layout = "main"
    page.Name = "user/view"
    page.Title = "User Details"
    page.Data = usr // Pass the user details to the template

    return u.Controller.RenderPage(ctx, page)
}