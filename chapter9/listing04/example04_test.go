// Sample unit test for testing an internal endpoint.
package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	ex "github.com/goinaction/code/chapter9/listing04"
)

const succeed = "\u2713"
const failed = "\u2717"

func init() {
	ex.Routes()
}

// TestDownload tests if download web content is working.
func TestSendJSON(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint.")
	{
		r, _ := http.NewRequest("GET", "/sendjson", nil)
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)

		if w.Code != 200 {
			t.Fatalf("\tShould received a status code of 200 for the response. Received[%d] %s", w.Code, failed)
		}
		t.Log("\tShould received a status code of 200 for the response.", succeed)

		u := struct {
			Name  string
			Email string
		}{}
		if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
			t.Fatal("\tShould be able to unmarshal the response.", failed)
		}
		t.Log("\tShould be able to unmarshal the response.", succeed)

		if u.Name == "Bill" {
			t.Log("\tShould have \"Bill\" for Name in the response.", succeed)
		} else {
			t.Error("\tShould have \"Bill\" for Name in the response.", failed, u.Name)
		}

		if u.Email == "bill@ardanstudios.com" {
			t.Log("\tShould have \"bill@ardanstudios.com\" for Email in the response.", succeed)
		} else {
			t.Error("\tShould have \"bill@ardanstudios.com\" for Email in the response.", failed, u.Email)
		}
	}
}