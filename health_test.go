package consuldiscovery

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHealth(t *testing.T) {
	Convey("ServiceHealth", t, func() {
		client := getClient(t)
		nodes, err := client.ServiceHealth("simple_service")
		So(err, ShouldEqual, nil)
		So(len(nodes), ShouldEqual, 1)
		node := nodes[0]
		So(node.Service.ServiceID, ShouldEqual, "simple_service")
		So(node.Service.ServiceName, ShouldEqual, "simple_service")
		So(len(node.Checks), ShouldEqual, 1)
		check := node.Checks[0]
		So(check.Status, ShouldEqual, "passing")
	})
}
