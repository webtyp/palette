//go:build !(wasm && browser)

package ui

import "webtyp.com/model"

func applyPaletteCSS(_ []model.Item) {}
