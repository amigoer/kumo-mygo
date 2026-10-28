package fill

import (
	"sync"

	"github.com/egoist/mygo/ui"
)

// lazyIcon parses its SVG on first use. Zero-value caches keep package
// initialization from rooting icons a program never draws.
type lazyIcon struct {
	once sync.Once
	svg  *ui.SVG
}

func (i *lazyIcon) get(data string) *ui.SVG {
	i.once.Do(func() { i.svg = ui.MustParseSVG([]byte(data)) })
	return i.svg
}
