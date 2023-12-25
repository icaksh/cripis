package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/url"
	"time"
)

func FetchDataFromApi(keyword string) (map[string]interface{}, error) {
	a := fiber.AcquireAgent()
	req := a.Request()
	var obj fiber.Map
	req.Header.SetMethod(fiber.MethodPost)
	req.Header.SetContentType(fiber.MIMEApplicationJSON)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Pdki-Signature", PrivateKey())
	req.SetRequestURI("https://pdki-indonesia.dgip.go.id/api/search?")
	a.QueryString(
		"keyword=" + url.QueryEscape(keyword) + "&order_state=asc&page=1&type=trademark")

	if err := a.Parse(); err != nil {
		return obj, err
	}

	code, body, errs := a.Bytes()

	if errs != nil {
		return obj, errs[0]
	}

	if code != 200 {
		return obj, errs[1]
	}

	if err := json.Unmarshal(body, &obj); err != nil {
		return obj, err
	}
	return obj, nil
}

func PrivateKey() string {

	currentTime := time.Now()
	futureTime := currentTime.Add(10 * time.Second)

	encryptedTime := encryptData([]byte(futureTime.String()))
	pdkiSignature := "PDKI/" + encryptedTime
	return pdkiSignature
}

func encryptData(data []byte) string {
	key := []byte("Ym4OlEZmDPSPZBCnRsjLQ2pGLRYyAfqT")
	ivBytes := []byte("1LYCgGZpZksSZuA4")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	data = pkcs7Pad(data, aes.BlockSize)

	ciphertext := make([]byte, len(data))

	mode := cipher.NewCBCEncrypter(block, ivBytes)
	mode.CryptBlocks(ciphertext, data)

	return hex.EncodeToString(ciphertext)
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}
