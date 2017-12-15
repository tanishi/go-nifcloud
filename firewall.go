package nifcloud

import (
	"context"
	"fmt"
)

func (c *Client) CreateSecurityGroup(ctx context.Context, param *CreateSecurityGroupInput) (*CreateSecurityGroupOutput, error) {
	if param.GroupName == "" {
		fmt.Errorf("Validation error: missing GroupName")
	}

	q := Query{
		"Action":    "CreateSecurityGroup",
		"GroupName": param.GroupName,
	}

	req, err := c.NewRequest(ctx, "POST", q)

	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body CreateSecurityGroupOutput

	if err := decodeBody(res.Body, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (c *Client) DeleteSecurityGroup(ctx context.Context, param *DeleteSecurityGroupInput) (*DeleteSecurityGroupOutput, error) {
	if param.GroupName == "" {
		fmt.Errorf("Validation error: missing GroupName")
	}

	q := Query{
		"Action":    "DeleteSecurityGroup",
		"GroupName": param.GroupName,
	}

	req, err := c.NewRequest(ctx, "POST", q)

	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body DeleteSecurityGroupOutput

	if err := decodeBody(res.Body, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (c *Client) DescribeSecurityGroups(ctx context.Context, param *DescribeSecurityGroupsInput) (*DescribeSecurityGroupsOutput, error) {
	q := Query{
		"Action":    "DescribeSecurityGroups",
		"GroupName": param.GroupName,
	}

	req, err := c.NewRequest(ctx, "POST", q)

	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body DescribeSecurityGroupsOutput

	if err := decodeBody(res.Body, &body); err != nil {
		return nil, err
	}

	return &body, nil
}
