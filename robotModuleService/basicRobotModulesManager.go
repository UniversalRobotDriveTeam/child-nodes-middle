package robotModuleService

import (
	_const12 "github.com/UniversalRobotDriveTeam/child-nodes-middle/const"
	"time"
)

// StartRobotModulesMonitor 开始子节点模块监控器
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) StartRobotModulesMonitor() {
	manager.robotMonitorIsRun = true
	go manager.robotModulesMonitor()
}

// StopRobotModulesMonitor 停止子节点模块监控器
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) StopRobotModulesMonitor() {
	*manager.stopRobotMonitorChannel <- 0
	manager.robotMonitorIsRun = false
}

// 子节点监控器 负责启动子模块
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) robotModulesMonitor() {
	for {
		time.Sleep(time.Millisecond * 50)
		select {
		case <-*manager.stopRobotMonitorChannel:
			break
		default:
			if manager.RobotBasicModulesExpectRunSet.ConfigServiceExpectRun && !manager.RobotBasicModulesSet.ConfigService.IsAlive() {
				go manager.RobotBasicModulesSet.ConfigService.Start()
			}
			if manager.RobotBasicModulesExpectRunSet.LocalDatabaseServiceExpectRun && !manager.RobotBasicModulesSet.LocalDatabaseService.IsAlive() {
				go manager.RobotBasicModulesSet.LocalDatabaseService.Start()
			}
			if manager.RobotBasicModulesExpectRunSet.RemoteDatabaseServiceExpectRun && !manager.RobotBasicModulesSet.RemoteDatabaseService.IsAlive() {
				go manager.RobotBasicModulesSet.RemoteDatabaseService.Start()
			}
			if manager.RobotBasicModulesExpectRunSet.SerialServiceExpectRun && !manager.RobotBasicModulesSet.SerialService.IsAlive() {
				go manager.RobotBasicModulesSet.SerialService.Start()
			}
			if manager.RobotBasicModulesExpectRunSet.WebAppExpectRun && !manager.RobotBasicModulesSet.WebApp.IsAlive() {
				go manager.RobotBasicModulesSet.WebApp.Start()
			}
			/*
				if manager.RobotBasicModulesExpectRunSet.VisualAppExpectRun && !manager.RobotBasicModulesSet.VisualApp.IsAlive() {
					go manager.RobotBasicModulesSet.VisualApp.Start()
				}
			*/
		}
	}
}

// InitBasicModule 子节点基础模块初始化器 负责初始化子节点基础模块
// 传入：底层模块名
// 传出：无
func (manager *BasicRobotModuleManager) InitBasicModule(moduleName string) {
	switch moduleName {
	case _const12.ConfigService:
		manager.cConfig()
		manager.iConfig()
	case _const12.HttpApp:
		manager.iHttpApp()
		manager.cHttpApp()
	case _const12.LocalDatabaseService:
		manager.cLocalDatabase()
		manager.iLocalDatabase()
	case _const12.SerialService:
		manager.cSerial()
		manager.iConfig()
	case _const12.RemoteDatabaseService:
		manager.cRemoteDatabase()
		manager.iRemoteDatabase()
	case _const12.WebApp:
		manager.cWebApp()
		manager.iWebApp()
	}
}

// InitAllBasicModels 初始化所有基层模块
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) InitAllBasicModels() {
	manager.InitBasicModule(_const12.ConfigService)
	manager.InitBasicModule(_const12.LocalDatabaseService)
	manager.InitBasicModule(_const12.RemoteDatabaseService)
	manager.InitBasicModule(_const12.SerialService)
	manager.InitBasicModule(_const12.HttpApp)
	manager.InitBasicModule(_const12.WebApp)
}