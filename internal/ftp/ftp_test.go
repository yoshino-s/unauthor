package ftp

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

func TestFtp(t *testing.T) {
	logger := zaptest.NewLogger(t, zaptest.Level(zapcore.DebugLevel))
	Convey("TestFtp", t, func() {
		r, err := Ftp(context.Background(), "www.pcoip.cn:21")
		So(err, ShouldBeNil)
		logger.Info("TestFtp", zap.Any("result", r))
	})
}
