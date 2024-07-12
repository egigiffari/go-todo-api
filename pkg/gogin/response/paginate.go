package response

type Paginate struct {
	Page     int `json:"page"`
	Size     int `json:"size"`
	LastPage int `json:"last_page"`
	Total    int `json:"total"`
	Data     any `json:"data"`
}
