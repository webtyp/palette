package migrate

import (
	"webtyp.com/ddl"
	"webtyp.com/palette"
)

// Migrate reconciles the palette_color table schema.
func Migrate(conn ddl.Execer, compiler ddl.Compiler) error {
	return ddl.New(conn, compiler).Sync(&palette.PaletteColor{})
}
