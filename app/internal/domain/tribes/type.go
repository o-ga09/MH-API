package Tribes

type TribeId struct{ value string }
type TribeName struct{ value string }
type TribeNameJA struct{ value string }
type TribeNameEN struct{ value string }
type TribeDescription struct{ value string }

func (f *TribeId) GetID() string                   { return f.value }
func (f *TribeName) GetName() string               { return f.value }
func (f *TribeNameJA) GetNameJA() string           { return f.value }
func (f *TribeNameEN) GetNameEN() string           { return f.value }
func (f *TribeDescription) GetDescription() string { return f.value }
