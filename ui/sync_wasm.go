//go:build wasm && browser

package ui

import (
	"syscall/js"

	"webtyp.com/model"
)

// applyPaletteCSS writes --palette-<id> CSS custom properties directly on
// document.documentElement.style via CSSOM setProperty, which is allowed under
// strict CSP (style-src 'self') without creating any inline <style> element.
func applyPaletteCSS(items []model.Item) {
	doc := js.Global().Get("document")
	if !doc.Truthy() {
		return
	}
	rootStyle := doc.Get("documentElement").Get("style")
	if !rootStyle.Truthy() {
		return
	}
	for _, it := range items {
		if it.ID != "" && it.Description != "" {
			rootStyle.Call("setProperty", "--palette-"+it.ID, it.Description)
		}
	}
}
