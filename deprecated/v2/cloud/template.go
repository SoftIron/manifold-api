package cloud

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
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

	switch v := value.(type) {
	case string:
		return fmt.Sprintf("%s%s = %q%s", strings.Repeat("\t", indent), pad(key, ljust), v, delim)
	case map[string]any:
		templates = []map[string]any{v}
	case []map[string]any:
		templates = v
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
