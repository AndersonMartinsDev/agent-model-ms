package response

type ResponseModel struct {
	Object     interface{} `json:"data"`
	StatusHttp int         `json:"status"`
}
