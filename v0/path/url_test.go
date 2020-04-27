package path_test

import (
	"testing"

	path2 "github.com/VEuPathDB/lib-go-wdk-api/v0/path"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewApiUrl(t *testing.T) {
	Convey("NewApiUrl", t, func() {
		sites := [][2]string{
			{"plasmodb.org", "plasmo"},
			{"toxodb.org", "toxo"},
		}
		prefixes := []string{"", "http://", "https://"}

		for _, site := range sites {
			for _, prefix := range prefixes {
				path := prefix + site[0]
				Convey(path, func() {
					Convey("Without app path", func() {
						a, b := path2.NewApiUrl(path)
						So(b, ShouldBeNil)
						So(a.String(), ShouldEqual, "https://"+site[0]+"/"+site[1]+"/")
					})
					Convey("Without app path with trailing slash", func() {
						a, b := path2.NewApiUrl(path + "/")
						So(b, ShouldBeNil)
						So(a.String(), ShouldEqual, "https://"+site[0]+"/"+site[1]+"/")
					})
					Convey("With app path", func() {
						a, b := path2.NewApiUrl(path + "/" + site[1])
						So(b, ShouldBeNil)
						So(a.String(), ShouldEqual, "https://"+site[0]+"/"+site[1]+"/")
					})
					Convey("With app path with trailing slash", func() {
						a, b := path2.NewApiUrl(path + "/" + site[1] + "/")
						So(b, ShouldBeNil)
						So(a.String(), ShouldEqual, "https://"+site[0]+"/"+site[1]+"/")
					})
					Convey("Without app path with query", func() {
						a, b := path2.NewApiUrl(path + "?query=yay")
						So(b, ShouldBeNil)
						So(a.String(), ShouldEqual, "https://"+site[0]+"/"+site[1]+"/"+"?query=yay")
					})
					Convey("Without app path with trailing slash and query", func() {
						a, b := path2.NewApiUrl(path + "/" + "?query=yay")
						So(b, ShouldBeNil)
						So(a.String(), ShouldEqual, "https://"+site[0]+"/"+site[1]+"/"+"?query=yay")
					})
					Convey("With app path with query", func() {
						a, b := path2.NewApiUrl(path + "/" + site[1] + "?query=yay")
						So(b, ShouldBeNil)
						So(a.String(), ShouldEqual, "https://"+site[0]+"/"+site[1]+"/"+"?query=yay")
					})
					Convey("With app path with trailing slash and query", func() {
						a, b := path2.NewApiUrl(path + "/" + site[1] + "/?query=yay")
						So(b, ShouldBeNil)
						So(a.String(), ShouldEqual, "https://"+site[0]+"/"+site[1]+"/"+"?query=yay")
					})
				})
			}
		}
	})
}
