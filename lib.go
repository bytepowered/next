package next

import (
    "context"
    "time"
)

// Kind 表示Event类型
type Kind uint16

type Header map[string]any

// Event 具体Event消息接口
type Event interface {
    // ID 返回事件ID
    ID() int64
    // Timestamp 返回事件发生时间
    Timestamp() time.Time
    // Kind 返回事件类型
    Kind() Kind
    // Header 返回事件头
    Header() Header
    // Data 返回事件数据
    Data() any
}

type EventSink func(ctx context.Context, event Event) error

// Input 事件输入源
type Input interface {
    OnRead(ctx context.Context, sink EventSink) error
}

// Output 事件输出源
type Output interface {
    OnWrite(ctx context.Context, events []Event) error
}

// FilterFunc 执行过滤原始Event的函数；
type FilterFunc func(ctx context.Context, event Event) error

// EventFilter 原始Event过滤接口
type EventFilter interface {
    OnFilter(next FilterFunc) FilterFunc
}

// EventTransformer 处理Event格式转换
type EventTransformer interface {
    OnTransform(ctx context.Context, in []Event) (out []Event, err error)
}
