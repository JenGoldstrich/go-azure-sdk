package iotconnectors

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = IotConnectorId{}

// IotConnectorId is a struct representing the Resource ID for a Iot Connector
type IotConnectorId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	IotConnectorName  string
}

// NewIotConnectorID returns a new IotConnectorId struct
func NewIotConnectorID(subscriptionId string, resourceGroupName string, workspaceName string, iotConnectorName string) IotConnectorId {
	return IotConnectorId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		IotConnectorName:  iotConnectorName,
	}
}

// ParseIotConnectorID parses 'input' into a IotConnectorId
func ParseIotConnectorID(input string) (*IotConnectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(IotConnectorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IotConnectorId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.IotConnectorName, ok = parsed.Parsed["iotConnectorName"]; !ok {
		return nil, fmt.Errorf("the segment 'iotConnectorName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseIotConnectorIDInsensitively parses 'input' case-insensitively into a IotConnectorId
// note: this method should only be used for API response data and not user input
func ParseIotConnectorIDInsensitively(input string) (*IotConnectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(IotConnectorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IotConnectorId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.IotConnectorName, ok = parsed.Parsed["iotConnectorName"]; !ok {
		return nil, fmt.Errorf("the segment 'iotConnectorName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateIotConnectorID checks that 'input' can be parsed as a Iot Connector ID
func ValidateIotConnectorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIotConnectorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Iot Connector ID
func (id IotConnectorId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.HealthcareApis/workspaces/%s/iotConnectors/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.IotConnectorName)
}

// Segments returns a slice of Resource ID Segments which comprise this Iot Connector ID
func (id IotConnectorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftHealthcareApis", "Microsoft.HealthcareApis", "Microsoft.HealthcareApis"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticIotConnectors", "iotConnectors", "iotConnectors"),
		resourceids.UserSpecifiedSegment("iotConnectorName", "iotConnectorValue"),
	}
}

// String returns a human-readable description of this Iot Connector ID
func (id IotConnectorId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Iot Connector Name: %q", id.IotConnectorName),
	}
	return fmt.Sprintf("Iot Connector (%s)", strings.Join(components, "\n"))
}
