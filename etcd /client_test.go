package etcd

import (
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"os"
	"testing"
)


func TestPersistence(t *testing.T) {
	c := NewClient(nil)


	fo, err := os.Create("config.json")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	c.SetPersistence(fo)
	
}
