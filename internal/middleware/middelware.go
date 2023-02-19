package middleware

import (
	"compress/gzip"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

type Middleware interface {
	GzipHandle() http.Header
	TestM() gin.HandlerFunc
}

func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func GzipHandle() gin.HandlerFunc {

	return func(c *gin.Context) {
		// проверяем, что клиент поддерживает gzip-сжатие
		if !strings.Contains(c.Request.Header.Get("Content-Encoding"), "gzip") {
			// если gzip не поддерживается, передаём управление
			// дальше без изменений
			c.Next()
			return
		}

		gz, err := gzip.NewReader(c.Request.Body)
		fmt.Println(err)

		defer c.Request.Body.Close()

		c.Request.Body = gz

		c.Next()

	}
}
