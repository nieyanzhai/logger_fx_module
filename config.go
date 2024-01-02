package logger_fx_module

type Log struct {
	Path       string
	Name       string
	MaxSize    int `json:"max_size"`
	MaxBackups int `json:"max_backups"`
	MaxAge     int `json:"max_age"`
	Compress   bool
	LocalTime  bool `json:"local_time"`
}
