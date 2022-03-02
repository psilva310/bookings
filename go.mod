module github.com/psilva310/bookings

go 1.17

require pkg/handlers v0.0.0

require pkg/render v0.0.0

require pkg/config v0.0.0

require (
	github.com/alexedwards/scs/v2 v2.5.0
	github.com/justinas/nosurf v1.1.1
	pkg/models v0.0.0 // indirect
)

require (
	github.com/go-chi/chi v1.5.4
	github.com/go-chi/chi/v5 v5.0.7
)

replace pkg/handlers => ./pkg/handlers

replace pkg/render => ./pkg/render

replace pkg/config => ./pkg/config

replace pkg/models => ./pkg/models
