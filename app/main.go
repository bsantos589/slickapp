package main

import (
	slick "github.com/bsantos589/slickapp"
    slickapp/view/profile
)

func main() {
	app := slick.New()

	app.Get("/profile", HandleProfileIndex)
	app.Get("/dashboard", HandleDashBoardIndex)

	println("Starting Slick server...") // Add this
	if err := app.Start(); err != nil {
		println("Slick failed to start:", err.Error())
	}

}
func HandleUserProfile(c *slick.Context) error {
	user := profile.User{
		FirstName: "Brandon",
		Lastname:  "Santos",
		Email:     "bs@bs.com",
	}
	return c.Render(profile.Index(user))
}

func HandleDashboardIndex(c *slick.Context) error {
	return c.Render(profile.Index(user))
}
