package serverazureadadministrators

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdministratorProperties struct {
	AdministratorType AdministratorType `json:"administratorType"`
	Login             string            `json:"login"`
	Sid               string            `json:"sid"`
	TenantId          *string           `json:"tenantId,omitempty"`
}
