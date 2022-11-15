package query

// TypeRef is used internally to set constraint for accepted types.
type typeRef interface {
	int64 | string
}

type Column interface {
	CheckValue() bool
}

// Column defines a data field in schema, including its type and
// value.
type TypedColumn[T typeRef] struct {
	name  string
	value T
}

func (TypedColumn[T]) CheckValue() bool {
	return true
}

func STRING(name string) TypedColumn[string] {
	return TypedColumn[string]{
		name: name,
	}
}

func INT64(name string) TypedColumn[int64] {
	return TypedColumn[int64]{
		name: name,
	}
}
