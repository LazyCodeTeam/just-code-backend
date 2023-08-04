package internal

import (
	"context"
	"net"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
	"golang.org/x/exp/slog"

	"github.com/LazyCodeTeam/just-code-backend/internal/api"
	"github.com/LazyCodeTeam/just-code-backend/internal/api/handler"
)

func StartServer() {
	fx.New(
		fx.Provide(
			newServer,
			newFirebaseApp,
			newFirebaseAuthClient,
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

func newFirebaseApp() (*firebase.App, error) {
	config := &firebase.Config{
		ProjectID: "just-code-dev",
	}
	app, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		slog.Error("Error initializing firebase app: %v\n", "error", err)

		return nil, err
	}

	return app, nil
}

func newFirebaseAuthClient(app *firebase.App) (*auth.Client, error) {
	client, err := app.Auth(context.Background())
	if err != nil {
		slog.Error("Error initializing firebase auth client: %v\n", "error", err)

		return nil, err
	}

	return client, nil
}
