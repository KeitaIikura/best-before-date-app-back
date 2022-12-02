package crypter

import (
	"bbdate/pkg/logging"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// 参考）
// https://github.com/golang/crypto
// https://gist.github.com/mickelsonm/e1bf365a149f3fe59119
// https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/09.6.html

type Crypter struct {
	cipherKeyText string
}

func NewCrypter(cipherKeyText string) ICrypter {
	return &Crypter{
		cipherKeyText: cipherKeyText,
	}
}

// 暗号化
func (c Crypter) Encrypt(s string) string {
	source := []byte(s)

	block, err := aes.NewCipher([]byte(c.cipherKeyText))
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(source))
	iv := ciphertext[:aes.BlockSize]

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], source)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

func (c Crypter) Decrypt(s string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(s)

	block, err := aes.NewCipher([]byte(c.cipherKeyText))
	if err != nil {
		panic(err)
	}

	// サイズ不足の場合はエラー
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

func (c Crypter) EncryptString(source string) (string, error) {
	return c.Encrypt(source), nil
}

func (c Crypter) EncryptData(data interface{}) ([]byte, error) {
	m, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	es, err := c.EncryptString(string(m))
	if err != nil {
		return nil, err
	}
	return []byte(es), nil
}

func (c Crypter) DecryptString(source string) (string, error) {
	return c.Decrypt(source), nil
}

func (c Crypter) DecryptData(data []byte, output interface{}) error {
	ds, err := c.DecryptString(string(data))
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(ds), output); err != nil {
		return err
	}
	return nil
}

// 参考）https://gobyexample.com/sha1-hashes

func (c Crypter) GenerateSha512Hash(str ...string) string {
	var hashStr string
	h := sha512.New()
	hashStr = strings.Join(str, "")
	_, err := h.Write([]byte(hashStr))
	if err != nil {
		logging.Info("system", fmt.Sprintf("hash write error: %v", err))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (c Crypter) GenerateSha256Hash(str ...string) string {
	var hashStr string
	h := sha256.New()
	hashStr = strings.Join(str, "")
	_, err := h.Write([]byte(hashStr))
	if err != nil {
		logging.Info("system", fmt.Sprintf("hash write error: %v", err))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
