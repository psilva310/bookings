module github.com/psilva310/bookings

go 1.17

require (
	github.com/alexedwards/scs/v2 v2.5.0
	github.com/justinas/nosurf v1.1.1
)

require (
	config v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi v1.5.4
	github.com/go-chi/chi/v5 v5.0.7
	handlers v0.0.0-00010101000000-000000000000
	render v0.0.0-00010101000000-000000000000
)

require models v0.0.0-00010101000000-000000000000 // indirect

replace handlers => ./cmd/internal/handlers

replace render => ./cmd/internal/render

replace config => ./cmd/internal/config

replace models => ./cmd/internal/models
