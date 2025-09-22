package main

import (
	"github.com/anthdm/slick"
)

func main() {
	app := slick.New()

	app.Get("/profile", HandleUserProfile)

	app.Start(":3000")

}
func HandleUserProfile(c *slick.Context) error {
	return c.Render(profile.Index())
}
