package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/congim/xpush/config"
	"github.com/congim/xpush/pkg/message"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// MySQL mysql
type MySQL struct {
	DB     *sql.DB
	conf   *config.Mysql
	logger *zap.Logger
}

// New new sql
func New(conf *config.Mysql, logger *zap.Logger) *MySQL {
	return &MySQL{
		conf:   conf,
		logger: logger,
	}
}

// Start connect mysql
func (m *MySQL) Init() error {
	// 初始化mysql连接
	sqlConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", m.conf.Acc, m.conf.Passwd, m.conf.Addr, m.conf.Port, m.conf.Database)
	db, err := sql.Open("mysql", sqlConn)
	if err != nil {
		m.logger.Error("mysql init failed", zap.Error(err), zap.String("sqlConn", sqlConn))
		return err
	}

	// 测试db是否正常
	if err := db.Ping(); err != nil {
		m.logger.Error("mysql ping failed", zap.Error(err))
		return err
	}
	m.DB = db

	return nil
}

// Store ...
func (m *MySQL) Store(msgs []*message.Message, msgIDs []string) error {
	if len(msgs) != len(msgIDs) {
		return fmt.Errorf("msgids长度和msgs长度不一致")
	}
	tx, err := m.DB.Begin()
	if err != nil {
		m.logger.Error("store failed", zap.Error(err))
		return err
	}
	defer tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.

	stmt, err := tx.Prepare(SQLStoreMsg)
	if err != nil {
		m.logger.Error("store failed", zap.Error(err))
		return err
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for index, msg := range msgs {
		if _, err := stmt.Exec(msg.Topic, msgIDs[index], msg.ID, msg.Type, msg.Payload, time.Now()); err != nil {
			m.logger.Error("statement exec failed", zap.Error(err))
			continue
		}
	}

	if err := tx.Commit(); err != nil {
		m.logger.Error("sql commit failed", zap.Error(err))
		return err
	}

	return nil
}

// Get ...
func (m *MySQL) Get(topic string, offset []byte, count int) ([]*message.Message, error) {

	return nil, nil
}

// Close stop sql
func (m *MySQL) Close() error {
	if m.DB != nil {
		if err := m.DB.Close(); err != nil {
			m.logger.Error("mysql db close failed", zap.Error(err))
			return err
		}
	}
	return nil
}
