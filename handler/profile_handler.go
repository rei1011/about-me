package handler

import (
	"about-me/domain/organization"
	"about-me/domain/period"
	"about-me/domain/profile"
	"about-me/handler/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	// TODO usecaseからプロフィール情報を取得できるようにする
	o1 := organization.NewOrganization("Wantedly, Inc")
	o2 := organization.NewOrganization("テスト大学")
	ui := "id"
	period1, _ := period.NewPeriod(time.Date(2019, 4, 1, 0, 0, 0, 0, time.Local), time.Date(2021, 2, 1, 0, 0, 0, 0, time.Local))
	period2, _ := period.NewPeriod(time.Date(2021, 3, 1, 0, 0, 0, 0, time.Local), time.Time{})
	period3, _ := period.NewPeriod(time.Date(2016, 4, 1, 0, 0, 0, 0, time.Local), time.Date(2019, 3, 1, 0, 0, 0, 0, time.Local))
	ps := profile.NewUserProfiles(ui, []profile.UserProfile{
		profile.NewWorkUserProfile(ui, o1, "Frontend Engineer", period1, profile.FullTime),
		profile.NewWorkUserProfile(ui, o1, "Backend Engineer", period2, profile.FullTime),
		profile.NewEducationalUserProfile(ui, o2, "理学部 情報学科", period3),
	})
	om := profile.NewOrganizationProfileMap(ps)
	list := om.ToListByPeriod()

	c.IndentedJSON(http.StatusOK, response.NewProfileListRes(list))
}
