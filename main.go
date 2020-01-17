package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	idFlagName         = "id"
	secretFlagName     = "secret"
	secretFileFlagName = "secret-file"
)

const (
	appName = "hmac-generator"
	version = "1.0.0"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = `Generates MAC authorization`
	app.Version = version
	app.UsageText = fmt.Sprintf(`%s --id <key id> --secret/--secret-file <value>`, appName)

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  fmt.Sprintf("%s", idFlagName),
			Usage: fmt.Sprintf("Authorization key id"),
		},
		cli.StringFlag{
			Name:  fmt.Sprintf("%s, s", secretFileFlagName),
			Usage: fmt.Sprintf("Authorization secret"),
		},
		cli.StringFlag{
			Name:  fmt.Sprintf("%s, sf", secretFileFlagName),
			Usage: "File from which to read authorization secret",
		},
	}
	app.Action = action

	app.Run(os.Args)
}

func action(c *cli.Context) error {
	id := c.String(idFlagName)
	if id == "" {
		return fmt.Errorf("id should be provided to identify the client")
	}
	secret := c.String(secretFlagName)
	generateMACAuthorization(id, secret)
	return nil
}

func generateMACAuthorization(id string, secret string) {
	rand.Seed(time.Now().UnixNano())
	nonce := rand.Int()
	mac := hmac.New(sha256.New, []byte(secret))
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	mac.Write([]byte(strconv.Itoa(int(timestamp)) + strconv.Itoa(nonce)))
	sum := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	fmt.Printf("Authorization: MAC ts=%d,id=%s,nonce=%d,mac=%s\n", timestamp, id, nonce, sum)
}
