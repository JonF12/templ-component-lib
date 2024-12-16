package functions

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/JonF12/templ-component-lib/src/cms"
	"github.com/a-h/templ"
)

func RenderTemplComponent(componentName string, propsJson string) (templ.Component, error) {
	res, found := cms.GetComponent(componentName)
	if !found {
		return nil, errors.New("didnt find component")
	}
	propsValue := reflect.New(reflect.TypeOf(res.PropsType).Elem())
	newProps := propsValue.Interface()
	if err := json.Unmarshal([]byte(propsJson), newProps); err != nil {
		return nil, err
	}
	return res.Render(newProps), nil
}
