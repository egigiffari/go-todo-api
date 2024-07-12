package response

type Ok struct {
	Data      any    `json:"data"`
	TraceId   string `json:"_trace_id"`
	Timestamp string `json:"timestamp"`
}
