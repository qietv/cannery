package storage

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"sync"
	"time"
)

const (
	OffsetTable = "offset"
	ProjectTable = "projects"
	InitSQL   = `CREATE TABLE IF NOT EXISTS "` + OffsetTable + `"
		(
			[id] INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			[offset] INTEGER  NOT NULL,
			[pid] INTEGER
					constraint offset___pid
						references projects (id)
							on delete set null
		);

		CREATE TABLE IF NOT EXISTS "` + ProjectTable + `"
		(
			[id] INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			[name] NVARCHAR(32) not null,
			[path] NVARCHAR(255),
			[host] NVARCHAR(128),
			[status] NVARCHAR(128)
		);
		create unique index projects_name_uindex
			on projects (name);
		`
)

type sqlite struct {
	filename string
	conn     *sql.Conn
	locker   *sync.Mutex
}

func (s *sqlite) Initialize() (err error) {
	var (
		db     *sql.DB
	)
	s.locker = &sync.Mutex{}
	db, err = sql.Open("sqlite3", fmt.Sprintf("file:%s?cache=shared", s.filename))
	if err != nil {
		return
	}
	//db exists?
	s.conn, err = db.Conn(context.Background())
	if err != nil {
		return err
	}
	_, err = s.conn.ExecContext(context.Background(), InitSQL)
	if err != nil {
		return
	}
	err = s.watch(db)
	return
}

func (s *sqlite) AddProject(ctx context.Context,host, project, path string) (err error) {
	var (
		tx *sql.Tx
		result sql.Result
		pid int64
	)
	tx, err = s.conn.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.IsolationLevel(sql.LevelRepeatableRead),
	})
	result, err = tx.Exec("INSERT INTO projects (name, path, host, status ) VALUES (?, ?, ?, ?)", project, path, host, "PR")
	if err != nil {
		tx.Rollback()
		return
	}
	pid, _ = result.LastInsertId()
	_, err = tx.Exec("INSERT INTO offset (pid, offset) VALUES (?, 0)", pid)
	if err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit()
}

func (s *sqlite) RmProject(ctx context.Context, project string) (err error) {
	_, err = s.conn.ExecContext(ctx, "DELETE FROM offset WHERE project = ? ", project)
	return
}

func (s *sqlite) Read(ctx context.Context, project string) (offset int, err error) {
	var (
		result *sql.Row
	)
	result = s.conn.QueryRowContext(ctx, "SELECT offset FROM offset WHERE  project = ?", project)
	if result == nil {
		return 0, UnknownProject
	}
	err = result.Scan(&offset)
	if err == sql.ErrNoRows {
		err = UnknownProject
	}
	return
}

func (s *sqlite) UpdateRead(ctx context.Context, project string, offset int, replace bool) (err error) {
	var (
		result   sql.Result
		affected int64
	)
	if replace {
		result, err = s.conn.ExecContext(ctx, "UPDATE offset SET offset = ? WHERE  project = ?", offset, project)
	}else {
		result, err = s.conn.ExecContext(ctx, "UPDATE offset SET offset = offset + ? WHERE  project = ?", offset, project)
	}
	if affected, err = result.RowsAffected(); affected == 0 || err != nil {
		err = UnknownProject
	}
	return
}

func (s *sqlite) watch(db *sql.DB) (err error) {
	var (
		retryTimes int
		ticker *time.Ticker
	)
	go func(s *sqlite, db *sql.DB) {
		ticker = time.NewTicker(time.Second * 5)
		for {
			select {
			case <- ticker.C:
				s.locker.Lock()
				pError := db.Ping()
				if pError != nil {
					s.conn, err = db.Conn(context.Background())
					if err != nil {
						retryTimes++
						if retryTimes > 3 {
							panic(fmt.Sprintf("sqlite fail, err:%s", err.Error()))
						}
					}
				}
				s.locker.Unlock()
			}
		}

	}(s, db)
	return
}
