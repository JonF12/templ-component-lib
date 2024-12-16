package cms

import (
	"github.com/JonF12/templ-component-lib/dist/src/cards/cardimage"
	"github.com/JonF12/templ-component-lib/src/address"
	"github.com/JonF12/templ-component-lib/src/cards/card3spot"
	"github.com/JonF12/templ-component-lib/src/checkbox"
	"github.com/JonF12/templ-component-lib/src/dateinput"
	"github.com/JonF12/templ-component-lib/src/dropzone"
	"github.com/JonF12/templ-component-lib/src/emailsignup"
	layoutgrid "github.com/JonF12/templ-component-lib/src/layout/grid"
	layoutsection "github.com/JonF12/templ-component-lib/src/layout/section"
	"github.com/JonF12/templ-component-lib/src/searchselect"
	"github.com/a-h/templ"
)

var componentMap = map[string]Component{
	"address": {
		Category:      "address",
		ComponentName: "address",
		DisplayText:   "Address input component with validation",
		PropsType:     &address.Props{},
		Render: func(p any) templ.Component {
			return address.Render(p.(*address.Props))
		},
	},

	"cards/card3spot": {
		Category:      "cards",
		ComponentName: "card3spot",
		DisplayText:   "Three-spot card layout component",
		PropsType:     &card3spot.Props{},
		Render: func(p any) templ.Component {
			return card3spot.Render(p.(*card3spot.Props))
		},
	},
	"cards/cardimage": {
		Category:      "cards",
		ComponentName: "cardimage",
		DisplayText:   "Card with image component",
		PropsType:     &cardimage.Props{},
		Render: func(p any) templ.Component {
			return cardimage.Render(p.(*cardimage.Props))
		},
	},
	"checkbox": {
		Category:      "checkbox",
		ComponentName: "checkbox",
		DisplayText:   "Customizable checkbox input",
		PropsType:     &checkbox.Props{},
		Render: func(p any) templ.Component {
			return checkbox.Render(p.(*checkbox.Props))
		},
	},
	"dateinput": {
		Category:      "dateinput",
		ComponentName: "dateinput",
		DisplayText:   "Date input with validation",
		PropsType:     &dateinput.Props{},
		Render: func(p any) templ.Component {
			return dateinput.Render(p.(*dateinput.Props))
		},
	},
	"dropzone": {
		Category:      "dropzone",
		ComponentName: "dropzone",
		DisplayText:   "File upload dropzone",
		PropsType:     &dropzone.Props{},
		Render: func(p any) templ.Component {
			return dropzone.Render(p.(*dropzone.Props))
		},
	},
	"emailsignup": {
		Category:      "emailsignup",
		ComponentName: "emailsignup",
		DisplayText:   "Email signup form component",
		PropsType:     &emailsignup.Props{},
		Render: func(p any) templ.Component {
			return emailsignup.Render(p.(*emailsignup.Props))
		},
	},
	"layout/grid": {
		Category:      "layout",
		ComponentName: "grid",
		DisplayText:   "Grid layout component",
		PropsType:     &layoutgrid.Props{},
		Render: func(p any) templ.Component {
			return layoutgrid.Render(p.(*layoutgrid.Props))
		},
	},
	"layout/section": {
		Category:      "layout",
		ComponentName: "section",
		DisplayText:   "Section layout component",
		PropsType:     &layoutsection.Props{},
		Render: func(p any) templ.Component {
			return layoutsection.Render(p.(*layoutsection.Props))
		},
	},
	"searchselect": {
		Category:      "searchselect",
		ComponentName: "searchselect",
		DisplayText:   "Searchable select component",
		PropsType:     &searchselect.Props{},
		Render: func(p any) templ.Component {
			return searchselect.Render(p.(*searchselect.Props))
		},
	},
}
