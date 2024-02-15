package cloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

// Template is nested map of string key x value pairs.
type Template map[string]any

// String implements the Stringer interface for t. Example output:
//
//	FOO = "bar"
//	BAZ = [
//	 QUX = "quux",
//	 CORGE = "grault"
//	]
//	BAZ = [
//	 GARPLY = "waldo"
//	]
//	GARPLY = "waldo"
func (t Template) String() string {
	var buf strings.Builder

	kps := make([]keyValuePair, 0, len(t))
	length := 0

	for key := range t {
		if len(key) > length {
			length = len(key)
		}
		kps = append(kps, keyValuePair{key, t[key]})
	}

	sort.Slice(kps, func(i, j int) bool {
		return kps[i].key < kps[j].key
	})

	for i := range kps {
		buf.WriteString(templateString(0, length, "\n", pad(kps[i].key, length), kps[i].value))
	}

	return buf.String()
}

// UnmarshalJSON implements the json.Unmarshaler interface for t.
func (t *Template) UnmarshalJSON(b []byte) error {
	m := map[string]any(*t)
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	// Cleanup nested []any types to more concrete []map[string]any types. Take
	// advantage of the known structure of the template.
	var cleanup func(map[string]any) error

	cleanup = func(m map[string]any) error {
		for key, value := range m {
			switch v := value.(type) {
			case []any:
				list := make([]map[string]any, 0)

				for _, m := range v {
					v, ok := m.(map[string]any)
					if !ok {
						return fmt.Errorf("expected map[string]any, got %T", m)
					}
					list = append(list, v)
				}

				m[key] = list
			case map[string]any:
				if err := cleanup(v); err != nil {
					return err
				}
			}
		}
		return nil
	}

	if err := cleanup(m); err != nil {
		return err
	}

	*t = Template(m)

	return nil
}

type keyValuePair struct {
	key   string
	value any
}

func pad(s string, length int) string {
	return fmt.Sprintf("%s%s", s, strings.Repeat(" ", length-len(s)))
}

func templateString(indent, ljust int, delim, key string, value any) string {
	var templates []map[string]any

	lhs := fmt.Sprintf("%s%s= ", strings.Repeat("\t", indent), pad(key, ljust))

	switch v := value.(type) {
	case map[string]any:
		templates = []map[string]any{v}
	case []map[string]any:
		templates = v
	case string, fmt.Stringer:
		return lhs + fmt.Sprintf("%q", v) + delim
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return lhs + fmt.Sprintf("%v", v) + delim
	case bool:
		if v {
			return lhs + "YES" + delim
		}
		return lhs + "NO" + delim
	default:
		return lhs + fmt.Sprintf("\"unknown value type: %v\"", v) + delim
	}

	var buf strings.Builder

	for _, template := range templates {

		buf.WriteString(fmt.Sprintf("%s%s = [\n", strings.Repeat("\t", indent), pad(key, ljust)))

		kps := make([]keyValuePair, 0, len(template))
		length := 0

		for key := range template {
			if len(key) > length {
				length = len(key)
			}
			kps = append(kps, keyValuePair{key, template[key]})
		}

		sort.Slice(kps, func(i, j int) bool {
			return kps[i].key < kps[j].key
		})

		delim = ",\n"
		for i := range kps {
			if i == len(kps)-1 {
				delim = "\n"
			}
			buf.WriteString(templateString(indent+1, length, delim, pad(kps[i].key, length), kps[i].value))
		}

		buf.WriteString(fmt.Sprintf("%s]\n", strings.Repeat("\t", indent)))

	}

	return buf.String()
}

// NewTemplate returns a new Template from the struct t. If t is not a struct
// (or a pointer to one) it panics.
func NewTemplate(t any) Template {
	rv := reflect.ValueOf(t)

	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		panic("NewTemplate: expected struct, got " + rv.Kind().String())
	}

	return newTemplate(rv.Interface())
}

// newTemplate returns a new map[string]any from the struct t. This exists so we
// can recursively call it and always return a map[string]any. And only deal
// with the Template type alias at the top level. This is done to simplify the
// constructed nested map, which keeps the Template String method "simple".
//
// This is not intended to be general purpose and only needs to handle types the
// are used in the Template structs (which we create).
func newTemplate(t any) map[string]any { //nolint:gocognit
	template := make(map[string]any)
	rv := reflect.ValueOf(t)
	rt := reflect.TypeOf(t)

	for i := 0; i < rv.NumField(); i++ {
		var options tagOptions

		field := rt.Field(i)
		fieldName := camelToSnake(field.Name)

		if tag := field.Tag.Get("template"); tag != "" {
			fieldName, options = parseTag(tag)
		}

		fv := rv.Field(i)

		switch field.Type.Kind() {
		case reflect.Struct:
			m := newTemplate(rv.Field(i).Interface())
			if len(m) > 0 || options.Contains(always) {
				template[fieldName] = newTemplate(rv.Field(i).Interface())
			}
		case reflect.Slice, reflect.Array:
			if fv.Len() == 0 && !options.Contains(always) {
				continue
			}

			m := make([]map[string]any, fv.Len())
			for i := 0; i < fv.Len(); i++ {
				m[i] = newTemplate(fv.Index(i).Interface())
			}
			template[fieldName] = m
		case reflect.Pointer:
			if fv.IsNil() { // only handles pointers to structs or simple types
				continue
			}

			fv = fv.Elem()
			if fv.Kind() == reflect.Struct {
				template[fieldName] = newTemplate(fv.Interface())
				continue
			}

			fallthrough

		default:
			if fv.IsZero() && !options.Contains(always) {
				continue
			}

			v := simpleValue(fv)

			if options.Contains("string") { // not sure about this, might not be needed or might want for other types
				template[fieldName] = fmt.Sprint(v)
				continue
			}

			template[fieldName] = v
		}
	}

	return template
}

func simpleValue(v reflect.Value) any {
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Bool:
		return v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint()
	case reflect.Float32, reflect.Float64:
		return v.Float()
	}

	return "<nil>"
}

const always = "always" // opposite of omitempty, default is to omit empty fields

// camelToSnake converts a camelCase string to SNAKE_CASE.
func camelToSnake(s string) string {
	var buffer bytes.Buffer

	for i, char := range s {
		if i > 0 && unicode.IsUpper(char) {
			var next rune

			if i+1 < len(s) {
				next = rune(s[i+1])
			}
			prev := rune(s[i-1])

			if (unicode.IsDigit(prev) || unicode.IsLower(prev)) || unicode.IsLower(next) {
				buffer.WriteRune('_')
			}
		}
		buffer.WriteRune(unicode.ToUpper(char))
	}

	return buffer.String()
}

// The following code is from the Go standard library and is licensed under a
// BSD-style license. Specifically, this is from the encoding/json package.

// tagOptions is the string following a comma in a struct field's "json" tag, or
// the empty string. It does not include the leading comma.
type tagOptions string

// parseTag splits a struct field's json tag into its name and comma-separated
// options.
func parseTag(tag string) (string, tagOptions) {
	tag, opt, _ := strings.Cut(tag, ",")
	return tag, tagOptions(opt)
}

// Contains reports whether a comma-separated list of options contains a
// particular substr flag. substr must be surrounded by a string boundary or
// commas.
func (o tagOptions) Contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}
	s := string(o)
	for s != "" {
		var name string
		name, s, _ = strings.Cut(s, ",")
		if name == optionName {
			return true
		}
	}
	return false
}
