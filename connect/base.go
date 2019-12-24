package connect

import (
	"gim/logic/db"
	"gim/public/imctx"
)

func Context() *imctx.Context {
	return imctx.NewContext(db.Factoty.GetSession())
}
