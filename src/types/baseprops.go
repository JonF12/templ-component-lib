package types

type PropsInterface interface {
	IsProps()
}

type BaseProps struct {
	NamePrefix string
	Subheading string
	ClassName  string
}

func (BaseProps) IsProps() bool { return true }
