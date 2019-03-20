package main

//go:generate enameg $GOFILE

// +enameg
type HogeType int

const (
	HogeTypeA HogeType = 0 // A
	HogeTypeB HogeType = 1 // B
	HogeTypeC HogeType = 2 // C
)
