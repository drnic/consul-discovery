package consuldiscovery

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func getCurrentNodeName() string {
	return "drnic.local"
}

func TestHealth(t *testing.T) {
	Convey("HealthByNode", t, func() {
		client := getClient(t)
		checks, err := client.HealthByNode(getCurrentNodeName())
		So(err, ShouldEqual, nil)
		So(len(checks), ShouldEqual, 3)
	})

	Convey("HealthByState", t, func() {
		client := getClient(t)
		checks, err := client.HealthByState("critical")
		So(err, ShouldEqual, nil)
		So(len(checks), ShouldEqual, 1)
	})

	Convey("HealthByService", t, func() {
		client := getClient(t)
		nodes, err := client.HealthByService("simple_service")
		So(err, ShouldEqual, nil)
		So(len(nodes), ShouldEqual, 1)
		node := nodes[0]
		So(node.Service.ServiceID, ShouldEqual, "simple_service")
		So(node.Service.ServiceName, ShouldEqual, "simple_service")
		So(len(node.Checks), ShouldEqual, 3)
		check := node.Checks[0]
		So(check.Status, ShouldEqual, "passing")
	})

}
