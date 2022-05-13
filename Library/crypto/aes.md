# aes
%%t220415~16%%
[aes package](https://pkg.go.dev/crypto/aes)

```go
import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "fmt"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}

func AesEncrypt(plaintext []byte, key, iv []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    blockSize := block.BlockSize()
    plaintext = PKCS7Padding(plaintext, blockSize)
    blockMode := cipher.NewCBCEncrypter(block, iv)
    crypted := make([]byte, len(plaintext))
    blockMode.CryptBlocks(crypted, plaintext)
    return crypted, nil
}

func AesDecrypt(ciphertext []byte, key, iv []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    blockSize := block.BlockSize()
    blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
    origData := make([]byte, len(ciphertext))
    blockMode.CryptBlocks(origData, ciphertext)
    origData = PKCS7UnPadding(origData)
    return origData, nil
}
```
[^sf]

[^sf]: [golang基础学习-AES加密 - SegmentFault 思否](https://segmentfault.com/a/1190000021267253)