package ethereum

type EventRaw struct {
	Timestamp   string `json:"timestamp"`
	BlockNumber string `json:"blockNumber"`
	Txhash      string `json:"transactionHash"`
	Data        string `json:"data"`
}
type EventHistory struct {
	ActualDestAmount string `json:"actualDestAmount"`
	ActualSrcAmount  string `json:"actualSrcAmount"`
	Dest             string `json:"dest"`
	Source           string `json:"source"`

	BlockNumber string `json:"blockNumber"`
	Txhash      string `json:"txHash"`
	Timestamp   string `json:"timestamp"`
}

type Rate struct {
	Source  string `json:"source"`
	Dest    string `json:"dest"`
	Rate    string `json:"rate"`
	Minrate string `json:"minRate"`
}

type GasPrice struct {
	Fast     string `json:"fast"`
	Standard string `json:"standard"`
	Low      string `json:"low"`
	Default  string `json:"default"`
}

type Token struct {
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Address string `json:"address"`
	Decimal int    `json:"decimal"`
	UsdId   string `json:"usd_id"`
}

type QuoInfo struct {
	MarketCap float64 `json:"market_cap"`
	Volume24h float64 `json:"volume_24h"`
}

type TokenGeneralInfo struct {
	CirculatingSupply float64            `json:"circulating_supply"`
	TotalSupply       float64            `json:"total_supply"`
	MaxSupply         float64            `json:"max_supply"`
	MarketCap         float64            `json:"market_cap"`
	Quotes            map[string]QuoInfo `json:"quotes`
}

// type TokenInfoData struct {
// 	Data TokenGeneralInfo `data`
// }

type CandleTicker struct {
	SellPrice string `json:"sell_price"`
	BuyPrice  string `json:"buy_price"`
	UnixTime  uint64 `json:"unix_time"`
}

type TokenInfo struct {
	TokenSymbol       string                  `json:"symbol"`
	CirculatingSupply string                  `json:"circulating_supply"`
	TotalSupply       string                  `json:"total_supply"`
	MaxSupply         string                  `json:"max_supply"`
	MarketCap         string                  `json:"market_cap"`
	SellPrice         string                  `json:"sell_price"`
	BuyPrice          string                  `json:"buy_price"`
	Last7days         map[uint64]CandleTicker `json:"last_7d"`
}

type RateHistory struct {
	SellPrice string `json:"sell_price"`
	BuyPrice  string `json:"buy_price"`
}
type RateInfo struct {
	LastSell      string                 `json:"last_sell"`
	LastBuy       string                 `json:"last_buy"`
	HistoryRecord map[int64]*RateHistory `json:"history_record"`
}

type Rates struct {
	R float64   `json:"r"`
	P []float64 `json:"p"`
}

type MarketInfo struct {
	Rates  *Rates             `json:"rates"`
	Quotes map[string]QuoInfo `json:"quotes"`
}

type RightMarketInfo struct {
	Rate   *float64           `json:"rate"`
	Quotes map[string]QuoInfo `json:"quotes"`
}

func NewMarketInfo(quotes map[string]QuoInfo, rates *Rates) *MarketInfo {
	return &MarketInfo{
		Rates:  rates,
		Quotes: quotes,
	}
}
