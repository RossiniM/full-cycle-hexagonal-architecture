package handler

import "encoding/json"

func jsonError(msg string) []byte {
	jsonError := struct {
		Message string `json:"message"`
	}{
		msg,
	}
	r, err := json.Marshal(jsonError)
	if err != nil {
		return []byte(err.Error())
	}
	return r

}
