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

// Trace() calls the corresponding function of the original logrus package
func (lw *Logger) Trace(args ...interface{}) {
	lw.logger.Trace(args...)
}

// Tracef() calls the corresponding function of the original logrus package
func (lw *Logger) Tracef(format string, args ...interface{}) {
	lw.logger.Tracef(format, args...)
}

// Traceln() calls the corresponding function of the original logrus package
func (lw *Logger) Traceln(format string, args ...interface{}) {
	lw.logger.Traceln(args...)
}

// Debug() calls the corresponding function of the original logrus package
func (lw *Logger) Debug(args ...interface{}) {
	lw.logger.Debug(args...)
}

// Debugf() calls the corresponding function of the original logrus package
func (lw *Logger) Debugf(format string, args ...interface{}) {
	lw.logger.Debugf(format, args...)
}

// Debugln() calls the corresponding function of the original logrus package
func (lw *Logger) Debugln(format string, args ...interface{}) {
	lw.logger.Debugln(args...)
}

// Info() calls the corresponding function of the original logrus package
func (lw *Logger) Info(args ...interface{}) {
	lw.logger.Info(args...)
}

// Infof() calls the corresponding function of the original logrus package
func (lw *Logger) Infof(format string, args ...interface{}) {
	lw.logger.Infof(format, args...)
}

// Infoln() calls the corresponding function of the original logrus package
func (lw *Logger) Infoln(format string, args ...interface{}) {
	lw.logger.Infoln(args...)
}

// Warn() calls the corresponding function of the original logrus package
func (lw *Logger) Warn(args ...interface{}) {
	lw.logger.Warn(args...)
	lw.logger.Warning()
}

// Warnf() calls the corresponding function of the original logrus package
func (lw *Logger) Warnf(format string, args ...interface{}) {
	lw.logger.Warnf(format, args...)
}

// Warnln() calls the corresponding function of the original logrus package
func (lw *Logger) Warnln(format string, args ...interface{}) {
	lw.logger.Warnln(args...)
}

// Error() calls the corresponding function of the original logrus package
func (lw *Logger) Error(args ...interface{}) {
	lw.logger.Error(args...)
}

// Errorf() calls the corresponding function of the original logrus package
func (lw *Logger) Errorf(format string, args ...interface{}) {
	lw.logger.Errorf(format, args...)
}

// Errorln() calls the corresponding function of the original logrus package
func (lw *Logger) Errorln(format string, args ...interface{}) {
	lw.logger.Errorln(args...)
}

// Fatal() calls the corresponding function of the original logrus package
func (lw *Logger) Fatal(args ...interface{}) {
	lw.logger.Fatal(args...)
}

// Fatalf() calls the corresponding function of the original logrus package
func (lw *Logger) Fatalf(format string, args ...interface{}) {
	lw.logger.Fatalf(format, args...)
}

// Fatalln() calls the corresponding function of the original logrus package
func (lw *Logger) Fatalln(format string, args ...interface{}) {
	lw.logger.Fatalln(args...)
}

// Panic() calls the corresponding function of the original logrus package
func (lw *Logger) Panic(args ...interface{}) {
	lw.logger.Panic(args...)
}

// Panicf() calls the corresponding function of the original logrus package
func (lw *Logger) Panicf(format string, args ...interface{}) {
	lw.logger.Panicf(format, args...)
}

// Panicln() calls the corresponding function of the original logrus package
func (lw *Logger) Panicln(format string, args ...interface{}) {
	lw.logger.Panicln(args...)
}

//

// WithField() calls the corresponding function of the original logrus package
func (lw *Logger) WithField(key string, value interface{}) *logrus.Entry {
	return lw.logger.WithField(key, value)
}

// WithFields() calls the corresponding function of the original logrus package
func (lw *Logger) WithFields(fields logrus.Fields) *logrus.Entry {
	return lw.logger.WithFields(fields)
}

//

// GetWriter() calls the corresponding function of the original logrus package
func (lw *Logger) GetWriter() *io.PipeWriter {
	return lw.logger.Writer()
}
