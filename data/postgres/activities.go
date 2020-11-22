package postgres

import (
	"database/sql"
	"github.com/stewie1520/pm/data/model"
	"net"
)

type ActivityModel struct {
	DB *sql.DB
}

func (a *ActivityModel) Insert(action string, ip net.IP) (int, error) {
	stmt, _ := a.DB.Prepare(`
		INSERT
		INTO activities
			(action, time, ip)
		VALUES
			(?, (now() at time zone 'utc'), ?)
	`)
	result, err :=stmt.Exec(action, ip)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (a *ActivityModel) Latest() *model.Activity {
	stmt, err := a.DB.Prepare(`
		SELECT action, time, ip
		FROM activities
		ORDER BY time DESC 
		FETCH FIRST ROW ONLY
	`)
	if err != nil {
		return nil
	}

	actv := &model.Activity{}
	var ip string

	row := stmt.QueryRow()
	err = row.Scan(&actv.Action, &actv.Time, &ip)
	if err != nil {
		return nil
	}

	actv.IP = net.ParseIP(ip)
	return actv
}
