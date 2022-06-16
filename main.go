package newtype

import "github.com/codyoss/orgtype"

func NewSomething() *Something {
	return &Something{}
}

type Something struct{}

func (s *Something) AMethod(_ *orgtype.AType) orgtype.AnotherType {
	return orgtype.AnotherType{}
}

var AVar bool

const AConst = 1

func AFunc() {}

func AnotherFunc(*AType) *AnotherType {
	return nil
}

type AType struct {
	AField bool
}

type AnotherType struct {
	AField int
}

type AInterface interface {
	AMethod(*AType) AnotherType
}
