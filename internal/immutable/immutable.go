package immutable

type MyImmutable[T any] struct {
	val T
}

func NewMyImmutable[T any](val T) MyImmutable[T] {
	return MyImmutable[T]{ val: val }
}

func (im MyImmutable[T]) Value() T {
	return im.val
}