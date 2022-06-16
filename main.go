package newtype

func NewSomething() *Something {
	return &Something{}
}

type Something struct{}

func (s *Something) AMethod(_ AType) AnotherType {
	return AnotherType{}
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
