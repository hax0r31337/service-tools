package service

import (
	"fmt"
	"unsafe"

	"service-tools/utils"

	"github.com/blono/win"
)

func Add(path, name string) error {
	scManager := win.OpenSCManager(nil, nil, win.SC_MANAGER_CREATE_SERVICE|win.SC_MANAGER_LOCK)
	if scManager == 0 {
		return fmt.Errorf("OpenSCManager failed: %v", win.GetLastError())
	}

	service := win.CreateService(scManager, name, name,
		win.SERVICE_ALL_ACCESS, win.SERVICE_WIN32_OWN_PROCESS, win.SERVICE_AUTO_START,
		win.SERVICE_ERROR_NORMAL, path, nil, nil, nil, nil, nil)
	if service == 0 {
		return fmt.Errorf("CreateService failed: %v", win.GetLastError())
	}
	defer win.CloseServiceHandle(service)

	desc := win.SERVICE_DESCRIPTION{
		LpDescription: utils.StringUTF16("サービスの説明"),
	}
	win.LockServiceDatabase(scManager)
	win.ChangeServiceConfig2(service, win.SERVICE_CONFIG_DESCRIPTION, uintptr(unsafe.Pointer(&desc)))
	win.UnlockServiceDatabase(scManager)
	win.StartService(service, 0, nil)

	return nil
}

func Remove(name string) error {
	scManager := win.OpenSCManager(nil, nil, win.SC_MANAGER_CONNECT)
	if scManager == 0 {
		return fmt.Errorf("OpenSCManager failed: %v", win.GetLastError())
	}

	service := win.OpenService(scManager, name, win.SERVICE_STOP|win.SERVICE_QUERY_STATUS|win.DELETE)
	if service == 0 {
		return fmt.Errorf("OpenService failed: %v", win.GetLastError())
	}
	defer win.CloseServiceHandle(service)

	var ss win.SERVICE_STATUS

	win.QueryServiceStatus(service, &ss)
	if ss.DwCurrentState == win.SERVICE_RUNNING {
		win.ControlService(service, win.SERVICE_CONTROL_STOP, &ss)
	}

	win.DeleteService(service)

	return nil
}
