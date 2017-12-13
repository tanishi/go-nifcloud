package nifcloud

import (
	"context"
	"fmt"
)

func (c *Client) AuthorizeSecurityGroupIngress() {}

func (c *Client) CreateSecurityGroup(ctx context.Context, param *CreateSecurityGroupInput) (*CreateSecurityGroupOutput, error) {
	if param.GroupName == "" {
		fmt.Errorf("Validation error: missing GroupName")
	}

	u := c.URL
	q := u.Query()
	q.Set("Action", "CreateSecurityGroup")
	q.Set("GroupName", param.GroupName)
	//q.Set("GroupDescription", param.GroupDescription)
	//q.Set("Placement.AvailabilityZone", param.AvailabilityZone)
	u.RawQuery = q.Encode()

	req, err := c.NewRequest(ctx, "POST", u, nil)

	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	var body CreateSecurityGroupOutput

	if err := decodeBody(res, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (c *Client) DeleteSecurityGroup()                         {}
func (c *Client) DeregisterInstancesFromSecurityGroup()        {}
func (c *Client) DescribeSecurityActivities()                  {}
func (c *Client) DescribeSecurityGroups()                      {}
func (c *Client) RegisterInstancesWithSecurityGroup()          {}
func (c *Client) RevokeSecurityGroupIngress()                  {}
func (c *Client) UpdateSecurityGroup()                         {}
func (c *Client) UpdateSecurityGroupOption()                   {}
func (c *Client) DescribeSecurityGroupOption()                 {}
func (c *Client) NiftyRegisterRoutersWithSecurityGroup()       {}
func (c *Client) NiftyDeregisterRoutersFromSecurityGroup()     {}
func (c *Client) NiftyRegisterVpnGatewaysWithSecurityGroup()   {}
func (c *Client) NiftyDeregisterVpnGatewaysFromSecurityGroup() {}
