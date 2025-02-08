package currency

type Currency struct {
	ID                            int       `json:id`
	Name                          string    `json:"name,omitempty"`
	Symbol                        string    `json:"symbol,omitempty"`
	Slug                          string    `json:"slug,omitempty"`
	CmcRank                       int       `json:"cmc_rank,omitempty"`
	NumMarketPairs                int       `json:"num_market_pairs,omitempty"`
	CyrculatingSupply             float64   `json:"cyrculating_supply,omitempty"`
	TotalSupply                   float64   `json:"total_supply,omitempty"`
	MarketCapByTotalSupply        float64   `json:"market_cap_by_total_supply,omitempty"`
	MaxSupply                     float64   `json:"max_supply,omitempty"`
	InfiniteSupply                bool      `json:"infinite_supply,omitempty"`
	LastUpdated                   string    `json:"last_updated,omitempty"`
	DateAdded                     string    `json:"date_added,omitempty"`
	Tags                          []string  `json:"tags,omitempty"`
	SelfReportedCirculatingSupply float64   `json:"self_reported_circulating_supply,omitempty"`
	SelfReportedMarketCap         float64   `json:"self_reported_market_cap,omitempty"`
	TvlRatio                      float64   `json:"tvl_ratio,omitempty"`
	Platform                      string    `json:"platform,omitempty"`
	Quote                         QuoteItem `json:"quote,omitempty"`
}

type Status struct {
	Timestamp    string `json:"timestamp,omitempty"`
	ErrorCode    int    `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	Elapsed      int    `json:"elapsed,omitempty"`
	CreditCount  int    `json:"credit_count,omitempty"`
	Notice       string `json:"notice,omitempty"`
}

type CurrencyName string
type CurrencyValue struct {
	Price                 float64 `json:"price,omitempty"`
	Volume24h             float64 `json:"volume_24h,omitempty"`
	VolumeChange24h       float64 `json:"volume_change_24h,omitempty"`
	Volume24hReported     float64 `json:"volume_24h_reported,omitempty"`
	Volume7d              float64 `json:"volume_7d,omitempty"`
	Volume7dReported      float64 `json:"volume_7d_reported,omitempty"`
	Volume30d             float64 `json:"volume_30d,omitempty"`
	Volume30dReported     float64 `json:"volume_30d_reported,omitempty"`
	MarketCap             float64 `json:"market_cap,omitempty"`
	MarkeCapDominance     float64 `json:"market_cap_dominance,omitempty"`
	FullyDilutedMarketCap float64 `json:"fully_diluted_market_cap,omitempty"`
	Tvl                   float64 `json:"tvl,omitempty"`
	PercentChange1h       float64 `json:"percent_change_1h,omitempty"`
	PercentChange24h      float64 `json:"percent_change_24h,omitempty"`
	PercentChange7d       float64 `json:"percent_change_7d,omitempty"`
	LastUpdated           string  `json:"last_updated,omitempty"`
}

type QuoteItem map[CurrencyName]CurrencyValue
