package main

import (
	"github.com/stretchr/testify/suite"
)

// ParkappSuite is a suite of autotests
type ParkappSuite struct {
	suite.Suite

	parkappServerAddress string
}

// SetupSuite bootstraps suite dependencies
func (suite *ParkappSuite) SetupSuite() {
	// check required flags
	suite.Require().NotEmpty(flagParkappHost, "-parkapp-host non-empty flag required")
	suite.Require().NotEmpty(flagParkappPort, "-parkapp-port non-empty flag required")

	suite.parkappServerAddress = "http://" + flagParkappHost + ":" + flagParkappPort
}
