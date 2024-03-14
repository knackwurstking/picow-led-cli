package picow

type Response struct {
	ID    int     `json:"id"`
	Error *string `json:"error"`
	Data  any     `json:"data"`
}
