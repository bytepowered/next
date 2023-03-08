package next

const (
    HEADER_REQUEST_ID  = "x-request-id"
    HEADER_REMOTE_ADDR = "x-remote-addr"
)

func NewHeader(capacity int) Header {
    return make(Header, capacity)
}

func (h Header) GetOrDefault(name string, def any) any {
    if v, ok := h[name]; ok {
        return v
    }
    return def
}

func (h Header) Get(name string) (any, bool) {
    v, ok := h[name]
    return v, ok
}

func (h Header) GetString(name string) string {
    return h.GetOrDefault(name, "").(string)
}

func (h Header) GetInt(name string) int {
    return h.GetOrDefault(name, 0).(int)
}

func (h Header) Put(name string, value any) {
    h[name] = value
}

func (h Header) PutIfAbsent(name string, value any) {
    if _, ok := h[name]; !ok {
        h[name] = value
    }
}
