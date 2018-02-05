package nifcloud

import (
	"context"
	"os"
	"testing"
)

func TestDescribeInstanceAttribute(t *testing.T) {
	u := os.Getenv("NIFCLOUD_ENDPOINT")
	accessKey := os.Getenv("NIFCLOUD_ACCESSKEY")
	secretAccessKey := os.Getenv("NIFCLOUD_SECRET_ACCESSKEY")

	c, err := NewClient(u, accessKey, secretAccessKey)

	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	params := &DescribeInstanceAttributeInput{
		InstanceID: "onishiTest",
	}

	if _, err := c.DescribeInstanceAttribute(ctx, params); err != nil {
		t.Error(err)
	}
}

func TestDescribeInstances(t *testing.T) {
	u := os.Getenv("NIFCLOUD_ENDPOINT")
	accessKey := os.Getenv("NIFCLOUD_ACCESSKEY")
	secretAccessKey := os.Getenv("NIFCLOUD_SECRET_ACCESSKEY")

	c, err := NewClient(u, accessKey, secretAccessKey)

	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	params := &DescribeInstancesInput{
		InstanceIDs: []string{"onishiTest"},
	}

	if _, err := c.DescribeInstances(ctx, params); err != nil {
		t.Error(err)
	}
}
