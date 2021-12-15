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
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger    *logrus.Entry
	filePath  string
	level     string
	formatter string
	file      *os.File
}

// New() creates and configures a new instance of the logger.
func New(logFilePath, logLevel, logFormatter string, logFields Fields) (*Logger, error) {
	// Create a new instance of the Logger
	logger := &Logger{
		logger:    &logrus.Entry{},
		filePath:  logFilePath,
		level:     logLevel,
		formatter: logFormatter,
		file:      nil,
	}

	// If not stated, a "info" logginf level is used, since an http.Server and httputil.ReverseProxy
	// use the "info" level when send messages to a given Writer.
	if logLevel == "" {
		logger.level = "info"
	}

	// Create a new instance of the logrus logger
	lr := logrus.New()

	// Set the logging level
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, fmt.Errorf("logger: new(): unable to set the logging level '%s': %w", logLevel, err)
	}
	lr.SetLevel(level)

	// Set the system logger formatter
	switch strings.ToLower(logFormatter) {
	// If not stated, a json will be used as a default formater
	case "":
		fallthrough
	case "json":
		lr.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		lr.SetFormatter(&logrus.TextFormatter{})
	default:
		return nil, fmt.Errorf("logger: new(): unknown logging formatter: '%s'", logFormatter)
	}

	// Set the os.Stdout or a file for writing the system log messages
	if logFilePath == "" {
		return nil, errors.New("logger: new(): a log file path is empty")
	}

	if strings.ToLower(logFilePath) == "stdout" {
		lr.SetOutput(os.Stdout)
	} else {
		// Open a file for the logger output
		logger.file, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, fmt.Errorf("logger: new(): unable to open log file '%s' for writing: %w", logFilePath, err)
		}
		// Redirect the logger output to the file
		lr.SetOutput(logger.file)
	}

	// Assign an entry of the logrus to the created Logger
	// ! ToDo: replace by function parameter values
	logger.logger = lr.WithFields(logrus.Fields{"type": "system"})

	return logger, nil
}

// Function for calling by http.Server or httputil.ReverseProxy ErrorLog
func (lw *Logger) Write(p []byte) (n int, err error) {
	// Customization of the line to be logged
	output := string(p)
	if !strings.Contains(output, ",success") {
		output = strings.TrimSuffix(output, "\n")
		lw.logger.WithFields(logrus.Fields{"result": "denied"}).Info(output)
	} else {
		output = strings.TrimSuffix(output, ",success")
		lw.logger.WithFields(logrus.Fields{"result": "success"}).Info(output)
	}
	return 1, nil
}

// The LogHTTPRequest() function logs an HTTP request details
func (lw *Logger) LogHTTPRequest(req *http.Request) {
	lw.Infof("%s,%s,%s,%t,%t,%s,success",
		req.RemoteAddr,
		req.TLS.ServerName,
		getTLSVersionName(req.TLS.Version),
		req.TLS.HandshakeComplete,
		req.TLS.DidResume,
		tls.CipherSuiteName(req.TLS.CipherSuite))
}

// func (lw *Logger) Terminate() {
// 	lw.logfile.Close()
// }
