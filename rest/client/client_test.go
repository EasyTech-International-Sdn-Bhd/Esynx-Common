package client

import "testing"

func TestNewApiClient(t *testing.T) {
	c, err := NewApiClient(&RestClientParams{
		URL:      "http://localhost:9090/",
		AppName:  "Esynx",
		Username: "vitsync",
		Password: "vitsyncprogram",
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.GetBulkAgent()
	if err != nil {
		t.Fatal(err)
	}
}
