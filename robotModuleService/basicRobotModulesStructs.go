package robotModuleService

import (
	"github.com/UniversalRobotDriveTeam/child-nodes-basic/robotBasicAPI/httpNetService"
	"github.com/UniversalRobotDriveTeam/child-nodes-basic/robotBasicAPI/websocketService"
	"github.com/UniversalRobotDriveTeam/child-nodes-config-service/configService"
	"github.com/UniversalRobotDriveTeam/child-nodes-database-service/databaseService"
	"github.com/UniversalRobotDriveTeam/child-nodes-device-service/deviceService"
)

// RobotBasicModules 构成子节点最基本的几个模块
type RobotBasicModules struct {
	// 配置服务
	ConfigService *configService.ConfigManager
	// 本地数据库服务
	LocalDatabaseService *databaseService.DatabaseAPP
	// 远程数据库服务
	RemoteDatabaseService *databaseService.DatabaseAPP
	// 串口服务
	SerialService *deviceService.SerialApp
	// Http传递服务
	HttpApp *httpNetService.HttpAPP
	// webSocket传递服务
	WebApp *websocketService.WebApp
	// 视觉服务 todo
	//VisualApp *visualService
}

// RobotBasicModulesExpectRun 子节点的服务是否需要运行
type RobotBasicModulesExpectRun struct {
	// 配置服务是否运行
	ConfigServiceExpectRun bool
	// 本地数据库服务是否运行
	LocalDatabaseServiceExpectRun bool
	// 远程数据库服务是否运行
	RemoteDatabaseServiceExpectRun bool
	// 串口服务是否运行
	SerialServiceExpectRun bool
	// 错误管理器是否运行
	ErrorManagerExpectRun bool
	// 本地方法调用HTTP是否运行
	LocalCallHttpExpectRun bool
	// 本地方法调用UNIX SOCKET是否运行
	LocalCallUnixSocketExpectRun bool
	// Http传递服务是否运行
	HttpAppExpectRun bool
	// webSocket传递服务是否运行
	WebAppExpectRun bool
	// 视觉服务是否运行
	VisualAppExpectRun bool
}

// BasicRobotModulesInitArgs 各个子节点的初始化参量
type BasicRobotModulesInitArgs struct {
	// 所有的配置都是指针
	// 配置服务初始化参数
	ConfigServiceInit interface{}
	// 本地数据库服务初始化参数
	LocalDatabaseServiceInit interface{}
	// 远程数据库服务初始化参数
	RemoteDatabaseServiceInit interface{}
	// 串口服务初始化参数
	SerialServiceInit interface{}
	// 错误管理器初始化参数
	ErrorManagerInit interface{}
	// 本地方法调用HTTP初始化参数
	LocalCallHttpInit interface{}
	// 本地方法调用UNIX SOCKET初始化参数
	LocalCallUnixSocketInit interface{}
	// Http传递服务初始化参数
	HttpAppInit interface{}
	// webSocket传递服务初始化参数
	WebAppInit interface{}
	// 视觉服务初始化参数
	VisualAppInit interface{}
}

// BasicRobotModuleManager 模块管理器 它负责在模块掉线以后根据要求启动模块 并监控模块状态
type BasicRobotModuleManager struct {
	// 子节点所有的服务模块
	RobotBasicModulesSet *RobotBasicModules
	// 子节点服务模块是否运行
	RobotBasicModulesExpectRunSet *RobotBasicModulesExpectRun
	// 各个子节点的初始化参量
	BasicRobotModulesInitArgsSet *BasicRobotModulesInitArgs
	// 子节点服务模块监控线程是否在运行
	robotMonitorIsRun bool
	// 终止子节点监控器运行的管道
	stopRobotMonitorChannel *chan int
}