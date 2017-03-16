package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestToyRobot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ToyRobot Suite")
}
