# 图标更新实施计划

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 将 `548c13b1-e6cf-4ddf-a452-6ecf2052556c.jpg` 设置为 ZoneTray 的 EXE 图标和系统托盘图标。

**Architecture:** 
1. 编写 Go 脚本将 JPG 转换为包含多尺寸（16, 32, 48, 256）的 ICO 文件。
2. 使用 `rsrc` 工具生成 Windows 资源文件 `icon.syso`。
3. 重新编译生成带图标的 `ZoneTray.exe`。

**Tech Stack:** Go, github.com/nfnt/resize, github.com/sergeymakinen/go-ico, github.com/akavel/rsrc

---

### Task 1: 准备工作

**Files:**
- Modify: `go.mod` (通过 go get 更新)

- [x] **Step 1: 安装依赖库**

运行: `go get github.com/nfnt/resize github.com/sergeymakinen/go-ico`

- [x] **Step 2: 确认依赖安装成功**

检查 `go.mod` 中是否包含上述包。

### Task 2: 编写转换脚本

**Files:**
- Create: `scripts/convert_icon.go`

- [x] **Step 1: 创建 scripts 目录**

运行: `mkdir scripts`

- [x] **Step 2: 编写转换代码并保存**

文件内容:
```go
package main

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
	"github.com/sergeymakinen/go-ico"
)

func main() {
	file, err := os.Open("548c13b1-e6cf-4ddf-a452-6ecf2052556c.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	sizes := []uint{16, 32, 48, 256}
	var images []image.Image
	for _, size := range sizes {
		images = append(images, resize.Resize(size, size, img, resize.Lanczos3))
	}

	outFile, err := os.Create("assets/icon.ico")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = ico.EncodeAll(outFile, images)
	if err != nil {
		panic(err)
	}
}
```

### Task 3: 执行转换并更新资源

**Files:**
- Modify: `assets/icon.ico`
- Modify: `icon.syso`

- [x] **Step 1: 运行转换脚本**

运行: `go run scripts/convert_icon.go`
预期: `assets/icon.ico` 被更新。

- [x] **Step 2: 安装 rsrc 工具**

运行: `go install github.com/akavel/rsrc@latest`

- [x] **Step 3: 生成 icon.syso**

运行: `rsrc -ico assets/icon.ico -o icon.syso`
预期: `icon.syso` 被更新。

### Task 4: 重新编译项目

**Files:**
- Modify: `ZoneTray.exe`

- [x] **Step 1: 执行编译命令**

运行: `go build -ldflags "-H windowsgui" -o ZoneTray.exe`
预期: 生成新的 `ZoneTray.exe`。

### Task 5: 验证

- [x] **Step 1: 确认编译成功**

检查 `ZoneTray.exe` 的修改时间。
确认托盘图标显示正确。
