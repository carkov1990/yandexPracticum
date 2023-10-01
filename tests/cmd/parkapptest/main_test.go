package main

//go:generate go test -c -o=../../bin/parkapptest

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestParkapp(t *testing.T) {
	suite.Run(t, new(ParkappSuite))
}
