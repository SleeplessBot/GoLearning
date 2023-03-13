package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

const (
	traceStr     = "%s [%.3fms] [rows:%v] %s"
	traceWarnStr = "%s %s [%.3fms] [rows:%v] %s"
	traceErrStr  = "%s %s [%.3fms] [rows:%v] %s"
)

type GoZeroGormLogger struct {
	logx.Logger
	LogLevel                  logger.LogLevel
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
}

func NewGoZeroGormLogger() GoZeroGormLogger {
	return GoZeroGormLogger{
		Logger:                    logx.WithContext(context.Background()),
		LogLevel:                  logger.Warn,
		SlowThreshold:             100 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: false,
	}
}

func (l GoZeroGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return GoZeroGormLogger{
		Logger:                    logx.WithContext(context.Background()),
		SlowThreshold:             l.SlowThreshold,
		LogLevel:                  level,
		SkipCallerLookup:          l.SkipCallerLookup,
		IgnoreRecordNotFoundError: l.IgnoreRecordNotFoundError,
	}
}

func (l GoZeroGormLogger) Info(ctx context.Context, format string, args ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}
	l.Logger.WithContext(ctx).Infof(format, args...)
}

func (l GoZeroGormLogger) Warn(ctx context.Context, format string, args ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	l.Logger.WithContext(ctx).Errorf(format, args...)
}

func (l GoZeroGormLogger) Error(ctx context.Context, format string, args ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}
	l.Logger.WithContext(ctx).Errorf(format, args...)
}

func (l GoZeroGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			s := fmt.Sprintf(traceErrStr, utils.FileWithLineNum(), err.Error(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			l.Logger.WithContext(ctx).Error(s)
		} else {
			s := fmt.Sprintf(traceErrStr, utils.FileWithLineNum(), err.Error(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			l.Logger.WithContext(ctx).Error(s)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Logger.WithContext(ctx).Errorf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Logger.WithContext(ctx).Errorf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.Logger.WithContext(ctx).Infof(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Logger.WithContext(ctx).Infof(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
