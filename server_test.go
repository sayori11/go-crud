package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type TestTable[T int | float64] struct {
	a, b, result T
}

func TestSum(t *testing.T) {
	intTests := []TestTable[int]{
		{1, 2, 3},
		{4, 5, 9},
	}

	floatTests := []TestTable[float64]{
		{5.5, 3.5, 9.0},
		{4.0, 6.0, 10},
	}
	RunSumTests(t, intTests)
	RunSumTests(t, floatTests)

}

func RunSumTests[T int | float64](t *testing.T, tests []TestTable[T]) {
	for _, tt := range tests {
		testname := fmt.Sprintf("%v+%v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := Sum(tt.a, tt.b)
			assert.Equal(t, ans, tt.result)
		})
	}
}

func TestValidation(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	productJSON := `{"price": 40}`
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(productJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()

	err := updateProduct(c)
	assert.Error(t, err.(*echo.HTTPError))
}
