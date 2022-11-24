package replicationprotecteditems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DisableProtectionProviderSpecificInput = InMageDisableProtectionProviderSpecificInput{}

type InMageDisableProtectionProviderSpecificInput struct {
	ReplicaVmDeletionStatus *string `json:"replicaVmDeletionStatus,omitempty"`

	// Fields inherited from DisableProtectionProviderSpecificInput
}

var _ json.Marshaler = InMageDisableProtectionProviderSpecificInput{}

func (s InMageDisableProtectionProviderSpecificInput) MarshalJSON() ([]byte, error) {
	type wrapper InMageDisableProtectionProviderSpecificInput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InMageDisableProtectionProviderSpecificInput: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InMageDisableProtectionProviderSpecificInput: %+v", err)
	}
	decoded["instanceType"] = "InMage"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InMageDisableProtectionProviderSpecificInput: %+v", err)
	}

	return encoded, nil
}
