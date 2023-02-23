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
	// w.Writer будет отвечать за gzip-сжатие, поэтому пишем в него
	return w.Writer.Write(b)
}

type responseBodyWriter struct {
	gin.ResponseWriter
	Writer io.Writer
}

func GzipHandle() gin.HandlerFunc {

	return func(c *gin.Context) {
		// проверяем, что клиент поддерживает gzip-сжатие
		if !strings.Contains(c.Request.Header.Get("Content-Encoding"), "gzip") {
			// если gzip не поддерживается, передаём управление
			// дальше без изменений
			c.Next()

			//next.ServeHTTP(c.Writer, c.Request)
			return
		}

		gz, err := gzip.NewReader(c.Request.Body)
		fmt.Println(err)

		defer c.Request.Body.Close()

		c.Request.Body = gz

		c.Next()
		//next.ServeHTTP(gzipWriter{ResponseWriter: c.Writer, Writer: gz}, c.Request)
	}
}
