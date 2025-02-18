package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"fmt"
)

func InitLogger() {
    consoleWriter := zerolog.ConsoleWriter{
        Out:        os.Stdout,
        TimeFormat: "2006-01-02 15:04:05",
        NoColor:    false,
        FormatLevel: func(i interface{}) string {
            var l string
            if ll, ok := i.(string); ok {
                switch ll {
                case "debug":
                    l = "\x1b[36mDEBUG\x1b[0m" // 青色
                case "info":
                    l = "\x1b[32mINFO\x1b[0m"  // 绿色
                case "warn":
                    l = "\x1b[33mWARN\x1b[0m"  // 黄色
                case "error":
                    l = "\x1b[31mERROR\x1b[0m" // 红色
                case "fatal":
                    l = "\x1b[35mFATAL\x1b[0m" // 紫色
                default:
                    l = ll
                }
            }
            return fmt.Sprintf("| %-6s|", l)
        },
        FormatMessage: func(i interface{}) string {
            return fmt.Sprintf("message: %s", i)
        },
        FormatCaller: func(i interface{}) string {
            return fmt.Sprintf("%s:", i)
        },
    }

    // 设置全局日志配置
    log.Logger = zerolog.New(consoleWriter).
        Level(zerolog.DebugLevel).
        With().
        Timestamp().
        Caller().
        Logger()
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Error(msg string) {
	log.Error().Msg(msg)
}

func Warn(msg string) {
	log.Warn().Msg(msg)
}

func Debug(msg string) {
	log.Debug().Msg(msg)
}

func Trace(msg string) {
	log.Trace().Msg(msg)
}