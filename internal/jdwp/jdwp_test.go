package jdwp

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

func TestJdwp(t *testing.T) {
	logger := zaptest.NewLogger(t, zaptest.Level(zapcore.DebugLevel))
	Convey("TestJdwp", t, func() {
		r, err := Jdwp(context.Background(), "appdev.com.magicflu.com:8000")
		So(err, ShouldBeNil)
		logger.Info("TestJdwp", zap.Any("result", r))
	})
}
