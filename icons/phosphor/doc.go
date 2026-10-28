// Package phosphor provides the regular weight of Phosphor Icons, the icon
// set Cloudflare's Kumo uses, as MyGo SVGs.
//
//	kumo.Icon(c, phosphor.MagnifyingGlass)
//	ui.Icon(c, phosphor.MagnifyingGlass())
//
// Each accessor parses its SVG once, on first use, and returns the same
// shared pointer; a program retains only the icons it calls. Icons fill with
// the color of the text, as MyGo's Icon draws any SVG. Names follow upstream
// kebab-case in PascalCase (caret-down is CaretDown); names starting with a
// digit take the prefix Icon. The fill weight is in package fill, and
// package catalog maps every name to its accessor for galleries.
//
// The generated sources are pinned in upstream.json and covered by the
// adjacent Phosphor LICENSE (MIT).
package phosphor
