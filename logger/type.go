package logger

// Logger common logger interface.
type Logger interface {
	// Info writes a information message.
	Info(args ...interface{})
	// Infof writes a formatted information message.
	Infof(template string, args ...interface{})
	// Infow writes a formatted information message with key-value pairs.
	Infow(template string, args ...interface{})
	// Warn writes a warning message.
	Warn(args ...interface{})
	// Warnf writes a formatted warning message.
	Warnf(template string, args ...interface{})
	// Warnw writes a formatted information message with key-value pairs.
	Warnw(template string, args ...interface{})
	// Error writes an error message.
	Error(args ...interface{})
	// Errorf writes a formatted error message.
	Errorf(template string, args ...interface{})
	// Errorw writes a formatted information message with key-value pairs.
	Errorw(template string, args ...interface{})
	// Debug writes a debug message.
	Debug(args ...interface{})
	// Debugf writes a formatted debug message.
	Debugf(template string, args ...interface{})
	// Debugw writes a formatted information message with key-value pairs.
	Debugw(template string, args ...interface{})
	// Fatal writes a fatal message.
	Fatal(args ...interface{})
	// Fatalf writes a formatted fatal message.
	Fatalf(template string, args ...interface{})
	// With add fields to be used for all logs.
	With(fields ...interface{}) Logger
	// Flush any buffered log entries.
	Flush() error
}
