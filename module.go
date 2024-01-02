package logger_fx_module

import "go.uber.org/fx"

// Add Logger to fx module
var module = fx.Module(
	"logger",
	fx.Provide(
		newLogger,
	),
)
