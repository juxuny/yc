package utils

import (
	"math/rand"
	"time"
)

var randInstance *rand.Rand

func init() {
	randInstance = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func getRandInstance() *rand.Rand {
	randInstance.Seed(time.Now().UnixNano())
	return randInstance
}

type randHelper struct{}

var RandHelper = randHelper{}

func (randHelper) RandString(n int) string {
	return StringHelper.RandString(n)
}

func (randHelper) RandNumString(n int) string {
	return StringHelper.RandNumString(n)
}

func (randHelper) Float64(ratio float64) float64 {
	return getRandInstance().Float64() * ratio
}

func (randHelper) Int63() int64 {
	return getRandInstance().Int63()
}

// return [0, max)
func (randHelper) Int63n(max int64) int64 {
	return getRandInstance().Int63n(max)
}
