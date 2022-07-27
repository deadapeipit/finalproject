package server

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"finalproject/entity"
	"fmt"
	"io"
	"net/http"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var LogonUser *entity.User

type response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

const (
	Success              int = 200
	Success201           int = 201
	ErrorBadRequest      int = 400
	ErrorUnauthorized    int = 401
	ErrorForbidden       int = 403
	ErrorNotFound        int = 404
	ErrorDataHandleError int = 500
)

func EncryptPassword(pwd string) (string, error) {
	// Hashing the password with the default cost of 10
	securePassword, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(securePassword), nil
}

func WriteJsonResp(w http.ResponseWriter, status int, obj interface{}) {
	resp := response{
		Status: status,
		Data:   obj,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}

//EncryptConnectionString
func EncryptConnectionString(data string) string {
	key := []byte(Config.EncryptionKey)
	dataConn := []byte(data)

	ciphertext, err := encryptToken(key, dataConn)
	if err != nil {
		fmt.Printf("Error generating connection string %v \n", err)
		return ""
	}

	return base64.URLEncoding.EncodeToString(ciphertext)
}

//DecryptConnectionString
func DecryptConnectionString(data string) string {
	key := []byte(Config.EncryptionKey)
	decodedByte, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		fmt.Printf("Decode base 64 connection string error: %v accessToken: %s \n", err, data)
		return ""
	}

	result, err := decryptToken(key, decodedByte)
	if err != nil {
		fmt.Printf("Error generating connection string %v \n", err)
		return ""
	}

	return string(result)
}

func encryptToken(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)
	return ciphertext, nil
}

func decryptToken(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return text, nil
}
