package response

type VotingInfos struct {
	Votes []string `json:"votes"`
}

type ResponseObject struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  *string     `json:"error"`
}
