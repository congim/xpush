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
		if _, err := stmt.Exec(msg.Topic, msgIDs[index], msg.ID, msg.Type, msg.Payload, time.Now().Unix()); err != nil {
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
func (m *MySQL) Get(topic string, offset int, checkTime int64) ([]*message.Message, error) {
	//"SELECT id, original_id, type, payload, insert_time FROM msgs WHERE topic=? AND insert_time <=? ORDER BY insert_time DESC LIMIT ?,?;"
	rows, err := m.DB.Query(SQLPullMsg, topic, checkTime, offset, 50)
	if err != nil {
		m.logger.Error("Query Context", zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	msgs := make([]*message.Message, 0)

	for rows.Next() {
		var id string
		var insertTime string
		msg := message.New()
		msg.Topic = topic
		if err := rows.Scan(&id, &msg.ID, &msg.Type, &msg.Payload, &insertTime); err != nil {
			m.logger.Error("get msg failed", zap.String("sql", SQLPullMsg),
				zap.String("topic", topic), zap.Int("offset", offset), zap.Int64("checkTime", checkTime),
				zap.Error(err))
			return nil, err
		}
		msgs = append(msgs, msg)
	}
	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	rerr := rows.Close()
	if rerr != nil {
		m.logger.Error("rows close failed", zap.String("sql", SQLPullMsg),
			zap.String("topic", topic), zap.Int("offset", offset), zap.Int64("checkTime", checkTime),
			zap.Error(err))
		return nil, err
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		m.logger.Error("rows err failed", zap.String("sql", SQLPullMsg),
			zap.String("topic", topic), zap.Int("offset", offset), zap.Int64("checkTime", checkTime),
			zap.Error(err))
		return nil, err
	}
	return msgs, nil
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
