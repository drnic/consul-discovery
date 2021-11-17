package consuldiscovery

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func getClient(t *testing.T) *Client {
	client, err := NewClient(DefaultConfig())
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	return client
}

func TestCatalog(t *testing.T) {
	Convey("CatalogServices", t, func() {
		client := getClient(t)
		services, err := client.CatalogServices()
		So(err, ShouldEqual, nil)
		So(len(services), ShouldEqual, 2)
		So(services[0].Name, ShouldEqual, "consul")
		So(services[1].Name, ShouldEqual, "simple_service")
	})

	Convey("CatalogServiceByName", t, func() {
		client := getClient(t)
		nodes, err := client.CatalogServiceByName("simple_service")
		So(err, ShouldEqual, nil)
		So(len(nodes), ShouldEqual, 1)
		So(nodes[0].ServiceID, ShouldEqual, "simple_service")
		So(nodes[0].ServiceName, ShouldEqual, "simple_service")
		So(nodes[0].ServicePort, ShouldEqual, 6666)
		// TODO: So(nodes[0].ServiceTags, ShouldEqual, []string{"tag1", "tag2"})
		// TODO: and other fields...
	})

}
