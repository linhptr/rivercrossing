package state

import "testing"

func TestPutInBoat(t *testing.T) {
	//Hva jeg forventer?
	wanted := "[V---\\___korn___\\ \\___rev+HS___/____kylling__/--Ø]"
	state := PutInBoat("rev"); //Hva jeg fikk
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)

	}
}
