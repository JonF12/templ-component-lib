package types

import "encoding/json"

type PropsInterface interface {
	IsProps()
}

type BaseProps struct {
	NamePrefix string
	Subheading string
	ClassName  string
}

func (b BaseProps) ToJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
func (b BaseProps) ToJSONString(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (BaseProps) IsProps() bool { return true }
