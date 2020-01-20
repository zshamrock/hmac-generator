package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateMACAuthorization(id string, secret string) string {
	rand.Seed(time.Now().UnixNano())
	nonce := rand.Int()
	mac := hmac.New(sha256.New, []byte(secret))
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	mac.Write([]byte(strconv.Itoa(int(timestamp)) + strconv.Itoa(nonce)))
	sum := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return fmt.Sprintf("Authorization: HMAC ts=%d,id=%s,nonce=%d,mac=%s", timestamp, id, nonce, sum)
}
