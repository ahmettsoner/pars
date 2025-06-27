package manager

import (
	"fmt"

	"parsdevkit.net/models"

	angularManager "parsdevkit.net/platforms/angular/managers"
	"parsdevkit.net/platforms/core"
	dotnetManager "parsdevkit.net/platforms/dotnet/managers"
	goManager "parsdevkit.net/platforms/go/managers"
	nodejsManager "parsdevkit.net/platforms/nodejs/managers"
	parsManager "parsdevkit.net/platforms/pars/managers"
)

func ManagerFactory(platform models.PlatformType) (core.ManagerInterface, error) {
	switch platform {
	case models.PlatformTypes.Pars:
		return parsManager.NewParsManager(), nil
	case models.PlatformTypes.Dotnet:
		return dotnetManager.NewDotnetManager(), nil
	case models.PlatformTypes.Angular:
		return angularManager.NewAngularManager(), nil
	case models.PlatformTypes.NodeJS:
		return nodejsManager.NewNodeJSManager(), nil
	case models.PlatformTypes.GO:
		return goManager.NewGoManager(), nil
	default:
		return nil, fmt.Errorf("xxx: Platorm '%s' tanımlı değil", platform)
	}
}
