package main

import (
	"fmt"

	kumo "github.com/amigoer/kumo-mygo"
	"github.com/amigoer/kumo-mygo/icons/phosphor"
	"github.com/amigoer/kumo-mygo/theme"
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

// page is one entry of the sidebar.
type page struct {
	name  string
	icon  kumo.IconFunc
	group string
	build func(g *gallery, c *ui.Context)
}

// gallery holds the state of every demonstration, which the components
// edit through pointers.
type gallery struct {
	pages     []page
	current   int
	collapsed bool
	appear    int // 0 system, 1 light, 2 dark

	name, token, search, note string
	region                    string
	plan                      string
	agree, analytics, mixed   bool
	notifications, beta       bool
	volume                    float64
	tab, tab2                 int
	crumb                     int
	pageNo                    int
	open, alert, pop, menu    bool
	starred                   bool
	more                      bool
	banner                    bool
	rows                      []worker
}

type worker struct {
	name, route, status string
	requests            int
	selected            bool
}

func newGallery() *gallery {
	g := &gallery{volume: 40, region: "weur", plan: "pro", pageNo: 1, banner: true, notifications: true}
	g.rows = []worker{
		{name: "api-gateway", route: "api.example.com/*", status: "已部署", requests: 182340},
		{name: "image-resizer", route: "img.example.com/*", status: "已部署", requests: 92810},
		{name: "auth-edge", route: "example.com/login", status: "构建中", requests: 1204},
		{name: "cron-cleanup", route: "—", status: "失败", requests: 0},
	}
	g.pages = []page{
		{"按钮 Button", phosphor.CursorClick, "基础", (*gallery).buttons},
		{"徽标 Badge", phosphor.Tag, "基础", (*gallery).badges},
		{"文本 Text", phosphor.TextAa, "基础", (*gallery).texts},
		{"加载 Loader", phosphor.CircleNotch, "基础", (*gallery).loaders},
		{"输入框 Input", phosphor.Textbox, "表单", (*gallery).inputs},
		{"选择器 Select", phosphor.CaretUpDown, "表单", (*gallery).selects},
		{"复选与单选", phosphor.CheckSquare, "表单", (*gallery).choices},
		{"开关 Switch", phosphor.ToggleRight, "表单", (*gallery).switches},
		{"滑块 Slider", phosphor.SlidersHorizontal, "表单", (*gallery).sliders},
		{"横幅 Banner", phosphor.Info, "反馈", (*gallery).banners},
		{"通知 Toast", phosphor.Bell, "反馈", (*gallery).toasts},
		{"进度 Meter", phosphor.Gauge, "反馈", (*gallery).meters},
		{"空状态 Empty", phosphor.Cube, "反馈", (*gallery).empties},
		{"对话框 Dialog", phosphor.AppWindow, "叠层", (*gallery).dialogs},
		{"弹出层 Popover", phosphor.ChatCenteredText, "叠层", (*gallery).popovers},
		{"标签页 Tabs", phosphor.Tabs, "导航", (*gallery).tabsPage},
		{"面包屑与分页", phosphor.Path, "导航", (*gallery).navigation},
		{"卡片 LayerCard", phosphor.Cards, "布局", (*gallery).cards},
		{"表格 Table", phosphor.Table, "布局", (*gallery).tables},
		{"代码 Code", phosphor.Code, "布局", (*gallery).code},
	}
	return g
}

func (g *gallery) view(c *ui.Context) {
	theme.Apply(c)
	k := theme.Of(c)
	kumo.Toaster(c)
	ui.Row(c).Fill().Background(k.Canvas).Children(func() {
		g.sidebar(c)
		ui.Column(c).Grow(1).MinWidth(0).FillHeight().Children(func() {
			g.topBar(c)
			ui.Scroll(c.Key(g.current)).Grow(1).MinHeight(0).FillWidth().Children(func() {
				ui.Column(c).Padding(32, 40).Gap(24).MaxWidth(920).FillWidth().Children(func() {
					p := g.pages[g.current]
					kumo.Text(c, p.name, kumo.TextOptions{Variant: kumo.TextHeading2})
					p.build(g, c)
				})
			})
		})
	})
}

func (g *gallery) sidebar(c *ui.Context) {
	kumo.Sidebar(c, kumo.SidebarOptions{
		Collapsed: &g.collapsed,
		Header: func() {
			if !g.collapsed {
				ui.Row(c).Gap(8).AlignItems(ui.Center).Grow(1).Children(func() {
					kumo.Icon(c, phosphor.Cloud).Size(22, 22).TextColor(theme.Of(c).TextBrand)
					kumo.Text(c, "Kumo", kumo.TextOptions{Bold: true, Size: kumo.SizeLG})
					kumo.Badge(c, "MyGo", kumo.BadgeOptions{Variant: kumo.BadgeBeta})
				})
			}
			kumo.SidebarTrigger(c, &g.collapsed, phosphor.Sidebar)
		},
		Footer: func() {
			if !g.collapsed {
				kumo.Text(c, "原生 · 无 WebView", kumo.TextOptions{Variant: kumo.TextSecondary, Size: kumo.SizeSM})
			}
		},
	}, func() {
		group := ""
		var items []int
		flush := func() {
			if len(items) == 0 {
				return
			}
			label := group
			if g.collapsed {
				label = ""
			}
			list := items
			kumo.SidebarGroup(c.Key(group), label, func() {
				for _, i := range list {
					p := g.pages[i]
					if kumo.SidebarItem(c.Key(i), p.name, p.icon, i == g.current, g.collapsed).Clicked() {
						g.current = i
					}
				}
			})
			items = nil
		}
		for i, p := range g.pages {
			if p.group != group {
				flush()
				group = p.group
			}
			items = append(items, i)
		}
		flush()
	})
}

func (g *gallery) topBar(c *ui.Context) {
	k := theme.Of(c)
	ui.Row(c).Height(58).Shrink(0).PaddingX(24).Gap(12).AlignItems(ui.Center).
		Background(k.Base).BorderWidth(0, 0, 1, 0).BorderColor(k.Line).Children(func() {
		kumo.Breadcrumbs(c, []string{"Kumo", g.pages[g.current].group, g.pages[g.current].name}, &g.crumb,
			kumo.BreadcrumbsOptions{Size: kumo.SizeSM}).Grow(1)
		modes := []int{0, 1, 2}
		before := g.appear
		kumo.RadioSegmented(c, &g.appear, modes, []string{"跟随系统", "浅色", "深色"})
		if g.appear != before {
			mygo.Theme.SetSource([]mygo.ThemeSource{mygo.ThemeSystem, mygo.ThemeLight, mygo.ThemeDark}[g.appear])
		}
	})
}

// section titles a group of demonstrations.
func section(c *ui.Context, title, desc string, fn func()) {
	ui.Column(c).Gap(12).FillWidth().Children(func() {
		ui.Column(c).Gap(4).Children(func() {
			kumo.Text(c, title, kumo.TextOptions{Variant: kumo.TextHeading3})
			if desc != "" {
				kumo.Text(c, desc, kumo.TextOptions{Variant: kumo.TextSecondary})
			}
		})
		kumo.Surface(c).Padding(20).Gap(16).FillWidth().Children(fn)
	})
}

func row(c *ui.Context, fn func()) {
	ui.Row(c).Gap(12).AlignItems(ui.Center).Wrap().Children(fn)
}

func (g *gallery) buttons(c *ui.Context) {
	section(c, "变体", "与 Kumo 相同的六种变体，默认为 secondary。", func() {
		row(c, func() {
			kumo.Button(c, "Primary", kumo.ButtonOptions{Variant: kumo.ButtonPrimary})
			kumo.Button(c, "Secondary")
			kumo.Button(c, "Ghost", kumo.ButtonOptions{Variant: kumo.ButtonGhost})
			kumo.Button(c, "Outline", kumo.ButtonOptions{Variant: kumo.ButtonOutline})
			kumo.Button(c, "Destructive", kumo.ButtonOptions{Variant: kumo.ButtonDestructive})
			kumo.Button(c, "Secondary destructive", kumo.ButtonOptions{Variant: kumo.ButtonSecondaryDestructive})
		})
	})
	section(c, "尺寸", "xs 20 · sm 26 · base 36 · lg 40 DIP。", func() {
		row(c, func() {
			for _, s := range []kumo.Size{kumo.SizeXS, kumo.SizeSM, kumo.SizeBase, kumo.SizeLG} {
				kumo.Button(c, "部署", kumo.ButtonOptions{Variant: kumo.ButtonPrimary, Size: s, Icon: phosphor.Rocket})
			}
		})
	})
	section(c, "图标、形状与状态", "", func() {
		row(c, func() {
			kumo.Button(c, "新建", kumo.ButtonOptions{Icon: phosphor.Plus})
			kumo.Button(c, "设置", kumo.ButtonOptions{Shape: kumo.ShapeSquare, Icon: phosphor.Gear})
			kumo.Button(c, "收藏", kumo.ButtonOptions{Shape: kumo.ShapeCircle, Icon: phosphor.Star, Variant: kumo.ButtonOutline})
			kumo.Button(c, "保存中", kumo.ButtonOptions{Variant: kumo.ButtonPrimary, Loading: true})
			kumo.Button(c, "不可用", kumo.ButtonOptions{Disabled: true})
		})
		row(c, func() {
			switch kumo.ButtonGroup(c, []kumo.ButtonGroupItem{
				{Label: "左对齐", Icon: phosphor.TextAlignLeft, IconOnly: true},
				{Label: "居中", Icon: phosphor.TextAlignCenter, IconOnly: true},
				{Label: "右对齐", Icon: phosphor.TextAlignRight, IconOnly: true},
			}, kumo.ButtonOptions{}) {
			case 0, 1, 2:
				kumo.Notify(c, "对齐方式已更改", "", kumo.ToastDefault)
			}
		})
	})
}

func (g *gallery) badges(c *ui.Context) {
	section(c, "语义变体", "", func() {
		row(c, func() {
			for _, b := range []struct {
				name string
				v    kumo.BadgeVariant
			}{{"Primary", kumo.BadgePrimary}, {"Secondary", kumo.BadgeSecondary}, {"Info", kumo.BadgeInfo},
				{"Success", kumo.BadgeSuccess}, {"Warning", kumo.BadgeWarning}, {"Error", kumo.BadgeError},
				{"Destructive", kumo.BadgeDestructive}, {"Outline", kumo.BadgeOutline}, {"Beta", kumo.BadgeBeta}} {
				kumo.Badge(c, b.name, kumo.BadgeOptions{Variant: b.v})
			}
		})
	})
	section(c, "颜色与状态点", "", func() {
		row(c, func() {
			for _, b := range []struct {
				name string
				v    kumo.BadgeVariant
			}{{"Red", kumo.BadgeRed}, {"Green", kumo.BadgeGreen}, {"Orange", kumo.BadgeOrange}, {"Purple", kumo.BadgePurple},
				{"Teal", kumo.BadgeTeal}, {"Blue", kumo.BadgeBlue}, {"Neutral", kumo.BadgeNeutral}} {
				kumo.Badge(c, b.name, kumo.BadgeOptions{Variant: b.v})
			}
		})
		row(c, func() {
			kumo.Badge(c, "运行中", kumo.BadgeOptions{Variant: kumo.BadgeSuccess, Dot: true})
			kumo.Badge(c, "降级", kumo.BadgeOptions{Variant: kumo.BadgeWarning, Dot: true})
			kumo.Badge(c, "离线", kumo.BadgeOptions{Variant: kumo.BadgeError, Dot: true})
			kumo.Badge(c, "已暂停", kumo.BadgeOptions{Variant: kumo.BadgeNeutral, Dot: true})
			kumo.Badge(c, "Workers", kumo.BadgeOptions{Variant: kumo.BadgeSecondary, Icon: phosphor.Lightning})
		})
	})
}

func (g *gallery) texts(c *ui.Context) {
	section(c, "排版", "", func() {
		kumo.Text(c, "Heading 1", kumo.TextOptions{Variant: kumo.TextHeading1})
		kumo.Text(c, "Heading 2", kumo.TextOptions{Variant: kumo.TextHeading2})
		kumo.Text(c, "Heading 3", kumo.TextOptions{Variant: kumo.TextHeading3})
		kumo.Text(c, "正文：Kumo 的组件直接由 MyGo 在 GPU 上绘制，没有 HTML、没有 JavaScript。")
		kumo.Text(c, "次要文本，用于说明。", kumo.TextOptions{Variant: kumo.TextSecondary})
		kumo.Text(c, "部署成功", kumo.TextOptions{Variant: kumo.TextSuccess})
		kumo.Text(c, "部署失败：超出配额", kumo.TextOptions{Variant: kumo.TextError})
		kumo.Text(c, "wrangler deploy --env production", kumo.TextOptions{Variant: kumo.TextMono})
	})
	section(c, "链接", "", func() {
		row(c, func() {
			kumo.Link(c, "Kumo 文档", "https://kumo-ui.com")
			kumo.Link(c, "随文字颜色", "https://kumo-ui.com", kumo.LinkOptions{Variant: kumo.LinkCurrent})
			kumo.Link(c, "无下划线", "https://kumo-ui.com", kumo.LinkOptions{Variant: kumo.LinkPlain})
			kumo.Link(c, "MyGo", "https://github.com/egoist/mygo", kumo.LinkOptions{External: true})
		})
	})
}

func (g *gallery) loaders(c *ui.Context) {
	section(c, "Loader", "sm 16 · base 24 · lg 32", func() {
		row(c, func() {
			kumo.Loader(c, kumo.LoaderSM)
			kumo.Loader(c, kumo.LoaderBase)
			kumo.Loader(c, kumo.LoaderLG, theme.Of(c).Brand)
		})
	})
}

func (g *gallery) inputs(c *ui.Context) {
	section(c, "Input", "", func() {
		kumo.Input(c, &g.name, kumo.InputOptions{
			FieldOptions: kumo.FieldOptions{Label: "Worker 名称", Description: "只能包含小写字母、数字和连字符。"},
			Placeholder:  "my-worker",
		})
		errMsg := ""
		if g.name != "" && len(g.name) < 3 {
			errMsg = "名称至少需要 3 个字符"
		}
		kumo.Input(c, &g.name, kumo.InputOptions{
			FieldOptions: kumo.FieldOptions{Label: "带校验", Error: errMsg, Optional: true, LabelTooltip: "与上方输入框共享同一个值"},
			Prefix:       "https://", Suffix: ".workers.dev",
		})
		kumo.Input(c, &g.search, kumo.InputOptions{Icon: phosphor.MagnifyingGlass, Placeholder: "搜索…", Size: kumo.SizeSM})
		kumo.SensitiveInput(c, &g.token, kumo.InputOptions{FieldOptions: kumo.FieldOptions{Label: "API Token"}, Placeholder: "粘贴令牌"})
		kumo.InputArea(c, &g.note, kumo.InputOptions{FieldOptions: kumo.FieldOptions{Label: "备注"}, Placeholder: "写点什么…"})
		kumo.Input(c, &g.region, kumo.InputOptions{FieldOptions: kumo.FieldOptions{Label: "只读"}, ReadOnly: true})
	})
	section(c, "ClipboardText", "", func() {
		kumo.ClipboardText(c, "0f3d9a2c-77b1-4c4e-9e21-5c8f1d2a9b44", kumo.ClipboardTextOptions{Label: "Account ID"})
	})
}

func (g *gallery) selects(c *ui.Context) {
	section(c, "Select", "", func() {
		kumo.Select(c, &g.region, []kumo.SelectOption[string]{
			{Group: "欧洲", Value: "weur", Label: "西欧", Icon: phosphor.GlobeHemisphereEast},
			{Value: "eeur", Label: "东欧", Icon: phosphor.GlobeHemisphereEast},
			{Group: "北美", Value: "wnam", Label: "北美西部", Icon: phosphor.GlobeHemisphereWest},
			{Value: "enam", Label: "北美东部", Icon: phosphor.GlobeHemisphereWest},
			{Group: "亚太", Value: "apac", Label: "亚太", Description: "包括新加坡、东京与悉尼", Icon: phosphor.GlobeHemisphereEast},
		}, kumo.SelectOptions{FieldOptions: kumo.FieldOptions{Label: "区域", Description: "Durable Object 的位置提示。"}})
		var empty string
		kumo.Select(c, &empty, []kumo.SelectOption[string]{{Value: "a", Label: "选项 A"}, {Value: "b", Label: "选项 B"}},
			kumo.SelectOptions{Placeholder: "未选择", Size: kumo.SizeSM})
	})
}

func (g *gallery) choices(c *ui.Context) {
	section(c, "Checkbox", "", func() {
		kumo.CheckboxGroup(c, kumo.CheckboxGroupOptions{Legend: "偏好设置"}, func() {
			kumo.Checkbox(c, &g.agree, "我同意服务条款")
			kumo.Checkbox(c, &g.analytics, "发送匿名使用数据", kumo.CheckboxOptions{Description: "帮助我们改进产品。"})
			kumo.Checkbox(c, &g.mixed, "部分选中", kumo.CheckboxOptions{Indeterminate: !g.mixed})
			kumo.Checkbox(c, new(bool), "不可用", kumo.CheckboxOptions{Disabled: true})
		})
	})
	section(c, "Radio", "", func() {
		kumo.RadioGroup(c, kumo.CheckboxGroupOptions{Legend: "套餐"}, func() {
			kumo.Radio(c, &g.plan, "free", "Free")
			kumo.Radio(c, &g.plan, "pro", "Pro", kumo.RadioOptions{Description: "$20 / 月"})
			kumo.Radio(c, &g.plan, "biz", "Business", kumo.RadioOptions{Description: "$200 / 月"})
		})
		ui.Row(c).Gap(12).Children(func() {
			kumo.Radio(c, &g.plan, "free", "Free", kumo.RadioOptions{Card: true, Description: "个人项目"}).Grow(1).Basis(0)
			kumo.Radio(c, &g.plan, "pro", "Pro", kumo.RadioOptions{Card: true, Description: "小团队"}).Grow(1).Basis(0)
			kumo.Radio(c, &g.plan, "biz", "Business", kumo.RadioOptions{Card: true, Description: "企业"}).Grow(1).Basis(0)
		})
		kumo.RadioSegmented(c, &g.plan, []string{"free", "pro", "biz"}, []string{"Free", "Pro", "Business"})
	})
}

func (g *gallery) switches(c *ui.Context) {
	section(c, "Switch", "", func() {
		kumo.Switch(c, &g.notifications, "邮件通知")
		kumo.Switch(c, &g.beta, "启用 Beta 功能", kumo.SwitchOptions{Size: kumo.SizeLG})
		kumo.Switch(c, &g.beta, "中性色", kumo.SwitchOptions{Neutral: true, Size: kumo.SizeSM})
		kumo.Switch(c, new(bool), "不可用", kumo.SwitchOptions{Disabled: true})
	})
}

func (g *gallery) sliders(c *ui.Context) {
	section(c, "Slider", "", func() {
		kumo.Slider(c, &g.volume, kumo.SliderOptions{Label: "并发上限", Format: func(v float64) string { return fmt.Sprintf("%.0f", v) }})
		kumo.Slider(c, &g.volume, kumo.SliderOptions{Size: kumo.SizeSM, HideBounds: true})
	})
}

func (g *gallery) banners(c *ui.Context) {
	section(c, "Banner", "", func() {
		if g.banner {
			kumo.Banner(c, kumo.BannerOptions{Icon: phosphor.Info, Title: "新版本可用",
				Description: "Kumo for MyGo 现已支持深色模式与高对比度。",
				Action: func() {
					kumo.Button(c, "了解更多", kumo.ButtonOptions{Size: kumo.SizeSM})
				},
				OnDismiss: func() { g.banner = false }})
		} else if kumo.Button(c, "重新显示横幅").Clicked() {
			g.banner = true
		}
		kumo.Banner(c, kumo.BannerOptions{Variant: kumo.BannerAlert, Icon: phosphor.Warning, Title: "用量接近上限", Description: "本月请求数已使用 92%。"})
		kumo.Banner(c, kumo.BannerOptions{Variant: kumo.BannerError, Icon: phosphor.WarningOctagon, Title: "部署失败", Description: "脚本超过 10 MB 的大小限制。"})
		kumo.Banner(c, kumo.BannerOptions{Variant: kumo.BannerSecondary, Size: kumo.SizeSM, Icon: phosphor.Lightbulb, Title: "提示", Description: "按 ⌘K 打开命令面板。"})
	})
}

func (g *gallery) toasts(c *ui.Context) {
	section(c, "Toast", "通知堆叠在窗口右下角，鼠标停留时暂停计时。", func() {
		row(c, func() {
			if kumo.Button(c, "默认").Clicked() {
				kumo.Notify(c, "已保存", "设置已同步到所有区域。", kumo.ToastDefault)
			}
			if kumo.Button(c, "成功").Clicked() {
				kumo.Notify(c, "部署成功", "api-gateway 已发布到 312 个城市。", kumo.ToastSuccess)
			}
			if kumo.Button(c, "信息").Clicked() {
				kumo.Notify(c, "维护通知", "周六 02:00 UTC 计划维护。", kumo.ToastInfo)
			}
			if kumo.Button(c, "警告").Clicked() {
				kumo.Notify(c, "证书即将过期", "还剩 7 天。", kumo.ToastWarning)
			}
			if kumo.Button(c, "错误").Clicked() {
				kumo.Notify(c, "构建失败", "第 42 行存在语法错误。", kumo.ToastError)
			}
			if kumo.Button(c, "带操作").Clicked() {
				c.AddToast(ui.Toast{Title: "已删除 worker", Description: "cron-cleanup", Action: "撤销",
					OnAction: func() { kumo.Notify(c, "已恢复", "", kumo.ToastSuccess) }})
			}
		})
	})
}

func (g *gallery) meters(c *ui.Context) {
	k := theme.Of(c)
	section(c, "Meter", "", func() {
		kumo.Meter(c, 42, kumo.MeterOptions{Label: "KV 读取"})
		kumo.Meter(c, 7.6, kumo.MeterOptions{Label: "R2 存储", Max: 10, Value: "7.6 GB / 10 GB"})
		kumo.Meter(c, 96, kumo.MeterOptions{Label: "请求配额", Color: k.Danger})
	})
}

func (g *gallery) empties(c *ui.Context) {
	section(c, "Empty", "", func() {
		kumo.Empty(c, kumo.EmptyOptions{
			Icon: phosphor.Cube, Title: "还没有 Worker",
			Description: "用命令行创建第一个 Worker，或从模板开始。",
			CommandLine: "npm create cloudflare@latest",
			Contents: func() {
				kumo.Button(c, "从模板创建", kumo.ButtonOptions{Variant: kumo.ButtonPrimary, Icon: phosphor.Plus})
				kumo.Button(c, "查看文档")
			},
		})
	})
}

func (g *gallery) dialogs(c *ui.Context) {
	section(c, "Dialog", "", func() {
		row(c, func() {
			if kumo.Button(c, "打开对话框").Clicked() {
				g.open = true
			}
			if kumo.Button(c, "删除…", kumo.ButtonOptions{Variant: kumo.ButtonDestructive}).Clicked() {
				g.alert = true
			}
		})
	})
	kumo.Dialog(c, &g.open, kumo.DialogOptions{
		Title: "重命名 Worker", Description: "新名称会立即生效，旧的路由将被保留。",
		Footer: func() {
			if kumo.Button(c, "取消").Clicked() {
				g.open = false
			}
			if kumo.Button(c, "保存", kumo.ButtonOptions{Variant: kumo.ButtonPrimary}).Clicked() {
				g.open = false
				kumo.Notify(c, "已重命名", g.name, kumo.ToastSuccess)
			}
		},
	}, func() {
		kumo.Input(c, &g.name, kumo.InputOptions{FieldOptions: kumo.FieldOptions{Label: "名称"}, AutoFocus: true})
	})
	kumo.Dialog(c, &g.alert, kumo.DialogOptions{
		Alert: true, Size: kumo.SizeSM, HideClose: true,
		Title: "删除 api-gateway？", Description: "此操作无法撤销，所有路由将停止响应。",
		Footer: func() {
			if kumo.Button(c, "取消").Clicked() {
				g.alert = false
			}
			if kumo.Button(c, "删除", kumo.ButtonOptions{Variant: kumo.ButtonDestructive}).Clicked() {
				g.alert = false
				kumo.Notify(c, "已删除", "api-gateway", kumo.ToastError)
			}
		},
	}, nil)
}

func (g *gallery) popovers(c *ui.Context) {
	section(c, "Popover、DropdownMenu 与 Tooltip", "", func() {
		row(c, func() {
			b := kumo.Button(c, "分享", kumo.ButtonOptions{Icon: phosphor.ShareNetwork})
			if b.Clicked() {
				g.pop = !g.pop
			}
			kumo.Popover(c, b, &g.pop, kumo.PopoverOptions{Title: "分享链接", Description: "任何拥有链接的人都能查看。", Width: 320}, func() {
				kumo.ClipboardText(c, "https://dash.example.com/s/8f2k", kumo.ClipboardTextOptions{Size: kumo.SizeSM})
			})
			m := kumo.Button(c, "操作", kumo.ButtonOptions{Icon: phosphor.DotsThree})
			if m.Clicked() {
				g.menu = !g.menu
			}
			kumo.DropdownMenu(c, m, &g.menu, []kumo.MenuItem{
				{Header: true, Label: "Worker"},
				{Label: "重命名", Icon: phosphor.PencilSimple, Shortcut: "⌘R", OnSelect: func() { g.open = true }},
				{Label: "复制", Icon: phosphor.Copy, Shortcut: "⌘D"},
				{Label: "收藏", Checked: &g.starred},
				{Separator: true},
				{Label: "删除", Icon: phosphor.Trash, Danger: true, OnSelect: func() { g.alert = true }},
			})
			t := kumo.Button(c, "悬停查看提示", kumo.ButtonOptions{Variant: kumo.ButtonGhost})
			kumo.Tooltip(c, t, "Tooltip 在指针停留 0.6 秒后出现")
		})
	})
}

func (g *gallery) tabsPage(c *ui.Context) {
	names := []string{"概览", "指标", "日志", "设置"}
	section(c, "Segmented（默认）", "", func() {
		kumo.Tabs(c, &g.tab, names)
		kumo.Text(c, "当前："+names[g.tab], kumo.TextOptions{Variant: kumo.TextSecondary})
		kumo.Tabs(c, &g.tab, names, kumo.TabsOptions{Size: kumo.SizeSM})
	})
	section(c, "Underline", "", func() {
		kumo.Tabs(c, &g.tab2, names, kumo.TabsOptions{Variant: kumo.TabsUnderline})
	})
}

func (g *gallery) navigation(c *ui.Context) {
	section(c, "Breadcrumbs", "", func() {
		kumo.Breadcrumbs(c, []string{"账户", "Workers & Pages", "api-gateway", "设置"}, &g.crumb)
	})
	section(c, "Pagination", "", func() {
		kumo.Pagination(c, &g.pageNo, kumo.PaginationOptions{PerPage: 20, Total: 238})
		kumo.Pagination(c, &g.pageNo, kumo.PaginationOptions{PerPage: 20, Total: 238, Simple: true, HideSummary: true})
	})
	section(c, "Collapsible", "", func() {
		kumo.Collapsible(c, "高级设置", &g.more, func() {
			kumo.Switch(c, &g.beta, "兼容性标志 nodejs_compat")
			kumo.Text(c, "这些设置会影响运行时行为。", kumo.TextOptions{Variant: kumo.TextSecondary})
		})
	})
}

func (g *gallery) cards(c *ui.Context) {
	ui.Row(c).Gap(16).FillWidth().Children(func() {
		ui.Column(c).Grow(1).Basis(0).Children(func() {
			kumo.LayerCard(c, kumo.LayerCardOptions{Title: "快速开始", Icon: phosphor.RocketLaunch, Actions: func() {
				kumo.Link(c, "查看全部", "")
			}}, func() {
				kumo.Text(c, "从模板创建 Worker", kumo.TextOptions{Bold: true})
				kumo.Text(c, "几秒钟内部署到全球网络。", kumo.TextOptions{Variant: kumo.TextSecondary})
			})
		})
		ui.Column(c).Grow(1).Basis(0).Children(func() {
			kumo.Surface(c).Padding(16).Gap(8).Children(func() {
				kumo.Text(c, "Surface", kumo.TextOptions{Bold: true})
				kumo.Text(c, "基础容器：base 底色、line 描边与轻阴影。", kumo.TextOptions{Variant: kumo.TextSecondary})
			})
		})
	})
}

func (g *gallery) tables(c *ui.Context) {
	cols := []kumo.TableColumn{{Header: "名称"}, {Header: "路由", Grow: 1.4}, {Header: "状态", Width: 110}, {Header: "请求数", Width: 110, Align: ui.End}}
	kumo.LayerCard(c, kumo.LayerCardOptions{Title: "Workers"}, func() {
		kumo.Table(c, cols, len(g.rows), func(r, col int) {
			w := g.rows[r]
			switch col {
			case 0:
				kumo.Text(c, w.name, kumo.TextOptions{Bold: true})
			case 1:
				kumo.Text(c, w.route, kumo.TextOptions{Variant: kumo.TextMonoSecondary, Size: kumo.SizeSM, Truncate: true})
			case 2:
				v := kumo.BadgeSuccess
				switch w.status {
				case "构建中":
					v = kumo.BadgeWarning
				case "失败":
					v = kumo.BadgeError
				}
				kumo.Badge(c, w.status, kumo.BadgeOptions{Variant: v, Dot: true})
			case 3:
				kumo.Text(c, fmt.Sprint(w.requests))
			}
		}, kumo.TableOptions{Selected: func(r int) *bool { return &g.rows[r].selected }})
	})
}

func (g *gallery) code(c *ui.Context) {
	section(c, "CodeBlock", "", func() {
		kumo.CodeBlock(c, `export default {
  async fetch(request, env) {
    return new Response("Hello from the edge!");
  },
};`, kumo.CodeBlockOptions{Copy: true})
		ui.Row(c).Gap(4).AlignItems(ui.Baseline).Wrap().Children(func() {
			kumo.Text(c, "行内代码：运行")
			kumo.Code(c, "go tool mygo dev")
			kumo.Text(c, "启动开发模式。")
		})
	})
}
