package newtype

import "github.com/codyoss/newtype/newtypepb"

func NewSomething() *Something {
	return &Something{}
}

type Something struct{}

func (s *Something) AMethod(_ *newtypepb.AType) newtypepb.AnotherType {
	return newtypepb.AnotherType{}
}
