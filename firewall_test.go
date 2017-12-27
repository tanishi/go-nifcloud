package nifcloud

import (
	"context"
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
		GroupName: "tanishi",
		IPPermissions: []IPPermission{
			IPPermission{
				IPProtocol: "HTTP",
				IPRanges:   []string{"0.0.0.0/0"},
			},
		},
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

	_, err = c.DeregisterInstancesFromSecurityGroup(ctx, params)

	if err != nil {
		t.Error(err)
	}
}

func TestDescribeSecurityActiviries(t *testing.T) {
	u := os.Getenv("NIFCLOUD_ENDPOINT")
	accessKey := os.Getenv("NIFCLOUD_ACCESSKEY")
	secretAccessKey := os.Getenv("NIFCLOUD_SECRET_ACCESSKEY")

	c, err := NewClient(u, accessKey, secretAccessKey)

	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	params := &DescribeSecurityActivitiesInput{
		GroupName: "tanishi",
	}

	_, err = c.DescribeSecurityActivities(ctx, params)

	if err != nil {
		t.Error(err)
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

	_, err = c.DescribeSecurityGroups(ctx, params)

	if err != nil {
		t.Error(err)
	}
}

func TestRegisterInstancesWithSecurityGroup(t *testing.T) {
	u := os.Getenv("NIFCLOUD_ENDPOINT")
	accessKey := os.Getenv("NIFCLOUD_ACCESSKEY")
	secretAccessKey := os.Getenv("NIFCLOUD_SECRET_ACCESSKEY")

	c, err := NewClient(u, accessKey, secretAccessKey)

	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	params := &RegisterInstancesWithSecurityGroupInput{
		GroupName:   "tanishi",
		InstanceIDs: []string{"onishiTest"},
	}

	_, err = c.RegisterInstancesWithSecurityGroup(ctx, params)

	if err != nil {
		t.Error(err)
	}
}

func TestRevokeSecurityGroupIngress(t *testing.T) {
	u := os.Getenv("NIFCLOUD_ENDPOINT")
	accessKey := os.Getenv("NIFCLOUD_ACCESSKEY")
	secretAccessKey := os.Getenv("NIFCLOUD_SECRET_ACCESSKEY")

	c, err := NewClient(u, accessKey, secretAccessKey)

	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	params := &RevokeSecurityGroupIngressInput{
		GroupName: "tanishi",
		IPPermissions: []IPPermission{
			IPPermission{
				IPProtocol: "HTTP",
				IPRanges:   []string{"0.0.0.0/0"},
			},
		},
	}

	_, err = c.RevokeSecurityGroupIngress(ctx, params)

	if err != nil {
		t.Error(err)
	}
}
