package zookeeper

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

func TestZookeeper(t *testing.T) {
	logger := zaptest.NewLogger(t, zaptest.Level(zapcore.DebugLevel))
	Convey("TestZookeeper", t, func() {
		r, err := Zookeeper(context.Background(), "pool.tokenstring.online:2181")
		So(err, ShouldBeNil)
		logger.Info("TestZookeeper", zap.Any("result", r))
	})
}
