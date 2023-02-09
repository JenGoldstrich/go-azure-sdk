package datasetmapping

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataSetMapping = BlobFolderDataSetMapping{}

type BlobFolderDataSetMapping struct {
	Properties BlobFolderMappingProperties `json:"properties"`

	// Fields inherited from DataSetMapping
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

var _ json.Marshaler = BlobFolderDataSetMapping{}

func (s BlobFolderDataSetMapping) MarshalJSON() ([]byte, error) {
	type wrapper BlobFolderDataSetMapping
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BlobFolderDataSetMapping: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BlobFolderDataSetMapping: %+v", err)
	}
	decoded["kind"] = "BlobFolder"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BlobFolderDataSetMapping: %+v", err)
	}

	return encoded, nil
}
