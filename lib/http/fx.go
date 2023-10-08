package http

import (
	"net/http"

	"go.uber.org/fx"
)

var Module = fx.Module("http", fx.Provide(
	fx.Annotate(
		NewServer,
		fx.As(new(http.Handler)),
	),
))
