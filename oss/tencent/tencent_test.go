package tencent

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jinzhu/configor"
	"github.com/qor5/x/v3/oss/tests"
)

var client *Client

func init() {
	cfg := Config{}
	configor.New(&configor.Config{ENVPrefix: "TENCENT"}).Load(&cfg)

	if len(cfg.AppID) == 0 {
		fmt.Println("No tencent configuration")
		return
	}

	client = New(&cfg)
}

func TestClient_Get(t *testing.T) {

}

func TestClient_Put(t *testing.T) {
	if client == nil {
		t.Skip(`skip because of no config: `)
	}

	f, err := os.ReadFile("/home/owen/Downloads/2.png")
	if err != nil {
		t.Error(err)
		return
	}

	client.Put(context.Background(), "test.png", bytes.NewReader(f))
}

func TestClient_Put2(t *testing.T) {
	if client == nil {
		t.Skip(`skip because of no config: `)
	}

	tests.TestAll(client, t)
}

func TestClient_Delete(t *testing.T) {
	if client == nil {
		t.Skip(`skip because of no config: `)
	}

	fmt.Println(client.Delete(context.Background(), "test.png"))
}
