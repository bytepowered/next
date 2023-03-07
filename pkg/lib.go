package next

import (
	"context"
	"time"
)

// Kind 表示Event类型
type Kind uint16

type Header map[string]any

// Event 具体Event消息接口
type Event[T any] interface {
	// ID 返回事件ID
	ID() int64
	// Timestamp 返回事件发生时间
	Timestamp() time.Time
	// Kind 返回事件类型
	Kind() Kind
	// Header 返回事件头
	Header() Header
	// Data 返回事件数据
	Data() T
}

type EventSink[T any] func(ctx context.Context, event Event[T]) error

// Input 事件输入源
type Input[T any] interface {
	OnRead(ctx context.Context, sink EventSink[T]) error
}

// Output 事件输出源
type Output[T any] interface {
	OnWrite(ctx context.Context, events []Event[T]) error
}

// FilterFunc 执行过滤原始Event的函数；
type FilterFunc[T any] func(ctx context.Context, event Event[T]) error

// EventFilter 原始Event过滤接口
type EventFilter[T any] interface {
	OnFilter(next FilterFunc[T]) FilterFunc[T]
}

// EventTransformer 处理Event格式转换
type EventTransformer[F any, T any] interface {
	OnTransform(ctx context.Context, in []Event[F]) (out []Event[T], err error)
}
