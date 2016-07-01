package svcparse

import (
	"strings"
	"testing"
)

func TestUnderscoreIdent(t *testing.T) {
	r := strings.NewReader("service Example_Service {}")
	lex := NewSvcLexer(r)
	svc, err := ParseService(lex)

	if err != nil {
		t.Error(err)
	}
	if svc == nil {
		t.Fatalf("Returned service is nil\n")
	}
	if len(svc.Methods) != 0 {
		t.Errorf("Parser found too many methods, expected 0, got %v\n", len(svc.Methods))
	}
}

func TestAdditionalHttpOpts(t *testing.T) {
	r := strings.NewReader(`
service Example_Service {
	rpc Example(Empty) returns (Empty) {
		option (google.api.http) = {
			// Some example comment
			get: "/ExampleGet"
			body: "*"

			additional_bindings {
				post: "/ExamplePost"
			}
		}
	}
}`)
	lex := NewSvcLexer(r)
	svc, err := ParseService(lex)

	if err != nil {
		t.Error(err)
	}
	if svc == nil {
		t.Fatalf("Returned service is nil\n")
	}
}
