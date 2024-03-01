package psql

import (
	"github.com/celestiaorg/celestia-core/state/indexer"
	"github.com/celestiaorg/celestia-core/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
