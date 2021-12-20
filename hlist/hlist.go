package hlist

type HList interface {
	sealed()
	// Cons[_,_] | Nil
}

type Header[HT any] interface {
	HList
	Head() HT
}

type Cons[H any, T HList] interface {
	HList
	Head() H
	Tail() T
}

type Nil struct {
}

func (r Nil) Head() Nil {
	return r
}

func (r Nil) Tail() Nil {
	return r
}

func (r Nil) sealed() {

}

type hlistImpl[H any, T HList] struct {
	head H
	tail T
}

func (r hlistImpl[H, T]) Head() H {
	return r.head
}

func (r hlistImpl[H, T]) Tail() T {
	return r.tail
}

func (r hlistImpl[H, T]) sealed() {

}

func Concat[H any, T HList](h H, t T) Cons[H, T] {
	return hlistImpl[H, T]{h, t}
}

func Empty() Nil {
	return Nil{}
}
