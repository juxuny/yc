package dt

import (
	"encoding/json"
	"math/rand"
	"testing"
	"time"
)

func TestNewID(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var data = struct {
		Id     ID `json:"id"`
		UserId ID `json:"userId"`
	}{
		Id:     NewID(1000),
		UserId: NewID(rand.Uint64()),
	}
	jsonData, _ := json.Marshal(data)
	t.Log(string(jsonData))
	var parseData = struct {
		Id     ID `json:"id"`
		UserId ID `json:"userId"`
	}{}
	err := json.Unmarshal(jsonData, &parseData)
	if err != nil {
		t.Fatal(err)
	}
	if !data.Id.Equal(parseData.Id) || !data.UserId.Equal(parseData.UserId) {
		t.Fatal("parse failed")
	}
	t.Log(parseData.UserId.String())
}
