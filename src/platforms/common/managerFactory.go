package common

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

func GetPlatformManager(platform models.PlatformType, platformRegistry map[models.PlatformType]func() core.ManagerInterface) (core.ManagerInterface, error) {
	factory, ok := platformRegistry[platform]
	if !ok {
		return nil, fmt.Errorf("xxx: Platform Manager bulunamadı '%s'", platform)
	}
	manager := factory()
	return manager, nil
}

var Registry = map[models.PlatformType]func() core.ManagerInterface{
	models.PlatformTypes.Pars:    func() core.ManagerInterface { return parsManager.NewParsManager() },
	models.PlatformTypes.Dotnet:  func() core.ManagerInterface { return dotnetManager.NewDotnetManager() },
	models.PlatformTypes.Angular: func() core.ManagerInterface { return angularManager.NewAngularManager() },
	models.PlatformTypes.NodeJS:  func() core.ManagerInterface { return nodejsManager.NewNodeJSManager() },
	models.PlatformTypes.GO:      func() core.ManagerInterface { return goManager.NewGoManager() },
}
