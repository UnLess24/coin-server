package currency

import (
	"context"
	"fmt"
	"net/http"

	"github.com/UnLess24/coin/server/config"
)

func Request(ctx context.Context, cfg *config.Config) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cfg.Host+cfg.LastListUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", cfg.APIKey)

	return req, nil
}
