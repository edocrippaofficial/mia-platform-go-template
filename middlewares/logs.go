package middlewares

import (
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const (
	forwardedHostHeaderKey = "x-forwarded-host"
	forwardedForHeaderKey  = "x-forwarded-for"
	requestIDHeaderName    = "x-request-id"

	IncomingRequestMessage  = "incoming request"
	RequestCompletedMessage = "request completed"
)

type HTTP struct {
	Request  *Request  `json:"request,omitempty"`
	Response *Response `json:"response,omitempty"`
}

type Request struct {
	Method    string    `json:"method"`
	UserAgent UserAgent `json:"user_agent"`
}

type UserAgent struct {
	Original string `json:"original"`
}

type Response struct {
	StatusCode int          `json:"status_code"`
	Body       ResponseBody `json:"body"`
}

type ResponseBody struct {
	Bytes int64 `json:"bytes"`
}

type URL struct {
	Path string `json:"path"`
}

type Host struct {
	ForwardedHost string `json:"forwarded_host,omitempty"`
	Hostname      string `json:"hostname"`
	IP            string `json:"ip,omitempty"`
}

type RequestLogger struct {
	logger  *logrus.Logger
	context echo.Context
}

func (rl RequestLogger) logIncomingRequest() {
	req := rl.context.Request()

	rl.logger.WithFields(logrus.Fields{
		"http": HTTP{
			Request: &Request{
				Method: req.Method,
				UserAgent: UserAgent{
					Original: req.Header.Get("User-Agent"),
				},
			},
		},
		"url": URL{Path: req.RequestURI},
		"host": Host{
			ForwardedHost: req.Header.Get(forwardedHostHeaderKey),
			Hostname:      removePort(req.Host),
			IP:            req.Header.Get(forwardedForHeaderKey),
		},
	}).Trace(IncomingRequestMessage)
}

func (rl RequestLogger) logRequestCompleted(start time.Time) {
	req := rl.context.Request()
	res := rl.context.Response()

	rl.logger.WithFields(logrus.Fields{
		"http": HTTP{
			Request: &Request{
				Method: req.Method,
				UserAgent: UserAgent{
					Original: req.Header.Get("User-Agent"),
				},
			},
			Response: &Response{
				StatusCode: res.Status,
				Body: ResponseBody{
					Bytes: res.Size,
				},
			},
		},
		"url": URL{Path: req.RequestURI},
		"host": Host{
			ForwardedHost: req.Header.Get(forwardedHostHeaderKey),
			Hostname:      removePort(req.Host),
			IP:            req.Header.Get(forwardedForHeaderKey),
		},
		"responseTime": float64(time.Since(start).Milliseconds()),
	}).Info(RequestCompletedMessage)
}

func RequestMiddlewareLogger(logger *logrus.Logger, excludedPrefix []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for _, prefix := range excludedPrefix {
				if strings.HasPrefix(c.Request().RequestURI, prefix) {
					return next(c)
				}
			}

			requestLogger := RequestLogger{
				logger:  logger,
				context: c,
			}
			start := time.Now()

			requestLogger.logIncomingRequest()
			err := next(c)
			requestLogger.logRequestCompleted(start)

			return err
		}
	}
}

func removePort(host string) string {
	if idx := strings.LastIndex(host, ":"); idx != -1 {
		return host[:idx]
	}
	return host
}
