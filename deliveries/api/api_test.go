package api_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danielfmelo/travel_finder/deliveries/api"
	"github.com/danielfmelo/travel_finder/entity"
	"github.com/danielfmelo/travel_finder/tests/mocks"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestGetBestPrice(t *testing.T) {

	var testCases = []struct {
		name               string
		expectedBody       string
		ori                string
		dest               string
		val                string
		path               string
		value              int
		expectedError      error
		expectedStatusCode int
		url                string
	}{
		{
			name:               "test should get correct best price",
			expectedBody:       `{"path":"GRU -> FLN","value":40}`,
			ori:                "GRU",
			dest:               "FLN",
			val:                "40",
			path:               `GRU -> FLN`,
			value:              40,
			expectedError:      nil,
			expectedStatusCode: http.StatusOK,
			url:                "/origins/GRU/destinations/FLN",
		},
		{
			name:               "test should return bad request",
			expectedBody:       "",
			path:               "",
			expectedError:      nil,
			expectedStatusCode: http.StatusBadRequest,
			url:                "/origins//destinations/FLN",
		},
		{
			name:               "test should return bad request",
			expectedBody:       "",
			path:               "",
			ori:                "GRU",
			dest:               "FLN",
			expectedError:      errors.New("Not found"),
			expectedStatusCode: http.StatusNotFound,
			url:                "/origins/GRU/destinations/FLN",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			router := httprouter.New()
			finder := &mocks.FinderMock{}
			cheapestRoute := entity.CheapestRoute{
				Path:  tc.path,
				Value: tc.value,
			}

			finder.On("GetSmallestPriceAndRoute", tc.ori, tc.dest).Return(cheapestRoute, tc.expectedError)

			a := api.New(finder)
			a.Handlers(router)

			url := tc.url
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				t.Error(err)
			}

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)
			assert.Equal(t, res.Code, tc.expectedStatusCode)
			if res.Code == http.StatusOK {
				cr := entity.CheapestRoute{}
				body, err := ioutil.ReadAll(res.Body)
				assert.NoError(t, err)
				err = json.Unmarshal(body, &cr)
				assert.NoError(t, err)
				assert.Equal(t, tc.value, cr.Value)
				assert.Equal(t, tc.path, cr.Path)
			}
		})
	}
}

func TestPostInsertRoute(t *testing.T) {

	var testCases = []struct {
		name               string
		expectedError      error
		ori                string
		dest               string
		val                string
		expectedStatusCode int
		url                string
	}{
		{
			name:               "test should get correct best price",
			expectedError:      nil,
			expectedStatusCode: http.StatusCreated,
			url:                "/origins/GRU/destinations/FLN/values/40",
			ori:                "GRU",
			dest:               "FLN",
			val:                "40",
		},
		{
			name:               "test should return bad request",
			expectedError:      nil,
			expectedStatusCode: http.StatusBadRequest,
			url:                "/origins/GRU/destinations//values/40",
		},
		{
			name:               "test should return internal server error",
			expectedError:      errors.New("Not found"),
			expectedStatusCode: http.StatusInternalServerError,
			url:                "/origins/GRU/destinations/FLN/values/40",
			ori:                "GRU",
			dest:               "FLN",
			val:                "40",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			router := httprouter.New()
			finder := &mocks.FinderMock{}
			rec := entity.Record{tc.ori, tc.dest, tc.val}
			finder.On("Save", rec).Return(tc.expectedError).Once()
			a := api.New(finder)
			a.Handlers(router)

			url := tc.url
			req, err := http.NewRequest("POST", url, nil)
			if err != nil {
				t.Error(err)
			}

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			if res.Code != tc.expectedStatusCode {
				t.Errorf("Status code is not same expected, actual %d expected: %d", res.Code, tc.expectedStatusCode)
			}
		})
	}
}
