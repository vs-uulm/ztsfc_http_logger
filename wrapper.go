// The logger is based basically on the logrus implementation
// (see https://github.com/sirupsen/logrus),
// but has extensions, that allow some customization of logging messages.
//
// The logger's wrapper.go contains:
//   - wrapping functions for writing log messages of all supported logging levels:
//     [Trace, Debug, Info, Warning, Error, Fatal, Panic]
//   - functions WithField and WithFields, that return logger entries, that write
//     log messages with additinal fields
//   - function GetWriter, that returns logger's writer for using as ErrorLog in
//     http.Server or httputil.ReverseProxy instances

package ztsfc_http_logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

type Fields map[string]interface{}

// Trace() calls the corresponding function of the original logrus package
func (logger *Logger) Trace(args ...interface{}) {
	logger.lr.Trace(args...)
}

// Tracef() calls the corresponding function of the original logrus package
func (logger *Logger) Tracef(format string, args ...interface{}) {
	logger.lr.Tracef(format, args...)
}

// Traceln() calls the corresponding function of the original logrus package
func (logger *Logger) Traceln(format string, args ...interface{}) {
	logger.lr.Traceln(args...)
}

// Debug() calls the corresponding function of the original logrus package
func (logger *Logger) Debug(args ...interface{}) {
	logger.lr.Debug(args...)
}

// Debugf() calls the corresponding function of the original logrus package
func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.lr.Debugf(format, args...)
}

// Debugln() calls the corresponding function of the original logrus package
func (logger *Logger) Debugln(format string, args ...interface{}) {
	logger.lr.Debugln(args...)
}

// Info() calls the corresponding function of the original logrus package
func (logger *Logger) Info(args ...interface{}) {
	logger.lr.Info(args...)
}

// Infof() calls the corresponding function of the original logrus package
func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.lr.Infof(format, args...)
}

// Infoln() calls the corresponding function of the original logrus package
func (logger *Logger) Infoln(format string, args ...interface{}) {
	logger.lr.Infoln(args...)
}

// Warn() calls the corresponding function of the original logrus package
func (logger *Logger) Warn(args ...interface{}) {
	logger.lr.Warn(args...)
	logger.lr.Warning()
}

// Warnf() calls the corresponding function of the original logrus package
func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.lr.Warnf(format, args...)
}

// Warnln() calls the corresponding function of the original logrus package
func (logger *Logger) Warnln(format string, args ...interface{}) {
	logger.lr.Warnln(args...)
}

// Error() calls the corresponding function of the original logrus package
func (logger *Logger) Error(args ...interface{}) {
	logger.lr.Error(args...)
}

// Errorf() calls the corresponding function of the original logrus package
func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.lr.Errorf(format, args...)
}

// Errorln() calls the corresponding function of the original logrus package
func (logger *Logger) Errorln(format string, args ...interface{}) {
	logger.lr.Errorln(args...)
}

// Fatal() calls the corresponding function of the original logrus package
func (logger *Logger) Fatal(args ...interface{}) {
	logger.lr.Fatal(args...)
}

// Fatalf() calls the corresponding function of the original logrus package
func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.lr.Fatalf(format, args...)
}

// Fatalln() calls the corresponding function of the original logrus package
func (logger *Logger) Fatalln(format string, args ...interface{}) {
	logger.lr.Fatalln(args...)
}

// Panic() calls the corresponding function of the original logrus package
func (logger *Logger) Panic(args ...interface{}) {
	logger.lr.Panic(args...)
}

// Panicf() calls the corresponding function of the original logrus package
func (logger *Logger) Panicf(format string, args ...interface{}) {
	logger.lr.Panicf(format, args...)
}

// Panicln() calls the corresponding function of the original logrus package
func (logger *Logger) Panicln(format string, args ...interface{}) {
	logger.lr.Panicln(args...)
}

//

// WithField() calls the corresponding function of the original logrus package
func (logger *Logger) WithField(key string, value interface{}) *logrus.Entry {
	return logger.lr.WithField(key, value)
}

// WithFields() calls the corresponding function of the original logrus package
func (logger *Logger) WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.lr.WithFields(fields)
}

//

// GetWriter() calls the corresponding function of the original logrus package
func (logger *Logger) GetWriter() *io.PipeWriter {
	return logger.lr.Writer()
}
