package pwapi

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"pathwar.land/go/internal/testutil"
)

func TestServer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := testutil.Logger(t)

	server, cleanup := TestingServer(t, ctx, ServerOpts{Logger: logger})
	defer cleanup()

	svc := fmt.Sprintf("http://%s", server.HTTPListenerAddr())
	//svc = strings.Replace(svc, "[::]", "127.0.0.1", -1)

	resp, err := http.Get(svc + "/status")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	expected := `{
  "everything_is_ok": true
}`
	if string(body) != expected {
		t.Fatalf("Expected %q, got %q instead.", expected, string(body))
	}
	// FIXME: check rest of the headers (CORS, Content-Type, etc)
}