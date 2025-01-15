package iterator

import "iter"

type Iterator[E any] interface {
	Iterate() iter.Seq2[int, E]
}
