package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	helper "github.com/KFBI1706/Txtdump/helper"
	model "github.com/KFBI1706/Txtdump/model"

	scrypt "github.com/elithrar/simple-scrypt"
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
	sha256hash := sha256encode(helper.HexToBytes(vals[4]))
	hash := vals[4]
	return salt, hash, sha256hash, nil
}

func checkPass(ps string, id int, perms int) bool {
	prop, err := getProp("salt", id)
	if err != nil {
		log.Println(err)
	}
	dk, err := _scrypt.Key([]byte(ps), prop, 16384, 8, 1, scrypt.DefaultParams.DKLen)

	if err != nil {
		log.Fatal(err)
	}
	prop, err = getProp("hash", id)
	if err != nil {
		log.Println(err)
	}
	if subtle.ConstantTimeCompare(hexToBytes(sha256encode(dk)), prop) == 1 || perms == 1 {
		return true
	}
	return false
}

func RequestDecrypt(post *model.PostData) bool {
	if checkPass(post.Hash, post.ID, post.PostPerms) {
		dk := getKey(post)
		tmp, _ := b64.StdEncoding.DecodeString(post.Key)
		key := [32]byte{}
		copy(key[:], dk[0:32])
		encKey, _ := decrypt(tmp, &key)
		copy(key[:], encKey[0:32])
		ct, _ := b64.StdEncoding.DecodeString(post.Content)
		pt, _ := decrypt(ct, &key)
		post.Content = string(pt)
		return true
	}
	return false
}

func getKey(post *model.PostData) (dk []byte) {
	prop, err := getProp("salt", post.ID)
	if err != nil {
		log.Println(err)
	}
	dk, _ = _scrypt.Key([]byte(post.Hash), prop, 16384, 8, 1, scrypt.DefaultParams.DKLen)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func encryptPost(content []byte, key *[32]byte) (string, string) {
	ct, _ := encrypt(content, key)
	encodedContent := b64.StdEncoding.EncodeToString(ct)
	tmp := make([]byte, 32)
	copy(tmp, key[:])
	encodedKey := b64.StdEncoding.EncodeToString(tmp)
	return encodedContent, encodedKey
}

func securePost(post *model.PostData, pass string) {
	rand.Seed(time.Now().UnixNano())
	post.ID = genFromSeed()
	if post.PostPerms > 1 {
		if salt, hash, sha256hash, err := securePass(pass); sha256hash != "" {
			post.Salt = salt
			post.Hash = sha256hash
			if post.PostPerms == 3 {
				encKey := newEncryptionKey()
				post.Content, post.Key = encryptPost([]byte(post.Content), encKey)
				tmpKey := hexToBytes(hash)
				key := [32]byte{}
				copy(key[:], tmpKey[0:32])
				tmpKey, _ = b64.StdEncoding.DecodeString(post.Key) //same as encKey
				tmpKey, _ = encrypt(tmpKey, &key)                  //encrypt the file key with the password hash
				post.Key = b64.StdEncoding.EncodeToString(tmpKey)

			}
		} else {
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func newEncryptionKey() *[32]byte {
	rnd := make([]byte, 32)
	rand.Read(rnd)
	key := [32]byte{}
	copy(key[:], rnd[0:32])
	//if _, err := rand.Read(key); err != nil {
	//	panic(err)
	//}
	return &key
}

func encrypt(plaintext []byte, key *[32]byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func decrypt(ciphertext []byte, key *[32]byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}
