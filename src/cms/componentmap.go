package cms

import (
	"maps"
	"slices"

	"github.com/a-h/templ"
)

type Component struct {
	Category      string
	ComponentName string
	DisplayText   string
	PropsType     any
	Render        func(any) templ.Component
}

func GetComponent(key string) (Component, bool) {
	comp, ok := componentMap[key]
	return comp, ok
}

func GetAllComponents() []Component {
	return slices.Collect(maps.Values(componentMap))
}
