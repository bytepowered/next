package next

import (
	"testing"
	"time"
)

func TestFrameEvent_DataString(t *testing.T) {
	type fields struct {
		id     int64
		kind   Kind
		time   time.Time
		data   []byte
		header Header
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestFrameEvent_DataString",
			fields: fields{
				id:     1,
				kind:   1,
				time:   time.Now(),
				data:   []byte("Hello, World!"),
				header: Header{},
			},
			want: "Hello, World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FrameEvent{
				id:     tt.fields.id,
				kind:   tt.fields.kind,
				time:   tt.fields.time,
				data:   tt.fields.data,
				header: tt.fields.header,
			}
			if got := f.DataString(); got != tt.want {
				t.Errorf("FrameEvent.DataString() = %v, want %v", got, tt.want)
			}
		})
	}
}
