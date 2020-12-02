package main

import "testing"

func TestProcessInput(t *testing.T) {
	p := PolicyChecker{}
	p.processInput("../input_test1.dat")

	go func() {
		wg.Wait()
		close(p.ValidPasswords)
	}()

	nValidPasswords := p.countValid()
	if nValidPasswords != 2 {
		t.Errorf("Test failed. Got %d and expected %d",
			nValidPasswords,
			2)
	}
}
