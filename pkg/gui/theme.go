package gui

func (gui *Gui) SetColorScheme() error {
	gui.g.FgColor = GetGocuiStyle([]string{"default"})
	gui.g.SelFgColor = GetGocuiStyle([]string{"bold", "green"})
	return nil
}
