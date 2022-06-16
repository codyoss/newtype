package newtype

import "github.com/codyoss/orgtype"

func NewSomething() *Something {
	return &Something{}
}

type Something struct{}

func (s *Something) AMethod(_ *orgtype.AType) orgtype.AnotherType {
	return orgtype.AnotherType{}
}
