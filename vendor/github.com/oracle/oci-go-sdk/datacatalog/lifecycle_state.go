// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

// LifecycleStateEnum Enum with underlying type: string
type LifecycleStateEnum string

// Set of constants representing the allowable values for LifecycleStateEnum
const (
	LifecycleStateCreating LifecycleStateEnum = "CREATING"
	LifecycleStateActive   LifecycleStateEnum = "ACTIVE"
	LifecycleStateInactive LifecycleStateEnum = "INACTIVE"
	LifecycleStateUpdating LifecycleStateEnum = "UPDATING"
	LifecycleStateDeleting LifecycleStateEnum = "DELETING"
	LifecycleStateDeleted  LifecycleStateEnum = "DELETED"
	LifecycleStateFailed   LifecycleStateEnum = "FAILED"
	LifecycleStateMoving   LifecycleStateEnum = "MOVING"
)

var mappingLifecycleState = map[string]LifecycleStateEnum{
	"CREATING": LifecycleStateCreating,
	"ACTIVE":   LifecycleStateActive,
	"INACTIVE": LifecycleStateInactive,
	"UPDATING": LifecycleStateUpdating,
	"DELETING": LifecycleStateDeleting,
	"DELETED":  LifecycleStateDeleted,
	"FAILED":   LifecycleStateFailed,
	"MOVING":   LifecycleStateMoving,
}

// GetLifecycleStateEnumValues Enumerates the set of values for LifecycleStateEnum
func GetLifecycleStateEnumValues() []LifecycleStateEnum {
	values := make([]LifecycleStateEnum, 0)
	for _, v := range mappingLifecycleState {
		values = append(values, v)
	}
	return values
}
