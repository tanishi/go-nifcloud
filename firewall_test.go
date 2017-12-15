package nifcloud

import (
	"context"
	"os"
	"testing"
)

func TestCreateSecurityGroup(t *testing.T) {
	u := os.Getenv("NIFCLOUD_ENDPOINT")
	accessKey := os.Getenv("NIFCLOUD_ACCESSKEY")
	secretAccessKey := os.Getenv("NIFCLOUD_SECRET_ACCESSKEY")

	c, err := NewClient(u, accessKey, secretAccessKey)

	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	params := &CreateSecurityGroupInput{
		GroupName: "tanishi",
	}

	res, err := c.CreateSecurityGroup(ctx, params)

	if err != nil {
		t.Error(err)
	}

	if !res.Return {
		t.Error(res.Return)
	}
}

func TestDeleteSecurityGroup(t *testing.T) {
	u := os.Getenv("NIFCLOUD_ENDPOINT")
	accessKey := os.Getenv("NIFCLOUD_ACCESSKEY")
	secretAccessKey := os.Getenv("NIFCLOUD_SECRET_ACCESSKEY")

	c, err := NewClient(u, accessKey, secretAccessKey)

	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	params := &DeleteSecurityGroupInput{
		GroupName: "tanishi",
	}

	res, err := c.DeleteSecurityGroup(ctx, params)

	if err != nil {
		t.Error(err)
	}

	if !res.Return {
		t.Error(res.Return)
	}
}

func TestDescribeSecurityGroups(t *testing.T) {
	u := os.Getenv("NIFCLOUD_ENDPOINT")
	accessKey := os.Getenv("NIFCLOUD_ACCESSKEY")
	secretAccessKey := os.Getenv("NIFCLOUD_SECRET_ACCESSKEY")

	c, err := NewClient(u, accessKey, secretAccessKey)

	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	params := &DescribeSecurityGroupsInput{}

	res, err := c.DescribeSecurityGroups(ctx, params)

	if err != nil {
		t.Error(err)
	}
}
