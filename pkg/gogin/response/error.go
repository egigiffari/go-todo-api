package response

type Error struct {
	Errors    map[string][]string `json:"errors"`
	TraceId   string              `json:"_trace_id"`
	Timestamp string              `json:"timestamp"`
}
