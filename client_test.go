package nifcloud

import (
	"context"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	u := os.Getenv("NIFCLOUD_ENDPOINT")
	accessKey := os.Getenv("NIFCLOUD_ACCESSKEY")
	secretAccessKey := os.Getenv("NIFCLOUD_SECRET_ACCESSKEY")

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
