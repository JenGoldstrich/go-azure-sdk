package virtualmachineimages

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVGenerationTypes string

const (
	HyperVGenerationTypesVOne HyperVGenerationTypes = "V1"
	HyperVGenerationTypesVTwo HyperVGenerationTypes = "V2"
)

func PossibleValuesForHyperVGenerationTypes() []string {
	return []string{
		string(HyperVGenerationTypesVOne),
		string(HyperVGenerationTypesVTwo),
	}
}

func parseHyperVGenerationTypes(input string) (*HyperVGenerationTypes, error) {
	vals := map[string]HyperVGenerationTypes{
		"v1": HyperVGenerationTypesVOne,
		"v2": HyperVGenerationTypesVTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HyperVGenerationTypes(input)
	return &out, nil
}

type OperatingSystemTypes string

const (
	OperatingSystemTypesLinux   OperatingSystemTypes = "Linux"
	OperatingSystemTypesWindows OperatingSystemTypes = "Windows"
)

func PossibleValuesForOperatingSystemTypes() []string {
	return []string{
		string(OperatingSystemTypesLinux),
		string(OperatingSystemTypesWindows),
	}
}

func parseOperatingSystemTypes(input string) (*OperatingSystemTypes, error) {
	vals := map[string]OperatingSystemTypes{
		"linux":   OperatingSystemTypesLinux,
		"windows": OperatingSystemTypesWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperatingSystemTypes(input)
	return &out, nil
}

type VMDiskTypes string

const (
	VMDiskTypesNone      VMDiskTypes = "None"
	VMDiskTypesUnmanaged VMDiskTypes = "Unmanaged"
)

func PossibleValuesForVMDiskTypes() []string {
	return []string{
		string(VMDiskTypesNone),
		string(VMDiskTypesUnmanaged),
	}
}

func parseVMDiskTypes(input string) (*VMDiskTypes, error) {
	vals := map[string]VMDiskTypes{
		"none":      VMDiskTypesNone,
		"unmanaged": VMDiskTypesUnmanaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VMDiskTypes(input)
	return &out, nil
}
