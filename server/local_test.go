package server

import (
	"context"
	"testing"

	"bitbucket.org/creachadair/jrpc2/caller"
	"bitbucket.org/creachadair/jrpc2/handler"
)

func TestLocal(t *testing.T) {
	loc := NewLocal(make(handler.Map), nil)

	ctx := context.Background()
	si, err := caller.RPCServerInfo(ctx, loc.Client)
	if err != nil {
		t.Fatalf("rpc.serverInfo failed: %v", err)
	}

	// A couple sanity checks on the server info.
	if nr := si.Counter["rpc.requests"]; nr != 1 {
		t.Errorf("rpc.serverInfo reports %d requests, wanted 1", nr)
	}
	if len(si.Methods) != 0 {
		t.Errorf("rpc.serverInfo reports methods %+q, wanted []", si.Methods)
	}

	// Close the client and wait for the server to stop.
	if err := loc.Close(); err != nil {
		t.Errorf("Server wait: got %v, want nil", err)
	}
}
