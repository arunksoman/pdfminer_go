package arcfour

import (
	"crypto/rc4"
	"fmt"
	"log"
)

// https://stackoverflow.com/questions/63578512/how-to-use-crypto-rc4
func ArcFourEncrypt(key string, data string) []byte {
	c, err := rc4.NewCipher([]byte(key))
	if err != nil {
		log.Fatalln(err)
	}
	src := []byte(data)
	fmt.Println("Plaintext: ", src)

	dst := make([]byte, len(src))
	c.XORKeyStream(dst, src)
	fmt.Println("Ciphertext: ", dst)
	return dst
}
