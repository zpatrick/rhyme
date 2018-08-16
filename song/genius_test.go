package song

import (
	"fmt"
	"testing"
)

var token = "2g7h3BzGUd-PU9UGy2c6l4qB8uDi97BMpUqe_MWp8h0BRByRaNBgAhwaAUzJw1hU"

func TestClient(t *testing.T) {
	genius := NewGeniusClient(token)
	song, err := genius.Search("paper planes", "mia")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", song)
}
