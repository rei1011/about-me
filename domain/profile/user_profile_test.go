package profile

import (
	"about-me/domain/organization"
	"about-me/domain/period"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 終了時期が遅い順で並び替えられること
func TestSortByPeriod0(t *testing.T) {
	// setup
	ui := "id"
	period1, _ := period.NewPeriod(time.Date(2019, time.February, 1, 1, 0, 0, 0, time.Local), time.Date(2020, time.February, 1, 1, 0, 0, 0, time.Local))
	p1 := NewWorkUserProfile(ui, organization.NewOrganization("first company"), "", period1, FullTime)
	period2, _ := period.NewPeriod(time.Date(2020, time.February, 1, 2, 0, 0, 0, time.Local), time.Date(2021, time.February, 1, 1, 0, 0, 0, time.Local))
	p2 := NewWorkUserProfile(ui, organization.NewOrganization("second company"), "", period2, FullTime)
	up := NewUserProfiles(ui, []UserProfile{p1, p2})

	// exercise
	up.SortByPeriod()

	// verify
	assert.Equal(t, "second company", up.ProfileList[0].Organization.Name)
	assert.Equal(t, "first company", up.ProfileList[1].Organization.Name)
}

// 終了時期がゼロ値の場合、先頭になること
func TestSortByPeriod1(t *testing.T) {
	// setup
	ui := "id"
	period1, _ := period.NewPeriod(time.Date(2019, time.February, 1, 1, 0, 0, 0, time.Local), time.Date(2020, time.February, 1, 1, 0, 0, 0, time.Local))
	p1 := NewWorkUserProfile(ui, organization.NewOrganization("first company"), "", period1, FullTime)
	period2, _ := period.NewPeriod(time.Date(2020, time.February, 1, 2, 0, 0, 0, time.Local), time.Time{})
	p2 := NewWorkUserProfile(ui, organization.NewOrganization("second company"), "", period2, FullTime)
	up := NewUserProfiles(ui, []UserProfile{p1, p2})

	// exercise
	up.SortByPeriod()

	//verify
	assert.Equal(t, "second company", up.ProfileList[0].Organization.Name)
	assert.Equal(t, "first company", up.ProfileList[1].Organization.Name)
}

// 終了時期が同じ場合、開始時期の降順になること
func TestSortByPeriod2(t *testing.T) {
	// setup
	ui := "id"
	period1, _ := period.NewPeriod(time.Date(2020, time.February, 1, 1, 0, 0, 0, time.Local), time.Date(2022, time.February, 1, 1, 0, 0, 0, time.Local))
	p1 := NewWorkUserProfile(ui, organization.NewOrganization("first company"), "", period1, FullTime)
	period2, _ := period.NewPeriod(time.Date(2020, time.February, 1, 2, 0, 0, 0, time.Local), time.Date(2022, time.February, 1, 1, 0, 0, 0, time.Local))
	p2 := NewWorkUserProfile(ui, organization.NewOrganization("second company"), "", period2, FullTime)
	up := NewUserProfiles(ui, []UserProfile{p1, p2})

	// exercise
	up.SortByPeriod()

	// verify
	assert.Equal(t, "second company", up.ProfileList[0].Organization.Name)
	assert.Equal(t, "first company", up.ProfileList[1].Organization.Name)
}

// 職歴と学歴でグルーピングされること
func TestGroupByType1(t *testing.T) {
	ui := "id"
	period1, _ := period.NewPeriod(time.Date(2019, time.April, 1, 0, 0, 0, 0, time.Local), time.Date(2020, time.March, 1, 0, 0, 0, 0, time.Local))
	wo := organization.NewOrganization("Wantedly, Inc")
	p1 := NewWorkUserProfile(ui, wo, "Backend Engineer", period1, FullTime)

	period2, _ := period.NewPeriod(time.Date(2020, time.April, 1, 0, 0, 0, 0, time.Local), time.Date(2021, time.April, 1, 0, 0, 0, 0, time.Local))
	p2 := NewWorkUserProfile(ui, wo, "Frontend Engineer", period2, FullTime)

	period3, _ := period.NewPeriod(time.Date(2016, time.April, 1, 0, 0, 0, 0, time.Local), time.Date(2019, time.March, 1, 0, 0, 0, 0, time.Local))
	eo := organization.NewOrganization("テスト大学")
	p3 := NewEducationalUserProfile(ui, eo, "理工学部 情報学科", period3)
	up := NewUserProfiles(ui, []UserProfile{p1, p2, p3})

	result := up.GroupByType()

	assert.Equal(t, len(result[WorkHistory].ProfileList), 2)
	assert.Equal(t, len(result[EducationalBackground].ProfileList), 1)
}

// グルーピングされてかつ、終了時期の降順になっていること
func TestGroupByType2(t *testing.T) {
	ui := "id"
	period1, _ := period.NewPeriod(time.Date(2019, time.April, 1, 0, 0, 0, 0, time.Local), time.Date(2020, time.March, 1, 0, 0, 0, 0, time.Local))
	wo := organization.NewOrganization("Wantedly, Inc")
	p1 := NewWorkUserProfile(ui, wo, "Backend Engineer", period1, FullTime)

	period2, _ := period.NewPeriod(time.Date(2020, time.April, 1, 0, 0, 0, 0, time.Local), time.Date(2021, time.April, 1, 0, 0, 0, 0, time.Local))
	p2 := NewWorkUserProfile(ui, wo, "Frontend Engineer", period2, FullTime)

	period3, _ := period.NewPeriod(time.Date(2016, time.April, 1, 0, 0, 0, 0, time.Local), time.Date(2019, time.March, 1, 0, 0, 0, 0, time.Local))
	eo := organization.NewOrganization("テスト大学")
	p3 := NewEducationalUserProfile(ui, eo, "理工学部 情報学科", period3)
	up := NewUserProfiles(ui, []UserProfile{p1, p2, p3})

	result := up.GroupByType()

	assert.Equal(t, len(result[WorkHistory].ProfileList), 2)
	assert.Equal(t, len(result[EducationalBackground].ProfileList), 1)
}

// プロフィールが空の時
func TestGroupByType3(t *testing.T) {
	up := UserProfiles{}
	result := up.GroupByType()
	assert.Equal(t, len(result), 0)
}
