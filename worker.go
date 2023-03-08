package next

import (
    "context"
    "fmt"
)

type Worker struct {
    name         string
    input        Input
    output       Output
    filters      []EventFilter
    transformers []EventTransformer
}

func NewWorker(name string) *Worker {
    return &Worker{name: name}
}

func (w *Worker) Name() string {
    return w.name
}

func (w *Worker) SetInput(input Input) *Worker {
    w.input = input
    return w
}

func (w *Worker) SetOutput(output Output) *Worker {
    w.output = output
    return w
}

func (w *Worker) AddFilter(filter EventFilter) *Worker {
    w.filters = append(w.filters, filter)
    return w
}

func (w *Worker) AddTransformer(transformer EventTransformer) *Worker {
    w.transformers = append(w.transformers, transformer)
    return w
}

func (w *Worker) Run(ctx context.Context) error {
    next := FilterFunc(func(context context.Context, event Event) (err error) {
        events := []Event{event}
        for _, tf := range w.transformers {
            if events, err = tf.OnTransform(context, events); err != nil {
                return fmt.Errorf("transformer %T error: %w", tf, err)
            }
        }
        return w.output.OnWrite(context, events)
    })
    sink := EventSink(func(context context.Context, event Event) error {
        return makeFilterChain(next, w.filters)(context, event)
    })
    return w.input.OnRead(ctx, sink)
}

func makeFilterChain(next FilterFunc, filters []EventFilter) FilterFunc {
    for i := len(filters) - 1; i >= 0; i-- {
        next = filters[i].OnFilter(next)
    }
    return next
}
