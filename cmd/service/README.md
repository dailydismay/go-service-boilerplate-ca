# [FX](https://github.com/uber-go/fx) Application initialization

Nothing interesting except handlers grouping. It's done with this construction:

Handlers depend on usecases and domain services:

```go
type Handler interface {
	Pattern() string
	Method() string
	Middleware() []fiber.Handler
	Handle(*fiber.Ctx) error
}

// interfaces/http_rest/auth/post_register/post_register.go
type handler struct {
	usecase register.UseCase
}

func New(usecase register.UseCase) common.Handler {
	return &handler{usecase}
}

```

Grouping gives us an ability to remove massive dependencies list in root http server (and other interface ports).

```go
fx.Provide(
    fx.Annotate(
        post_login.New,
        fx.ResultTags(`group:"handlers"`),
    ),
    fx.Annotate(
        post_register.New,
        fx.ResultTags(`group:"handlers"`),
    ),
),
fx.Invoke(
    // annotations are struct tags alternative
    // we need to group handlers into slice and pass to api router
    fx.Annotate(httprest.New, fx.ParamTags(``, ``, `group:"handlers"`)),
),
```

```go
func New(lc fx.Lifecycle, config *config.Config, handlers []common.Handler) {
	f := fiber.New(
		fiber.Config{
			ErrorHandler: common.ErrorHandler,
			ReadTimeout:  time.Second * 3,
		},
	)
	for _, h := range handlers {
		f.Add(h.Method(), h.Pattern(), append(h.Middleware(), h.Handle)...)
	}

    // ... see more in internal/interfaces/http_rest/http_rest.go
}
```
