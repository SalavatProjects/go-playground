package calculator

type CalculateMethod int

const (
	Add CalculateMethod = iota
	Subtract
	Multiply
	Divide
)

func Calculate(method CalculateMethod) func(int, int) int {
	switch method {
	case Add:
		return func(i int, i2 int) int {
			return i + i2
		}
	case Subtract:
		return func(i int, i2 int) int {
			return i - i2
		}
	case Multiply:
		return func(i int, i2 int) int {
			return i * i2
		}
	case Divide:
		return func(i int, i2 int) int {
			return i / i2
		}
	default:
		panic("unknown calculate method")
	}
}
