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
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	lr        *logrus.Entry
	level     string
	logFile   *os.File
	filePath  string
	formatter string
}

// New() creates and configures a new instance of the logger.
func New(logFilePath, logLevel, logFormatter string, logFields Fields) (*Logger, error) {
	// Create a new instance of the Logger
	logger := &Logger{
		lr:        &logrus.Entry{},
		filePath:  logFilePath,
		level:     logLevel,
		formatter: logFormatter,
		logFile:   nil,
	}

	// If not stated, a "info" logging level is used, since an http.Server and httputil.ReverseProxy
	// use the "info" level when send messages to a given Writer.
	if logLevel == "" {
		logger.level = "info"
	}

	// Create a new instance of the logrus logger
	lr := logrus.New()

	// Set the logging level
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, fmt.Errorf("logger: New(): unable to set the logging level '%s': %s", logLevel, err.Error())
	}
	lr.SetLevel(level)

	// Set the logger formatter
	switch strings.ToLower(logFormatter) {
	// If not stated, a json will be used as a default formater
	case "":
		fallthrough
	case "json":
		lr.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		lr.SetFormatter(&logrus.TextFormatter{})
	default:
		return nil, fmt.Errorf("logger: New(): unknown logging formatter '%s'", logFormatter)
	}

	// Set the os.Stdout or a file for writing the log messages
	if len(logFilePath) == 0 || strings.ToLower(logFilePath) == "stdout" {
		lr.SetOutput(os.Stdout)
	} else {
		// Open a file for the logger output
		logger.logFile, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, fmt.Errorf("logger: New(): unable to open the file '%s' for writing: %s", logFilePath, err.Error())
		}
		// Redirect the logger output to the file
		lr.SetOutput(logger.logFile)
	}

	// Assign logger logFields to logrus Fields variable
	fields := logrus.Fields(logFields)

	// Assign an entry of the logrus to the created Logger
	logger.lr = lr.WithFields(fields)

	return logger, nil
}

// Function for calling by http.Server or httputil.ReverseProxy ErrorLog
func (logger *Logger) Write(p []byte) (n int, err error) {
	// Customization of the line to be logged
	output := string(p)
	if !strings.Contains(output, ",success") {
		output = strings.TrimSuffix(output, "\n")
		logger.lr.WithFields(logrus.Fields{"result": "denied"}).Info(output)
	} else {
		output = strings.TrimSuffix(output, ",success")
		logger.lr.WithFields(logrus.Fields{"result": "success"}).Info(output)
	}
	return 1, nil
}

// The LogHTTPRequest() function logs an HTTP request details
func (logger *Logger) LogHTTPRequest(req *http.Request) {
	logger.Infof("%s,%s,%s,%t,%t,%s,success",
		req.RemoteAddr,
		req.TLS.ServerName,
		getTLSVersionName(req.TLS.Version),
		req.TLS.HandshakeComplete,
		req.TLS.DidResume,
		tls.CipherSuiteName(req.TLS.CipherSuite))
}

// Terminate() gracefully shuts the logger down
func (logger *Logger) Terminate() {
	// Close the log file if it was open
	if logger.logFile != nil {
		logger.logFile.Close()
	}
}
