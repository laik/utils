package utils

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
)

func TestQueryAll(t *testing.T) {

	sqlStmt := "select id,username,workunitid from etbasedata.pub_user where username=?"

	type user struct {
		Id         int
		Username   string
		Workunitid int
	}

	db, err := sql.Open("mysql", "dxp:dxpplatform@tcp(192.168.4.113:31863)/etbasedata")
	if err != nil {
		t.Fatalf("initial database error %s...\n", err)
	}

	Convey("test initial done", t, func() {
		_user := &user{}

		rs, err := db.QueryContext(context.Background(), sqlStmt, "cgz")

		So(err, ShouldBeNil)

		err = One(_user, rs)

		So(err, ShouldBeNil)

		So(_user.Id, ShouldEqual, 2852)
		So(_user.Username, ShouldEqual, "cgz")
		So(_user.Workunitid, ShouldEqual, 3543)

	})

	// many
	Convey("test initial done", t, func() {
		users := new([]user)

		rs, err := db.Query(sqlStmt, "cgz")

		So(err, ShouldBeNil)

		err = All(users, rs)

		So(err, ShouldBeNil)

		for _, v := range *users {

			So(v.Id, ShouldEqual, 2852)
			So(v.Username, ShouldEqual, "cgz")
			So(v.Workunitid, ShouldEqual, 3543)
		}

	})
}
