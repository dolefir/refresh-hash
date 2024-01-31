package logger

import (
	"fmt"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ModDev = iota
	ModProd
)

const (
	FormatJSON = iota
	FormatConsole
)

const (
	prodENV = "prod"
	devENV  = "dev"

	jsonFormat = "json"
	textFormat = "text"

	infoLvl  = "info"
	debugLvl = "debug"
	errorLvl = "error"
)

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = zapcore.DebugLevel
	// InfoLevel is the default logging priority.
	InfoLevel = zapcore.InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = zapcore.WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = zapcore.ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel = zapcore.DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel = zapcore.PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = zapcore.FatalLevel
)

// Option options for zap setup
type Option func(z *zap.Logger) *zap.Logger

// SelectLevel ...
func SelectLevel(level int) Option {
	return func(z *zap.Logger) *zap.Logger {
		return z.WithOptions(zap.IncreaseLevel(zapcore.Level(level)))
	}
}

// Name ...
func Name(name string) Option {
	return func(z *zap.Logger) *zap.Logger {
		return z.Named(name)
	}
}

// Tags ...
func Tags(fields map[string]string) Option {
	return func(z *zap.Logger) *zap.Logger {
		if len(fields) > 0 {
			zapFields := make([]zap.Field, 0, len(fields))
			for k, field := range fields {
				zapFields = append(zapFields, zap.String(k, field))
			}
			return z.With(zapFields...)
		}
		return z
	}
}

// Output ...
func Output(o io.Writer) Option {
	return func(z *zap.Logger) *zap.Logger {
		if o != nil {
			return z.WithOptions(zap.ErrorOutput(zapcore.AddSync(o)))
		}
		return z
	}
}

type zapLog struct {
	log *zap.SugaredLogger
}

// LoggerEnv ...
func LoggerEnv(mode, format int, options ...Option) Logger {
	var cfg zapcore.EncoderConfig
	var enc zapcore.Encoder

	switch mode {
	case ModDev:
		cfg = zap.NewDevelopmentEncoderConfig()
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	case ModProd:
		cfg = zap.NewProductionEncoderConfig()
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		cfg = zap.NewDevelopmentEncoderConfig()
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	cfg.FunctionKey = "F"

	switch format {
	case FormatJSON:
		enc = zapcore.NewJSONEncoder(cfg)
	case FormatConsole:
		enc = zapcore.NewConsoleEncoder(cfg)
	default:
		enc = zapcore.NewJSONEncoder(cfg)
	}

	coreJSON := zapcore.NewCore(
		enc,
		zapcore.AddSync(os.Stdout),
		zap.DebugLevel,
	)

	log := zap.New(coreJSON)
	log = log.WithOptions(zap.AddCallerSkip(1))

	if len(options) > 0 {
		for _, option := range options {
			log = option(log)
		}
	}

	return zapLog{log: log.Sugar()}
}

type CFGLogger struct {
	// working mode dev/prod
	Mode string
	// format mode text/json
	LogFormat string
	// log level debug/error
	LogLevel string
}

func logMod(mode string) int {
	var m int

	switch mode {
	case prodENV:
		m = ModProd
	case devENV:
		m = ModDev
	default:
		fmt.Println("Logger mode not set, default: dev mode")
		m = ModDev
	}

	return m
}

func logFormat(format string) int {
	var f int

	switch format {
	case jsonFormat:
		f = FormatJSON
	case textFormat:
		f = FormatConsole
	default:
		fmt.Println("Logger log format not set, default: text mode")
		f = FormatConsole
	}

	return f
}

func logLeven(level string) int {
	var l int

	switch level {
	case infoLvl:
		l = int(InfoLevel)
	case debugLvl:
		l = int(DebugLevel)
	case errorLvl:
		l = int(ErrorLevel)
	default:
		fmt.Println("Logger log level not set, default: debug level")
		l = int(DebugLevel)
	}

	return l
}

func NewLogger(cfg *CFGLogger, tags map[string]string) Logger {
	return LoggerEnv(
		logMod(cfg.Mode),
		logFormat(cfg.LogFormat),
		Tags(tags),
		SelectLevel(logLeven(cfg.LogLevel)),
	)
}

// Info writes a information message.
func (z zapLog) Info(args ...interface{}) {
	z.log.Info(args...)
}

// Infof writes a formatted information message.
func (z zapLog) Infof(template string, args ...interface{}) {
	z.log.Infof(template, args...)
}

// Infow writes a formatted information message with key, val pairs
func (z zapLog) Infow(template string, args ...interface{}) {
	z.log.Infow(template, args...)
}

// Warn writes a warning message.
func (z zapLog) Warn(args ...interface{}) {
	z.log.Warn(args...)
}

// Warnf writes a formatted warning message.
func (z zapLog) Warnf(template string, args ...interface{}) {
	z.log.Warnf(template, args...)
}

// Warnw writes a formatted information message with key, val pairs
func (z zapLog) Warnw(template string, args ...interface{}) {
	z.log.Warnw(template, args...)
}

// Error writes an error message.
func (z zapLog) Error(args ...interface{}) {
	z.log.Error(args...)
}

// Errorf writes a formatted error message.
func (z zapLog) Errorf(template string, args ...interface{}) {
	z.log.Errorf(template, args...)
}

// Errorw writes a formatted information message with key, val pairs
func (z zapLog) Errorw(template string, args ...interface{}) {
	z.log.Errorw(template, args...)
}

// Debug writes a debug message.
func (z zapLog) Debug(args ...interface{}) {
	z.log.Debug(args...)
}

// Debugf writes a formatted debug message.
func (z zapLog) Debugf(template string, args ...interface{}) {
	z.log.Debugf(template, args...)
}

// Debugw writes a formatted information message with key, val pairs
func (z zapLog) Debugw(template string, args ...interface{}) {
	z.log.Debugw(template, args...)
}

// Fatal writes a fatal message.
func (z zapLog) Fatal(args ...interface{}) {
	z.log.Fatal(args...)
}

// Fatalf writes a formatted fatal message.
func (z zapLog) Fatalf(template string, args ...interface{}) {
	z.log.Fatalf(template, args...)
}

// With add fields to be used for all logs
func (z zapLog) With(fields ...interface{}) Logger {
	return zapLog{log: z.log.With(fields...)}
}

// Flush any buffered log entries
func (z zapLog) Flush() error {
	return z.log.Sync()
}
