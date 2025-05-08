package main

//go:generate stringer -type=Status

type Status int

const (
	Unknown Status = iota
	Active
	Inactive
)
