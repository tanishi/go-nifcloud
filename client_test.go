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

	_, err = c.NewRequest(ctx, "GET", c.URL, nil)

	if err != nil {
		t.Error(err)
	}
}
