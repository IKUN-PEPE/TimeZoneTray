# ZoneTray 设计文档

> 日期: 2026-05-15
> 状态: 已更新 (新增设置功能)

## 1. 项目概述
ZoneTray 是一个 Windows 托盘小工具，用户可以通过输入美国邮编快速查询该地区的当前本地时间及对应的时区。

## 2. 核心功能
- **托盘常驻**: 程序运行后在系统托盘显示。
- **邮编查询**: 点击菜单弹出原生输入框。
- **时间转换**: 自动通过 API 将邮编转换为经纬度，再获取当地时间。
- **系统通知**: 通过 Windows Toast 通知显示查询结果。
- **设置中心 (新)**: 集成在托盘菜单中的设置项，目前支持“开机自启动”。

## 3. 技术栈
- **语言**: Go (Golang)
- **UI 库**: github.com/getlantern/systray
- **输入框**: github.com/gen2brain/dlgs
- **通知库**: github.com/go-toast/toast
- **系统集成**: golang.org/x/sys/windows/registry (用于注册表操作)
- **API**: 
  - 经纬度: pi.zippopotam.us
  - 时区时间: 	imeapi.io

## 4. 架构设计
采用模块化设计，将业务逻辑、配置管理与 UI 逻辑分离。

### 目录结构
`
TimeZoneTray/
├── main.go          # 托盘管理、UI 交互、通知触发
├── api/             # 核心业务逻辑：API 请求
├── config/          # 新增：配置管理与系统集成 (自启)
│   ├── config.go    # 持久化配置 (JSON)
│   └── autostart.go # Windows 注册表自启实现
├── assets/          # 静态资源
└── go.mod           # 依赖管理
`

## 5. 详细逻辑设计

### 5.1 设置功能 (Autostart)
- **实现方式**: 
  - 注册表路径: HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run
  - 键值名: ZoneTray
  - 键值: 当前可执行文件的绝对路径。
- **持久化**:
  - 文件: %APPDATA%/ZoneTray/config.json (或程序同级目录，视权限而定)
  - 字段: {"autostart": true}

### 5.2 UI 流程更新
1. systray.OnReady: 初始化菜单，包括“设置”子菜单。
2. Load Config: 启动时读取 config.json 并根据状态执行 Check() 或 Uncheck()。
3. Menu Toggle: 点击“开机自启动”菜单项，切换勾选状态，同步修改注册表并保存配置。

## 6. 异常处理
- **权限不足**: 如果操作注册表失败，通过 Toast 提示用户“设置失败，请尝试以管理员身份运行”。

## 7. 编译与交付
- **命令**: go build -ldflags "-H windowsgui" -o ZoneTray.exe
- **目标**: 隐藏终端黑窗口，生成独立执行文件。
