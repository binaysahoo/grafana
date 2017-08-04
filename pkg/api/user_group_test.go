package api

import (
	"testing"

	"github.com/grafana/grafana/pkg/bus"
	"github.com/grafana/grafana/pkg/components/simplejson"
	"github.com/grafana/grafana/pkg/models"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUserGroupApiEndpoint(t *testing.T) {
	Convey("Given two user groups", t, func() {
		mockResult := models.SearchUserGroupQueryResult{
			UserGroups: []*models.UserGroup{
				{Name: "userGroup1"},
				{Name: "userGroup2"},
			},
			TotalCount: 2,
		}

		Convey("When searching with no parameters", func() {
			loggedInUserScenario("When calling GET on", "/api/user-groups/search", func(sc *scenarioContext) {
				var sentLimit int
				var sendPage int
				bus.AddHandler("test", func(query *models.SearchUserGroupsQuery) error {
					query.Result = mockResult

					sentLimit = query.Limit
					sendPage = query.Page

					return nil
				})

				sc.handlerFunc = SearchUserGroups
				sc.fakeReqWithParams("GET", sc.url, map[string]string{}).exec()

				So(sentLimit, ShouldEqual, 1000)
				So(sendPage, ShouldEqual, 1)

				respJSON, err := simplejson.NewJson(sc.resp.Body.Bytes())
				So(err, ShouldBeNil)

				So(respJSON.Get("totalCount").MustInt(), ShouldEqual, 2)
				So(len(respJSON.Get("userGroups").MustArray()), ShouldEqual, 2)
			})
		})

		Convey("When searching with page and perpage parameters", func() {
			loggedInUserScenario("When calling GET on", "/api/user-groups/search", func(sc *scenarioContext) {
				var sentLimit int
				var sendPage int
				bus.AddHandler("test", func(query *models.SearchUserGroupsQuery) error {
					query.Result = mockResult

					sentLimit = query.Limit
					sendPage = query.Page

					return nil
				})

				sc.handlerFunc = SearchUserGroups
				sc.fakeReqWithParams("GET", sc.url, map[string]string{"perpage": "10", "page": "2"}).exec()

				So(sentLimit, ShouldEqual, 10)
				So(sendPage, ShouldEqual, 2)
			})
		})
	})
}
