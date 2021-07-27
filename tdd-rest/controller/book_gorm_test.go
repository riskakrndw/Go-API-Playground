package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"tdd/rest/config"
	"tdd/rest/database"
	"tdd/rest/model"
	"testing"

	"github.com/labstack/echo"
)

func TestDBPostBookController(t *testing.T) {
	//bikin db
	db, err := database.CreateDB(config.TEST_DB_CONNECTION_STRING)
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&model.Book{})
	db.Delete(&model.Book{}, "1=1")
	m := model.NewGormBookModel(db)

	//bikin controller
	postBookController := CreatePostBookController(m)
	if err != nil {
		t.Error(err)
	}

	//test controller
	testPostBookController(t, postBookController)
	db.Delete(&model.Book{}, "1=1")
}

func testPostBookController(t *testing.T, postBookController echo.HandlerFunc) {
	//request
	jsonContent := `{"title": "Sebuah Seni untuk Bersikap Bodo Amat"}`
	req := httptest.NewRequest(http.MethodPost, "/book", strings.NewReader(jsonContent))
	// req := httptest.NewRequest(http.MethodGet, "/", nil)
	v := url.Values{}
	v.Set("title", "tes judul")
	req.Form = v
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	postBookController(c)

	//test
	statusCode := rec.Result().StatusCode
	if statusCode != 200 {
		t.Errorf("Response is not 200: %d", statusCode)
	}

	body := rec.Body.Bytes()
	var book model.Book
	if err := json.Unmarshal(body, &book); err != nil {
		t.Error(err)
	}
	if book.Title != "tes judul" {
		t.Errorf("expected Harry Potter, got: %#v", book.Title)
	}
}

func testGetBookController(t *testing.T, bookController echo.HandlerFunc) {
	// coba request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	bookController(c)

	// test
	statusCode := rec.Result().StatusCode
	if statusCode != 200 {
		t.Errorf("Response is not 200: %d", statusCode)
	}
	body := rec.Body.Bytes()
	var books []model.Book
	if err := json.Unmarshal(body, &books); err != nil {
		t.Error(err)
	}
	if len(books) != 1 {
		t.Errorf("expected one book, got: %#v", books)
		return
	}
	if books[0].Title != "Harry Potter" {
		t.Errorf("expected Harry Potter, got: %#v", books[0].Title)
	}
}

func TestDBGetBookController(t *testing.T) {
	// bikin db
	db, err := database.CreateDB(config.TEST_DB_CONNECTION_STRING)
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&model.Book{})
	db.Delete(&model.Book{}, "1=1")
	m := model.NewGormBookModel(db)

	// bikin controller
	bookController := CreateGetBookController(m)
	if err != nil {
		t.Error(err)
	}

	// insert data baru
	m.Insert(model.Book{Title: "Harry Potter"})

	// test controller
	testGetBookController(t, bookController)
	db.Delete(&model.Book{}, "1=1")
}

func TestMockGetBookController(t *testing.T) {
	m := model.NewMockBookModel()
	bookController := CreateGetBookController(m)
	// insert data baru
	m.Insert(model.Book{Title: "Harry Potter"})
	// test controller
	testGetBookController(t, bookController)
}

func TestMockPostBookController(t *testing.T) {
	m := model.NewMockBookModel()
	postBookController := CreatePostBookController(m)
	testPostBookController(t, postBookController)
}
