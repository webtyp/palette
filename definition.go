package palette

import (
	"webtyp.com/color"
	"webtyp.com/input"
	"webtyp.com/model"
)

// PaletteColorModel defines a database-stored colour swatch that any module
// (specialties, roles, tags, etc.) can reference without hardcoding colours in code.
var PaletteColorModel = model.Definition{
	Name:  "palette_color",
	Allow: model.Create | model.Read | model.Update | model.Delete,
	Rules: model.Rules{
		Create: model.AddMyBusinessRule,
		Read:   model.HaveNoBusinessRule,
		Update: model.AddMyBusinessRule,
		Delete: model.HaveNoBusinessRule,
	},
	Fields: model.Fields{
		{Name: "id", Type: model.Text(), DB: &model.FieldDB{PK: true}, OmitEmpty: true, ServerOnly: true},
		{Name: "tenant_id", Type: model.Text(), NotNull: true, ServerOnly: true},
		{Name: "name", Type: input.Text(), Label: "Name", NotNull: true,
			Permitted: model.Permitted{Letters: true, Tilde: true, Numbers: true, Spaces: true, Extra: []rune{'-', '(', ')'}, Minimum: 1, Maximum: 60}},
		{Name: "hex", Type: color.Input(), Label: "Color", Help: "Base hex colour (#RRGGBB). Translucent fill and legible text adapt automatically to light and dark themes.", NotNull: true},
		{Name: "updated_at", Type: model.NanoTime(), OmitEmpty: true, ServerOnly: true},
	},
}

// Definition is the module declaration mounted by appkit.App.Modules.
var Definition = model.ModuleDefinition{
	Name:     "palette",
	Entities: []model.Definition{PaletteColorModel},
}

// Domain errors.
const (
	ErrNoDB          model.Invalid  = "palette: database is required"
	ErrNoIDs         model.Invalid  = "palette: id generator is required"
	ErrNoTenant      model.Invalid  = "palette: tenant is required"
	ErrDuplicateName model.Conflict = "palette: a colour with that name already exists"
)
