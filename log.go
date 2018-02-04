package log

import (
	"fmt"
	stdlog "log"
)

//go:generate counterfeiter -o ../fakes/fake_logger.go . Logger

// Logger interface allows you to maintain a unified interface while using a
// custom logger. This allows you to write log statements without dictating
// the specific underlying library used for logging. You can avoid vendoring
// of logging libraries, which is especially useful when writing shared code
// such as a library. This package contains a simple logger and a no-op logger
// which both implement this interface. It is also supplemented with some
// additional helpers/shims for other common logging libraries such as logrus
type Logger interface {
	Debug(msg ...interface{})
	Info(msg ...interface{})
	Warn(msg ...interface{})
	Error(msg ...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})

	WithFields(fields map[string]interface{}) Logger
}

/**************
 Simple Logger
**************/

type simple struct {
	fields map[string]interface{}
}

// NewSimple creates a basic logger that wraps the core log library.
func NewSimple() Logger {
	return &simple{}
}

// WithFields will return a new logger based on the original logger
// with the additional supplied fields
func (b *simple) WithFields(fields map[string]interface{}) Logger {
	cp := &simple{}

	if b.fields == nil {
		cp.fields = fields
		return cp
	}

	cp.fields = make(map[string]interface{}, len(b.fields)+len(fields))
	for k, v := range b.fields {
		cp.fields[k] = v
	}

	for k, v := range fields {
		cp.fields[k] = v
	}

	return cp
}

// Debug log message
func (b *simple) Debug(msg ...interface{}) {
	stdlog.Printf("[DEBUG] %s %s", fmt.Sprint(msg...), pretty(b.fields))
}

// Info log message
func (b *simple) Info(msg ...interface{}) {
	stdlog.Printf("[INFO] %s %s", fmt.Sprint(msg...), pretty(b.fields))
}

// Warn log message
func (b *simple) Warn(msg ...interface{}) {
	stdlog.Printf("[WARN] %s %s", fmt.Sprint(msg...), pretty(b.fields))
}

// Error log message
func (b *simple) Error(msg ...interface{}) {
	stdlog.Printf("[ERROR] %s %s", fmt.Sprint(msg...), pretty(b.fields))
}

// Debugf log message with formatting
func (b *simple) Debugf(format string, args ...interface{}) {
	stdlog.Print(fmt.Sprintf("[DEBUG] "+format, args...), " ", pretty(b.fields))
}

// Infof log message with formatting
func (b *simple) Infof(format string, args ...interface{}) {
	stdlog.Print(fmt.Sprintf("[INFO] "+format, args...), " ", pretty(b.fields))
}

// Warnf log message with formatting
func (b *simple) Warnf(format string, args ...interface{}) {
	stdlog.Print(fmt.Sprintf("[WARN] "+format, args...), " ", pretty(b.fields))
}

// Errorf log message with formatting
func (b *simple) Errorf(format string, args ...interface{}) {
	stdlog.Print(fmt.Sprintf("[ERROR] "+format, args...), " ", pretty(b.fields))
}

// helper for pretty printing of fields
func pretty(m map[string]interface{}) string {
	if len(m) < 1 {
		return ""
	}

	s := ""
	for k, v := range m {
		s += fmt.Sprintf("%s=%v ", k, v)
	}

	return s[:len(s)-1]
}

/*************
 No-Op Logger
*************/

type noop struct{}

// NewNoop creates a no-op logger that can be used to silence all logging from this library.
func NewNoop() Logger {
	return &noop{}
}

// Debug log message no-op
func (n *noop) Debug(msg ...interface{}) {}

// Info log message no-op
func (n *noop) Info(msg ...interface{}) {}

// Warn log message no-op
func (n *noop) Warn(msg ...interface{}) {}

// Error log message no-op
func (n *noop) Error(msg ...interface{}) {}

// Debugf log message with formatting no-op
func (n *noop) Debugf(format string, args ...interface{}) {}

// Infof log message with formatting no-op
func (n *noop) Infof(format string, args ...interface{}) {}

// Warnf log message with formatting no-op
func (n *noop) Warnf(format string, args ...interface{}) {}

// Errorf log message with formatting no-op
func (n *noop) Errorf(format string, args ...interface{}) {}

// WithFields no-op
func (n *noop) WithFields(fields map[string]interface{}) Logger { return n }
