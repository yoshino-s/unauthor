package redis

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

func TestRedis(t *testing.T) {
	logger := zaptest.NewLogger(t, zaptest.Level(zapcore.DebugLevel))
	Convey("TestRedis", t, func() {
		r, err := Redis(context.Background(), "127.0.0.1:6379")
		So(err, ShouldBeNil)
		logger.Info("TestRedis", zap.Any("result", r))
	})
}
