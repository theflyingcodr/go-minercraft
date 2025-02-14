package minercraft

import "time"

const (

	// version is the current package version
	version = "v0.8.3"

	// defaultUserAgent is the default user agent for all requests
	defaultUserAgent string = "go-minercraft: " + version

	// defaultFastQuoteTimeout is used for the FastestQuote timeout
	defaultFastQuoteTimeout = 20 * time.Second
)

const (
	// routePolicyQuote is the route for getting a policy quote
	routePolicyQuote = "/mapi/policyQuote"

	// routeFeeQuote is the route for getting a fee quote
	routeFeeQuote = "/mapi/feeQuote"

	// routeQueryTx is the route for querying a transaction
	routeQueryTx = "/mapi/tx/"

	// routeSubmitTx is the route for submit a transaction
	routeSubmitTx = "/mapi/tx"

	// routeSubmitTxs is the route for submit batched transactions
	routeSubmitTxs = "/mapi/txs"
)

const (
	// MinerTaal is the name of the known miner for "Taal"
	MinerTaal = "Taal"

	// MinerMempool is the name of the known miner for "Mempool"
	MinerMempool = "Mempool"

	// MinerMatterpool is the name of the known miner for "Matterpool"
	MinerMatterpool = "Matterpool"

	// MinerGorillaPool is the name of the known miner for "GorillaPool"
	MinerGorillaPool = "GorillaPool"
)

// KnownMiners is a pre-filled list of known miners
// Any pre-filled tokens are for free use only
// update your custom token with client.MinerUpdateToken("name", "token")
const KnownMiners = `
[
  {
   "name": "Taal",
   "miner_id": "03e92d3e5c3f7bd945dfbf48e7a99393b1bfb3f11f380ae30d286e7ff2aec5a270",
   "token": "",
   "url": "https://merchantapi.taal.com"
  },
  {
   "name": "Mempool",
   "miner_id": "03e92d3e5c3f7bd945dfbf48e7a99393b1bfb3f11f380ae30d286e7ff2aec5a270",
   "token": "561b756d12572020ea9a104c3441b71790acbbce95a6ddbf7e0630971af9424b",
   "url": "https://www.ddpurse.com/openapi"
  },
  {
   "name": "Matterpool",
   "miner_id": "0253a9b2d017254b91704ba52aad0df5ca32b4fb5cb6b267ada6aefa2bc5833a93",
   "token": "",
   "url": "https://merchantapi.matterpool.io"
  },
  {
   "name": "GorillaPool",
   "miner_id": "03ad780153c47df915b3d2e23af727c68facaca4facd5f155bf5018b979b9aeb83",
   "token": "",
   "url": "https://merchantapi.gorillapool.io"
  }
]
`
