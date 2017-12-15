package nifcloud

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"net/http"
	"net/url"
	"strings"
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
	parsedURL, err := url.Parse(u)

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

func (c *Client) NewRequest(ctx context.Context, method string, query Query) (*http.Request, error) {
	query["AccessKeyId"] = c.AccessKeyID
	query["SignatureMethod"] = SignatureMethod
	query["SignatureVersion"] = SignatureVersion

	u := c.URL

	sign := generateStringToSign(method, u.Host, u.Path, query)
	signature := generateSignature(c.SecretAccessKey, sign)
	query["Signature"] = signature

	values := url.Values{}
	for key, value := range query {
		values.Set(key, value)
	}

	req, err := http.NewRequest(method, u.String(), strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	return req, nil
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

func generateStringToSign(method, endpoint, path string, q Query) string {
	sq := NewSortedQuery(q)
	return method + "\n" + endpoint + "\n" + path + "\n" + sq.String()
}
