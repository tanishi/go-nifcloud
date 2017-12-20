package nifcloud

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestAuthorizeSecurityGroupIngress(t *testing.T) {
	u := os.Getenv("NIFCLOUD_ENDPOINT")
	accessKey := os.Getenv("NIFCLOUD_ACCESSKEY")
	secretAccessKey := os.Getenv("NIFCLOUD_SECRET_ACCESSKEY")

	c, err := NewClient(u, accessKey, secretAccessKey)

	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	params := &AuthorizeSecurityGroupIngressInput{
		GroupName:  "tanishi",
		IPProtocol: "HTTP",
		IPRange:    "0.0.0.0/0",
	}

	res, err := c.AuthorizeSecurityGroupIngress(ctx, params)

	if err != nil {
		t.Error(err)
	}

	if !res.Return {
		t.Errorf("Request Failed")
	}
}

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

func TestDeregisterInstancesFromSecurityGroup(t *testing.T) {
	u := os.Getenv("NIFCLOUD_ENDPOINT")
	accessKey := os.Getenv("NIFCLOUD_ACCESSKEY")
	secretAccessKey := os.Getenv("NIFCLOUD_SECRET_ACCESSKEY")

	c, err := NewClient(u, accessKey, secretAccessKey)

	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	params := &DeregisterInstancesFromSecurityGroupInput{
		GroupName:   "tanishi",
		InstanceIDs: []string{"onishiTest"},
	}

	res, err := c.DeregisterInstancesFromSecurityGroup(ctx, params)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
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

	_, err = c.DescribeSecurityGroups(ctx, params)

	if err != nil {
		t.Error(err)
	}
}
