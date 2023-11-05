package period

import (
	"errors"
	"time"
)

// 在籍期間を表現するドメイン
type Period struct {
	StartDate time.Time
	EndDate   time.Time
}

func NewPeriod(s time.Time, e time.Time) (period Period, err error) {
	if s.IsZero() {
		return Period{
			StartDate: time.Time{},
			EndDate:   time.Time{},
		}, errors.New("star date is zero value")
	}

	if s.After(e) && !e.IsZero() {
		return Period{
			StartDate: time.Time{},
			EndDate:   time.Time{},
		}, errors.New("start date is later than end date")
	}

	return Period{
		StartDate: s,
		EndDate:   e,
	}, nil
}

func (p Period) DisplayPeriod() (startDate string, endDate string) {
	if p.EndDate.IsZero() {
		return p.StartDate.Format("2006/01"), "now"
	}

	return p.StartDate.Format("2006/01"), p.EndDate.Format("2006/01")
}
