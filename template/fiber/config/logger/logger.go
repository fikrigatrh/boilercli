package logger

import (
	"boilerplate/config/infra"
	zerolog "github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"time"
)

type ThirdPartyLog struct {
	ID          string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Level       string    `gorm:"type:varchar(10)"`
	ReferenceId string    `gorm:"type:varchar(255)"`
	RequestId   string    `gorm:"type:varchar(50)"`
	Url         string    `gorm:"column:url;type:varchar(255)"`
	Request     string    `gorm:"type:jsonb;not null"`
	Response    string    `gorm:"type:jsonb;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	Channel     string    `gorm:"column:channel;type:varchar(100)"`
	Endpoint    string    `gorm:"column:endpoint;type:varchar(100)"`
}

type LoggerDb struct {
	LogChan chan ThirdPartyLog
	Db      *gorm.DB
}

func NewLoggerDb(infra *infra.Infra) *LoggerDb {
	logger := &LoggerDb{
		LogChan: make(chan ThirdPartyLog, 100),
		Db:      infra.DbPsql.DB,
	}

	go logger.ProcessLogs() // Start async log processing

	return logger
}

func (l *LoggerDb) ProcessLogs() {
	for logEntry := range l.LogChan {
		if err := l.Db.Create(&logEntry).Error; err != nil {
			zerolog.Error().Err(err).Msg("failed to save log")
		}
	}
}

func (l *LoggerDb) Log(logEntry ThirdPartyLog) {
	l.LogChan <- logEntry
}

func (l *LoggerDb) Close() {
	close(l.LogChan)
}
