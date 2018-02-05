package nifcloud

import (
	"context"
	"fmt"
	"strconv"
)

func (c *Client) DescribeInstanceAttribute(ctx context.Context, param *DescribeInstanceAttributeInput) (*DescribeInstanceAttributeOutput, error) {
	if param.InstanceID == "" {
		return nil, fmt.Errorf("Validation error: missing InstanceID")
	}

	q := Query{
		"Action":     "DescribeInstanceAttribute",
		"InstanceId": param.InstanceID,
		"Attribute":  param.Attribute,
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

	var body DescribeInstanceAttributeOutput

	if err := decodeBody(res.Body, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (c *Client) DescribeInstances(ctx context.Context, param *DescribeInstancesInput) (*DescribeInstancesOutput, error) {
	q := Query{
		"Action": "DescribeInstances",
	}

	for i, v := range param.InstanceIDs {
		n := strconv.Itoa(i + 1)
		q.Set("InstanceId."+n, v)
	}

	for i, v := range param.Tenancies {
		n := strconv.Itoa(i + 1)
		q.Set("Tenancy."+n, v)
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

	var body DescribeInstancesOutput

	if err := decodeBody(res.Body, &body); err != nil {
		return nil, err
	}

	return &body, nil
}
