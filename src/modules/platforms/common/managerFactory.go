package manager

import (
	"parsdevkit.net/models"

	angularManager "parsdevkit.net/platforms/angular/managers"
	"parsdevkit.net/platforms/core"
	dotnetManager "parsdevkit.net/platforms/dotnet/managers"
	goManager "parsdevkit.net/platforms/go/managers"
	nodejsManager "parsdevkit.net/platforms/nodejs/managers"
	parsManager "parsdevkit.net/platforms/pars/managers"
)

func ManagerFactory(platform models.PlatformType) core.ManagerInterface {
	switch platform {
	case models.PlatformTypes.Pars:
		return parsManager.NewParsManager()
	case models.PlatformTypes.Dotnet:
		return dotnetManager.NewDotnetManager()
	case models.PlatformTypes.Angular:
		return angularManager.NewAngularManager()
	case models.PlatformTypes.NodeJS:
		return nodejsManager.NewNodeJSManager()
	case models.PlatformTypes.GO:
		return goManager.NewGoManager()
	default:
		return nil
	}
}
