package timeseriesdatabaseconnections

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionType string

const (
	ConnectionTypeAzureDataExplorer ConnectionType = "AzureDataExplorer"
)

func PossibleValuesForConnectionType() []string {
	return []string{
		string(ConnectionTypeAzureDataExplorer),
	}
}

func (s *ConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectionType(input string) (*ConnectionType, error) {
	vals := map[string]ConnectionType{
		"azuredataexplorer": ConnectionTypeAzureDataExplorer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectionType(input)
	return &out, nil
}

type TimeSeriesDatabaseConnectionState string

const (
	TimeSeriesDatabaseConnectionStateCanceled     TimeSeriesDatabaseConnectionState = "Canceled"
	TimeSeriesDatabaseConnectionStateDeleted      TimeSeriesDatabaseConnectionState = "Deleted"
	TimeSeriesDatabaseConnectionStateDeleting     TimeSeriesDatabaseConnectionState = "Deleting"
	TimeSeriesDatabaseConnectionStateDisabled     TimeSeriesDatabaseConnectionState = "Disabled"
	TimeSeriesDatabaseConnectionStateFailed       TimeSeriesDatabaseConnectionState = "Failed"
	TimeSeriesDatabaseConnectionStateMoving       TimeSeriesDatabaseConnectionState = "Moving"
	TimeSeriesDatabaseConnectionStateProvisioning TimeSeriesDatabaseConnectionState = "Provisioning"
	TimeSeriesDatabaseConnectionStateRestoring    TimeSeriesDatabaseConnectionState = "Restoring"
	TimeSeriesDatabaseConnectionStateSucceeded    TimeSeriesDatabaseConnectionState = "Succeeded"
	TimeSeriesDatabaseConnectionStateSuspending   TimeSeriesDatabaseConnectionState = "Suspending"
	TimeSeriesDatabaseConnectionStateUpdating     TimeSeriesDatabaseConnectionState = "Updating"
	TimeSeriesDatabaseConnectionStateWarning      TimeSeriesDatabaseConnectionState = "Warning"
)

func PossibleValuesForTimeSeriesDatabaseConnectionState() []string {
	return []string{
		string(TimeSeriesDatabaseConnectionStateCanceled),
		string(TimeSeriesDatabaseConnectionStateDeleted),
		string(TimeSeriesDatabaseConnectionStateDeleting),
		string(TimeSeriesDatabaseConnectionStateDisabled),
		string(TimeSeriesDatabaseConnectionStateFailed),
		string(TimeSeriesDatabaseConnectionStateMoving),
		string(TimeSeriesDatabaseConnectionStateProvisioning),
		string(TimeSeriesDatabaseConnectionStateRestoring),
		string(TimeSeriesDatabaseConnectionStateSucceeded),
		string(TimeSeriesDatabaseConnectionStateSuspending),
		string(TimeSeriesDatabaseConnectionStateUpdating),
		string(TimeSeriesDatabaseConnectionStateWarning),
	}
}

func (s *TimeSeriesDatabaseConnectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeSeriesDatabaseConnectionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeSeriesDatabaseConnectionState(input string) (*TimeSeriesDatabaseConnectionState, error) {
	vals := map[string]TimeSeriesDatabaseConnectionState{
		"canceled":     TimeSeriesDatabaseConnectionStateCanceled,
		"deleted":      TimeSeriesDatabaseConnectionStateDeleted,
		"deleting":     TimeSeriesDatabaseConnectionStateDeleting,
		"disabled":     TimeSeriesDatabaseConnectionStateDisabled,
		"failed":       TimeSeriesDatabaseConnectionStateFailed,
		"moving":       TimeSeriesDatabaseConnectionStateMoving,
		"provisioning": TimeSeriesDatabaseConnectionStateProvisioning,
		"restoring":    TimeSeriesDatabaseConnectionStateRestoring,
		"succeeded":    TimeSeriesDatabaseConnectionStateSucceeded,
		"suspending":   TimeSeriesDatabaseConnectionStateSuspending,
		"updating":     TimeSeriesDatabaseConnectionStateUpdating,
		"warning":      TimeSeriesDatabaseConnectionStateWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeSeriesDatabaseConnectionState(input)
	return &out, nil
}
