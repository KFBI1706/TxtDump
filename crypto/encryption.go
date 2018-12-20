package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"log"
	"math/rand"

	"github.com/KFBI1706/TxtDump/model"
	"github.com/KFBI1706/TxtDump/sql"
	scrypt "github.com/elithrar/simple-scrypt"
	_scrypt "golang.org/x/crypto/scrypt"
)

/*RequestDecrypt decrypts the post if the field post.Crypto.Hash is the correct password.
It takes a pointer to the model.Post as the only argument
and returns a bool based on if the post is successfully decrypted*/
func RequestDecrypt(post *model.Post) bool {
	if CheckPass(post.Crypto.Hash, post.ID, post.Data.PostPerms) {
		key := GetEncKey(post)
		ct, _ := base64.StdEncoding.DecodeString(post.Data.Content)
		pt, err := decrypt(ct, &key)
		if err != nil {
			log.Fatal(err)
		}
		post.Data.Content = string(pt)
		return true
	}
	return false
}

/*GetEncKey gets the encryption key used for the file by decrypting the stored-key with the passord scrypt-hash
takes a pointer to model.Post as the only argument
 and returns a 32 length byte array*/
func GetEncKey(post *model.Post) (key [32]byte) {
	dk := getKey(post)
	key = [32]byte{}
	copy(key[:], dk[0:32])
	tmp, _ := base64.StdEncoding.DecodeString(post.Crypto.Key)
	encKey, _ := decrypt(tmp, &key)
	copy(key[:], encKey[0:32])
	return
}

/*GetToken generates a random token used to prefent CSRF attacks*/
func GetToken() string {
	data := make([]byte, 64)
	rand.Read(data)
	return base64.StdEncoding.EncodeToString(data)
}

func getKey(post *model.Post) (dk []byte) {
	prop, err := sql.GetProp("salt", post.ID)
	if err != nil {
		log.Println(err)
	}
	dk, _ = _scrypt.Key([]byte(post.Crypto.Hash), prop, 16384, 8, 1, scrypt.DefaultParams.DKLen)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func encryptPost(content []byte, key *[32]byte) (string, string) {
	ct, _ := encrypt(content, key)
	encodedContent := base64.StdEncoding.EncodeToString(ct)
	tmp := make([]byte, 32)
	copy(tmp, key[:])
	encodedKey := base64.StdEncoding.EncodeToString(tmp)
	return encodedContent, encodedKey
}

/*SecurePost generates a secure post for a unencrypted post. encryption varies based upon PostPerms
takes a pointer to model.Post and a unencrypted password-string as arguments
returns nothing*/
func SecurePost(post *model.Post, pass string) {
	if post.Data.PostPerms > 1 {
		if salt, hash, sha256hash, err := securePass(pass); sha256hash != "" {
			post.Crypto.Salt = salt
			post.Crypto.Hash = sha256hash
			if post.Data.PostPerms == 3 {
				encKey := newencryptionKey()
				post.Data.Content, post.Crypto.Key = encryptPost([]byte(post.Data.Content), encKey)
				tmpKey := sql.HexToBytes(hash)
				key := [32]byte{}
				copy(key[:], tmpKey[0:32])
				tmpKey, _ = base64.StdEncoding.DecodeString(post.Crypto.Key) //same as encKey
				tmpKey, _ = encrypt(tmpKey, &key)                            //encrypt the file key with the password hash
				post.Crypto.Key = base64.StdEncoding.EncodeToString(tmpKey)

			}
		} else {
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func newencryptionKey() *[32]byte {
	rnd := make([]byte, 32)
	rand.Read(rnd)
	key := [32]byte{}
	copy(key[:], rnd[0:32])
	return &key
}

/*EncryptBytes is used to encrypt bytes with a key, pretty self-explanatory..
takes a byte slice b and a pointer to a 32 byte array key
returns a byte slice (ct) and a error (err)*/
func EncryptBytes(b []byte, key *[32]byte) (ct []byte, err error) {
	ct, err = encrypt(b, key)
	if err != nil {
		log.Fatal(err)
	}
	return
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
