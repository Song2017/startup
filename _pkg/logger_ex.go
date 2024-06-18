package pkg

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogFormat struct {
	Platform   string      `json:"smklog.platform"`
	Latency    int64       `json:"smklog.latency"`
	Category   string      `json:"smklog.category"`
	Timing     time.Time   `json:"smklog.timing"`
	Label      string      `json:"smklog.label"`
	RemoteAddr string      `json:"smklog.remote_addr"`
	Input      interface{} `json:"smklog.input"`
	Output     interface{} `json:"smklog.output"`
}

var (
	ZapLogger *zap.Logger
	ZapOnce   sync.Once
)

func InitZapLogger() *zap.Logger {
	ZapOnce.Do(
		func() {
			if ZapLogger == nil {
				writeSyncer := zapcore.AddSync(os.Stdout)
				encoderConfig := zap.NewProductionEncoderConfig()
				encoderConfig.TimeKey = "smklog.timing"
				encoderConfig.LevelKey = "smklog.level"
				encoderConfig.MessageKey = "smklog.data"
				encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
				encoder := zapcore.NewJSONEncoder(encoderConfig)
				levelEnabler := zap.NewAtomicLevelAt(zap.InfoLevel)
				core := zapcore.NewCore(encoder, writeSyncer, levelEnabler)
				ZapLogger = zap.New(core)
			}
		},
	)
	return ZapLogger
}

func (l LogFormat) ToString() string {
	if data, err := json.Marshal(l); err != nil {
		return err.Error()
	} else {
		return string(data)
	}
}

func (log LogFormat) LoggerInfo() {
	InitZapLogger().Info(
		log.ToString(), zap.Field{
			Key:    "smklog.label",
			String: log.Label,
			Type:   zapcore.StringType,
		},
	)
}

func (log LogFormat) LoggerError() {
	InitZapLogger().Error(
		log.ToString(),
		zap.Field{
			Key:    "smklog.label",
			String: log.Label,
			Type:   zapcore.StringType,
		})
}

func LoggerInfo(platform, label string, input, output interface{}) {
	nomadLog := LogFormat{
		Platform: platform,
		Label:    label,
		Input:    input,
		Output:   output,
	}

	nomadLog.LoggerInfo()
}

func LoggerError(platform, label string, input, output interface{}) {
	nomadLog := LogFormat{
		Platform: platform,
		Label:    label,
		Input:    input,
		Output:   output,
	}

	nomadLog.LoggerError()
}
