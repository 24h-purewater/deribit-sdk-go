package deribit

type JSONRpcResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Usin    int64  `json:"usIn"`
	Usout   int64  `json:"usOut"`
	Usdiff  int    `json:"usDiff"`
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

type OrderListResult struct {
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

type TradeListResult struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Trades  []Trade `json:"trades"`
		HasMore bool    `json:"has_more"`
	} `json:"result"`
}

type InstrumentListResult struct {
	Jsonrpc string       `json:"jsonrpc"`
	ID      int          `json:"id"`
	Result  []Instrument `json:"result"`
}

type IndexPriceResult struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		IndexPrice             float64 `json:"index_price"`
		EstimatedDeliveryPrice float64 `json:"estimated_delivery_price"`
	} `json:"result"`
}

type Instrument struct {
	TickSize            float64 `json:"tick_size"`
	TakerCommission     float64 `json:"taker_commission"`
	SettlementPeriod    string  `json:"settlement_period"`
	QuoteCurrency       string  `json:"quote_currency"`
	MinTradeAmount      float64 `json:"min_trade_amount"`
	MakerCommission     float64 `json:"maker_commission"`
	Leverage            int     `json:"leverage"`
	Kind                string  `json:"kind"`
	IsActive            bool    `json:"is_active"`
	InstrumentName      string  `json:"instrument_name"`
	ExpirationTimestamp int64   `json:"expiration_timestamp"`
	CreationTimestamp   int64   `json:"creation_timestamp"`
	ContractSize        float64 `json:"contract_size"`
	BaseCurrency        string  `json:"base_currency"`
}

type Position struct {
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

type PositionListResult struct {
	Jsonrpc string     `json:"jsonrpc"`
	ID      int        `json:"id"`
	Result  []Position `json:"result"`
}

type InstrumentResult struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
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
	} `json:"result"`
	Usin    int64 `json:"usIn"`
	Usout   int64 `json:"usOut"`
	Usdiff  int   `json:"usDiff"`
	Testnet bool  `json:"testnet"`
}

type OrderBookResult struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
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
	} `json:"result"`
}

type BuyResult struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Trades []struct {
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
		} `json:"trades"`
		Order struct {
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
	} `json:"result"`
}
type SellResult struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Trades []struct {
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
		} `json:"trades"`
		Order struct {
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
	} `json:"result"`
}

type CancelResult struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
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
	} `json:"result"`
}

type GetPositionResult struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
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
	} `json:"result"`
}

type GetTransactionLogResult struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Logs         []TransactionLog `json:"logs"`
		Continuation int              `json:"continuation"`
	} `json:"result"`
}

type TransactionLog struct {
	Username         string      `json:"username"`
	UserSeq          int         `json:"user_seq"`
	UserID           int         `json:"user_id"`
	Type             string      `json:"type"`
	TradeID          string      `json:"trade_id"`
	Timestamp        int64       `json:"timestamp"`
	Side             string      `json:"side"`
	Price            interface{} `json:"price"`
	Position         interface{} `json:"position"`
	OrderID          string      `json:"order_id"`
	InterestPl       interface{} `json:"interest_pl"`
	InstrumentName   string      `json:"instrument_name"`
	Info             string      `json:"info,omitempty"`
	ID               int         `json:"id"`
	Equity           float64     `json:"equity"`
	Currency         string      `json:"currency"`
	Commission       float64     `json:"commission"`
	Change           float64     `json:"change"`
	Cashflow         float64     `json:"cashflow"`
	Balance          float64     `json:"balance"`
	TotalInterestPl  float64     `json:"total_interest_pl,omitempty"`
	SessionUpl       float64     `json:"session_upl,omitempty"`
	SessionRpl       float64     `json:"session_rpl,omitempty"`
	PriceCurrency    string      `json:"price_currency,omitempty"`
	UserRole         string      `json:"user_role,omitempty"`
	ProfitAsCashflow bool        `json:"profit_as_cashflow,omitempty"`
	MarkPrice        float64     `json:"mark_price,omitempty"`
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

type GetAccountSummaryResult struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
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
	} `json:"result"`
	UsIn    int64 `json:"usIn"`
	UsOut   int64 `json:"usOut"`
	UsDiff  int   `json:"usDiff"`
	Testnet bool  `json:"testnet"`
}
