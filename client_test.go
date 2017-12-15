package nifcloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestClient(t *testing.T) {
	u := "http://example.com/api"
	accessKey := "YOUR_ACCSESSKEY"
	secretAccessKey := "YOUR_SECRET_ACCESSKEY"

	actual, err := NewClient(u, accessKey, secretAccessKey)

	if err != nil {
		t.Error(err)
	}

	parsedURL, err := url.Parse(u)

	if err != nil {
		t.Error(err)
	}

	expected := &Client{
		URL:             parsedURL,
		HTTPClient:      http.DefaultClient,
		AccessKeyID:     accessKey,
		SecretAccessKey: secretAccessKey,
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, but: %v", expected, actual)
	}
}

func TestRequest(t *testing.T) {
	u := "http://example.com/api"
	accessKey := "YOUR_ACCSESSKEY"
	secretAccessKey := "YOUR_SECRET_ACCESSKEY"

	c, err := NewClient(u, accessKey, secretAccessKey)

	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err = c.NewRequest(ctx, "GET", Query{})

	if err != nil {
		t.Error(err)
	}
}

func TestGenerateSignature(t *testing.T) {
	method := "GET"
	endpoint := "west-1.cp.cloud.nifty.com"
	path := "/api/"
	accessKeyID := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	secretAccessKey := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	query := Query{
		"Action":           "DescribeInstances",
		"AccessKeyId":      accessKeyID,
		"SignatureMethod":  "HmacSHA256",
		"SignatureVersion": "2",
		"InstanceId":       "test001",
		"Description":      "/",
	}

	expected := "dHOoGcBgO14Roaioryic9IdFPg7G+lihZ8Wyoa25ok4="

	sign := generateStringToSign(method, endpoint, path, query)

	actual := generateSignature(secretAccessKey, sign)

	if actual != expected {
		t.Errorf("expected: %v, but: %v", expected, actual)
	}
}

func TestGenerateStringToSign(t *testing.T) {
	method := "GET"
	endpoint := "http://api.example.com"
	path := "/api/"
	query := Query{
		"Action":           "DescribeInstances",
		"AccessKeyId":      "YOUR_ACCSESSKEY",
		"SignatureMethod":  "HmacSHA256",
		"SignatureVersion": "2",
		"InstanceId":       "test001",
		"Description":      "/",
	}

	actual := generateStringToSign(method, endpoint, path, query)
	fmt.Println(actual)
}
