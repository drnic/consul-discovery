package consuldiscovery

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func getClient(t *testing.T) *Client {
	client, err := NewClient(DefaultConfig())
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	return client
}

func TestClient(t *testing.T) {
	Convey("Client", t, func() {
		client := getClient(t)
		services := client.ServiceList()
		So(len(services), ShouldEqual, 2)
		So(services[0].Name, ShouldEqual, "consul")
		So(services[1].Name, ShouldEqual, "simple_service")
	})
}
