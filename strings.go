package kumo

// Strings holds the words the components show or tell assistive technology
// on their own, in English. An app in another language sets them once,
// before its first window, as Kumo's locale provider does:
//
//	kumo.Strings = kumo.StringsZhHans
var Strings = StringsEn

// Messages are the components' own words.
type Messages struct {
	Optional    string // after a label whose field is optional
	Loading     string // a Loader's accessible name
	Close       string // the close buttons of dialogs, banners and toasts
	Copy        string // ClipboardText's button
	Copied      string // ClipboardText's button after copying
	Show        string // SensitiveInput's reveal button
	Hide        string // SensitiveInput's conceal button
	Previous    string // Pagination
	Next        string // Pagination
	First       string // Pagination
	Last        string // Pagination
	PageOf      string // Pagination: "Page %d of %d"
	ShowingOf   string // Pagination: "Showing %d-%d of %d"
	MoreInfo    string // a label tooltip's button
	NoResults   string // an empty select or combobox
	SelectEmpty string // a select's placeholder when nothing is chosen
}

// StringsEn are the English words, the default.
var StringsEn = Messages{
	Optional:    "(optional)",
	Loading:     "Loading",
	Close:       "Close",
	Copy:        "Copy",
	Copied:      "Copied",
	Show:        "Show",
	Hide:        "Hide",
	Previous:    "Previous page",
	Next:        "Next page",
	First:       "First page",
	Last:        "Last page",
	PageOf:      "Page %d of %d",
	ShowingOf:   "Showing %d-%d of %d",
	MoreInfo:    "More information",
	NoResults:   "No results",
	SelectEmpty: "Select…",
}

// StringsZhHans are the words in Simplified Chinese.
var StringsZhHans = Messages{
	Optional:    "（可选）",
	Loading:     "加载中",
	Close:       "关闭",
	Copy:        "复制",
	Copied:      "已复制",
	Show:        "显示",
	Hide:        "隐藏",
	Previous:    "上一页",
	Next:        "下一页",
	First:       "第一页",
	Last:        "最后一页",
	PageOf:      "第 %d 页，共 %d 页",
	ShowingOf:   "第 %d-%d 条，共 %d 条",
	MoreInfo:    "更多信息",
	NoResults:   "无结果",
	SelectEmpty: "请选择…",
}
