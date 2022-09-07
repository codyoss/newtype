package newtypepb

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
