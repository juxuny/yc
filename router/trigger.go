package router

import "context"

type Trigger int

const (
	TriggerSignValidator = Trigger(1)
)

type TriggerFunc func(ctx context.Context, body []byte) error

type TriggerConfig map[Trigger]TriggerFunc

func (t TriggerConfig) Merge(v TriggerConfig) {
	for k, v := range v {
		t[k] = v
	}
}
