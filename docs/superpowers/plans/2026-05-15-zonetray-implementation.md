# ZoneTray Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (- [ ]) syntax for tracking.

**Goal:** 构建一个 Windows 托盘应用，通过邮编查询当地时间并弹出通知，支持设置中心。

**Architecture:** 采用模块化分层架构。api/ 包负责业务逻辑，config/ 包负责配置持久化与自启管理，main.go 负责 UI 集成。

**Tech Stack:** Go, systray, dlgs, toast, registry.

---

### Task 1: 项目初始化 [DONE]
### Task 2: 实现经纬度获取逻辑 [DONE]
### Task 3: 实现时间获取逻辑 [DONE]
### Task 4: 实现系统托盘基础框架 [DONE]
### Task 5: 集成输入框与通知系统 [DONE]
### Task 6: 最终构建与 GUI 隐藏验证 [DONE]

---

### Task 7: 实现设置中心 (开机自启)

**Files:**
- Create: config/config.go
- Create: config/autostart.go
- Modify: main.go

- [ ] **Step 1: 创建配置管理包**
实现 Config 结构体和 Load/Save 方法，将配置保存为 config.json。
- [ ] **Step 2: 实现 Windows 自启逻辑**
使用 golang.org/x/sys/windows/registry 实现 EnableAutostart 和 DisableAutostart 函数。
- [ ] **Step 3: 在 main.go 中集成设置菜单**
  - 添加“设置”子菜单。
  - 添加“开机自启动”勾选项。
  - 启动时根据配置自动勾选，并确保注册表同步。
- [ ] **Step 4: 测试自启切换**
运行程序，切换勾选，验证注册表键值是否正确增删。
- [ ] **Step 5: 最终构建验证**
重新生成 ZoneTray.exe 并进行全流程测试。
