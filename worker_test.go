package next

import (
    "context"
    "testing"
)

type input struct{}

func (i *input) OnRead(ctx context.Context, sink EventSink[ByteFrame]) error {
    for i := 0; i < 10; i++ {
        err := sink(ctx, NewSimpleFrameEvent(int64(i), 0, []byte("hello")))
        if err != nil {
            return err
        }
    }
    return nil
}

type output struct{}

func (o *output) OnWrite(ctx context.Context, events []Event[ByteFrame]) error {
    for _, event := range events {
        println(string(event.Data()))
    }
    return nil
}

type filter struct {
}

func (f *filter) OnFilter(next FilterFunc[ByteFrame]) FilterFunc[ByteFrame] {
    return func(ctx context.Context, event Event[ByteFrame]) error {
        if event.ID()%2 == 0 {
            return nil
        }
        return next(ctx, event)
    }
}

func TestWorker_Run(t *testing.T) {
    worker := Worker[ByteFrame]{
        name: "test",
    }
    worker.SetInput(&input{})
    worker.SetOutput(&output{})
    worker.AddFilter(&filter{})
    err := worker.Run(context.Background())
    if err != nil {
        t.Fatal(err)
    }
}
