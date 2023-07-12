package pluginhelper

import (
	"fmt"
	"plugin"

	"github.com/solsw/generichelper"
)

// Value returns value of type 'T' of a symbol named 'symName' in [plugin] 'p'.
//
// [plugin]: https://pkg.go.dev/plugin#Plugin
func Value[T any](p *plugin.Plugin, symName string) (T, error) {
	symbol, err := p.Lookup(symName)
	if err != nil {
		return generichelper.ZeroValue[T](), err
	}
	t, ok := symbol.(T)
	if !ok {
		return generichelper.ZeroValue[T](), fmt.Errorf("symbol %s has wrong type: %T", symName, symbol)
	}
	return t, nil
}

// ValuePath returns value of type 'T' of a symbol named 'symName' in [plugin] with 'path'.
//
// [plugin]: https://pkg.go.dev/plugin#Plugin
func ValuePath[T any](path, symName string) (T, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return generichelper.ZeroValue[T](), err
	}
	return Value[T](p, symName)
}
