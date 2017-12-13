package nifcloud

type CreateSecurityGroupInput struct {
	GroupName        string
	GroupDescription string
	AvailabilityZone string
}

type CreateSecurityGroupOutput struct {
	RequestID string `xml:"requestId"`
	Return    bool   `xml:return`
}
