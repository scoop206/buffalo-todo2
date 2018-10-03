package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	if c.Value("current_user") != nil {
		return c.Redirect(302, "/items")
	} else {
		return c.Redirect(302, "/signin")
	}
}
