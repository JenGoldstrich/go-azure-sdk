package servergroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NameAvailability struct {
	Message       *string `json:"message,omitempty"`
	Name          *string `json:"name,omitempty"`
	NameAvailable *bool   `json:"nameAvailable,omitempty"`
	Type          *string `json:"type,omitempty"`
}
