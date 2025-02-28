package httpport

type errorRes struct {
	Errors struct {
		Body []string `json:"body"`
	} `json:"errors"`
}

func newErrorRes(args ...error) errorRes {
	var res errorRes
	for _, err := range args {
		res.Errors.Body = append(res.Errors.Body, err.Error())
	}
	return res
}
