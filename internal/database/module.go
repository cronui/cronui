package database

import (
	"context"

	"go.uber.org/fx"

	"github.com/cronui/cronui/ent"
)

var Module = fx.Options(
	fx.Provide(NewDatabase),
	fx.Invoke(register),
)

func register(lifecycle fx.Lifecycle, db *ent.Client) {
	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})
}
