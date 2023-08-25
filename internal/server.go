package internal

import (
	"context"
	"log/slog"
	"net"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	openapi "github.com/go-openapi/runtime/middleware"
	"go.uber.org/fx"

	"github.com/LazyCodeTeam/just-code-backend/internal/api"
	"github.com/LazyCodeTeam/just-code-backend/internal/api/handler"
	appMiddleware "github.com/LazyCodeTeam/just-code-backend/internal/api/middleware"
	"github.com/LazyCodeTeam/just-code-backend/internal/config"
	"github.com/LazyCodeTeam/just-code-backend/internal/core"
	"github.com/LazyCodeTeam/just-code-backend/internal/data"
)

func StartServer() {
	fx.New(
		fx.Provide(
			newServer,
			newFirebaseApp,
			newFirebaseAuthClient,
			config.New,
			fx.Annotate(
				newMux,
				fx.ParamTags(`group:"routes"`),
			),
		),
		fx.Provide(api.Providers()...),
		fx.Provide(data.Providers()...),
		fx.Provide(core.Providers()...),
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
	cmsHandler *handler.AdminContentHandler,
	healthHandler *handler.HealthHandler,
	authTokenValidator *appMiddleware.AuthTokenValidator,
) *chi.Mux {
	mux := chi.NewRouter()

	redocOpts := openapi.RedocOpts{
		SpecURL:  "swagger.yaml",
		BasePath: "/api",
		Path:     "/docs",
	}
	redoc := openapi.Redoc(redocOpts, nil)

	swatterUIOpts := openapi.SwaggerUIOpts{
		SpecURL:  "swagger.yaml",
		BasePath: "/api",
		Path:     "/swagger-ui",
	}
	swaggerUI := openapi.SwaggerUI(swatterUIOpts, nil)

	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Recoverer)

	mux.Method(http.MethodGet, "/api/docs", redoc)
	mux.Method(http.MethodGet, "/api/swagger-ui", swaggerUI)
	mux.Method(http.MethodGet, "/api/swagger.yaml", http.FileServer(http.Dir("./")))

	healthHandler.Register(mux)

	mux.Route("/admin/api", func(router chi.Router) {
		router.Use(authTokenValidator.Handle)
		cmsHandler.Register(router)
	})

	mux.Route("/api", func(router chi.Router) {
		router.Use(authTokenValidator.Handle)
		for _, h := range handlers {
			h.Register(router)
		}
	})

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
