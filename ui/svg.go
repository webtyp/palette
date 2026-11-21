//go:build !wasm

package ui

import (
	"webtyp.com/svg"
	"webtyp.com/svg/sprite"
)

// Icons declares the palette SVG symbol for sprite extraction.
type Icons struct{}

// IconSvg provides the palette icon symbol for the app sprite sheet.
func (*Icons) IconSvg() *sprite.Sprite {
	return sprite.NewSprite(
		sprite.Define(svg.Icon(ID), "0 0 24 24",
			sprite.Path("M12 2a10 10 0 0 0 0 20 2 2 0 0 0 2-2v-.5a1.5 1.5 0 0 1 1.5-1.5H17a5 5 0 0 0 0-10h-1a10 10 0 0 0-4-6Z"),
		),
	)
}
