package services

import (
	"context"
	"strings"

	"github.com/UnLess24/coin/server/internal/cache"
	"github.com/UnLess24/coin/server/internal/currency"
	"github.com/UnLess24/coin/server/pkg/api/coin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CoinService struct {
	coin.UnimplementedCoinServiceServer

	cdb cache.Cache
}

func NewCoinService(cdb cache.Cache) *CoinService {
	return &CoinService{cdb: cdb}
}

func (s *CoinService) ListCoins(ctx context.Context, req *coin.ListCoinsRequest) (*coin.ListCoinsResponse, error) {
	list, err := cache.List(ctx, s.cdb)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list coins: %v", err)
	}

	coins := make([]*coin.Currency, 0)
	for _, datum := range list.Data {
		quote := make(map[string]*coin.Quote)
		for k, v := range datum.Quote {
			quote[string(k)] = &coin.Quote{
				Price:       float32(v.Price),
				LastUpdated: v.LastUpdated,
			}
		}
		coin := &coin.Currency{
			Id:     int64(datum.ID),
			Name:   datum.Name,
			Symbol: datum.Symbol,
			Quote:  quote,
		}
		coins = append(coins, coin)
	}

	resp := &coin.ListCoinsResponse{
		Currencies: coins,
	}
	return resp, nil
}

func (s *CoinService) GetQuote(ctx context.Context, req *coin.GetQuoteRequest) (*coin.GetQuoteResponse, error) {
	list, err := cache.List(ctx, s.cdb)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list coins: %v", err)
	}

	quotes := make([]*coin.Quote, 0)
	name := strings.ToUpper(req.GetName())
	currName := currency.CurrencyName(name)
	for _, datum := range list.Data {
		if val, ok := datum.Quote[currName]; ok {
			quote := &coin.Quote{
				Price:       float32(val.Price),
				LastUpdated: val.LastUpdated}
			quotes = append(quotes, quote)
		}
	}

	return &coin.GetQuoteResponse{
		Name:   name,
		Quotes: quotes,
	}, nil
}
