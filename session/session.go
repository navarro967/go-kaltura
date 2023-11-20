package session

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	SESSION_BYTE_ENCODING   = "UTF-8"
	SESSION_KS_PADDING_SIZE = 16
	SESSION_KS_PADDING_CHAR = 0
)

type Session struct {
	Secret      string `json:"secret"`
	UserId      string `json:"userId"`
	PartnerId   int    `json:"partnerId"`
	Expiry      int    `json:"expiry"`
	Privileges  string `json:"privileges"`
	Type        int    `json:"type"`
	SessionType int    `json:"sessionType"`

	ImpersonatedPartnerId int `json:"impersonatedPartnerId"`

	Ks string `json:"ks"`
}

func NewSession(secret string, userId string, _type int, partnerId int, expiry int, privileges string) *Session {
	if expiry == 0 {
		expiry = 86400
	}
	return &Session{
		Secret:     secret,
		PartnerId:  partnerId,
		Expiry:     expiry,
		UserId:     userId,
		Privileges: privileges,
		Type:       _type,
	}
}

func (s *Session) GenerateSession() *Session {
	curTime := time.Now()
	expiry := float64(curTime.UnixNano()+int64(s.Expiry)) / float64(time.Second)
	info := fmt.Sprintf("%d;%d;%d;%d;%s;%s;%s;;",
		s.PartnerId,
		s.PartnerId,
		curTime.Unix()+int64(s.Expiry),
		s.Type,
		strconv.FormatFloat(expiry, 'f', 4, 64),
		s.UserId,
		s.Privileges,
	)
	sig := sha1.Sum([]byte(s.Secret + info))
	s.Ks = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%x|%s", sig, info)))

	return s
}

func (s *Session) GenerateSessionV2() *Session {
	fields := map[string]string{}

	for _, privilige := range strings.Split(s.Privileges, ",") {
		privilige = strings.TrimSpace(privilige)
		if privilige == "" {
			continue
		}
		if privilige = strings.ToLower(privilige); privilige == "*" {
			privilige = "all:*"
		}
		if strings.Contains(privilige, ":") {
			split := strings.SplitN(privilige, ":", 2)
			fields[split[0]] = split[1]
		} else {
			fields[privilige] = ""
		}
	}

	fields["_e"] = strconv.FormatInt(time.Now().Unix()+int64(s.Expiry), 10)
	fields["_t"] = strconv.Itoa(s.Type)
	fields["_u"] = s.UserId

	values := url.Values{}
	for key, value := range fields {
		values.Add(key, value)
	}

	//query string
	var queryString string = values.Encode()

	//prepend 16 random binary bites to the query string
	b := make([]byte, SESSION_KS_PADDING_SIZE)
	padding, _ := rand.Read(b)
	buffer := append(b[:padding], []byte(queryString)...)

	//prepend the binary sha1 hash of the string to the string
	fieldHash := sha1.Sum(buffer)
	buffer = append(fieldHash[:], buffer...)

	//hash the account’s API secret
	hashKey := sha1.Sum([]byte(s.Secret))

	//Encrypt the string with the SHA1 hash of the account’s API secret using AES128/CBC/Zero bytes padding.
	block, _ := aes.NewCipher(hashKey[:16])
	if fieldStrLen := len(buffer); fieldStrLen%block.BlockSize() != 0 {
		paddingSize := block.BlockSize() - fieldStrLen%block.BlockSize()
		buffer = append(buffer, bytes.Repeat([]byte{SESSION_KS_PADDING_CHAR}, paddingSize)...)
	}
	ciphertext := make([]byte, len(buffer))
	cipher := cipher.NewCBCEncrypter(block, make([]byte, block.BlockSize()-len(buffer)%block.BlockSize()))
	cipher.CryptBlocks(ciphertext, []byte(buffer))

	//Build the KS string in the following format: v2|{partnerId}|{encrypted string}
	header := fmt.Sprintf("v2|%d|", s.PartnerId)
	rawKs := append([]byte(header), ciphertext...)

	//Base64 encode the encrypted string
	s.Ks = base64.URLEncoding.EncodeToString([]byte(rawKs))

	return s
}
