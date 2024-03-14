package structs

type PicoWResponse struct {
	ID    int     `json:"id"`
	Error *string `json:"error"`
	Data  any     `json:"data"`
}
