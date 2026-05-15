package main

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gen2brain/dlgs"
	"github.com/getlantern/systray"
	"github.com/go-toast/toast"
	"github.com/user/zonetray/api"
	"github.com/user/zonetray/config"
)

//go:embed assets/icon.ico
var iconData []byte

func main() {
	setupLogging()
	systray.Run(onReady, onExit)
}

func setupLogging() {
	execPath, err := os.Executable()
	if err != nil {
		return
	}
	logPath := filepath.Join(filepath.Dir(execPath), "zonetray.log")

	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("无法打开日志文件: %v\n", err)
		return
	}

	// 同时输出到文件和终端
	multi := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multi)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("--- ZoneTray 启动 ---")
}

func onReady() {
	// 加载配置
	conf := config.LoadConfig()
	log.Printf("配置已加载: %+v", conf)

	// 设置图标
	systray.SetIcon(getIcon())

	systray.SetTitle("ZoneTray")
	systray.SetTooltip("ZoneTray - 时区查询")

	// 菜单项
	mQuery := systray.AddMenuItem("查询邮编时间", "触发邮编时间查询")
	
	mSettings := systray.AddMenuItem("设置", "配置选项")
	mAutostart := mSettings.AddSubMenuItemCheckbox("开机自启动", "设置程序随系统启动", conf.Autostart)
	
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "退出应用程序")

	// 退出逻辑
	go func() {
		<-mQuit.ClickedCh
		log.Println("用户点击退出")
		systray.Quit()
	}()

	// 设置监听逻辑
	go func() {
		for {
			<-mAutostart.ClickedCh
			if mAutostart.Checked() {
				mAutostart.Uncheck()
				conf.Autostart = false
			} else {
				mAutostart.Check()
				conf.Autostart = true
			}
			
			log.Printf("切换自启动: %v", conf.Autostart)
			err := config.SetAutostart(conf.Autostart)
			if err != nil {
				log.Printf("注册表操作失败: %v", err)
			}
			
			_ = conf.SaveConfig()
		}
	}()

	// 查询逻辑
	go func() {
		for {
			<-mQuery.ClickedCh
			// 每次查询启动独立协程，并锁定系统线程
			go func() {
				// 锁定 OS 线程是处理 Windows GUI 调用的最佳实践
				runtime.LockOSThread()
				defer runtime.UnlockOSThread()

				log.Println("[UI] 准备弹出输入对话框...")
				
				// 尝试调用对话框
				zip, ok, err := dlgs.Entry("ZoneTray", "请输入美国邮编：", "")
				
				log.Printf("[UI] 对话框返回: ok=%v, err=%v, zip=%s", ok, err, zip)
				
				if err != nil {
					log.Printf("对话框错误: %v", err)
					return
				}
				if !ok || zip == "" {
					return
				}

				// API 请求部分
				log.Printf("[API] 开始请求邮编: %s", zip)
				ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
				defer cancel()

				result, err := api.GetTimeByZip(ctx, zip)
				
				var title, message string
				if err != nil {
					log.Printf("[API] 请求失败: %v", err)
					title = "查询失败"
					message = fmt.Sprintf("无法获取邮编 %s 的时间: %v", zip, err)
				} else {
					log.Printf("[API] 请求成功: %s", result)
					title = "邮编时间查询结果"
					message = fmt.Sprintf("邮编 %s 的当前时间为:\n%s", zip, result)
				}

				notification := toast.Notification{
					AppID:   "ZoneTray",
					Title:   title,
					Message: message,
				}
				_ = notification.Push()
			}()
		}
	}()
}

func onExit() {
	// 在此处执行清理操作
	fmt.Println("应用程序已退出")
}

// getIcon 返回托盘图标的字节数组
func getIcon() []byte {
	return iconData
}
