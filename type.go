package nifcloud

type AuthorizeSecurityGroupIngressInput struct {
	GroupName     string
	IPPermissions []IPPermission
}

type IPPermission struct {
	IPProtocol  string
	FromPort    string
	ToPort      string
	InOut       string
	Groups      []string
	IPRanges    []string
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

type DescribeSecurityActivitiesInput struct {
	GroupName    string
	ActivityDate string
	Range        string
	StartNumber  string
	EndNumber    string
}

type DescribeSecurityActivitiesOutput struct {
	RequestID string `xml:"requestId"`
	GroupName string `xml:"groupName"`
	Log       string `xml:"log"`
}

type DescribeSecurityGroupsInput struct {
	GroupNames   []string
	FilterNames  []string
	FilterValues []string
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

type RegisterInstancesWithSecurityGroupInput struct {
	GroupName   string
	InstanceIDs []string
}

type RegisterInstancesWithSecurityGroupOutput struct {
	RequestID    string         `xml:"requestId"`
	InstancesSet []InstanceItem `xml:instancesSet>item`
}

type RevokeSecurityGroupIngressInput struct {
	GroupName     string
	IPPermissions []IPPermission
}

type RevokeSecurityGroupIngressOutput struct {
	RequestID string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

type UpdateSecurityGroupInput struct {
	GroupName              string
	GroupNameUpdate        string
	GroupDescriptionUpdate string
	GroupLogLimitUpdate    string
}

type UpdateSecurityGroupOutput struct {
	RequestID string `xml:"requestId"`
	Return    bool   `xml:"return"`
}

type DescribeInstancesInput struct {
	InstanceIDs []string
	Tenancies   []string
}

type DescribeInstancesOutput struct {
	RequestID    string          `xml:"requestId"`
	GroupSet     []Group         `xml:"groupSet>item"`
	InstancesSet []InstancesItem `xml:"reservationSet>item>instancesSet>item"`
}

type Group struct {
	GroupID string `xml:"groupId"`
}

type InstancesItem struct {
	InstanceID                string                 `xml:"instanceId"`
	InstanceUniqueID          string                 `xml:"instanceUniqueId"`
	ImageID                   string                 `xml:"imageId"`
	InstanceState             InstanceStateStruct    `xml:"instanceState"`
	PrivateDNSName            string                 `xml:"privateDnsName"`
	DNSName                   string                 `xml:"dnsName"`
	KeyName                   string                 `xml:"keyName"`
	InstanceType              string                 `xml:"instanceType"`
	LaunchTime                string                 `xml:"launchTime"`
	AvailabilityZone          string                 `xml:"placement>availabilityZone"`
	Platform                  string                 `xml:"platform"`
	ImageName                 string                 `xml:"imageName"`
	Monitoring                string                 `xml:"monitoring>state"`
	PrivateIPAddress          string                 `xml:"privateIpAddress"`
	IPAddress                 string                 `xml:"IpAddress"`
	PrivateIPAddressv6        string                 `xml:"privateIpAddressV6"`
	IPAddressv6               string                 `xml:"IpAddressV6"`
	Architecture              string                 `xml:"architecture"`
	RootDeviceType            string                 `xml:"rootDeviceType"`
	BlockDeviceMapping        string                 `xml:"blockDeviceMapping>item"`
	AccountingType            string                 `xml:"accountingType"`
	NextMonthAccountingType   string                 `xml:"nextMonthAccountingType"`
	LoadBalancing             []LoadBalancer         `xml:"loadBalancing>item"`
	State                     string                 `xml:"State"`
	CopyInfo                  string                 `xml:"copyInfo"`
	AutoScaling               AutoScalingStruct      `xml:"autoscaling>item"`
	IPType                    string                 `xml:"ipType"`
	NifryPrivateIPType        string                 `xml:"niftyPrivateIpType"`
	HotAdd                    string                 `xml:"hotAdd"`
	NiftySnapShotting         string                 `xml:"niftySnapshotting>item>state"`
	NiftyPrivateNetworkType   string                 `xml:"niftyPrivateNetworkType"`
	Tenancy                   string                 `xml:"tenancy"`
	NetworkInterfaceSet       []NetworkInterfaceItem `xml:"networkInterfaceSet>item"`
	NiftyElasticLoadBalancing string                 `xml:"niftyElasticLoadBalancing>item"`
}

type InstanceStateStruct struct {
	Code string `xml:"code"`
	Name string `xml:"name"`
}

type BlockDeviceMappingItem struct {
	DeviceName string    `xml:"deviceName"`
	EBS        EBSStruct `xml:"ebs"`
}

type EBSStruct struct {
	VolumeID            string `xml:"volumeId"`
	Status              string `xml:"status"`
	AttachTime          string `xml:"attachTime"`
	DeleteOnTermination bool   `xml:"deleteOnTermination"`
}

type LoadBalancer struct {
	LoadBalancerName string `xml:"loadBalancerName"`
	LoadBalancerPort string `xml:"loadBalancerPort"`
	InstancePort     string `xml:"instancePort"`
}

type AutoScalingStruct struct {
	AutoScalingGroupName string `xml:"autoScalingGroupName"`
	ExpireTime           string `xml:"expireTime"`
}

type NetworkInterfaceItem struct {
	NetworkInterfaceID string            `xml:"networkInterfaceId"`
	NiftyNetworkID     string            `xml:"niftyNetworkId"`
	NiftyNeteorkName   string            `xml:"niftyNetworkName"`
	Status             string            `xml:"status"`
	MacAddress         string            `xml:"macAddress"`
	PrivateIPAddress   string            `xml:"privateIpAddress"`
	PrivateIPAddressv6 string            `xml:"pruvateIpAddressV6"`
	SourceDestCheck    string            `xml:"sourceDestCheck"`
	Attachment         AttachmentStruct  `xml:"attachment"`
	Association        AssociationStruct `xml:"association"`
}

type AttachmentStruct struct {
	AttachmentID        string `xml:"attachmentId"`
	DeviceIndex         string `xml:"deviceIndex"`
	Status              string `xml:"status"`
	AttachTime          string `xml:"attachTime"`
	DeleteOnTermination string `xml:"deleteOnTermination"`
}

type AssociationStruct struct {
	PublicIP        string `xml:"publicIp"`
	PublicIPv6      string `xml:"publicIpV6"`
	PublicDNSName   string `xml:"publicDnsName"`
	PublicDNSNamev6 string `xml:"publicDnsNameV6"`
}

type NiftyElasticLoadBalancingItem struct {
	ElasticLoadBalancerID   string `xml:"elasticLoadBalancerId"`
	ElasticLoadBalancerName string `xml:"elasticLoadBalancerName"`
	Protocol                string `xml:"protocol"`
	ElasticLoadBalancerPort string `xml:"elasticLoadBalancerPort"`
	InstancePort            string `xml:"instancePort"`
}

type DescribeInstanceAttributeInput struct {
	InstanceID string
	Attribute  string
}

type DescribeInstanceAttributeOutput struct {
	RequestID                 string                          `xml:"requestId"`
	InstanceID                string                          `xml:"instanceId"`
	InstanceUniqueID          string                          `xml:"instanceUniqueId"`
	InstanceType              string                          `xml:"instanceType>value"`
	DisableAPITermination     bool                            `xml:"disableApiTermination>value"`
	BlockDeviceMapping        []BlockDeviceMappingItem        `xml:"blockDeviceMapping>item"`
	AccountingType            string                          `xml:"accountingType>value"`
	NextMonthAccountingType   string                          `xml:"nextMonthAccountingType>value"`
	LoadBalancing             []LoadBalancer                  `xml:"loadBalancing>item"`
	CopyInfo                  string                          `xml:"copyInfo>value"`
	AutoScaling               []AutoScalingStruct             `xml:"autoscaling>item"`
	IPType                    string                          `xml:"ipType"`
	NifryPrivateIPType        string                          `xml:"niftyPrivateIpType"`
	GroupID                   string                          `xml:"groupId>value"`
	Description               string                          `xml:"desctiption>value"`
	NetworkInterfaceSet       []NetworkInterfaceItem          `xml:"networkInterfaceSet>item"`
	NiftyElasticLoadBalancing []NiftyElasticLoadBalancingItem `xml:"niftyElasticLoadBalancing>item"`
}

type Query map[string]string

func (q Query) Set(k, v string) {
	q[k] = v
}
