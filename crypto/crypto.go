package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"github.com/tjfoc/gmsm/sm4"
)

func AESCTRDecrypt(cipherText string, keySpec, iv []byte) (decryptMsg string, err error) {
	if cipherText == "" {
		err = errors.New("cipherText is empty")
		return
	}
	decodedCipherText, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return
	}
	block, err := aes.NewCipher(keySpec)
	if err != nil {
		return
	}
	if len(cipherText) < aes.BlockSize {
		err = errors.New("cipherText block is too short")
		return
	}
	// ctr decrypt
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(decodedCipherText, decodedCipherText)
	return string(decodedCipherText), nil
}

func AESCBCDecrypt(cipherText, key, iv []byte) (plaintext []byte, err error) {
	var block cipher.Block
	if block, err = aes.NewCipher(key); err != nil {
		return
	}
	if len(cipherText) < aes.BlockSize {
		err = errors.New("cipherText block is too short")
		return
	}
	cbc := cipher.NewCBCDecrypter(block, iv)
	plaintext = make([]byte, len(cipherText))
	cbc.CryptBlocks(plaintext, cipherText)
	return pkcs5UnPadding(plaintext), nil
}

//sm4 ecb decrypt

type ecbDecrypter ecb

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func SM4Decrypt(cipherText string, key []byte) (plaintext string, err error) {
	decodedCipherText, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return
	}
	block, err := sm4.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := NewECBDecrypter(block)
	origData := make([]byte, len(decodedCipherText))
	blockMode.CryptBlocks(origData, decodedCipherText)
	return string(pkcs5UnPadding(origData)), nil
}

func pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	if length == 0 {
		return nil
	}
	unPadding := int(src[length-1])
	return src[:(length - unPadding)]
}
