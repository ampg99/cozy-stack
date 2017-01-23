package permissions

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

// Set is a Set of rule
type Set []Rule

// MarshalJSON implements json.Marshaller on Set
// see docs/permission for structure
func (ps Set) MarshalJSON() ([]byte, error) {

	m := make(map[string]*json.RawMessage)

	for i, r := range ps {
		b, err := json.Marshal(r)
		if err != nil {
			return nil, err
		}
		rm := json.RawMessage(b)
		key := r.Title
		if key == "" {
			key = "rule" + strconv.Itoa(i)
		}
		m[key] = &rm
	}

	return json.Marshal(m)
}

// MarshalScopeString transforms a Set into a string for Oauth Scope
// (a space separated concatenation of its rules)
func (ps Set) MarshalScopeString() (string, error) {
	out := ""
	if len(ps) == 0 {
		return "", nil
	}
	for _, r := range ps {
		s, err := r.MarshalScopeString()
		if err != nil {
			return "", err
		}
		out += " " + s
	}
	return out[1:], nil
}

// UnmarshalJSON parses a json formated permission set
func (ps *Set) UnmarshalJSON(j []byte) error {

	var m map[string]*json.RawMessage
	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	for title, rulejson := range m {
		var r Rule
		err := json.Unmarshal(*rulejson, &r)
		if err != nil {
			return err
		}
		r.Title = title
		*ps = append(*ps, r)
	}

	return nil
}

// UnmarshalScopeString parse a Scope string into a permission Set
func UnmarshalScopeString(in string) (Set, error) {
	parts := strings.Split(in, ruleSep)
	out := make(Set, len(parts))

	if len(parts) == 0 {
		return nil, errors.New("Empty scope string")
	}

	for i, p := range parts {
		s, err := UnmarshalRuleString(p)
		if err != nil {
			return nil, err
		}
		out[i] = s
	}

	return out, nil
}