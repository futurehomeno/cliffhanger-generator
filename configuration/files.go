package configuration

import (
	"github.com/futurehomeno/cliffhanger-generator/generator"
)

var FileSet = map[string][]*generator.File{
	"common": {
		{Path: "package/build/.gitkeep"},
		{Path: "package/debian/DEBIAN/control"},
		{Path: "package/debian/DEBIAN/preinst"},
		{Path: "package/debian/DEBIAN/postinst"},
		{Path: "package/debian/DEBIAN/prerm"},
		{Path: "package/debian/DEBIAN/postrm"},
		{Path: "package/debian/usr/lib/systemd/system/{{.PackageName}}.service", Permissions: 0644},
		{Path: "src/cmd/root.go"},
		{Path: "src/cmd/factory.go"},
		{Path: "src/internal/config/config.go"},
		{Path: "src/internal/routing/discovery.go"},
		{Path: "src/internal/routing/routing.go"},
		{Path: "src/internal/tasks/tasks.go"},
		{Path: "src/.golangci.yaml"},
		{Path: "src/main.go"},
		{Path: "testdata/data/.gitkeep"},
		{Path: "testdata/defaults/config.json"},
		{Path: "testdata/hub/hub.json"},
		{Path: "testdata/log/.gitkeep"},
		{Path: ".gitignore"},
		{Path: "docker-compose.yaml"},
		{Path: "Makefile"},
		{Path: "README.MD"},
	},
	"edge": {
		{Path: "package/debian/opt/thingsplex/{{.PackageName}}/VERSION", Permissions: 0644},
		{Path: "package/debian/opt/thingsplex/{{.PackageName}}/defaults/config.json", Permissions: 0644},
		{Path: "package/debian/opt/thingsplex/{{.PackageName}}/defaults/app-manifest.json", Permissions: 0644},
		{Path: "src/internal/app/app.go"},
		{Path: "testdata/defaults/app-manifest.json"},
	},
	"edge_adapter": {
		{Path: "package/debian/opt/thingsplex/{{.PackageName}}/defaults/adapter.json", Permissions: 0644},
		{Path: "src/internal/adapter/thing.go"},
		{Path: "testdata/defaults/adapter.json"},
	},
	"core": {
		{Path: "package/debian/etc/futurehome/{{.PackageName}}/config.json", Permissions: 0644},
		{Path: "package/debian/var/lib/futurehome/{{.PackageName}}/VERSION", Permissions: 0644},
	},
	"testing": {
		{Path: "src/cmd/testing.go"},
		{Path: "src/test/helper.go"},
		{Path: "src/{{.ServiceName}}_test.go"},
		{Path: "testdata/testing/not_configured/defaults/config.json"},
	},
	"testing_edge": {
		{Path: "testdata/testing/not_configured/defaults/app-manifest.json"},
	},
	"testing_edge_adapter": {
		{Path: "testdata/testing/not_configured/defaults/adapter.json"},
	},
}
