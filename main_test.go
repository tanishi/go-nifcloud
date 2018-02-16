package nifcloud

import (
	"flag"
	"os"
	"testing"
)

var c Nifcloud

func TestMain(m *testing.M) {
	flag.Parse()

	if testing.Short() {
		c = &Mock{}
	} else {
		u := os.Getenv("NIFCLOUD_ENDPOINT")
		accessKey := os.Getenv("NIFCLOUD_ACCESSKEY")
		secretAccessKey := os.Getenv("NIFCLOUD_SECRET_ACCESSKEY")

		c, _ = NewClient(u, accessKey, secretAccessKey)
	}

	code := m.Run()
	os.Exit(code)
}

type Mock struct {
	Nifcloud
}
