package openaitype

import "encoding/json"

type StringOrArray []string

func (soa *StringOrArray) MarshalJSON() ([]byte, error) {
	if len(*soa) == 0 {
		return []byte("null"), nil
	} else if len(*soa) == 1 {
		return json.Marshal((*soa)[0])
	} else {
		return json.Marshal([]string(*soa))
	}
}

func (soa *StringOrArray) UnmarshalJSON(data []byte) error {
	var s string
	var ss []string

	err := json.Unmarshal(data, &s)
	if err == nil {
		*soa = []string{s}
		return nil
	}

	err = json.Unmarshal(data, &ss)
	if err == nil {
		*soa = StringOrArray(ss)
		return nil
	}

	return err
}
