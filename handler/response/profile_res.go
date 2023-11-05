package response

import (
	"about-me/domain/profile"
)

type ProfileListRes struct {
	List []Profiles `json:"list"`
}

func NewProfileListRes(op []profile.OrganizationProfile) ProfileListRes {
	var profiles []Profiles
	for _, v := range op {
		pl := NewProfiles(v)
		profiles = append(profiles, pl)
	}

	return ProfileListRes{
		List: profiles,
	}
}

type Profiles struct {
	OrganizationName string    `json:"organizationName"`
	StartDate        string    `json:"startDate"`
	EndDate          string    `json:"endDate"`
	Profiles         []Profile `json:"profiles"`
}

func NewProfiles(o profile.OrganizationProfile) Profiles {
	period, _ := o.CalcPeriod()
	start, end := period.DisplayPeriod()
	var profiles []Profile
	for _, v := range o.UserProfiles.ProfileList {
		p := NewProfile(v)
		profiles = append(profiles, p)
	}
	p := Profiles{
		OrganizationName: o.Organization.Name,
		StartDate:        start,
		EndDate:          end,
		Profiles:         profiles,
	}
	return p
}

type Profile struct {
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	Specialization string `json:"specialization"`
	JobType        string `json:"jobType"`
}

func NewProfile(up profile.UserProfile) Profile {
	period := up.Period
	jt, _ := up.JobType()
	start, end := period.DisplayPeriod()
	return Profile{
		StartDate:      start,
		EndDate:        end,
		Specialization: up.Specialization,
		JobType:        jt,
	}
}
