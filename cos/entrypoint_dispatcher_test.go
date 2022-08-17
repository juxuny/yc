package cos

import "testing"

func TestEntrypointDispatcher_SelectOne(t *testing.T) {
	dispatcher := NewEntrypointDispatcher("activity", Options{
		CosEntrypoint: testEnv.Entrypoint,
		ConfigId:      testEnv.ConfigId,
		AccessKey:     testEnv.AccessKey,
		Secret:        testEnv.Secret,
	})
	entrypointCandidate := dispatcher.SelectOne()
	t.Log(entrypointCandidate)
}
