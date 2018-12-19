package logging

import (
	"bufio"
	"encoding/json"
	"github.com/fatih/color"
	"github.com/gorilla/handlers"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

var (
	service, packageName, parentName, gopath, templatePath string
	out                                                    = colorable.NewColorableStdout()
	logger                                                 = logrus.New()
)

func NewHandlerLogger() handlers.LogFormatter {
	// Setup logrus
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
		},
	})
	logger.SetLevel(logrus.InfoLevel)

	return func(writer io.Writer, params handlers.LogFormatterParams) {

		host, _, err := net.SplitHostPort(params.Request.RemoteAddr)
		if err != nil {
			host = params.Request.RemoteAddr
		}

		uri := params.Request.RequestURI

		// Requests using the CONNECT method over HTTP/2.0 must use
		// the authority field (aka r.Host) to identify the target.
		// Refer: https://httpwg.github.io/specs/rfc7540.html#CONNECT
		if params.Request.ProtoMajor == 2 && params.Request.Method == "CONNECT" {
			uri = params.Request.Host
		}
		if uri == "" {
			uri = params.URL.RequestURI()
		}

		duration := int64(time.Now().Sub(params.TimeStamp) / time.Millisecond)

		fields := logrus.Fields{
			"host":       host,
			"url":        uri,
			"duration":   duration,
			"status":     params.StatusCode,
			"method":     params.Request.Method,
			"request":    params.Request.RequestURI,
			"remote":     params.Request.RemoteAddr,
			"size":       params.Size,
			"referer":    params.Request.Referer(),
			"user_agent": params.Request.UserAgent(),
			"request_id": params.Request.Header.Get("x-request-id"),
		}

		if headers, err := json.Marshal(params.Request.Header); err == nil {
			fields["headers"] = string(headers)
		} else {
			fields["header_error"] = err.Error()
		}

		logger.WithFields(fields).WithTime(params.TimeStamp).Info(out, "%s %s %d", color.GreenString(params.Request.Method),
			color.YellowString(uri), color.RedString(string(params.StatusCode)))
	}
}

func IfErr(msg string, err error) {
	if err != nil {
		logrus.Fatal(out,
			color.RedString(msg),
			logger.WithError(err),
		)
	}
}

func IfNoErr(msg string, err error) {
	if err == nil {
		logger.Fatal(out, "%s: %s \n",
			color.GreenString(msg),
		)
	}
}

func OK() bool {
	logger.Infoln(out, "Is this OK? %ses/%so\n",
		color.GreenString("[y]"),
		color.RedString("[n]"),
	)
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return strings.Contains(strings.ToLower(scan.Text()), "y")
}

func Exit(msg string) {
	logger.Warn(out, "Error:", color.YellowString(msg))
	os.Exit(1)
}
