package cos

import (
	"context"
	"github.com/juxuny/yc/log"
	"testing"
)

func TestClient_Health(t *testing.T) {
	resp, err := DefaultClient.Health(context.Background(), &HealthRequest{})
	if err != nil {
		log.Error(err)
		t.Fatal(err)
	}
	t.Log(resp.CurrentTime)
}
