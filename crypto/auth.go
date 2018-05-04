package crypto

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"log"
	"strings"

	"github.com/KFBI1706/Txtdump/sql"
	scrypt "github.com/elithrar/simple-scrypt"
	_scrypt "golang.org/x/crypto/scrypt"
)

func sha256encode(b []byte) string {
	h := sha256.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func securePass(ps string) (string, string, string, error) {
	sscrypt, err := scrypt.GenerateFromPassword([]byte(ps), scrypt.DefaultParams)
	if err != nil {
		log.Fatal(err)
		return "", "", "", err
	}
	vals := strings.Split(string(sscrypt), "$")
	salt := vals[3]
	sha256hash := sha256encode(sql.HexToBytes(vals[4]))
	hash := vals[4]
	return salt, hash, sha256hash, nil
}

/*CheckPass takes argument ps(string),id(int), perms(int). Where:
ps is the password in string form, id is the id of the post and perms is the post permissions
it returns a bool based on it's success*/
func CheckPass(ps string, id int, perms int) bool {
	prop, err := sql.GetProp("salt", id)
	if err != nil {
		log.Println(err)
	}
	dk, err := _scrypt.Key([]byte(ps), prop, 16384, 8, 1, scrypt.DefaultParams.DKLen)
	if err != nil {
		log.Fatal(err)
	}
	prop, err = sql.GetProp("hash", id)
	if err != nil {
		log.Println(err)
	}
	if subtle.ConstantTimeCompare(sql.HexToBytes(sha256encode(dk)), prop) == 1 || perms == 1 {
		return true
	}
	return false
}
