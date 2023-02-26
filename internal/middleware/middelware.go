package middleware

import (
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func generateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

const cookieName = "myCookie"

//const cookieKey = "mySecretKey12345678901234567890123456789"

var key, _ = generateRandom(2 * aes.BlockSize) // ключ шифрования

//cookieKey, _ = generateRandom(2 * aes.BlockSize) // ключ шифрования

// authenticate middleware checks if the user is authenticated or not
// and issues a signed cookie with a unique user ID if not
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie(cookieName)

		if err != nil || cookie.Value == "" {
			// Generate a unique ID for the user
			userID := generateUserID()

			// Encrypt the ID and sign the cookie
			encryptedCookie, err := encryptAndSignCookie([]byte(userID), key)
			if err != nil {
				fmt.Println("Failed to encrypt and sign cookie:", err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			// Set the cookie with a 1-hour expiration time
			c.SetCookie(cookieName, encryptedCookie, 3600, "/", "", false, true)

			// Set the user ID in the request context for future use
			c.Set("userID", userID)

			// Call the next middleware
			c.Next()
			return
		}

		// If the cookie exists, decrypt and verify its signature
		userID, err := decryptAndVerifyCookie(cookie.Value, key)
		if err != nil {
			userID := generateUserID()

			// Encrypt the ID and sign the cookie
			encryptedCookie, err := encryptAndSignCookie([]byte(userID), key)
			if err != nil {
				fmt.Println("Failed to encrypt and sign cookie:", err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			// Set the cookie with a 1-hour expiration time
			c.SetCookie(cookieName, encryptedCookie, 3600, "/", "", false, true)

			// Set the user ID in the request context for future use
			c.Set("userID", userID)

			// Call the next middleware
			c.Next()
		}

		// Set the user ID in the request context for future use
		c.Set("userID", userID)

		// Call the next middleware
		c.Next()
	}
}

func generateUserID() string {
	// generate a new UUID
	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Println("Failed to generate UUID:", err)
		return ""
	}
	// convert the UUID to a string and return it
	return id.String()
}

func encryptAndSignCookie(data []byte, key []byte) (string, error) {
	// generate a new AES cipher block using the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// generate a new random initialization vector
	iv := make([]byte, aes.BlockSize)
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return "", err
	}

	// encrypt the data using the AES cipher block and the random IV
	ciphertext := make([]byte, len(data))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, data)

	// append the IV to the encrypted data
	encryptedData := append(iv, ciphertext...)

	// generate a new HMAC using the encrypted data and the provided key
	mac := hmac.New(sha256.New, key)
	mac.Write(encryptedData)
	signature := mac.Sum(nil)

	// encode the encrypted data and the signature as a hex string and return it
	return hex.EncodeToString(encryptedData) + hex.EncodeToString(signature), nil
}
func decryptAndVerifyCookie(cookieValue string, key []byte) (string, error) {
	// decode the cookie value from a hex string to a byte slice
	decoded, err := hex.DecodeString(cookieValue)
	if err != nil {
		return "", err
	}

	// split the decoded byte slice into the encrypted data and the signature
	encryptedData := decoded[:len(decoded)-sha256.Size]
	signature := decoded[len(decoded)-sha256.Size:]

	// generate a new HMAC using the encrypted data and the provided key
	mac := hmac.New(sha256.New, key)
	mac.Write(encryptedData)
	expectedSignature := mac.Sum(nil)

	// compare the expected and actual signatures to verify the cookie's authenticity
	if !hmac.Equal(signature, expectedSignature) {
		return "", fmt.Errorf("cookie signature mismatch")
	}

	// extract the initialization vector and the ciphertext from the encrypted data
	iv := encryptedData[:aes.BlockSize]
	ciphertext := encryptedData[aes.BlockSize:]

	// generate a new AES cipher block using the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// decrypt the ciphertext using the AES cipher block and the initialization vector
	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	// convert the plaintext byte slice to a string and return it
	return string(plaintext), nil
}
