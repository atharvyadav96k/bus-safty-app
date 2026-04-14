package firestore

import (
	"sync"

	"cloud.google.com/go/firestore"
)

type Firestore struct {
	FirestoreClient *firestore.Client
	Once            sync.Once
}
