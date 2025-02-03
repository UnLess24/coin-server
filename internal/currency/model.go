package currency

type Currency struct {
	Id                            int
	Name                          string
	Symbol                        string
	Slug                          string
	CmcRank                       int     // cmc_rank
	NumMarketPairs                int     // num_market_pairs
	CyrculatingSupply             float64 // circulating_supply
	TotalSupply                   float64 // total_supply
	MarketCapByTotalSupply        float64 // market_cap_by_total_supply
	MaxSupply                     float64 // max_supply
	InfiniteSupply                bool    // infinite_supply
	LastUpdated                   string  // last_updated
	DateAdded                     string  // date_added
	Tags                          []string
	SelfReportedCirculatingSupply float64 // self_reported_circulating_supply
	SelfReportedMarketCap         float64 // self_reported_market_cap
	TvlRatio                      float64 // tvl_ratio
	Platform                      string
	Quote                         map[string]float64
}

type Status struct {
	Timestamp    string
	ErrorCode    int    // error_code
	ErrorMessage string // error_message
	Elapsed      int
	CreditCount  int // credit_count
	Notice       string
}
