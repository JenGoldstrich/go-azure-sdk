package blobcontainers

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Action string

const (
	ActionAcquire Action = "Acquire"
	ActionBreak   Action = "Break"
	ActionChange  Action = "Change"
	ActionRelease Action = "Release"
	ActionRenew   Action = "Renew"
)

func PossibleValuesForAction() []string {
	return []string{
		string(ActionAcquire),
		string(ActionBreak),
		string(ActionChange),
		string(ActionRelease),
		string(ActionRenew),
	}
}

func parseAction(input string) (*Action, error) {
	vals := map[string]Action{
		"acquire": ActionAcquire,
		"break":   ActionBreak,
		"change":  ActionChange,
		"release": ActionRelease,
		"renew":   ActionRenew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Action(input)
	return &out, nil
}

type ImmutabilityPolicyName string

const (
	ImmutabilityPolicyNameDefault ImmutabilityPolicyName = "default"
)

func PossibleValuesForImmutabilityPolicyName() []string {
	return []string{
		string(ImmutabilityPolicyNameDefault),
	}
}

func parseImmutabilityPolicyName(input string) (*ImmutabilityPolicyName, error) {
	vals := map[string]ImmutabilityPolicyName{
		"default": ImmutabilityPolicyNameDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImmutabilityPolicyName(input)
	return &out, nil
}

type ImmutabilityPolicyState string

const (
	ImmutabilityPolicyStateLocked   ImmutabilityPolicyState = "Locked"
	ImmutabilityPolicyStateUnlocked ImmutabilityPolicyState = "Unlocked"
)

func PossibleValuesForImmutabilityPolicyState() []string {
	return []string{
		string(ImmutabilityPolicyStateLocked),
		string(ImmutabilityPolicyStateUnlocked),
	}
}

func parseImmutabilityPolicyState(input string) (*ImmutabilityPolicyState, error) {
	vals := map[string]ImmutabilityPolicyState{
		"locked":   ImmutabilityPolicyStateLocked,
		"unlocked": ImmutabilityPolicyStateUnlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImmutabilityPolicyState(input)
	return &out, nil
}

type ImmutabilityPolicyUpdateType string

const (
	ImmutabilityPolicyUpdateTypeExtend ImmutabilityPolicyUpdateType = "extend"
	ImmutabilityPolicyUpdateTypeLock   ImmutabilityPolicyUpdateType = "lock"
	ImmutabilityPolicyUpdateTypePut    ImmutabilityPolicyUpdateType = "put"
)

func PossibleValuesForImmutabilityPolicyUpdateType() []string {
	return []string{
		string(ImmutabilityPolicyUpdateTypeExtend),
		string(ImmutabilityPolicyUpdateTypeLock),
		string(ImmutabilityPolicyUpdateTypePut),
	}
}

func parseImmutabilityPolicyUpdateType(input string) (*ImmutabilityPolicyUpdateType, error) {
	vals := map[string]ImmutabilityPolicyUpdateType{
		"extend": ImmutabilityPolicyUpdateTypeExtend,
		"lock":   ImmutabilityPolicyUpdateTypeLock,
		"put":    ImmutabilityPolicyUpdateTypePut,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImmutabilityPolicyUpdateType(input)
	return &out, nil
}

type LeaseDuration string

const (
	LeaseDurationFixed    LeaseDuration = "Fixed"
	LeaseDurationInfinite LeaseDuration = "Infinite"
)

func PossibleValuesForLeaseDuration() []string {
	return []string{
		string(LeaseDurationFixed),
		string(LeaseDurationInfinite),
	}
}

func parseLeaseDuration(input string) (*LeaseDuration, error) {
	vals := map[string]LeaseDuration{
		"fixed":    LeaseDurationFixed,
		"infinite": LeaseDurationInfinite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LeaseDuration(input)
	return &out, nil
}

type LeaseState string

const (
	LeaseStateAvailable LeaseState = "Available"
	LeaseStateBreaking  LeaseState = "Breaking"
	LeaseStateBroken    LeaseState = "Broken"
	LeaseStateExpired   LeaseState = "Expired"
	LeaseStateLeased    LeaseState = "Leased"
)

func PossibleValuesForLeaseState() []string {
	return []string{
		string(LeaseStateAvailable),
		string(LeaseStateBreaking),
		string(LeaseStateBroken),
		string(LeaseStateExpired),
		string(LeaseStateLeased),
	}
}

func parseLeaseState(input string) (*LeaseState, error) {
	vals := map[string]LeaseState{
		"available": LeaseStateAvailable,
		"breaking":  LeaseStateBreaking,
		"broken":    LeaseStateBroken,
		"expired":   LeaseStateExpired,
		"leased":    LeaseStateLeased,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LeaseState(input)
	return &out, nil
}

type LeaseStatus string

const (
	LeaseStatusLocked   LeaseStatus = "Locked"
	LeaseStatusUnlocked LeaseStatus = "Unlocked"
)

func PossibleValuesForLeaseStatus() []string {
	return []string{
		string(LeaseStatusLocked),
		string(LeaseStatusUnlocked),
	}
}

func parseLeaseStatus(input string) (*LeaseStatus, error) {
	vals := map[string]LeaseStatus{
		"locked":   LeaseStatusLocked,
		"unlocked": LeaseStatusUnlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LeaseStatus(input)
	return &out, nil
}

type ListContainersInclude string

const (
	ListContainersIncludeDeleted ListContainersInclude = "deleted"
)

func PossibleValuesForListContainersInclude() []string {
	return []string{
		string(ListContainersIncludeDeleted),
	}
}

func parseListContainersInclude(input string) (*ListContainersInclude, error) {
	vals := map[string]ListContainersInclude{
		"deleted": ListContainersIncludeDeleted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ListContainersInclude(input)
	return &out, nil
}

type MigrationState string

const (
	MigrationStateCompleted  MigrationState = "Completed"
	MigrationStateInProgress MigrationState = "InProgress"
)

func PossibleValuesForMigrationState() []string {
	return []string{
		string(MigrationStateCompleted),
		string(MigrationStateInProgress),
	}
}

func parseMigrationState(input string) (*MigrationState, error) {
	vals := map[string]MigrationState{
		"completed":  MigrationStateCompleted,
		"inprogress": MigrationStateInProgress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationState(input)
	return &out, nil
}

type PublicAccess string

const (
	PublicAccessBlob      PublicAccess = "Blob"
	PublicAccessContainer PublicAccess = "Container"
	PublicAccessNone      PublicAccess = "None"
)

func PossibleValuesForPublicAccess() []string {
	return []string{
		string(PublicAccessBlob),
		string(PublicAccessContainer),
		string(PublicAccessNone),
	}
}

func parsePublicAccess(input string) (*PublicAccess, error) {
	vals := map[string]PublicAccess{
		"blob":      PublicAccessBlob,
		"container": PublicAccessContainer,
		"none":      PublicAccessNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublicAccess(input)
	return &out, nil
}
