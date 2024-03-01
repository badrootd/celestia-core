package client_test

import (
	"os"
	"testing"

	"github.com/badrootd/celestia-core/abci/example/kvstore"
	nm "github.com/badrootd/celestia-core/node"
	rpctest "github.com/badrootd/celestia-core/rpc/test"
)

var node *nm.Node

func TestMain(m *testing.M) {
	// start a CometBFT node (and kvstore) in the background to test against
	dir, err := os.MkdirTemp("/tmp", "rpc-client-test")
	if err != nil {
		panic(err)
	}

	app := kvstore.NewPersistentKVStoreApplication(dir)
	// If testing block event generation
	// app.SetGenBlockEvents() needs to be called here
	node = rpctest.StartTendermint(app)

	code := m.Run()

	// and shut down proper at the end
	rpctest.StopTendermint(node)
	_ = os.RemoveAll(dir)
	os.Exit(code)
}
