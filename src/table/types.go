package table

import "github.com/a-h/templ"

type Formatter int

const (
  StringFormat Formatter = iota
  DateFormat
  NumberFormat
)

type Column struct {
  Header       string
  Field        string
  DefaultValue string
  IsAction     bool
  CustomHTML   func(interface{}) templ.Component
  Format       Formatter
}

type TableConfig struct {
  IdField string
}
