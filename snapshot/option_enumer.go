// Code generated by "enumer -type Option -transform title-upper -text"; DO NOT EDIT.

package snapshot

import (
	"fmt"
	"strings"
)

const _OptionName = "PauseVM"

var _OptionIndex = [...]uint8{0, 7}

const _OptionLowerName = "pausevm"

func (i Option) String() string {
	i -= 1
	if i < 0 || i >= Option(len(_OptionIndex)-1) {
		return fmt.Sprintf("Option(%d)", i+1)
	}
	return _OptionName[_OptionIndex[i]:_OptionIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _OptionNoOp() {
	var x [1]struct{}
	_ = x[PauseVM-(1)]
}

var _OptionValues = []Option{PauseVM}

var _OptionNameToValueMap = map[string]Option{
	_OptionName[0:7]:      PauseVM,
	_OptionLowerName[0:7]: PauseVM,
}

var _OptionNames = []string{
	_OptionName[0:7],
}

// OptionString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func OptionString(s string) (Option, error) {
	if val, ok := _OptionNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _OptionNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Option values", s)
}

// OptionValues returns all values of the enum
func OptionValues() []Option {
	return _OptionValues
}

// OptionStrings returns a slice of all String values of the enum
func OptionStrings() []string {
	strs := make([]string, len(_OptionNames))
	copy(strs, _OptionNames)
	return strs
}

// IsAOption returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Option) IsAOption() bool {
	for _, v := range _OptionValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for Option
func (i Option) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Option
func (i *Option) UnmarshalText(text []byte) error {
	var err error
	*i, err = OptionString(string(text))
	return err
}
