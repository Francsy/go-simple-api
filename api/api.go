package api

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	Code    int   // Success Code, usually 200
	Balance int64 // Account Balance
}

// Error Response
type Error struct {
	// Error Code
	Code int
	// Error message
	Message string
}
