package localrules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = LocalRuleId{}

// LocalRuleId is a struct representing the Resource ID for a Local Rule
type LocalRuleId struct {
	SubscriptionId     string
	ResourceGroupName  string
	LocalRuleStackName string
	LocalRuleName      string
}

// NewLocalRuleID returns a new LocalRuleId struct
func NewLocalRuleID(subscriptionId string, resourceGroupName string, localRuleStackName string, localRuleName string) LocalRuleId {
	return LocalRuleId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		LocalRuleStackName: localRuleStackName,
		LocalRuleName:      localRuleName,
	}
}

// ParseLocalRuleID parses 'input' into a LocalRuleId
func ParseLocalRuleID(input string) (*LocalRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(LocalRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := LocalRuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocalRuleStackName, ok = parsed.Parsed["localRuleStackName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "localRuleStackName", *parsed)
	}

	if id.LocalRuleName, ok = parsed.Parsed["localRuleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "localRuleName", *parsed)
	}

	return &id, nil
}

// ParseLocalRuleIDInsensitively parses 'input' case-insensitively into a LocalRuleId
// note: this method should only be used for API response data and not user input
func ParseLocalRuleIDInsensitively(input string) (*LocalRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(LocalRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := LocalRuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocalRuleStackName, ok = parsed.Parsed["localRuleStackName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "localRuleStackName", *parsed)
	}

	if id.LocalRuleName, ok = parsed.Parsed["localRuleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "localRuleName", *parsed)
	}

	return &id, nil
}

// ValidateLocalRuleID checks that 'input' can be parsed as a Local Rule ID
func ValidateLocalRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLocalRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Local Rule ID
func (id LocalRuleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/PaloAltoNetworks.CloudNGFW/localRuleStacks/%s/localRules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LocalRuleStackName, id.LocalRuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Local Rule ID
func (id LocalRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticPaloAltoNetworksCloudNGFW", "PaloAltoNetworks.CloudNGFW", "PaloAltoNetworks.CloudNGFW"),
		resourceids.StaticSegment("staticLocalRuleStacks", "localRuleStacks", "localRuleStacks"),
		resourceids.UserSpecifiedSegment("localRuleStackName", "localRuleStackValue"),
		resourceids.StaticSegment("staticLocalRules", "localRules", "localRules"),
		resourceids.UserSpecifiedSegment("localRuleName", "localRuleValue"),
	}
}

// String returns a human-readable description of this Local Rule ID
func (id LocalRuleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Local Rule Stack Name: %q", id.LocalRuleStackName),
		fmt.Sprintf("Local Rule Name: %q", id.LocalRuleName),
	}
	return fmt.Sprintf("Local Rule (%s)", strings.Join(components, "\n"))
}
