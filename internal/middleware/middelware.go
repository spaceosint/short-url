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
type middleware struct {
}

type Middleware interface {
	GzipHandle() gin.HandlerFunc
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

func (m middleware) GzipHandle() gin.HandlerFunc {

	return func(c *gin.Context) {
		// проверяем, что клиент поддерживает gzip-сжатие
		if !strings.Contains(c.Request.Header.Get("Accept-Encoding"), "gzip") {
			// если gzip не поддерживается, передаём управление
			// дальше без изменений
			c.Next()

			//next.ServeHTTP(c.Writer, c.Request)
			return
		}

		// создаём gzip.Writer поверх текущего w
		gz, err := gzip.NewWriterLevel(c.Writer, gzip.BestSpeed)
		if err != nil {
			io.WriteString(c.Writer, err.Error())
			return
		}
		defer gz.Close()

		c.Writer.Header().Set("Content-Encoding", "gzip")
		// передаём обработчику страницы переменную типа gzipWriter для вывода данных
		w := &responseBodyWriter{ResponseWriter: c.Writer, Writer: gz}
		c.Writer = w

		c.Next()
		//next.ServeHTTP(gzipWriter{ResponseWriter: c.Writer, Writer: gz}, c.Request)
	}
}

func (m middleware) TestM() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.Request.Header.Get("Accept-Encoding"), "gzip") {
			fmt.Println("True")
			c.Next()
			return
		}
	}
}
