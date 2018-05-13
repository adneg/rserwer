package tools

//  narzedzia
import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"
)

func Passhash(pass_text string) (pass string) {
	hash := sha256.New()
	io.WriteString(hash, pass_text)
	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func CheckErr(e error) {

	if e != nil {
		log.Printf(e.Error())
	}

}
