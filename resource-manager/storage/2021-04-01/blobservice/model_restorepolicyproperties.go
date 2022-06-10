package blobservice

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorePolicyProperties struct {
	Days            *int64  `json:"days,omitempty"`
	Enabled         bool    `json:"enabled"`
	LastEnabledTime *string `json:"lastEnabledTime,omitempty"`
	MinRestoreTime  *string `json:"minRestoreTime,omitempty"`
}

func (o *RestorePolicyProperties) GetLastEnabledTimeAsTime() (*time.Time, error) {
	if o.LastEnabledTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastEnabledTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RestorePolicyProperties) SetLastEnabledTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastEnabledTime = &formatted
}

func (o *RestorePolicyProperties) GetMinRestoreTimeAsTime() (*time.Time, error) {
	if o.MinRestoreTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.MinRestoreTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RestorePolicyProperties) SetMinRestoreTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.MinRestoreTime = &formatted
}
