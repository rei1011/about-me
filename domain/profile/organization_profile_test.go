package profile

import (
	"about-me/domain/organization"
	"about-me/domain/period"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 組織に在籍していた期間を計算する
func TestCalcPeriod0(t *testing.T) {
	// setup
	ui := "id"
	p, _ := period.NewPeriod(time.Date(2019, 4, 1, 0, 0, 0, 0, time.Local), time.Time{})
	u := []UserProfile{NewEducationalUserProfile(ui, organization.NewOrganization(""), "", p)}
	o := NewOrganizationProfile(organization.NewOrganization(""), NewUserProfiles(ui, u))

	// exercise
	period, _ := o.CalcPeriod()

	// verify
	assert.Equal(t, period.StartDate, time.Date(2019, 4, 1, 0, 0, 0, 0, time.Local))
	assert.Equal(t, period.EndDate, time.Time{})
}

// 組織への在籍期間開始日がゼロ値だった場合
func TestCalcPeriod1(t *testing.T) {
	// setup
	ui := "id"
	p, _ := period.NewPeriod(time.Time{}, time.Time{})
	u := []UserProfile{NewEducationalUserProfile(ui, organization.NewOrganization(""), "", p)}
	o := NewOrganizationProfile(organization.NewOrganization(""), NewUserProfiles(ui, u))

	// exercise
	_, err := o.CalcPeriod()

	// verify
	assert.NotNil(t, err)
}

// 職歴を組織ごとにグルーピングする
func TestNewOrganizationProfileMap0(t *testing.T) {
	// setup
	ui := "id"
	on1 := organization.NewOrganization("Wantedly, Inc")
	on2 := organization.NewOrganization("テスト大学")
	p1, _ := period.NewPeriod(time.Time{}, time.Time{})
	p2, _ := period.NewPeriod(time.Time{}, time.Time{})
	p3, _ := period.NewPeriod(time.Time{}, time.Time{})
	up := NewUserProfiles(ui, []UserProfile{NewWorkUserProfile(ui, on1, "", p1, FullTime), NewWorkUserProfile(ui, on1, "", p2, FullTime), NewEducationalUserProfile(ui, on2, "", p3)})

	// exercise
	result := NewOrganizationProfileMap(up)

	// verify
	assert.Equal(t, len(result.Map[on1].UserProfiles.ProfileList), 2)
	assert.Equal(t, len(result.Map[on2].UserProfiles.ProfileList), 1)
}

// 職歴が無い場合
func TestNewOrganizationProfileMap1(t *testing.T) {
	// setup
	up := UserProfiles{}

	// exercise
	result := NewOrganizationProfileMap(up)

	// verify
	assert.Equal(t, len(result.Map), 0)
}

// 終了時期の降順で並び替える
func TestToListByPeriod0(t *testing.T) {
	// setup
	ui := "id"
	p1, _ := period.NewPeriod(time.Date(2016, 4, 1, 0, 0, 0, 0, time.Local), time.Date(2019, 4, 1, 0, 0, 0, 0, time.Local))
	p2, _ := period.NewPeriod(time.Date(2019, 4, 1, 0, 0, 0, 0, time.Local), time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local))
	up1 := NewEducationalUserProfile(ui, organization.NewOrganization("テスト大学"), "", p1)
	up2 := NewWorkUserProfile(ui, organization.NewOrganization("Wantedly, Inc"), "", p2, FullTime)
	up := NewUserProfiles(ui, []UserProfile{up1, up2})

	// exercise
	result := NewOrganizationProfileMap(up).ToListByPeriod()

	// verify
	assert.Equal(t, "Wantedly, Inc", result[0].Organization.Name)
	assert.Equal(t, "テスト大学", result[1].Organization.Name)
}

// 終了時期がゼロ値だった場合、先頭に配置する
func TestToListByPeriod1(t *testing.T) {
	// setup
	ui := "id"
	period1, _ := period.NewPeriod(time.Date(2016, 4, 1, 0, 0, 0, 0, time.Local), time.Date(2019, 4, 1, 0, 0, 0, 0, time.Local))
	up1 := NewEducationalUserProfile(ui, organization.NewOrganization("テスト大学"), "", period1)
	period2, _ := period.NewPeriod(time.Date(2019, 4, 1, 0, 0, 0, 0, time.Local), time.Time{})
	up2 := NewWorkUserProfile(ui, organization.NewOrganization("Wantedly, Inc"), "", period2, FullTime)
	up := NewUserProfiles(ui, []UserProfile{up1, up2})

	// exercise
	result := NewOrganizationProfileMap(up).ToListByPeriod()

	// verify
	assert.Equal(t, "Wantedly, Inc", result[0].Organization.Name)
	assert.Equal(t, "テスト大学", result[1].Organization.Name)
}

// 終了時期が同じ場合、開始時期の降順で並び替える
func TestToListByPeriod2(t *testing.T) {
	// setup
	ui := "id"
	endDate := time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local)
	period1, _ := period.NewPeriod(time.Date(2016, 4, 1, 0, 0, 0, 0, time.Local), endDate)
	up1 := NewEducationalUserProfile(ui, organization.NewOrganization("テスト大学"), "", period1)
	period2, _ := period.NewPeriod(time.Date(2019, 4, 1, 0, 0, 0, 0, time.Local), endDate)
	up2 := NewWorkUserProfile(ui, organization.NewOrganization("Wantedly, Inc"), "", period2, FullTime)
	up := NewUserProfiles(ui, []UserProfile{up1, up2})

	// exercise
	result := NewOrganizationProfileMap(up).ToListByPeriod()

	// verify
	assert.Equal(t, "Wantedly, Inc", result[0].Organization.Name)
	assert.Equal(t, "テスト大学", result[1].Organization.Name)
}

// プロフィールが空の時
func TestToListByPeriod3(t *testing.T) {
	result := OrganizationProfileMap{}.ToListByPeriod()
	assert.Equal(t, len(result), 0)
}
