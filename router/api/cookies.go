package api

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	secret   string // len 32 example: "TWTom13LovesPythonBecauseHeSucks"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secret = os.Getenv("SECRET")
}

func encrypt(id int) string {
	idByte := []byte(strconv.Itoa(id))
	key := secret
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
		os.Exit(1)
	}
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(idByte))
	cfb.XORKeyStream(ciphertext, idByte)
	return fmt.Sprintf("%x", ciphertext)
}

func isCookieValid(id int, __yumm__ string) bool {
	return __yumm__ == encrypt(id)
}

func setCookie(c *gin.Context, id int) {
	c.SetCookie("__yumm__", encrypt(id), 3600, "/", "http://127.0.0.1", false, true)
}
