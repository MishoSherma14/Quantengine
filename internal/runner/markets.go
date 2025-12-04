package runner

// Market groups:
// 1) Top Tech Stocks
// 2) Top Crypto
// 3) Top Finance
// 4) Future: Forex, Commodities

var Markets = []string{
    // --- Tech ---
    "AAPL",   // Apple
    "MSFT",   // Microsoft
    "NVDA",   // Nvidia
    "AMZN",   // Amazon
    "GOOGL",  // Alphabet
    "META",   // Meta Platforms
    "TSLA",   // Tesla

    // --- Crypto ---
    "BTCUSDT",
    "ETHUSDT",
    "SOLUSDT",
    "XRPUSDT",
    "BNBUSDT",
    "ADAUSDT",
    "DOTUSDT",

    // --- Finance ---
    "JPM",
    "GS",
    "BAC",
    "WFC",
    "MS",
    "C",
    "BLK",
}

// Timeframe is fixed (for now)  
// Later can expand: 1h, 4h, 1w
var Timeframe = "1D"
