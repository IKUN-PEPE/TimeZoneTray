# ZoneTray - 美国邮编时区查询工具

ZoneTray 是一个轻量级的 Windows 托盘工具，旨在帮助用户快速查询美国任意邮编对应的当前本地时间和时区。

## ✨ 核心功能

- **托盘驻留**：运行后常驻系统托盘，不占用任务栏空间。
- **快速查询**：点击托盘菜单弹出原生输入框，输入美国邮编即可一键查询。
- **系统通知**：通过 Windows Toast 通知即时显示查询结果。
- **开机自启**：支持在设置菜单中一键开启或关闭开机自动启动（通过注册表实现）。
- **优化底层**：
  - 采用 **Goroutine 并发解耦**，确保查询时 UI 界面永不卡死。
  - 强制 **OS 线程锁定**，解决 Windows GUI API 兼容性问题。
  - 引入 **Windows Manifest**，确保使用现代系统控件样式。
- **本地日志**：自动生成 `zonetray.log`，记录运行状态与异常，方便排错。

## 🛠️ 技术栈

- **语言**: [Go (Golang)](https://golang.org/)
- **UI 库**: `github.com/getlantern/systray` (托盘管理)
- **对话框**: `github.com/gen2brain/dlgs` (原生输入框)
- **通知**: `github.com/go-toast/toast` (Windows 通知)
- **配置**: `golang.org/x/sys/windows/registry` (注册表操作)
- **API**:
  - 位置获取: [Zippopotam.us](http://api.zippopotam.us)
  - 时间获取: [TimeAPI.io](https://timeapi.io)

## 🚀 快速开始

### 编译

如果你已安装 Go 环境，可以使用以下命令编译出无黑窗口、带图标的 `.exe` 文件：

```powershell
# 1. 整理依赖
go mod tidy

# 2. 编译最终成品
go build -ldflags "-H windowsgui" -o ZoneTray.exe
```

### 运行

双击 `ZoneTray.exe` 即可启动。在系统右下角托盘找到图标，右键点击即可看到功能菜单。

## 📂 项目结构

- `main.go`: 程序入口，负责 UI 事件循环与逻辑集成。
- `api/`: 核心业务逻辑，包含 API 请求处理及单元测试。
- `config/`: 配置管理模块，负责 `config.json` 持久化与注册表自启设置。
- `assets/`: 静态资源，包含程序图标。
- `zonetray.log`: (运行后生成) 详细的运行日志。

## 📝 许可证

MIT License
