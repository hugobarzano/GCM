package api

// Nature define api nature
type Nature int32

const (
	// Base defines nothing
	Base Nature = iota
	// Go defines golang generator nature
	Go
	// Python
	Python
	// JS
	JS
)

// Generator interface definition
type Generator interface {
	Init() Generator
	WithName(name string) Generator
	WithPort(port int) Generator
	WithInputSpec(spec string) Generator
	GenerateApi(des string)
	GetFiles() map[string][]byte
}

type defaultNature struct{}

func (g *defaultNature) Init() Generator                     { return g }
func (g *defaultNature) WithName(name string) Generator      { return g }
func (g *defaultNature) WithPort(port int) Generator         { return g }
func (g *defaultNature) WithInputSpec(spec string) Generator { return g }
func (g *defaultNature) GenerateApi(des string)              {}
func (g *defaultNature) GetFiles() map[string][]byte         { return nil }

var DefaultNature = &defaultNature{}

type customNature struct {
	name  string
	port  int
	spec  []byte
	files map[string][]byte
	model map[string]interface{}
}

// New builds a generator from a given kind
func New(nature Nature) Generator {
	switch nature {
	case Go:
		return &goApi{}
	case Python:
		return &python{}
	case JS:
		return &javascript{}
	default:
		return DefaultNature
	}
}
