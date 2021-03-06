package onet

import (
	"fmt"
	"testing"

	"go.dedis.ch/onet/v3/ciphersuite"
	"go.dedis.ch/onet/v3/log"
)

//var testSuite = ciphersuite.NewEd25519CipherSuite()

var testSuite = ciphersuite.NewHope()
var testRegistry = ciphersuite.NewRegistry()

// To avoid setting up testing-verbosity in all tests
func TestMain(m *testing.M) {
	fmt.Println("Bjorn")
	testRegistry.RegisterCipherSuite(testSuite)
	log.MainTest(m)
}
