package tools

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func EncodeSha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	s := hex.EncodeToString(sum)
	return string(s)
}

func EncodeMd5(data string) string {
	w := md5.New()
	_, _ = io.WriteString(w, data)
	byData := w.Sum(nil)
	result := fmt.Sprintf("%x", byData)
	return result
}

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
