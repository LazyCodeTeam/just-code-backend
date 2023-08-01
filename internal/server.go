package internal

import (
	"context"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"

	"github.com/LazyCodeTeam/just-code-backend/internal/api"
	"github.com/LazyCodeTeam/just-code-backend/internal/api/handler"
)

func StartServer() {
	fx.New(
		fx.Provide(
			newServer,
			fx.Annotate(
				newMux,
				fx.ParamTags(`group:"routes"`),
			),
		),
		fx.Provide(api.Providers()...),
		fx.Invoke(startListener),
	).Run()
}

func startListener(lc fx.Lifecycle, server *http.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}
			go server.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}

func newMux(
	handlers []handler.Handler,
) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Recoverer)

	for _, h := range handlers {
		h.Register(mux)
	}

	return mux
}

func newServer(mux *chi.Mux) *http.Server {
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return &server
}
