package middleware

import (
	"compress/gzip"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Middleware interface {
	GzipHandle() http.Header
	TestM() gin.HandlerFunc
}

func (w gzipWriter) Write(data []byte) (int, error) {
	return w.gw.Write(data)
}

type gzipWriter struct {
	gw *gzip.Writer
	gin.ResponseWriter
}

func GzipWriterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !strings.Contains(c.Request.Header.Get("Accept-Encoding"), "gzip") {
			// если gzip не поддерживается, передаём управление
			// дальше без изменений
			c.Next()
			return
		}

		// Создание gzip.Writer с настройкой BestSpeed
		gz, err := gzip.NewWriterLevel(c.Writer, gzip.BestSpeed)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer gz.Close()

		c.Writer.Header().Set("Content-Encoding", "gzip")
		// Создание нового gzipWriter
		gw := gzipWriter{gw: gz, ResponseWriter: c.Writer}

		// Замена c.Writer на gzipWriter для записи сжатых данных
		c.Writer = &gw

		// Вызов следующего обработчика в цепочке
		c.Next()
	}
}

func GzipReaderHandle() gin.HandlerFunc {

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
