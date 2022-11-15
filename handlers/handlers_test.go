package handlers_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	handlers "github.com/ilievski-david/theheadhunter-backend/handlers"
	"github.com/ilievski-david/theheadhunter-backend/mocks/crud"
	"github.com/ilievski-david/theheadhunter-backend/models"
	"gotest.tools/assert"
)

func TestAddColor(t *testing.T) {
	t.Run("Correct", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCrud := crud.NewMockDatabase(ctrl)
		handler := handlers.NewHandler(mockCrud)

		color := models.ColorPost{
			UserToken: "test",
			Hex:       "test",
			Name:      "test",
		}

		mockCrud.EXPECT().InsertColor(color).Return(nil).Times(1)
		mockCrud.EXPECT().IgnoreOrInsertUser(color.UserToken).Return(nil).Times(1)
		mockCrud.EXPECT().QueryColors(color.UserToken).Return([]models.Color{}, nil).Times(1)

		router := gin.Default()
		rr := httptest.NewRecorder()

		router.POST("/addColor", handler.AddColor)
		reqBody, _ := json.Marshal(color)
		request, _ := http.NewRequest(http.MethodPost, "/addColor", bytes.NewBuffer(reqBody))
		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)
		//read response and validate response
		assert.Equal(t, 200, rr.Code)

		bodyBytes, err := ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Fail()
		}
		var body struct {
			Code int
		}
		err = json.Unmarshal(bodyBytes, &body)
		if err != nil {
			t.Fail()
		}
		assert.Equal(t, 200, body.Code)
	})
}
