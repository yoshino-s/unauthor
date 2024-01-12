package memcached

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

func TestMemcached(t *testing.T) {
	logger := zaptest.NewLogger(t, zaptest.Level(zapcore.DebugLevel))
	Convey("TestMemcached", t, func() {
		r, err := Memcached(context.Background(), "xinss.gnway.cc:11211")
		So(err, ShouldBeNil)
		logger.Info("TesTestMemcachedtZookeeper", zap.Any("result", r))
	})
}
