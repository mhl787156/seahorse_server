package db

import (
	"time"
)

const purgePeriod time.Duration = time.Hour //Run every hour

func (c *DbConn) purgeExpired() {
	// now := time.Now().UTC().Unix()
	// query := SessionTable.Filter(gorethink.Row.Field("expires_at").Lt(now))
	// _, err := query.Delete().RunWrite(c.Session)
	// if err != nil {
	// 	panic("Database cleanup failed...")
	// }
}

func (c *DbConn) CullSessions() {
	// for {
	// 	// c.purgeExpired()
	// 	// time.Sleep(purgePeriod)
	// }
}
