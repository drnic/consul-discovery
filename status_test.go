package consuldiscovery

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStatus(t *testing.T) {
	Convey("StatusLeader", t, func() {
		client := getClient(t)
		_, err := client.StatusLeader()
		So(err, ShouldEqual, nil)
	})

	Convey("StatusPeers", t, func() {
		client := getClient(t)
		peers, err := client.StatusPeers()
		So(err, ShouldEqual, nil)
		So(len(peers), ShouldEqual, 1)
	})
}
