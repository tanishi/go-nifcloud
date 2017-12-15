package nifcloud

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"sort"
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

	encodedQuery := encodeQuery(query)

	req, err := http.NewRequest(method, u.String(), strings.NewReader(encodedQuery))
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeBody(body io.Reader, out interface{}) error {
	decoder := xml.NewDecoder(body)
	return decoder.Decode(out)
}

func generateSignature(key, msg string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(msg))
	signature := base64.StdEncoding.EncodeToString([]byte(hash.Sum(nil)))

	return signature
}

func generateStringToSign(method, endpoint, path string, q Query) string {
	encodedQuery := encodeQuery(q)

	return method + "\n" + endpoint + "\n" + path + "\n" + encodedQuery
}

func encodeQuery(q Query) string {
	keys := make([]string, 0, len(q))
	for key, _ := range q {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	a := make([]string, 0, len(q))
	for _, key := range keys {
		a = append(a, url.QueryEscape(key)+"="+url.QueryEscape(q[key]))
	}
	s := strings.Join(a, "&")

	return s
}
