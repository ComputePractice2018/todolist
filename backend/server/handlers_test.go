package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ComputePractice2018/todolist/backend/data"
)

var testlist = "{\"id\":122,\"name\":\"te\",\"success\":false}"

func TestCrud(t *testing.T) {
	lt := data.NewListTask()
	router := NewRouter(lt)

	//GetList
	req, err := http.NewRequest("GET", "/api/todolist/task/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Error("Expected 200 code")
	}
	if resp.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Error("Expected json MIME-type")
	}

	testData := strings.NewReader(testlist)

	//AddList
	req, err = http.NewRequest("POST", "/api/todolist/task/", testData)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Error("Expected 201 code")
	}
	location := resp.Header.Get("Location")
	if location[len(location)-1:] != "1" {
		t.Error("Expected another location")
	}

	if len(lt.GetList()) != 1 {
		t.Error("Expected new value")
	}

	//EditList
	testData = strings.NewReader(testlist)
	req, err = http.NewRequest("PUT", "/api/todolist/task/1", testData)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusAccepted {
		t.Error("Expected 201 code")
	}
	location = resp.Header.Get("Location")
	if location[len(location)-1:] != "1" {
		t.Error("Expected another location")
	}

	if len(lt.GetList()) != 1 {
		t.Error("Expected new value")
	}

	//DeleteList
	req, err = http.NewRequest("DELETE", "/api/todolist/task/1", nil)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusNoContent {
		t.Error("Expected 204 code ")
	}
	if len(lt.GetList()) != 0 {
		t.Error("Expected new value")
	}

}
