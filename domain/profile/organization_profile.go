package profile

import (
	"about-me/domain/organization"
	"about-me/domain/period"
	"sort"
	"time"
)

// 所属組織ごとに経歴を保持するためのmap
type OrganizationProfileMap struct {
	Map map[organization.Organization]OrganizationProfile
}

func NewOrganizationProfileMap(u UserProfiles) OrganizationProfileMap {
	mp := map[organization.Organization]OrganizationProfile{}
	for _, v := range u.ProfileList {
		name := v.Organization
		value, found := mp[name]
		if !found {
			mp[name] = NewOrganizationProfile(name, NewUserProfiles(v.UserId, []UserProfile{v}))
		} else {
			pl := append(value.UserProfiles.ProfileList, v)
			op := NewOrganizationProfile(value.Organization, NewUserProfiles(value.UserId, pl))
			mp[name] = op
		}
	}

	return OrganizationProfileMap{
		Map: mp,
	}
}

// 組織への所属終了時期が遅い順でlistへ変換する
func (om OrganizationProfileMap) ToListByPeriod() []OrganizationProfile {
	var list []OrganizationProfile
	for _, v := range om.Map {
		v.UserProfiles.SortByPeriod()
		op := NewOrganizationProfile(v.Organization, v.UserProfiles)
		list = append(list, op)
	}

	sort.SliceStable(list, func(i, j int) bool {
		periodI, _ := list[i].CalcPeriod()
		periodJ, _ := list[j].CalcPeriod()

		if periodI.EndDate.IsZero() {
			return true
		}

		if periodJ.EndDate.IsZero() {
			return false
		}

		if periodI.EndDate != periodJ.EndDate {
			return periodI.EndDate.After(periodJ.EndDate)
		}

		if periodI.StartDate != periodJ.StartDate {
			return periodI.StartDate.After(periodJ.StartDate)
		}

		return list[i].Organization.Id.ID() > list[j].Organization.Id.ID()
	})
	return list
}

type OrganizationProfile struct {
	Organization organization.Organization
	UserId       string
	UserProfiles UserProfiles
}

func NewOrganizationProfile(on organization.Organization, up UserProfiles) OrganizationProfile {
	op := OrganizationProfile{
		Organization: on,
		UserId:       up.UserId,
		UserProfiles: NewUserProfiles(up.UserId, up.ProfileList),
	}

	return op
}

func (o *OrganizationProfile) CalcPeriod() (p period.Period, err error) {
	var minStartDate time.Time
	var maxEndDate time.Time
	endDateIsZeroExists := false

	for _, v := range o.UserProfiles.ProfileList {
		if minStartDate.IsZero() {
			minStartDate = v.Period.StartDate
		} else {
			if v.Period.StartDate.Before(minStartDate) {
				minStartDate = v.Period.StartDate
			}
		}

		if !endDateIsZeroExists {
			if v.Period.EndDate.IsZero() {
				maxEndDate = v.Period.EndDate
				endDateIsZeroExists = true
			} else if maxEndDate.IsZero() {
				maxEndDate = v.Period.EndDate
			} else if v.Period.EndDate.After(maxEndDate) {
				maxEndDate = v.Period.EndDate
			}
		}
	}

	return period.NewPeriod(minStartDate, maxEndDate)
}
