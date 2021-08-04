package main

import (
	"net/http"
	"reflect"
	"testing"
)

func TestCreateBaseUrl(t *testing.T) {
	var tests = []struct {
		input string
	}{
		{

			"testdomain",
		},
		{

			"testdomain1",
		},
	}
	for _, test := range tests {
		got := createBaseUrl(test.input)
		if reflect.String != reflect.TypeOf(got).Kind() {
			t.Errorf("createBaseUrl(string) failed to return a string")
		}

	}
}

func TestBasicAuth(t *testing.T) {
	tests := []struct {
		InputA string
		InputB string
	}{
		{
			"a",
			"2vefeefxwx",
		}, {
			"test",
			"testing",
		},
	}
	for _, test := range tests {
		got := basicAuth(test.InputA, test.InputB)
		if reflect.String != reflect.TypeOf(got).Kind() {
			t.Errorf("basicAuth(string,string) failed to return a string")
		}

	}
}
func TestLogin(t *testing.T) {
	var test = struct {
		want   int
		input  string
		method string
	}{
		http.StatusOK,
		"",
		http.MethodGet,
	}
	mock := NewAuth()
	got, err := mock.Login(test.method, test.input)
	if err != nil {
		t.Errorf("login(string,string) returned an error %v with url %v and method %v.\nExpected status code %v but received %v", err, got.Request.URL.String(), got.Request.Method, test.want, got.Status)
	}
	if got.StatusCode == test.want {
		t.Errorf("login(string,string) returned an error %v with url %v and method %v.\nExpected status code %v but received %v", err, got.Request.URL.String(), got.Request.Method, test.want, got.Status)
	}
}
