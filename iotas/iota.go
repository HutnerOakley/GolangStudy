package iotas

type Allergen int

const (
	CCVisa            = "Visa"
	CCMasterCard      = "MasterCard"
	CCAmericanExpress = "American Express"
)

const (
	CategoryBooks = iota
	CategoryHealth
	CategoryClothing
)

const (
	IgEggs Allergen = 1 << iota
	IgChocolate
	IgNuts
	IgStrawberries
	IgShellfish
)
