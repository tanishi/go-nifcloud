package nifcloud

import (
	"context"
	"fmt"
	"strconv"
)

func (c *Client) AuthorizeSecurityGroupIngress(ctx context.Context, param *AuthorizeSecurityGroupIngressInput) (*AuthorizeSecurityGroupIngressOutput, error) {
	if param.GroupName == "" {
		return nil, fmt.Errorf("Validation error: missing GroupName")
	}

	q := Query{
		"Action":    "AuthorizeSecurityGroupIngress",
		"GroupName": param.GroupName,
	}

	for i, v := range param.IPPermissions {
		if len(v.Groups) == 0 && len(v.IPRanges) == 0 {
			return nil, fmt.Errorf("Validation error: missing IPGroupName or IPRange")
		}

		if v.IPProtocol == "TCP" || v.IPProtocol == "UDP" {
			if v.FromPort == "" {
				return nil, fmt.Errorf("Validation error: missing FromPort")
			}
		}

		n := strconv.Itoa(i + 1)
		q.Set("IpPermissions."+n+".IpProtocol", v.IPProtocol)
		q.Set("IpPermissions."+n+".FromPort", v.FromPort)
		q.Set("IpPermissions."+n+".ToPort", v.ToPort)
		q.Set("IpPermissions."+n+".InOut", v.InOut)
		q.Set("IpPermissions."+n+".Description", v.Description)

		for j, g := range v.Groups {
			m := strconv.Itoa(j + 1)
			q.Set("IpPermissions."+n+".Groups."+m+".GroupName", g)
		}

		for j, ip := range v.IPRanges {
			m := strconv.Itoa(j + 1)
			q.Set("IpPermissions."+n+".IpRanges."+m+".CidrIp", ip)
		}
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

	if body.Error != nil {
		return nil, body.Error
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

func (c *Client) DeregisterInstancesFromSecurityGroup(ctx context.Context, param *DeregisterInstancesFromSecurityGroupInput) (*DeregisterInstancesFromSecurityGroupOutput, error) {
	if param.GroupName == "" {
		return nil, fmt.Errorf("Validation error: missing GroupName")
	}

	if len(param.InstanceIDs) == 0 {
		return nil, fmt.Errorf("Validation error: missing InstanceID")
	}

	q := Query{
		"Action":    "DeregisterInstancesFromSecurityGroup",
		"GroupName": param.GroupName,
	}

	for i, v := range param.InstanceIDs {
		n := strconv.Itoa(i + 1)
		q.Set("InstanceId."+n, v)
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

	var body DeregisterInstancesFromSecurityGroupOutput

	if err := decodeBody(res.Body, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (c *Client) DescribeSecurityActivities(ctx context.Context, param *DescribeSecurityActivitiesInput) (*DescribeSecurityActivitiesOutput, error) {
	q := Query{
		"Action":            "DescribeSecurityActivities",
		"GroupName":         param.GroupName,
		"ActivityDate":      param.ActivityDate,
		"Range.All":         param.Range,
		"Range.StartNumber": param.StartNumber,
		"Range.EndNumber":   param.EndNumber,
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

	var body DescribeSecurityActivitiesOutput

	if err := decodeBody(res.Body, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (c *Client) DescribeSecurityGroups(ctx context.Context, param *DescribeSecurityGroupsInput) (*DescribeSecurityGroupsOutput, error) {
	q := Query{
		"Action": "DescribeSecurityGroups",
	}

	for i, v := range param.GroupNames {
		n := strconv.Itoa(i + 1)
		q.Set("GroupName."+n, v)
	}

	for i, name := range param.FilterNames {
		n := strconv.Itoa(i + 1)
		q.Set("Filter."+n+".Name", name)

		for j, v := range param.GroupNames {
			m := strconv.Itoa(j + 1)
			q.Set("Filter."+n+".Value."+m, v)
		}
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

func (c *Client) RegisterInstancesWithSecurityGroup(ctx context.Context, param *RegisterInstancesWithSecurityGroupInput) (*RegisterInstancesWithSecurityGroupOutput, error) {
	q := Query{
		"Action":    "RegisterInstancesWithSecurityGroup",
		"GroupName": param.GroupName,
	}

	for i, v := range param.InstanceIDs {
		n := strconv.Itoa(i + 1)
		q.Set("InstanceId."+n, v)
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

	var body RegisterInstancesWithSecurityGroupOutput

	if err := decodeBody(res.Body, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (c *Client) RevokeSecurityGroupIngress(ctx context.Context, param *RevokeSecurityGroupIngressInput) (*RevokeSecurityGroupIngressOutput, error) {
	if param.GroupName == "" {
		return nil, fmt.Errorf("Validation error: missing GroupName")
	}

	q := Query{
		"Action":    "RevokeSecurityGroupIngress",
		"GroupName": param.GroupName,
	}

	for i, v := range param.IPPermissions {
		if len(v.Groups) == 0 && len(v.IPRanges) == 0 {
			return nil, fmt.Errorf("Validation error: missing IPGroupName or IPRange")
		}

		if v.IPProtocol == "TCP" || v.IPProtocol == "UDP" {
			if v.FromPort == "" {
				return nil, fmt.Errorf("Validation error: missing FromPort")
			}
		}

		n := strconv.Itoa(i + 1)
		q.Set("IpPermissions."+n+".IpProtocol", v.IPProtocol)
		q.Set("IpPermissions."+n+".FromPort", v.FromPort)
		q.Set("IpPermissions."+n+".ToPort", v.ToPort)
		q.Set("IpPermissions."+n+".InOut", v.InOut)
		q.Set("IpPermissions."+n+".Description", v.Description)

		for j, g := range v.Groups {
			m := strconv.Itoa(j + 1)
			q.Set("IpPermissions."+n+".Groups."+m+".GroupName", g)
		}

		for j, ip := range v.IPRanges {
			m := strconv.Itoa(j + 1)
			q.Set("IpPermissions."+n+".IpRanges."+m+".CidrIp", ip)
		}
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

	var body RevokeSecurityGroupIngressOutput

	if err := decodeBody(res.Body, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (c *Client) UpdateSecurityGroup(ctx context.Context, param *UpdateSecurityGroupInput) (*UpdateSecurityGroupOutput, error) {
	if param.GroupName == "" {
		return nil, fmt.Errorf("Validation error: missing GroupName")
	}

	q := Query{
		"Action":                 "UpdateSecurityGroup",
		"GroupName":              param.GroupName,
		"GroupNameUpdate":        param.GroupNameUpdate,
		"GroupDescriptionUpdate": param.GroupDescriptionUpdate,
		"GroupLogLimitUpdate":    param.GroupLogLimitUpdate,
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

	var body UpdateSecurityGroupOutput

	if err := decodeBody(res.Body, &body); err != nil {
		return nil, err
	}

	return &body, nil
}
