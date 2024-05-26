package session

import (
	"os"
	"time"

	"github.com/AdluAghnia/xhater/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/mysql"
)

var store *session.Store

func createStorage() *mysql.Storage {
	return mysql.New(mysql.Config{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: os.Getenv("DBUSER"),
		Password: os.Getenv("DBPASS"),
		Database: "xhaters",
	},
	)
}

func InitSessionStore() {
	store = session.New(session.Config{
		Storage:    createStorage(),
		Expiration: 1 * time.Hour,
		KeyLookup:  "cookie:myapp_session",
	})
}

func CreateUserSession(c *fiber.Ctx, uid string) error {
	// Get Or Create Storage
	s, _ := store.Get(c)

	// if this new session
	if s.Fresh() {
		sid := s.ID()

		// get uid
		// uid := c.Params("uid")

		// save session data
		s.Set("uid", uid)
		s.Set("sid", sid)
		s.Set("ip", c.Context().RemoteIP().String())
		s.Set("login", time.Unix(time.Now().Unix(), 0).UTC().String())
		s.Set("ua", string(c.Request().Header.UserAgent()))

		err := s.Save()
		if err != nil {
			return err
		}

		// Save user reference
		stmt, err := db.Db.Prepare("UPDATE sessions SET u = ? WHERE k = ?")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(uid, sid)
		if err != nil {
			return err
		}
	}

	return nil
}
