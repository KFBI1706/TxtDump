package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"

	"github.com/elithrar/simple-scrypt"
	_scrypt "golang.org/x/crypto/scrypt"
)

//Probably not the best way to do this
func genFromSeed() int {
	num := rand.Intn(9999999-1000000) + 1000000
	if checkForDuplicateID(num) {
		return num
	}
	num = rand.Intn(9999999-1000000) + 1000000
	log.Println(num)
	return num
}
func findpostfortest() (int, error) {
	var post int
	db, err := establishConn()
	err = db.QueryRow("SELECT ID FROM TEXT LIMIT 1;").Scan(&post)
	if err != nil {
		return 0, err
	}
	db.Close()
	return post, err
}
func setupDB() error {
	db, err := establishConn()
	if err != nil {
		return err
	}
	sql, err := readDBstring("sql/db.sql")
	if err != nil {
		return err
	}
	res, err := db.Exec(sql)
	if err != nil {
		return err
	}
	fmt.Println(res)
	db.Close()
	return nil
}
func sha256encode(b []byte) string {
	h := sha256.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func securePass(ps string) (string, string, error) {
	hash, err := scrypt.GenerateFromPassword([]byte(ps), scrypt.DefaultParams)
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}
	vals := strings.Split(string(hash), "$")
	salt := vals[3]
	sha256hash := sha256encode(hexToBytes(vals[4]))
	fmt.Println(vals[4])
	return salt, sha256hash, nil
}

func checkPass(ps string, id int, perms int) bool {
	dk, err := _scrypt.Key([]byte(ps), getSalt(id), 16384, 8, 1, scrypt.DefaultParams.DKLen)

	if err != nil {
		log.Fatal(err)
	}
	if subtle.ConstantTimeCompare(hexToBytes(sha256encode(dk)), getHashedPS(id)) == 1 || perms == 1 {
		return true
	}
	return false
}

func NewEncryptionKey() *[]byte {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	return &key
}

func Encrypt(plaintext []byte, key *[32]byte) (ciphertext []byte, err error) {
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

func Decrypt(ciphertext []byte, key *[32]byte) (plaintext []byte, err error) {
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

func clearOutDB() error {
	db, err := establishConn()
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE text")
	if err != nil {
		return err
	}
	return err
}
func determinePerms(postperm string) (int, error) {
	num, err := strconv.Atoi(postperm)
	if err != nil {
		return 0, err
	}
	return num, err
}
