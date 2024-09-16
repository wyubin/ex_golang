package test

import (
	"example.com/ex_mock/Mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCARSTRUCT_Start(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	MockInterface := Mocks.NewMockCAR(controller)
	var ctx *gin.Context
	MockInterface.EXPECT().EngineCheck(ctx, MockInterface).Return(true)
	result := CARSTRUCT{}.Start(ctx, "carkey", MockInterface)
	assert.Equal(t, false, result)
}
