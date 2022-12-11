package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"time"

	"contrib.go.opencensus.io/integrations/ocsql"
	"github.com/Pranc1ngPegasus/trial-field/domain/configuration"
	"github.com/Pranc1ngPegasus/trial-field/domain/logger"
	domain "github.com/Pranc1ngPegasus/trial-field/domain/persistence"
	"github.com/google/wire"
	_ "github.com/lib/pq"
)

var _ domain.DB = (*DB)(nil)

var NewDBSet = wire.NewSet(
	wire.Bind(new(domain.DB), new(*DB)),
	NewDB,
)

type DB struct {
	db *sql.DB
}

func NewDB(
	ctx context.Context,
	logger logger.Logger,
	config configuration.Configuration,
) (*DB, error) {
	cfg := config.Config()

	logger.Info(ctx, "Start RDB connector")

	driver, err := ocsql.Register("postgres", ocsql.WithAllTraceOptions())
	if err != nil {
		return nil, fmt.Errorf("failed to register db driver: %w", err)
	}

	dsn := buildDSN(cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseName)

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open RDB connection: %w", err)
	}

	db.SetMaxIdleConns(cfg.DatabaseMaxIdle)
	db.SetMaxOpenConns(cfg.DatabaseMaxOpen)
	db.SetConnMaxLifetime(cfg.DatabaseMaxLifetime)

	if cfg.Debug {
		go func() {
			for {
				stats := db.Stats()
				logger.Info(ctx, "db stats",
					logger.Field("max open", stats.MaxOpenConnections),
					logger.Field("open", stats.OpenConnections),
					logger.Field("idle", stats.Idle),
					logger.Field("inuse", stats.InUse),
					logger.Field("wait", stats.WaitCount),
				)

				time.Sleep(cfg.DebugLogFrequency)
			}
		}()
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return &DB{
		db: db,
	}, nil
}

func buildDSN(host string, port int, user string, password string, dbName string) string {
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, password),
		Host:   net.JoinHostPort(host, strconv.Itoa(port)),
		Path:   dbName,
		RawQuery: url.Values{
			"sslmode": []string{"disable"},
		}.Encode(),
	}

	return dsn.String()
}

func (p *DB) Conn() *sql.DB {
	return p.db
}
