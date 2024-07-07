package api

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hsiang086/intel-fest/database"
	_ "github.com/joho/godotenv/autoload"
)

var (
	commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	secret   = os.Getenv("SECRET") // len 32 example: "TWTom13LovesPythonBecauseHeSucks"
)

func encrypt(text any) (string, int) {
	var plaintext []byte
	switch v := text.(type) {
	case int:
		plaintext = []byte(strconv.Itoa(v))
	case string:
		plaintext = []byte(v)
	}
	key := secret
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
		os.Exit(1)
	}
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return fmt.Sprintf("%x", ciphertext), len(plaintext)
}

func decrypt(encrypted string, plainlen string) string {
	key := secret
	plainlenInt, _ := strconv.Atoi(plainlen)
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
		os.Exit(1)
	}
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	ciphertext, _ := hex.DecodeString(encrypted)
	plaintextCopy := make([]byte, plainlenInt)
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	return string(plaintextCopy)
}

func IsCookieValid(__yumm__ string) (bool, string) {
	s := strings.Split(__yumm__, ".")
	encryptEmail, plainEmailLen, encryptId, plainIdLen := s[0], s[1], s[2], s[3]
	email := decrypt(encryptEmail, plainEmailLen)
	idFromCookie, _ := strconv.Atoi(decrypt(encryptId, plainIdLen))
	if exist, idFromEmail := database.IsUserExist(email); exist && idFromEmail == idFromCookie {
		fmt.Printf("Email: %s, ID: %d\n", email, idFromCookie)
		return true, email
	}
	return false, ""
}

func setCookie(c *gin.Context, email string, id int) {
	encryptEmail, plainEmailLen := encrypt(email)
	encryptId, plainIdLen := encrypt(id)
	cookie := fmt.Sprintf("%s.%d.%s.%d", encryptEmail, plainEmailLen, encryptId, plainIdLen)
	c.SetCookie("__yumm__", cookie, 3600, "/", "http://127.0.0.1", false, true)
}
