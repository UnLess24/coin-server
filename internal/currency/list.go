package currency

import "encoding/json"

type ListResponse struct {
	Data   []Currency `json:"data,omitempty"`
	Status Status     `json:"status,omitempty"`
}

func MakeList(data []byte) ListResponse {
	list := ListResponse{}
	json.Unmarshal(data, &list)
	return list
}
