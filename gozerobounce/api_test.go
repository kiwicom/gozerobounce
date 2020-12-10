package gozerobounce

import "testing"

const KEY = ""

func TestIfIsDisposable(t *testing.T) {

	APIKey = KEY

	email := "disposable@example.com"
	resp, err := Validate(email, "99.110.204.1")
	if err != nil {
		t.Errorf(err.Error())
	}

	if !resp.IsDisposable() {
		t.Errorf("Email " + email + "should be disposable")
	}
}

func TestIfIsValid(t *testing.T) {

	APIKey = KEY

	email := "valid@example.com"
	resp, err := Validate(email, "99.110.204.1")
	if err != nil {
		t.Errorf(err.Error())
	}

	if !resp.IsValid() {
		t.Errorf("Email " + email + "should be valid")
	}

	email = "invalid@example.com"
	resp, err = Validate(email, "99.110.204.1")
	if err != nil {
		t.Errorf(err.Error())
	}

	if resp.IsValid() {
		t.Errorf("Email " + email + "should not be valid")
	}
}
