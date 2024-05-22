package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gojwt/helper"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func makeLogEntry(c *gin.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	var payload map[string]interface{}
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err == nil {
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	json.Unmarshal(bodyBytes, &payload)

	cwd, _ := os.Getwd()
	logLocation := filepath.Join(cwd + "/log/api.log")
	fmt.Println(logLocation)
	logFile, err := os.OpenFile(logLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	helper.IsError(err)

	log.SetOutput(io.MultiWriter(os.Stderr, logFile))

	log.RegisterExitHandler(func() {
		if logFile == nil {
			return
		}
		logFile.Close()
	})

	return log.WithFields(log.Fields{
		"at":      time.Now().Format("2006-01-02 15:04:05"),
		"method":  c.Request.Method,
		"uri":     c.Request.RequestURI,
		"ip":      c.Request.RemoteAddr,
		"payload": payload,
	})
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		makeLogEntry(c).Info()
		c.Next()
	}
}
