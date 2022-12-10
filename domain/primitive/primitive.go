package primitive

// Primitive interface.
type Plimitive[T any] interface {
	New(T) (Plimitive[T], error)
	Validate(T) error
	Value() T
	String() string
	Reconstruct(T) Plimitive[T]
}
