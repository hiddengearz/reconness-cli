package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"
)

// hashTo32Bytes will turn any input into a 32byte hash
func HashTo32Bytes(input string) []byte {

	data := sha256.Sum256([]byte(input))
	return data[0:]
}

//EncryptString will encrypt a string with the hash of the keyString and return a ciphertext
func EncryptString(plainText string, keyString string) (cipherTextString string, err error) {
	if keyString == "" {
		log.Fatal("Key to ddecrypt data can't be blank")
	}

	key := HashTo32Bytes(keyString)
	encrypted, err := encryptAES(key, []byte(plainText))
	if err != nil {
		log.Debug(err)
		return "", err
	}

	return base64.URLEncoding.EncodeToString(encrypted), nil
}

//encryptAES encryptes data using AES256
func encryptAES(key, data []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Debug("Unable to generate cipher")
		return nil, err
	}

	// create two 'windows' in to the output slice.
	output := make([]byte, aes.BlockSize+len(data))
	iv := output[:aes.BlockSize]
	encrypted := output[aes.BlockSize:]

	// populate the IV slice with random data.
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		log.Debug("Unable to populate the IV slice with random data")
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)

	// note that encrypted is still a window in to the output slice
	stream.XORKeyStream(encrypted, data)
	return output, nil
}

//DecryptString will decrypt a string with the hash of the keyString and return the cleartext
func DecryptString(cryptoText string, keyString string) (plainTextString string, err error) {
	if keyString == "" {
		log.Fatal("Key to decrypt data can't be blank")
	}

	encrypted, err := base64.URLEncoding.DecodeString(cryptoText)
	if err != nil {
		log.Debug("Unable to b64decode string")
		return "", err
	}
	if len(encrypted) < aes.BlockSize {
		log.Debug(fmt.Errorf("cipherText too short. It decodes to %v bytes but the minimum length is 16", len(encrypted)))
		return "", fmt.Errorf("Data can't be decrypted")
	}

	decrypted, err := decryptAES(HashTo32Bytes(keyString), encrypted)
	if err != nil {
		log.Debug("Unable to descrypt string")
		return "", err
	}

	return string(decrypted), nil
}

//decryptAES decrypts data using AES256
func decryptAES(key, data []byte) ([]byte, error) {
	// split the input up in to the IV seed and then the actual encrypted data.
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Debug("Unable to generate cipher")
		return nil, err
	}
	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(data, data)
	return data, nil
}
