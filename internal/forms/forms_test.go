package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	// to use the built-in PostForm struct,
	// must create a request type.
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	// if form is not valid
	if !isValid {
		t.Error("got invalid, wanted valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	// create new form
	form := New(r.PostForm)

	// for this form, require the fields, a,b,c
	form.Required("a", "b", "c")
	// if form valid, test should fail because these fields don't exist.
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	// if form is not valid
	if !form.Valid() {
		t.Error("shows that it does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	// situation when there are not values
	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not")
	}

	// situation when there are values
	postedValues = url.Values{}
	postedValues.Add("a", "apple")
	form = New(postedValues)

	// check the existence of the field "a" in the post request
	has = form.Has("a")
	if !has {
		t.Error("form shows does not have error when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	// 1st. check that MinLength does not work for a non existing field.
	// situation when there are values
	form.MinLength("does-not-exist", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	// if this is reached,
	// then there should be an error in the form.Errors struct
	isError := form.Errors.Get("does-not-exist")
	// must check both conditions (ie return statements) in Get()
	if isError == "" { // didnt work
		t.Error("should have an error but did not get one")
	}

	// 2nd test. check that MinLength does not show valid
	// when field is smaller than the minimum length specified.
	postedValues = url.Values{}
	postedValues.Add("a", "apple")

	form = New(postedValues)

	// "a" is < 100. form should not be valid
	form.MinLength("a", 100)
	if form.Valid() {
		t.Error("shows MinLength of 100 met when data is shorter")
	}

	// 3rd test.
	// check
	postedValues = url.Values{}
	postedValues.Add("b", "banana")
	form = New(postedValues)

	// len(b) >= 1. This should be valid.
	form.MinLength("b", 1)
	if !form.Valid() {
		t.Error("shows min of 1 is not met when it is")
	}

	// if this is reached, then Error struct is empty and should pass.
	// because we know "b" field exists.
	isError = form.Errors.Get("b")
	if isError != "" { // didnt work
		t.Error("should not have an error but got one")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	// Test 1. when x doesn't exist
	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	// Test 2. check valid email
	postedValues = url.Values{}
	postedValues.Add("email", "me@here.com")
	form = New(postedValues) // reinitialize form
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}

	// Test 3. check invalid email
	// add value to form 'apple'
	postedValues = url.Values{}
	postedValues.Add("email", "x")
	form = New(postedValues)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid for invalid email address")
	}
}
