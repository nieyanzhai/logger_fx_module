package logger_fx_module

import "go.uber.org/fx"

var Module = fx.Module(
	"logger",
	fx.Provide(
		newLogger,
	),
)
