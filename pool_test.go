package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestWritePoolFormat(t *testing.T) {
	var buf bytes.Buffer
	if err := WritePool(&buf, 2358, 2); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(buf.String(), "\r\n")

	wantFirst := "2358,,2358,novm,181,,,,,,,,,,,default,2358,pjsip,PJSIP/2358,fixed,2358,2358,,,disabled,dontcare,dontcare,dontcare,dontcare,disabled,191,disabled,enabled,2358,3,2358,,yes,,no,no,,from-internal,,181,yes,,rfc4914,yes,no,,1,1,1,7200,no,no,no,,241,auto,,,,241,yes,yes,yes,no,yes,0,0,REGEN,yes,pai,chan_pjsip,yes,271,,yes,no,,,,,ENABLED,ringallv2-prim,201,,2358,181,\"ext-local,2358,dest\",,,181,,Ring,188,novm,,yes,default,,,182"
	if lines[1] != wantFirst {
		t.Fatalf("row mismatch:\ngot:  %q\nwant: %q", lines[1], wantFirst)
	}

	wantSecond := "2359,,2359,novm,182,,,,,,,,,,,default,2359,pjsip,PJSIP/2359,fixed,2359,2359,,,disabled,dontcare,dontcare,dontcare,dontcare,disabled,192,disabled,enabled,2359,3,2359,,yes,,no,no,,from-internal,,182,yes,,rfc4915,yes,no,,1,1,1,7200,no,no,no,,242,auto,,,,242,yes,yes,yes,no,yes,0,0,REGEN,yes,pai,chan_pjsip,yes,272,,yes,no,,,,,ENABLED,ringallv2-prim,202,,2359,182,\"ext-local,2359,dest\",,,182,,Ring,189,novm,,yes,default,,,183"
	if lines[2] != wantSecond {
		t.Fatalf("row mismatch:\ngot:  %q\nwant: %q", lines[2], wantSecond)
	}
}

func TestWritePoolCount(t *testing.T) {
	var buf bytes.Buffer
	if err := WritePool(&buf, 1, 50); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimRight(buf.String(), "\r\n"), "\r\n")
	if len(lines) != 51 { // header + 50 rows
		t.Fatalf("got %d lines, want 51", len(lines))
	}
}
