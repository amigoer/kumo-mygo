# kumo-mygo

把 Cloudflare 的 [Kumo](https://github.com/cloudflare/kumo) 设计系统做成 **纯原生** 的 [MyGo](https://github.com/egoist/mygo) 组件库：同样的语义 token、变体和尺寸，由 MyGo 直接在 GPU 上绘制——没有 WebView、没有 HTML、没有 JavaScript，整个应用是一个几 MB 的 Go 二进制。

思路参考 [Onevra](https://github.com/oneo-me/onevra)：**行为交给 MyGo，外观交给我们**。每个组件都建立在 MyGo 的无样式基座（`ButtonBase`、`SelectBase`、`DialogBase`、`TabsBase`、`PopoverBase`、`ToastViewportBase`……）之上，键盘、焦点、无障碍和输入法由 MyGo 负责；本库只按 Kumo 的设计规范给它们上色、定尺寸。

```go
package main

import (
	"log"

	kumo "github.com/amigoer/kumo-mygo"
	"github.com/amigoer/kumo-mygo/icons/phosphor"
	"github.com/amigoer/kumo-mygo/theme"
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

type app struct{ name string }

func (a *app) view(c *ui.Context) {
	theme.Apply(c) // 根视图第一行：安装 Kumo 主题，跟随系统明暗
	kumo.Toaster(c)
	ui.Column(c).Fill().Padding(32).Gap(16).Background(theme.Of(c).Canvas).Children(func() {
		kumo.Input(c, &a.name, kumo.InputOptions{
			FieldOptions: kumo.FieldOptions{Label: "Worker 名称"},
			Placeholder:  "my-worker",
		})
		if kumo.Button(c, "部署", kumo.ButtonOptions{Variant: kumo.ButtonPrimary, Icon: phosphor.Rocket}).Clicked() {
			kumo.Notify(c, "部署成功", a.name, kumo.ToastSuccess)
		}
	})
}

func main() {
	a := &app{}
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{Title: "Hello Kumo", Width: 480, Height: 320, Content: ui.View(a.view)})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
```

## 安装

需要 Go 1.27.3+（MyGo 的要求）。

```sh
go get github.com/amigoer/kumo-mygo@latest
```

导入路径是 `github.com/amigoer/kumo-mygo`，代码里的包名是 `kumo`（Go 包名不能带连字符），所以照常写 `kumo.Button(...)`。

运行组件画廊：

```sh
go run ./examples/gallery                # 跟随系统
go run ./examples/gallery -theme dark    # 强制深色
go tool mygo dev                         # 热重载开发
go tool mygo build                       # 打包 .app / 安装包（读取 mygo.json）
```

## 包结构

| 包 | 内容 |
| --- | --- |
| `kumo` | 全部组件 |
| `kumo/theme` | Kumo 语义 token（`Tokens`）、明暗两套取值、`Apply` / `Of` / `Use` |
| `kumo/icons/phosphor` | Phosphor 图标 regular 字重，1512 个，按需解析（Kumo 用的就是 Phosphor） |
| `kumo/icons/phosphor/fill` | Phosphor fill 字重，同名 |
| `kumo/icons/phosphor/catalog` | 名称 → 访问器的全量表，仅供图标浏览器使用 |
| `tools/phosphorgen` | 从固定提交的 phosphor-icons/core 重新生成图标包 |

## 组件

| Kumo | 本库 | 说明 |
| --- | --- | --- |
| Button | `Button`、`StyleButton` | 6 种变体、xs/sm/base/lg、square/circle、loading；primary/destructive 的渐变面按 oklch color-mix 精确还原 |
| ButtonGroup | `ButtonGroup` | 相邻按钮直角拼接 |
| Badge | `Badge` | 17 种颜色 + `Dot` 状态点外观 |
| Banner | `Banner` | info/alert/error/secondary，base 与紧凑 sm，可关闭、可带操作 |
| Input / InputArea / InputGroup | `Input`、`InputArea` | 前缀/后缀/图标、错误态、只读 |
| SensitiveInput | `SensitiveInput` | 眼睛按钮切换显示 |
| Field / Label | `Field`、`Label` | 标签、描述、错误、(optional)、信息提示 |
| Checkbox | `Checkbox`、`CheckboxGroup` | 半选、错误、卡片外观 |
| Radio | `Radio`、`RadioGroup`、`RadioSegmented` | 卡片与分段外观 |
| Switch | `Switch` | sm/base/lg，蓝色与 neutral 两种 |
| Select | `Select` | 分组、图标、描述、占位符 |
| Slider | `Slider` | Kumo 的凹槽轨道 + 浮起填充 + 数值气泡 |
| Tabs | `Tabs` | segmented（默认）与 underline，指示器滑动动画 |
| Dialog | `Dialog` | sm/base/lg/自定义宽度，Alert 模式 |
| Popover | `Popover` | 标题、描述，四个方向 |
| DropdownMenu | `DropdownMenu` | 图标、快捷键、勾选项、危险项、分组标题、分隔线，方向键导航 |
| Tooltip | `Tooltip` | 四个方向 |
| Toast | `Toaster`、`Notify` | 默认/成功/信息/警告/错误，可带操作 |
| Loader | `Loader` | 与 Kumo SVG 动画一致的弧线 |
| Meter | `Meter` | |
| Empty | `Empty` | 图标、标题、可复制命令行、操作区 |
| Breadcrumbs | `Breadcrumbs` | |
| Pagination | `Pagination` | full / simple |
| LayerCard / Surface | `LayerCard`、`Surface` | |
| Table | `Table` | 斑马纹、紧凑表头、多选列、行点击 |
| Collapsible | `Collapsible` | |
| Link | `Link` | inline / current / plain、外链图标 |
| Text | `Text` | heading1-3、body、secondary、success、error、mono |
| Code / CodeBlock | `Code`、`CodeBlock` | 复制按钮 |
| ClipboardText | `ClipboardText` | |
| Sidebar | `Sidebar`、`SidebarGroup`、`SidebarItem`、`SidebarTrigger` | 可折叠到图标栏 |

尚未移植：Combobox、TagInput、DatePicker / DateRangePicker、CommandPalette、MenuBar、Toolbar、Grid、Chart、Flow 等。MyGo 本身已有 Combobox、Calendar、Chart 的无样式基座，可以按同样方式补上。

## 设计约定

- **语义 token，不写死颜色。** 组件只读 `theme.Of(c)` 里的 `Base`、`Line`、`Tint`、`Brand`、`TextSubtle`……，与 Kumo 的 `bg-kumo-base`、`ring-kumo-line` 一一对应（去掉 `kumo-` 前缀）。取值是 Kumo CSS 实际解析出来的 oklch 值，通过 `ui.Oklch` 保留广色域。
- **明暗自动切换。** `theme.Of` 按帧的 `Dark` 选 token，并像 MyGo 一样跟随系统的“增强对比度”和文字缩放。
- **Options 结构体 = Kumo 的 props。** 零值就是 Kumo 的默认值（如 `ButtonSecondary`、`SizeBase`、`TabsSegmented`）。
- **状态归应用。** 组件通过指针修改值，用 `Clicked()` / `Changed()` 报告事件，和 MyGo 自带控件一样。
- **组件返回 `ui.Element`**，所以还能继续链式调整：`kumo.Button(c, "保存").FillWidth()`。
- Kumo 的 Tailwind 环（`ring`）用不占位的外扩阴影实现；聚焦环是 2px 品牌色，与 Kumo 的 `focus-visible:ring-2 ring-kumo-brand` 一致。

## 定制品牌

```go
l, d := theme.Light(), theme.Dark()
l.Brand, d.Brand = ui.Hex("#7c3aed"), ui.Hex("#8b5cf6")
l.BrandHover, d.BrandHover = ui.Hex("#6d28d9"), ui.Hex("#7c3aed")
theme.Use(l, d) // 在创建窗口前调用
```

主按钮的渐变、描边、高光都由 `Brand` 在 oklch 空间推导，改一个颜色即可。

## 本地化

组件自带的少量文字（“关闭”“复制”“(optional)”“上一页”……）集中在 `kumo.Strings`：

```go
kumo.Strings = kumo.StringsZhHans
```

## 开发

```sh
gofmt -l .
go vet ./...
go test ./...          # 用 ui.NewTester 无窗口运行视图，明暗两种外观
```

重新生成图标：

```sh
git clone https://github.com/phosphor-icons/core ../phosphor-core
git -C ../phosphor-core checkout 2b75f3ad12b420c9504ef05df8d2564a28f8500e
go run ./tools/phosphorgen -upstream ../phosphor-core -commit 2b75f3ad12b420c9504ef05df8d2564a28f8500e
```

## 许可

MIT。设计 token 与规格移植自 Cloudflare Kumo（MIT）；Phosphor Icons 保留其 MIT 许可（`icons/phosphor/LICENSE`）。
