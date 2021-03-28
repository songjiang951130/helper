package mapper_test

import (
	"helper/dao/mapper"
	"testing"
)

func TestConnect(t *testing.T) {
	t.Log("sfadfsafas")
	mapper.Connect()
}

func TestList(t *testing.T) {
	mapper.GetSubscribeList()
}

func TestLowList(t *testing.T) {
	symbol := "SZ002511"
	mapper.GetDailyList(symbol)
}
