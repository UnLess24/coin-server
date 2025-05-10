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

// LatestList send request to CoinMarketCap and get latest list of currencies
func LatestList(ctx context.Context, cfg *config.Config) ([]byte, error) {
	client := &http.Client{}
	req, err := Request(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to send request %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get response %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	body, _ := io.ReadAll(resp.Body)

	return body, nil
}

// MakeList unmarshal data to ListResponse
func MakeList(data []byte) (ListResponse, error) {
	list := ListResponse{}
	err := json.Unmarshal(data, &list)
	if err != nil {
		return ListResponse{}, fmt.Errorf("failed to unmarshal data %w", err)
	}
	return list, nil
}
