package dubbo

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

func TestDubbo(t *testing.T) {
	logger := zaptest.NewLogger(t, zaptest.Level(zapcore.DebugLevel))
	Convey("TestDubbo", t, func() {
		r, err := Dubbo(context.Background(), "3489c0d3.dcrencai.org.cn:20880")
		So(err, ShouldBeNil)
		logger.Info("TestDubbo", zap.Any("result", r))
	})
}
