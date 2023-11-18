package types

// BITCOINTxIndexer defines the interface of custom bitcoin tx indexer.
type BITCOINTxIndexer interface {
	// ParseBlock parse bitcoin block tx
	ParseBlock(int64) ([]*BitcoinTxParseResult, error)
}
