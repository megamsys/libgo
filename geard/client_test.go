package geard

import (
	"os"
	"testing"
)


func TestPersistence(t *testing.T) {
	_ = NewClient("", "")


	fo, err := os.Create("config.json")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

}
