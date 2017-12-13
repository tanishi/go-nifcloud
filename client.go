package nifcloud

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	encodedQuery := encodeQuery(u)

	reqStr := fmt.Sprintf("SignatureMethod=%s&SignatureVersion=%s&AccessKeyId=%s&%s",
		SignatureMethod, SignatureVersion, c.AccessKeyID, encodedQuery)

	sign := fmt.Sprintf("%s\n%s\n%s\n%s", method, u.Host, u.Path, reqStr)

	hash := hmac.New(sha256.New, []byte(sign))
	hash.Write([]byte(c.SecretAccessKey))
	signature := base64.StdEncoding.EncodeToString([]byte(hash.Sum(nil)))

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Nifty-Authorization", signature)

	return req, nil
}

func encodeQuery(u *url.URL) string {
	query := url.Values{}
	for key, values := range u.Query() {
		for _, v := range values {
			query.Add(key, v)
		}
	}

	return query.Encode()
}
