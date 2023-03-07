package next

import "context"

type Worker[T any] struct {
	name    string
	input   Input[T]
	output  Output[T]
	filters []EventFilter[T]
}

func (w *Worker[T]) Name() string {
	return w.name
}

func (w *Worker[T]) SetInput(input Input[T]) {
	w.input = input
}

func (w *Worker[T]) SetOutput(output Output[T]) {
	w.output = output
}

func (w *Worker[T]) AddFilter(filter EventFilter[T]) {
	w.filters = append(w.filters, filter)
}

func (w *Worker[T]) Run(ctx context.Context) error {
	next := FilterFunc[T](func(context context.Context, event Event[T]) error {
		return w.output.OnWrite(context, []Event[T]{event})
	})
	sink := EventSink[T](func(context context.Context, event Event[T]) error {
		return makeFilterChain(next, w.filters)(context, event)
	})
	return w.input.OnRead(ctx, sink)
}

func makeFilterChain[T any](next FilterFunc[T], filters []EventFilter[T]) FilterFunc[T] {
	for i := len(filters) - 1; i >= 0; i-- {
		next = filters[i].OnFilter(next)
	}
	return next
}
