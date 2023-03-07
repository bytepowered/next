package next

import "time"

type ByteFrame []byte

var _ Event[ByteFrame] = new(FrameEvent)

type FrameEvent struct {
	id     int64
	kind   Kind
	time   time.Time
	data   []byte
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

func (f *FrameEvent) Data() ByteFrame {
	return f.data
}

func (f *FrameEvent) DataString() string {
	return string(f.data)
}

func NewFrameEvent(id int64, kind Kind, time time.Time, header Header, data []byte) *FrameEvent {
	return &FrameEvent{
		id:     id,
		kind:   kind,
		time:   time,
		header: header,
		data:   ByteFrame(data),
	}
}

func NewSimpleFrameEvent(id int64, kind Kind, data []byte) *FrameEvent {
	return &FrameEvent{
		id:     id,
		kind:   kind,
		time:   time.Now(),
		data:   ByteFrame(data),
		header: Header{},
	}
}
