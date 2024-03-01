package psql

import (
	"github.com/badrootd/celestia-core/state/indexer"
	"github.com/badrootd/celestia-core/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
