package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfileHandler(t *testing.T) {
	ts := httptest.NewServer(SetUpServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/profile/test", ts.URL))

	expect := `{
    "list": [
        {
            "organizationName": "Wantedly, Inc",
            "startDate": "2019/04",
            "endDate": "now",
            "profiles": [
                {
                    "startDate": "2021/03",
                    "endDate": "now",
                    "specialization": "Backend Engineer",
                    "jobType": "FullTime"
                },
                {
                    "startDate": "2019/04",
                    "endDate": "2021/02",
                    "specialization": "Frontend Engineer",
                    "jobType": "FullTime"
                }
            ]
        },
        {
            "organizationName": "テスト大学",
            "startDate": "2016/04",
            "endDate": "2019/03",
            "profiles": [
                {
                    "startDate": "2016/04",
                    "endDate": "2019/03",
                    "specialization": "理学部 情報学科",
                    "jobType": ""
                }
            ]
        }
    ]
}`
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	responseData, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, expect, string(responseData))
}
