package jwt

import (
	"math"
	"testing"
)

func TestJWT(t *testing.T) {

	id := uint32(math.MaxUint32)

	jwt, err := GenerateJWT(id)
	if err != nil {
		t.Fail()
		return
	}

	t.Logf("jwt generate ok!\njwt: %v\n", jwt)

	claims, err := ParseToken(jwt)
	if err != nil {
		t.Fail()
		return
	}

	t.Logf("jwt parse ok!\nuser id: %v\n", claims)
}

func TestParseToken(t *testing.T) {

	id := uint32(123)

	jwt, err := GenerateJWT(id)
	if err != nil {
		t.Fail()
		return
	}

	t.Logf("jwt generate ok!\njwt: %v\n", jwt)

	claims, err := ParseToken(jwt + "1")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	t.Logf("jwt parse ok!\nuser id: %v\n", claims)

}
