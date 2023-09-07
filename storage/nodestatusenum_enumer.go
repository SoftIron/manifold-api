// Code generated by "enumer -type NodeStatusEnum -linecomment -text"; DO NOT EDIT.

package storage

import (
	"fmt"
	"strings"
)

const _NodeStatusEnumName = "DISCOVERINGMISSINGOFFCRASHEDDOWNONPENDING_ONPENDING_OFFPENDING_REBOOTPENDING_CHANGE"

var _NodeStatusEnumIndex = [...]uint8{0, 11, 18, 21, 28, 32, 34, 44, 55, 69, 83}

const _NodeStatusEnumLowerName = "discoveringmissingoffcrasheddownonpending_onpending_offpending_rebootpending_change"

func (i NodeStatusEnum) String() string {
	if i < 0 || i >= NodeStatusEnum(len(_NodeStatusEnumIndex)-1) {
		return fmt.Sprintf("NodeStatusEnum(%d)", i)
	}
	return _NodeStatusEnumName[_NodeStatusEnumIndex[i]:_NodeStatusEnumIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _NodeStatusEnumNoOp() {
	var x [1]struct{}
	_ = x[NodeStatusDiscovering-(0)]
	_ = x[NodeStatusMissing-(1)]
	_ = x[NodeStatusOff-(2)]
	_ = x[NodeStatusCrashed-(3)]
	_ = x[NodeStatusBooting-(4)]
	_ = x[NodeStatusOn-(5)]
	_ = x[NodeStatusPendingOn-(6)]
	_ = x[NodeStatusPendingOff-(7)]
	_ = x[NodeStatusPendingReboot-(8)]
	_ = x[NodeStatusPendingChange-(9)]
}

var _NodeStatusEnumValues = []NodeStatusEnum{NodeStatusDiscovering, NodeStatusMissing, NodeStatusOff, NodeStatusCrashed, NodeStatusBooting, NodeStatusOn, NodeStatusPendingOn, NodeStatusPendingOff, NodeStatusPendingReboot, NodeStatusPendingChange}

var _NodeStatusEnumNameToValueMap = map[string]NodeStatusEnum{
	_NodeStatusEnumName[0:11]:       NodeStatusDiscovering,
	_NodeStatusEnumLowerName[0:11]:  NodeStatusDiscovering,
	_NodeStatusEnumName[11:18]:      NodeStatusMissing,
	_NodeStatusEnumLowerName[11:18]: NodeStatusMissing,
	_NodeStatusEnumName[18:21]:      NodeStatusOff,
	_NodeStatusEnumLowerName[18:21]: NodeStatusOff,
	_NodeStatusEnumName[21:28]:      NodeStatusCrashed,
	_NodeStatusEnumLowerName[21:28]: NodeStatusCrashed,
	_NodeStatusEnumName[28:32]:      NodeStatusBooting,
	_NodeStatusEnumLowerName[28:32]: NodeStatusBooting,
	_NodeStatusEnumName[32:34]:      NodeStatusOn,
	_NodeStatusEnumLowerName[32:34]: NodeStatusOn,
	_NodeStatusEnumName[34:44]:      NodeStatusPendingOn,
	_NodeStatusEnumLowerName[34:44]: NodeStatusPendingOn,
	_NodeStatusEnumName[44:55]:      NodeStatusPendingOff,
	_NodeStatusEnumLowerName[44:55]: NodeStatusPendingOff,
	_NodeStatusEnumName[55:69]:      NodeStatusPendingReboot,
	_NodeStatusEnumLowerName[55:69]: NodeStatusPendingReboot,
	_NodeStatusEnumName[69:83]:      NodeStatusPendingChange,
	_NodeStatusEnumLowerName[69:83]: NodeStatusPendingChange,
}

var _NodeStatusEnumNames = []string{
	_NodeStatusEnumName[0:11],
	_NodeStatusEnumName[11:18],
	_NodeStatusEnumName[18:21],
	_NodeStatusEnumName[21:28],
	_NodeStatusEnumName[28:32],
	_NodeStatusEnumName[32:34],
	_NodeStatusEnumName[34:44],
	_NodeStatusEnumName[44:55],
	_NodeStatusEnumName[55:69],
	_NodeStatusEnumName[69:83],
}

// NodeStatusEnumString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NodeStatusEnumString(s string) (NodeStatusEnum, error) {
	if val, ok := _NodeStatusEnumNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _NodeStatusEnumNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to NodeStatusEnum values", s)
}

// NodeStatusEnumValues returns all values of the enum
func NodeStatusEnumValues() []NodeStatusEnum {
	return _NodeStatusEnumValues
}

// NodeStatusEnumStrings returns a slice of all String values of the enum
func NodeStatusEnumStrings() []string {
	strs := make([]string, len(_NodeStatusEnumNames))
	copy(strs, _NodeStatusEnumNames)
	return strs
}

// IsANodeStatusEnum returns "true" if the value is listed in the enum definition. "false" otherwise
func (i NodeStatusEnum) IsANodeStatusEnum() bool {
	for _, v := range _NodeStatusEnumValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for NodeStatusEnum
func (i NodeStatusEnum) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for NodeStatusEnum
func (i *NodeStatusEnum) UnmarshalText(text []byte) error {
	var err error
	*i, err = NodeStatusEnumString(string(text))
	return err
}