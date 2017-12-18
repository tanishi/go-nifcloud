package nifcloud

import (
	"context"
	"fmt"
)

func (c *Client) AuthorizeSecurityGroupIngress(ctx context.Context, param *AuthorizeSecurityGroupIngressInput) (*AuthorizeSecurityGroupIngressOutput, error) {
	if param.GroupName == "" {
		return nil, fmt.Errorf("Validation error: missing GroupName")
	}

	if param.IPProtocol == "TCP" || param.IPProtocol == "UDP" {
		if param.FromPort == "" {
			return nil, fmt.Errorf("Validation error: missing FromPort")
		}
	}

	if param.IPGroupName == "" && param.IPRange == "" {
		return nil, fmt.Errorf("Validation error: missing IPGroupName or IPRange")
	}

	q := Query{
		"Action":                            "AuthorizeSecurityGroupIngress",
		"GroupName":                         param.GroupName,
		"IpPermissions.1.IpProtocol":        param.IPProtocol,
		"IpPermissions.1.IpRanges.1.CidrIp": param.IPRange,
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

	var body AuthorizeSecurityGroupIngressOutput

	if err := decodeBody(res.Body, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (c *Client) CreateSecurityGroup(ctx context.Context, param *CreateSecurityGroupInput) (*CreateSecurityGroupOutput, error) {
	if param.GroupName == "" {
		fmt.Errorf("Validation error: missing GroupName")
	}

	q := Query{
		"Action":                     "CreateSecurityGroup",
		"GroupName":                  param.GroupName,
		"GroupDescription":           param.GroupDescription,
		"Placement.AvailabilityZone": param.AvailabilityZone,
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
