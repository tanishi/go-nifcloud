package nifcloud

type AuthorizeSecurityGroupIngressInput struct {
	GroupName   string
	IPProtocol  string
	FromPort    string
	ToPort      string
	InOut       string
	IPGroupName string
	IPRange     string
	Description string
}

type AuthorizeSecurityGroupIngressOutput struct {
	RequestID string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

type CreateSecurityGroupInput struct {
	GroupName        string
	GroupDescription string
	AvailabilityZone string
}

type CreateSecurityGroupOutput struct {
	RequestID string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

type DeleteSecurityGroupInput struct {
	GroupName string
}

type DeleteSecurityGroupOutput struct {
	RequestID string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

type DeregisterInstancesFromSecurityGroupInput struct {
	GroupName   string
	InstanceIDs []string
}

type DeregisterInstancesFromSecurityGroupOutput struct {
	RequestID    string         `xml:"requestId"`
	InstancesSet []InstanceItem `xml:instancesSet>item`
}

type DescribeSecurityGroupsInput struct {
	GroupName   string
	FilterName  string
	FilterValue string
}

type DescribeSecurityGroupsOutput struct {
	RequestID         string                  `xml:"requestId"`
	SecurityGroupInfo []SecurityGroupInfoItem `xml:"securityGroupInfo>item"`
}

type SecurityGroupInfoItem struct {
	ownerID                       string                       `xml:"ownerID"`
	GroupName                     string                       `xml:"groupName"`
	GroupDescription              string                       `xml:"groupDescription"`
	GroupStatus                   string                       `xml:"groupStatus"`
	IPPermission                  []IPPermissionItem           `xml:"ipPermissions>item"`
	Instances                     []InstanceItem               `xml:"instancesSet>item"`
	InstanceUniqueIdsSet          []InstanceUniqueID           `xml:"instanceUniqueIdsSet>item"`
	InstancesNetworkInterfaceSet  []InstanceNetworkInterfaces  `xml:"instancesNetworkInterfaceSet>item"`
	Routers                       []Router                     `xml:"routerSet>item"`
	RouterNetworkInterfaceSet     []RouterNetworkInterface     `xml:"routerNetworkInterfaceSet>item"`
	VPNGatewaySet                 []VPNGateWay                 `xml:"vpnGatewaySet>item"`
	VPNGatewayNetworkInterfaceSet []VPNGatewayNetworkInterface `xml:"vpnGatewayNetworkInterfaceSet>item"`
	GroupRuleLimit                int                          `xml:"groupRuleLimit"`
	GroupLogLimit                 int                          `xml:"groupLogLimit"`
	GroupLogFilterNetBios         bool                         `xml:"groupLogFilterNetBios"`
	GroupLogFilterBroadcast       bool                         `xml:"groupLogFilterBroadcast"`
	AvailabilityZone              string                       `xml:"availabilityZone"`
}

type IPPermissionItem struct {
	IPProtocol  string                   `xml:"ipProtocol"`
	FromPort    int                      `xml:"fromPort"`
	ToPort      int                      `xml:"toPort"`
	InOut       string                   `xml:"inOut"`
	Groups      []UserIDGroupPairSetType `xml:"groups>item"`
	IPRanges    []IPRange                `xml:"ipRanges>item"`
	Description string                   `xml:"description"`
	AddDateTime string                   `xml:"addDatetime"`
}

type UserIDGroupPairSetType struct {
	UserID    string `xml:"userId"`
	GroupName string `xml:"groupName"`
}

type IPRange struct {
	CidrIP string `xml:"cidrIp"`
}

type InstanceItem struct {
	InstanceID string `xml:"instanceId"`
}

type InstanceUniqueID struct {
	InstanceIDUniqueID string `xml:"instanceUniqueId"`
}

type InstanceNetworkInterfaces struct {
	InstanceID       string `xml:"instanceId"`
	InstanceUniqueID string `xml:"instanceUniqueId"`
	NetworkID        string `xml:"networkId"`
	deviceIntdex     string `xml:"deviceIndex"`
	IPAddress        string `xml:"ipAddredd"`
}

type Router struct {
	RouterID   string `xml:"routerId"`
	RouterName string `xml:"routerName"`
}

type RouterNetworkInterface struct {
	RouterID     string `xml:"routerId"`
	RouterName   string `xml:"routerName"`
	NetworkID    string `xml:"networkId"`
	deviceIntdex string `xml:"deviceIndex"`
	IPAddress    string `xml:"ipAddredd"`
}

type VPNGateWay struct {
	VPNGateWayID      string `xml:"vpnGatewayId"`
	NiftyVPNGateWayID string `xml:"niftyVpnGatewayName"`
}

type VPNGatewayNetworkInterface struct {
	VPNGateWayID      string `xml:"vpnGatewayId"`
	NiftyVPNGateWayID string `xml:"niftyVpnGatewayName"`
	NetworkID         string `xml:"networkId"`
	deviceIntdex      string `xml:"deviceIndex"`
	IPAddress         string `xml:"ipAddredd"`
}

type Query map[string]string
