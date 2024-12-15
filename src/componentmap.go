package src

import (
	"github.com/JonF12/templ-component-lib/address"
	"github.com/JonF12/templ-component-lib/article"
	"github.com/JonF12/templ-component-lib/cards"
	"github.com/JonF12/templ-component-lib/checkbox"
	"github.com/JonF12/templ-component-lib/dateinput"
	"github.com/JonF12/templ-component-lib/dropzone"
	"github.com/JonF12/templ-component-lib/emailsignup"
	"github.com/JonF12/templ-component-lib/layout"
	"github.com/JonF12/templ-component-lib/searchselect"
	"github.com/a-h/templ"
)

type Component[T any] struct {
	Category      string
	ComponentName string
	DisplayText   string
	PropsType     T
	Render        func(T) templ.Component
}

var ComponentMap = map[string]any{
	"address": Component[address.Props]{
		Category:      "address",
		ComponentName: "address",
		DisplayText:   "Address input component with validation",
		PropsType:     address.Props{},
		Render:        address.Render,
	},
	"article/heading": Component[article.Props]{
		Category:      "article",
		ComponentName: "heading",
		DisplayText:   "Article heading with optional subheading",
		PropsType:     article.Props{},
		Render:        article.Render,
	},
	"cards/card3spot": Component[cards.Props]{
		Category:      "cards",
		ComponentName: "card3spot",
		DisplayText:   "Three-spot card layout component",
		PropsType:     cards.Props{},
		Render:        cards.Render,
	},
	"cards/cardimage": Component[cards.Props]{
		Category:      "cards",
		ComponentName: "cardimage",
		DisplayText:   "Card with image component",
		PropsType:     cards.Props{},
		Render:        cards.Render,
	},
	"checkbox": Component[checkbox.Props]{
		Category:      "checkbox",
		ComponentName: "checkbox",
		DisplayText:   "Customizable checkbox input",
		PropsType:     checkbox.Props{},
		Render:        checkbox.Render,
	},
	"dateinput": Component[dateinput.Props]{
		Category:      "dateinput",
		ComponentName: "dateinput",
		DisplayText:   "Date input with validation",
		PropsType:     dateinput.Props{},
		Render:        dateinput.Render,
	},
	"dropzone": Component[dropzone.Props]{
		Category:      "dropzone",
		ComponentName: "dropzone",
		DisplayText:   "File upload dropzone",
		PropsType:     dropzone.Props{},
		Render:        dropzone.Render,
	},
	"emailsignup": Component[emailsignup.Props]{
		Category:      "emailsignup",
		ComponentName: "emailsignup",
		DisplayText:   "Email signup form component",
		PropsType:     emailsignup.Props{},
		Render:        emailsignup.Render,
	},
	"layout/grid": Component[layout.Props]{
		Category:      "layout",
		ComponentName: "grid",
		DisplayText:   "Grid layout component",
		PropsType:     layout.Props{},
		Render:        layout.Render,
	},
	"layout/section": Component[layout.Props]{
		Category:      "layout",
		ComponentName: "section",
		DisplayText:   "Section layout component",
		PropsType:     layout.Props{},
		Render:        layout.Render,
	},
	"searchselect": Component[searchselect.Props]{
		Category:      "searchselect",
		ComponentName: "searchselect",
		DisplayText:   "Searchable select component",
		PropsType:     searchselect.Props{},
		Render:        searchselect.Render,
	},
}

// Helper function to get a typed component
func GetComponent[T any](key string) (Component[T], bool) {
	if comp, ok := ComponentMap[key]; ok {
		if typed, ok := comp.(Component[T]); ok {
			return typed, true
		}
	}
	return Component[T]{}, false
}
