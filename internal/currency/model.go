package currency

type Currency struct {
	Id                            int       `json:id`
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
type CurrencyValue float64

type QuoteItem map[CurrencyName]CurrencyValue
