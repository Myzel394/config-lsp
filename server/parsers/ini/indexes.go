package ini

type IniPropertyLocationIndex struct {
	// Which section this property belongs to
	// If `nil` = section is root section
	Section *Section
	// The Property itself
	Property Property
}
