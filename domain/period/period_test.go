package period

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 開始時期と終了時期を受け取って在籍期間を作成する
func TestNewPeriod0(t *testing.T) {
	s := time.Date(2018, time.April, 1, 0, 0, 0, 0, time.Local)
	e := time.Date(2019, time.March, 31, 0, 0, 0, 0, time.Local)
	_, err := NewPeriod(s, e)
	assert.Nil(t, err)
}

// 開始時期がゼロ値だった場合errorを返す
func TestNewPeriod1(t *testing.T) {
	s := time.Time{}
	e := time.Date(2019, time.March, 31, 0, 0, 0, 0, time.Local)
	_, err := NewPeriod(s, e)
	assert.NotNil(t, err)
}

// 開始時期が終了時期よりも後だった場合errorを返す
func TestNewPeriod2(t *testing.T) {
	s := time.Date(2019, time.March, 31, 0, 0, 0, 0, time.Local)
	e := time.Date(2019, time.March, 30, 0, 0, 0, 0, time.Local)
	_, err := NewPeriod(s, e)
	assert.NotNil(t, err)
}

// クライアント用のフォーマットに整形する
func TestDisplayPeriod0(t *testing.T) {
	s := time.Date(2018, time.April, 1, 0, 0, 0, 0, time.Local)
	e := time.Date(2019, time.March, 31, 0, 0, 0, 0, time.Local)
	period, _ := NewPeriod(s, e)
	start, end := period.DisplayPeriod()
	assert.Equal(t, "2018/04", start)
	assert.Equal(t, "2019/03", end)
}

// 在籍中の場合は"now"を返却する
func TestDisplayPeriod1(t *testing.T) {
	s := time.Date(2018, time.April, 1, 0, 0, 0, 0, time.Local)
	e := time.Time{}
	period, _ := NewPeriod(s, e)
	start, end := period.DisplayPeriod()
	assert.Equal(t, "2018/04", start)
	assert.Equal(t, "now", end)
}
