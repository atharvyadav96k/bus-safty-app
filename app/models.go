package app

import (
	"github.com/atharvyadav96k/SPOTNEARR_SHARED/app/models/firestore"
	"github.com/atharvyadav96k/SPOTNEARR_SHARED/app/models/secrets"
)

type App struct {
	Env       secrets.Env
	FireStore *firestore.Firestore
}
