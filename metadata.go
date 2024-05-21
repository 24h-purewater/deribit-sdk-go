package deribit

type JSONRpcResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	UsIn    int64  `json:"usIn"`
	UsOut   int64  `json:"usOut"`
	UsDiff  int    `json:"usDiff"`
	Testnet bool   `json:"testnet"`
}

type Order struct {
	TimeInForce         string  `json:"time_in_force"`
	ReduceOnly          bool    `json:"reduce_only"`
	ProfitLoss          float64 `json:"profit_loss"`
	Price               float64 `json:"price"`
	PostOnly            bool    `json:"post_only"`
	OrderType           string  `json:"order_type"`
	OrderState          string  `json:"order_state"`
	OrderID             string  `json:"order_id"`
	MaxShow             float64 `json:"max_show"`
	LastUpdateTimestamp int64   `json:"last_update_timestamp"`
	Label               string  `json:"label"`
	IsLiquidation       bool    `json:"is_liquidation"`
	InstrumentName      string  `json:"instrument_name"`
	FilledAmount        float64 `json:"filled_amount"`
	Direction           string  `json:"direction"`
	CreationTimestamp   int64   `json:"creation_timestamp"`
	Commission          float64 `json:"commission"`
	AveragePrice        float64 `json:"average_price"`
	API                 bool    `json:"api"`
	Amount              float64 `json:"amount"`
}

type OrderListResp struct {
	JSONRpcResponse
	Result []Order `json:"result"`
}

type Trade struct {
	UnderlyingPrice float64     `json:"underlying_price"`
	TradeSeq        int         `json:"trade_seq"`
	TradeID         string      `json:"trade_id"`
	Timestamp       int64       `json:"timestamp"`
	TickDirection   int         `json:"tick_direction"`
	State           string      `json:"state"`
	SelfTrade       bool        `json:"self_trade"`
	ReduceOnly      bool        `json:"reduce_only"`
	Price           float64     `json:"price"`
	PostOnly        bool        `json:"post_only"`
	OrderType       string      `json:"order_type"`
	OrderID         string      `json:"order_id"`
	MatchingID      interface{} `json:"matching_id"`
	MarkPrice       float64     `json:"mark_price"`
	Liquidity       string      `json:"liquidity"`
	Iv              float64     `json:"iv"`
	InstrumentName  string      `json:"instrument_name"`
	IndexPrice      float64     `json:"index_price"`
	FeeCurrency     string      `json:"fee_currency"`
	Fee             float64     `json:"fee"`
	Direction       string      `json:"direction"`
	Amount          float64     `json:"amount"`
}

type TradeListResp struct {
	JSONRpcResponse
	Result struct {
		Trades  []Trade `json:"trades"`
		HasMore bool    `json:"has_more"`
	} `json:"result"`
}

//InstrumentListResp InstrumentList Response
type InstrumentListResp struct {
	JSONRpcResponse
	Result []Instrument `json:"result"`
}
type Instrument struct {
	TickSize             float64 `json:"tick_size"`
	TakerCommission      float64 `json:"taker_commission"`
	Strike               int     `json:"strike"`
	SettlementPeriod     string  `json:"settlement_period"`
	QuoteCurrency        string  `json:"quote_currency"`
	OptionType           string  `json:"option_type"`
	MinTradeAmount       float64 `json:"min_trade_amount"`
	MakerCommission      float64 `json:"maker_commission"`
	Kind                 string  `json:"kind"`
	IsActive             bool    `json:"is_active"`
	InstrumentName       string  `json:"instrument_name"`
	ExpirationTimestamp  int64   `json:"expiration_timestamp"`
	CreationTimestamp    int64   `json:"creation_timestamp"`
	ContractSize         int     `json:"contract_size"`
	BlockTradeCommission float64 `json:"block_trade_commission"`
	BaseCurrency         string  `json:"base_currency"`
}

// IndexPriceResp  IndexPriceResp
type IndexPriceResp struct {
	JSONRpcResponse
	Result IndexPriceResult `json:"result"`
}
type IndexPriceResult struct {
	IndexPrice             float64 `json:"index_price"`
	EstimatedDeliveryPrice float64 `json:"estimated_delivery_price"`
}

// PositionListResp PositionList Response
type PositionListResp struct {
	JSONRpcResponse
	Result []PositionListResult `json:"result"`
}
type PositionListResult struct {
	Vega                  float64 `json:"vega"`
	TotalProfitLoss       float64 `json:"total_profit_loss"`
	Theta                 float64 `json:"theta"`
	Size                  float64 `json:"size"`
	SettlementPrice       float64 `json:"settlement_price"`
	RealizedProfitLoss    float64 `json:"realized_profit_loss"`
	OpenOrdersMargin      float64 `json:"open_orders_margin"`
	MarkPrice             float64 `json:"mark_price"`
	MaintenanceMargin     float64 `json:"maintenance_margin"`
	Kind                  string  `json:"kind"`
	InstrumentName        string  `json:"instrument_name"`
	InitialMargin         float64 `json:"initial_margin"`
	IndexPrice            float64 `json:"index_price"`
	Gamma                 float64 `json:"gamma"`
	FloatingProfitLossUsd float64 `json:"floating_profit_loss_usd"`
	FloatingProfitLoss    float64 `json:"floating_profit_loss"`
	Direction             string  `json:"direction"`
	Delta                 float64 `json:"delta"`
	AveragePriceUsd       float64 `json:"average_price_usd"`
	AveragePrice          float64 `json:"average_price"`
}

type InstrumentResp struct {
	JSONRpcResponse
	Result Instrument `json:"result"`
}

// OrderBookResp OrderBook Response
type OrderBookResp struct {
	JSONRpcResponse
	Result OrderBook `json:"result"`
}

type OrderBook struct {
	Timestamp int64 `json:"timestamp"`
	Stats     struct {
		Volume      float64 `json:"volume"`
		PriceChange float64 `json:"price_change"`
		Low         float64 `json:"low"`
		High        float64 `json:"high"`
	} `json:"stats"`
	State           string      `json:"state"`
	SettlementPrice float64     `json:"settlement_price"`
	OpenInterest    float64     `json:"open_interest"`
	MinPrice        float64     `json:"min_price"`
	MaxPrice        float64     `json:"max_price"`
	MarkPrice       float64     `json:"mark_price"`
	LastPrice       float64     `json:"last_price"`
	InstrumentName  string      `json:"instrument_name"`
	IndexPrice      float64     `json:"index_price"`
	Funding8H       float64     `json:"funding_8h"`
	CurrentFunding  float64     `json:"current_funding"`
	ChangeID        int         `json:"change_id"`
	Bids            [][]float64 `json:"bids"`
	BestBidPrice    float64     `json:"best_bid_price"`
	BestBidAmount   float64     `json:"best_bid_amount"`
	BestAskPrice    float64     `json:"best_ask_price"`
	BestAskAmount   float64     `json:"best_ask_amount"`
	Asks            [][]float64 `json:"asks"`
}

type OrderTrade struct {
	TradeSeq       int         `json:"trade_seq"`
	TradeID        string      `json:"trade_id"`
	Timestamp      int64       `json:"timestamp"`
	TickDirection  int         `json:"tick_direction"`
	State          string      `json:"state"`
	SelfTrade      bool        `json:"self_trade"`
	ReduceOnly     bool        `json:"reduce_only"`
	Price          float64     `json:"price"`
	PostOnly       bool        `json:"post_only"`
	OrderType      string      `json:"order_type"`
	OrderID        string      `json:"order_id"`
	MatchingID     interface{} `json:"matching_id"`
	MarkPrice      float64     `json:"mark_price"`
	Liquidity      string      `json:"liquidity"`
	Label          string      `json:"label"`
	InstrumentName string      `json:"instrument_name"`
	IndexPrice     float64     `json:"index_price"`
	FeeCurrency    string      `json:"fee_currency"`
	Fee            float64     `json:"fee"`
	Direction      string      `json:"direction"`
	Amount         float64     `json:"amount"`
}

type BuyResult struct {
	Trades []OrderTrade `json:"trades"`
	Order  struct {
		Web                 bool    `json:"web"`
		TimeInForce         string  `json:"time_in_force"`
		Replaced            bool    `json:"replaced"`
		ReduceOnly          bool    `json:"reduce_only"`
		ProfitLoss          float64 `json:"profit_loss"`
		Price               float64 `json:"price"`
		PostOnly            bool    `json:"post_only"`
		OrderType           string  `json:"order_type"`
		OrderState          string  `json:"order_state"`
		OrderID             string  `json:"order_id"`
		MaxShow             float64 `json:"max_show"`
		LastUpdateTimestamp int64   `json:"last_update_timestamp"`
		Label               string  `json:"label"`
		IsLiquidation       bool    `json:"is_liquidation"`
		InstrumentName      string  `json:"instrument_name"`
		FilledAmount        float64 `json:"filled_amount"`
		Direction           string  `json:"direction"`
		CreationTimestamp   int64   `json:"creation_timestamp"`
		Commission          float64 `json:"commission"`
		AveragePrice        float64 `json:"average_price"`
		API                 bool    `json:"api"`
		Amount              float64 `json:"amount"`
	} `json:"order"`
}
type BuyResp struct {
	JSONRpcResponse
	Result BuyResult `json:"result"`
}

type SellResult struct {
	Trades []OrderTrade `json:"trades"`
	Order  struct {
		Triggered           bool    `json:"triggered"`
		Trigger             string  `json:"trigger"`
		TimeInForce         string  `json:"time_in_force"`
		TriggerPrice        float64 `json:"trigger_price"`
		ReduceOnly          bool    `json:"reduce_only"`
		ProfitLoss          float64 `json:"profit_loss"`
		Price               float64 `json:"price"`
		PostOnly            bool    `json:"post_only"`
		OrderType           string  `json:"order_type"`
		OrderState          string  `json:"order_state"`
		OrderID             string  `json:"order_id"`
		MaxShow             float64 `json:"max_show"`
		LastUpdateTimestamp int64   `json:"last_update_timestamp"`
		Label               string  `json:"label"`
		IsLiquidation       bool    `json:"is_liquidation"`
		InstrumentName      string  `json:"instrument_name"`
		Direction           string  `json:"direction"`
		CreationTimestamp   int64   `json:"creation_timestamp"`
		API                 bool    `json:"api"`
		AveragePrice        float64 `json:"average_price"`
		Amount              float64 `json:"amount"`
	} `json:"order"`
}

type SellResp struct {
	JSONRpcResponse
	Result SellResult `json:"result"`
}

type CancelResult struct {
	Triggered           bool    `json:"triggered"`
	Trigger             string  `json:"trigger"`
	TimeInForce         string  `json:"time_in_force"`
	TriggerPrice        float64 `json:"trigger_price"`
	ReduceOnly          bool    `json:"reduce_only"`
	ProfitLoss          float64 `json:"profit_loss"`
	Price               float64 `json:"price"`
	PostOnly            bool    `json:"post_only"`
	OrderType           string  `json:"order_type"`
	OrderState          string  `json:"order_state"`
	OrderID             string  `json:"order_id"`
	MaxShow             float64 `json:"max_show"`
	LastUpdateTimestamp int64   `json:"last_update_timestamp"`
	Label               string  `json:"label"`
	IsLiquidation       bool    `json:"is_liquidation"`
	InstrumentName      string  `json:"instrument_name"`
	Direction           string  `json:"direction"`
	CreationTimestamp   int64   `json:"creation_timestamp"`
	API                 bool    `json:"api"`
	Amount              float64 `json:"amount"`
}

type CancelResp struct {
	JSONRpcResponse
	Result CancelResult `json:"result"`
}

type GetPositionResp struct {
	JSONRpcResponse
	Result GetPositionResult `json:"result"`
}

type GetPositionResult struct {
	AveragePrice              float64 `json:"average_price"`
	Delta                     float64 `json:"delta"`
	Direction                 string  `json:"direction"`
	EstimatedLiquidationPrice float64 `json:"estimated_liquidation_price"`
	FloatingProfitLoss        float64 `json:"floating_profit_loss"`
	IndexPrice                float64 `json:"index_price"`
	InitialMargin             float64 `json:"initial_margin"`
	InstrumentName            string  `json:"instrument_name"`
	Leverage                  int     `json:"leverage"`
	Kind                      string  `json:"kind"`
	MaintenanceMargin         float64 `json:"maintenance_margin"`
	MarkPrice                 float64 `json:"mark_price"`
	OpenOrdersMargin          float64 `json:"open_orders_margin"`
	RealizedProfitLoss        float64 `json:"realized_profit_loss"`
	SettlementPrice           float64 `json:"settlement_price"`
	Size                      float64 `json:"size"`
	SizeCurrency              int     `json:"size_currency"`
	TotalProfitLoss           float64 `json:"total_profit_loss"`
}

type GetTransactionLogResp struct {
	JSONRpcResponse
	Result struct {
		Logs         []TransactionLog `json:"logs"`
		Continuation int              `json:"continuation"`
	} `json:"result"`
}

type TransactionLog struct {
	SessionUpl     float64     `json:"session_upl,omitempty"`
	SessionRpl     float64     `json:"session_rpl,omitempty"`
	PriceCurrency  string      `json:"price_currency,omitempty"`
	TradeID        string      `json:"trade_id"`
	InterestPl     interface{} `json:"interest_pl"`
	Side           string      `json:"side"`
	UserSeq        int         `json:"user_seq"`
	Equity         float64     `json:"equity"`
	FeeBalance     float64     `json:"fee_balance"`
	InstrumentName string      `json:"instrument_name"`
	OrderID        string      `json:"order_id"`
	Amount         float64     `json:"amount"`

	Username        string      `json:"username"`
	MarkPrice       float64     `json:"mark_price,omitempty"`
	IndexPrice      float64     `json:"index_price"`
	Cashflow        float64     `json:"cashflow"`
	Commission      float64     `json:"commission"`
	UserID          int         `json:"user_id"`
	Price           float64     `json:"price"`
	Change          float64     `json:"change"`
	Currency        string      `json:"currency"`
	Balance         float64     `json:"balance"`
	Type            string      `json:"type"`
	Timestamp       int64       `json:"timestamp"`
	Position        interface{} `json:"position"`
	Info            string      `json:"info,omitempty"`
	ID              int         `json:"id"`
	TotalInterestPl float64     `json:"total_interest_pl,omitempty"`

	UserRole         string `json:"user_role,omitempty"`
	ProfitAsCashflow bool   `json:"profit_as_cashflow,omitempty"`
}

type TxLogInfo struct {
	TransferType    string  `json:"transfer_type"`
	OtherUserID     int     `json:"other_user_id"`
	OtherUser       string  `json:"other_user"`
	Transaction     string  `json:"transaction"`
	DepositType     string  `json:"deposit_type"`
	Addr            string  `json:"addr"`
	SettlementPrice float64 `json:"settlement_price"`
	FloatingPl      float64 `json:"floating_pl"`
}

type AccountSummary struct {
	TotalPl                    float64 `json:"total_pl"`
	SessionUpl                 float64 `json:"session_upl"`
	SessionRpl                 float64 `json:"session_rpl"`
	ProjectedMaintenanceMargin float64 `json:"projected_maintenance_margin"`
	ProjectedInitialMargin     float64 `json:"projected_initial_margin"`
	ProjectedDeltaTotal        float64 `json:"projected_delta_total"`
	PortfolioMarginingEnabled  bool    `json:"portfolio_margining_enabled"`
	OptionsVega                float64 `json:"options_vega"`
	OptionsValue               float64 `json:"options_value"`
	OptionsTheta               float64 `json:"options_theta"`
	OptionsSessionUpl          float64 `json:"options_session_upl"`
	OptionsSessionRpl          float64 `json:"options_session_rpl"`
	OptionsPl                  float64 `json:"options_pl"`
	OptionsGamma               float64 `json:"options_gamma"`
	OptionsDelta               float64 `json:"options_delta"`
	MarginBalance              float64 `json:"margin_balance"`
	MaintenanceMargin          float64 `json:"maintenance_margin"`
	Limits                     struct {
		NonMatchingEngine struct {
			Rate  float64 `json:"rate"`
			Burst float64 `json:"burst"`
		} `json:"non_matching_engine"`
		MatchingEngine struct {
			Rate  float64 `json:"rate"`
			Burst float64 `json:"burst"`
		} `json:"matching_engine"`
	} `json:"limits"`
	InitialMargin                float64 `json:"initial_margin"`
	FuturesSessionUpl            float64 `json:"futures_session_upl"`
	FuturesSessionRpl            float64 `json:"futures_session_rpl"`
	FuturesPl                    float64 `json:"futures_pl"`
	FeeBalance                   float64 `json:"fee_balance"`
	EstimatedLiquidationRatioMap struct {
	} `json:"estimated_liquidation_ratio_map"`
	EstimatedLiquidationRatio float64 `json:"estimated_liquidation_ratio"`
	Equity                    float64 `json:"equity"`
	DeltaTotalMap             struct {
	} `json:"delta_total_map"`
	DeltaTotal               float64 `json:"delta_total"`
	Currency                 string  `json:"currency"`
	Balance                  float64 `json:"balance"`
	AvailableWithdrawalFunds float64 `json:"available_withdrawal_funds"`
	AvailableFunds           float64 `json:"available_funds"`
}

type GetAccountSummaryResp struct {
	JSONRpcResponse
	Result AccountSummary `json:"result"`
}
