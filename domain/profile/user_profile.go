package profile

import (
	"about-me/domain/organization"
	"about-me/domain/period"
	"errors"
	"sort"
)

type UserProfiles struct {
	UserId      string
	ProfileList []UserProfile
}

func NewUserProfiles(ui string, p []UserProfile) UserProfiles {
	return UserProfiles{
		UserId:      ui,
		ProfileList: p,
	}
}

// 所属終了時期が遅い順で並び替える
// 終了時期がゼロ値の場合は先頭に配置する
// 所属終了時期が同じ場合、所属開始時期が遅い順で並び替える
// 所属開始時期および所属終了時期が同じ場合、組織のIDで並び替える
func (up UserProfiles) SortByPeriod() {
	sort.SliceStable(up.ProfileList, func(i, j int) bool {
		eleI := up.ProfileList[i]
		eleJ := up.ProfileList[j]
		endDateI := eleI.Period.EndDate
		endDateJ := eleJ.Period.EndDate

		if endDateI.IsZero() {
			return true
		}

		if endDateJ.IsZero() {
			return false
		}

		if endDateI != endDateJ {
			return endDateI.After(endDateJ)
		}

		startDateI := eleI.Period.StartDate
		startDateJ := eleJ.Period.StartDate
		if startDateI != startDateJ {
			return startDateI.After(startDateJ)
		}

		return eleI.Organization.Id.ID() > eleJ.Organization.Id.ID()
	})
}

// 職歴と学歴でgroupingする
func (up UserProfiles) GroupByType() map[string]UserProfiles {
	g := map[string]UserProfiles{}
	for _, v := range up.ProfileList {
		t := v.ProfileType
		value, found := g[t]
		if !found {
			switch t {
			case WorkHistory:
				g[t] = NewUserProfiles(v.UserId, []UserProfile{v})
			case EducationalBackground:
				g[t] = NewUserProfiles(v.UserId, []UserProfile{NewEducationalUserProfile(v.UserId, v.Organization, v.Specialization, v.Period)})
			}
		} else {
			pl := append(value.ProfileList, v)
			up := NewUserProfiles(v.UserId, pl)
			g[t] = up
		}
	}

	for _, e := range g {
		e.SortByPeriod()
	}

	return g
}

type UserProfile struct {
	UserId         string
	ProfileType    string
	Organization   organization.Organization
	Specialization string
	Period         period.Period
	jobType        string
}

func (u UserProfile) JobType() (string, error) {
	// 学歴に雇用形態は紐づかないためerrorを投げる
	if u.ProfileType == EducationalBackground {
		return "", errors.New("EducationalBackground has no JobType")
	}

	return u.jobType, nil
}

func NewEducationalUserProfile(ui string, or organization.Organization, special string, p period.Period) UserProfile {
	return UserProfile{
		UserId:         ui,
		ProfileType:    EducationalBackground,
		Organization:   or,
		Specialization: special,
		Period:         p,
		jobType:        "",
	}
}

func NewWorkUserProfile(ui string, or organization.Organization, special string, p period.Period, jt string) UserProfile {
	return UserProfile{
		UserId:         ui,
		ProfileType:    WorkHistory,
		Organization:   or,
		Specialization: special,
		Period:         p,
		jobType:        jt,
	}
}
