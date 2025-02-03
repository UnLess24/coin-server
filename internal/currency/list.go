package currency

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/UnLess24/coin/server/config"
)

type ListResponse struct {
	Data   []Currency `json:"data,omitempty"`
	Status Status     `json:"status,omitempty"`
}

func SendRequest(ctx context.Context, cfg *config.Config) (ListResponse, error) {
	client := &http.Client{}
	req, err := Request(context.Background(), cfg)
	if err != nil {
		return ListResponse{}, fmt.Errorf("failed to send request %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return ListResponse{}, fmt.Errorf("failed to get response %v", err)
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	list, err := MakeList(body)
	if err != nil {
		return ListResponse{}, fmt.Errorf("failed to make list %v", err)
	}
	return list, nil
}

func MakeList(data []byte) (ListResponse, error) {
	list := ListResponse{}
	err := json.Unmarshal(data, &list)
	if err != nil {
		return ListResponse{}, fmt.Errorf("failed to unmarshal data %v", err)
	}
	return list, nil
}
