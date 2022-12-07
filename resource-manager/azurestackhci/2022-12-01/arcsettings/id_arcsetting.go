package arcsettings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ArcSettingId{}

// ArcSettingId is a struct representing the Resource ID for a Arc Setting
type ArcSettingId struct {
	SubscriptionId    string
	ResourceGroupName string
	ClusterName       string
	ArcSettingName    string
}

// NewArcSettingID returns a new ArcSettingId struct
func NewArcSettingID(subscriptionId string, resourceGroupName string, clusterName string, arcSettingName string) ArcSettingId {
	return ArcSettingId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ClusterName:       clusterName,
		ArcSettingName:    arcSettingName,
	}
}

// ParseArcSettingID parses 'input' into a ArcSettingId
func ParseArcSettingID(input string) (*ArcSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(ArcSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ArcSettingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ClusterName, ok = parsed.Parsed["clusterName"]; !ok {
		return nil, fmt.Errorf("the segment 'clusterName' was not found in the resource id %q", input)
	}

	if id.ArcSettingName, ok = parsed.Parsed["arcSettingName"]; !ok {
		return nil, fmt.Errorf("the segment 'arcSettingName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseArcSettingIDInsensitively parses 'input' case-insensitively into a ArcSettingId
// note: this method should only be used for API response data and not user input
func ParseArcSettingIDInsensitively(input string) (*ArcSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(ArcSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ArcSettingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ClusterName, ok = parsed.Parsed["clusterName"]; !ok {
		return nil, fmt.Errorf("the segment 'clusterName' was not found in the resource id %q", input)
	}

	if id.ArcSettingName, ok = parsed.Parsed["arcSettingName"]; !ok {
		return nil, fmt.Errorf("the segment 'arcSettingName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateArcSettingID checks that 'input' can be parsed as a Arc Setting ID
func ValidateArcSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseArcSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Arc Setting ID
func (id ArcSettingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AzureStackHCI/clusters/%s/arcSettings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterName, id.ArcSettingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Arc Setting ID
func (id ArcSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAzureStackHCI", "Microsoft.AzureStackHCI", "Microsoft.AzureStackHCI"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterValue"),
		resourceids.StaticSegment("staticArcSettings", "arcSettings", "arcSettings"),
		resourceids.UserSpecifiedSegment("arcSettingName", "arcSettingValue"),
	}
}

// String returns a human-readable description of this Arc Setting ID
func (id ArcSettingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Arc Setting Name: %q", id.ArcSettingName),
	}
	return fmt.Sprintf("Arc Setting (%s)", strings.Join(components, "\n"))
}
