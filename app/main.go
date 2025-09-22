package main

import (
	"fmt"
	"slickapp/view/dashboard"
	"slickapp/view/profile"

	"github.com/anthdm/slick"
	"github.com/google/uuid"
)

func main() {
	app := slick.New()

	app.Plug(WithRequestID, WithAunth)
	ph := NewProfileHandler(&NOOPSB{})

	app.Get("/profile", ph.HandleProfileIndex)
	app.Get("/dashboard", HandleDashboardIndex)

	println("Starting Slick server...") // Add this
	if err := app.Start(); err != nil {
		println("Slick failed to start:", err.Error())
	}

}

func WithAunth(h slick.Handler) slick.Handler {
	return func(c *slick.Context) error {
		fmt.Println("auth")
		c.Set("email", "jannine@hr.com")
		return h(c)
	}
}

func WithRequestID(h slick.Handler) slick.Handler {
	return func(c *slick.Context) error {
		fmt.Println("request")
		c.Set("requestID", uuid.New())
		return h(c)
	}
}

type SupabaseClient interface {
	Auth(foo string) error
}

type NOOPSB struct{}

func (NOOPSB) Auth(foo string) error { return nil }

type ProfileHandler struct {
	sbClient SupabaseClient
	//..
}

func NewProfileHandler(sb SupabaseClient) *ProfileHandler {
	return &ProfileHandler{
		sbClient: sb,
	}
}

func (h *ProfileHandler) HandleProfileIndex(c *slick.Context) error {
	user := profile.User{
		FirstName: "Brandon",
		Lastname:  "Santos",
		Email:     "bs@bs.com",
	}
	return c.Render(profile.Index(user))
}

func HandleDashboardIndex(c *slick.Context) error {
	fmt.Println(c.Get("requestID"))
	return c.Render(dashboard.Index())
}
