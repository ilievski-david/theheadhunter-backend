package handlers_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	handlers "github.com/ilievski-david/theheadhunter-backend/mocks"
)

func TestAddColor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHandler := handlers.NewMockHandler(ctrl)

	mockHandler.EXPECT().AddColor(gomock.Any()).Return()
}
