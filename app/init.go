package app

import (
	"github.com/atharvyadav96k/SPOTNEARR_SHARED/app/models/firestore"
	"github.com/atharvyadav96k/SPOTNEARR_SHARED/app/models/secrets"
)

func Init() App {
	return App{
		FireStore: &firestore.Firestore{},
	}
}

func (a *App) InitEnvironmentVariables() {
	a.Env = secrets.NewSecrets()
}

func (a *App) InitFirestore() error {
	var err error
	a.FireStore.Once.Do(func() {
		client, initErr := firestore.InitFirestore(a.Env.GCP_PROJECT_ID)
		if initErr != nil {
			err = initErr
			return
		}
		a.FireStore.FirestoreClient = client
	})
	return err
}
