package basic

type IDType = uint64

// BaseInterface used only for internal purposes
type BaseInterface interface {
	GetID() IDType
	GetName() string
	SetName(string)
}

var baseNextID IDType = 1

// Base used only for internal purposes
// Base have to be initialized with MakeBase function
type Base struct {
	id   IDType
	name string
}

// MakeBase used only for internal purposes
func MakeBase() Base {
	b := Base{
		id:   baseNextID,
		name: "",
	}
	baseNextID++
	return b
}

func (b *Base) GetID() IDType    { return b.id }
func (b *Base) GetName() string  { return b.name }
func (b *Base) SetName(n string) { b.name = n }
