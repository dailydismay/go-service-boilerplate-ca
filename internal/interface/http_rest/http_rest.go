package httprest

import (
	"authsvc/internal/core/config"
	"authsvc/internal/interface/http_rest/common"
	"context"
	"fmt"
	"time"

	"go.uber.org/fx"

	"github.com/gofiber/fiber/v2"
)

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

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				err := f.Listen(fmt.Sprintf("localhost:%d", config.HttpPort))
				if err != nil {
					panic(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return f.ShutdownWithContext(ctx)
		},
	})
}
