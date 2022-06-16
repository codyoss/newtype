package newtype

import (
	"github.com/codyoss/newtype/pb"
)

func NewSomething() *Something {
	return &Something{}
}

type Something struct{}

func (s *Something) AMethod(_ *pb.AType) pb.AnotherType {
	return pb.AnotherType{}
}
