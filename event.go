package next

import "time"

type ByteFrame []byte

var _ Event = new(FrameEvent)

type FrameEvent struct {
    id     int64
    kind   Kind
    time   time.Time
    frame  []byte
    header Header
}

func (f *FrameEvent) Timestamp() time.Time {
    return f.time
}

func (f *FrameEvent) ID() int64 {
    return f.id
}

func (f *FrameEvent) Kind() Kind {
    return f.kind
}

func (f *FrameEvent) Header() Header {
    return f.header
}

func (f *FrameEvent) Data() any {
    return f.frame
}

func (f *FrameEvent) DataString() string {
    return string(f.frame)
}

func (f *FrameEvent) DataBytes() []byte {
    return f.frame
}

func NewFrameEvent(id int64, kind Kind, time time.Time, header Header, data []byte) *FrameEvent {
    return &FrameEvent{
        id:     id,
        kind:   kind,
        time:   time,
        header: header,
        frame:  ByteFrame(data),
    }
}

func NewSimpleFrameEvent(id int64, kind Kind, data []byte) *FrameEvent {
    return &FrameEvent{
        id:     id,
        kind:   kind,
        time:   time.Now(),
        frame:  ByteFrame(data),
        header: Header{},
    }
}
