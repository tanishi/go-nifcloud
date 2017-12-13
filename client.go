package nifcloud

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	SignatureVersion = "2"
	SignatureMethod  = "HmacSHA256"
)

type Client struct {
	URL             *url.URL
	HTTPClient      *http.Client
	AccessKeyID     string
	SecretAccessKey string
}

func NewClient(u, accessKeyID, secretAccessKey string) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(u)

	if err != nil {
		return nil, err
	}

	return &Client{
		URL:             parsedURL,
		HTTPClient:      http.DefaultClient,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
	}, nil
}

func (c *Client) NewRequest(ctx context.Context, method string, u *url.URL, body io.Reader) (*http.Request, error) {
	encodedQuery := encodeQuery(*u)

	t := time.Now().UTC().Format(time.RFC3339)

	reqStr := fmt.Sprintf("SignatureMethod=%s&SignatureVersion=%s&AccessKeyId=%s&%s&%s",
		SignatureMethod, SignatureVersion, c.AccessKeyID, t, encodedQuery)

	sign := fmt.Sprintf("%s\n%s\n%s\n%s", method, u.Host, u.Path, reqStr)

	signature := generateSignature(c.SecretAccessKey, sign)

	q := u.Query()
	q.Set("Signature", signature)
	q.Set("AccessKeyId", c.AccessKeyID)
	q.Set("SignatureMethod", SignatureMethod)
	q.Set("SignatureVersion", SignatureVersion)
	q.Set("Timestamp", t)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func encodeQuery(u url.URL) string {
	query := url.Values{}
	for key, values := range u.Query() {
		for _, v := range values {
			query.Add(key, v)
		}
	}

	return query.Encode()
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := xml.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func generateSignature(key, msg string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(msg))
	signature := base64.StdEncoding.EncodeToString([]byte(hash.Sum(nil)))

	return signature
}
