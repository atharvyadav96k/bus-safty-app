package app

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"cloud.google.com/go/firestore"
)

// return the query for the unique constrain
// for the struct record contains unique="true"
func GetUniqueFields(data interface{}) []firestore.PropertyFilter {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil
	}

	typ := val.Type()
	filters := make([]firestore.PropertyFilter, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)

		if field.Tag.Get("unique") != "true" {
			continue
		}

		fieldName := field.Tag.Get("firestore")

		switch {
		case fieldName == "" || fieldName == "-":
			fieldName = field.Name
		case strings.Contains(fieldName, ","):
			fieldName = strings.Split(fieldName, ",")[0]
		}

		value := val.Field(i)

		if value.IsZero() {
			continue
		}

		filters = append(filters, firestore.PropertyFilter{
			Path:     fieldName,
			Operator: "==",
			Value:    value.Interface(),
		})
	}

	return filters
}

// Make sure consistency in the database for the unique values should not be added duplicate in the
// database
// example:
//
//	type user struct{
//		email string `unique="true"`
//	}
//
// record with unique constrain are not stored again
func (a *App) CheckForDuplicate(ctx context.Context, collection string, data interface{}) error {
	filters := GetUniqueFields(data)
	if len(filters) == 0 {
		return nil
	}

	entityFilters := make([]firestore.EntityFilter, len(filters))
	for i, f := range filters {
		entityFilters[i] = f
	}

	query := a.StoreDoc(collection).WhereEntity(
		firestore.OrFilter{
			Filters: entityFilters,
		},
	)

	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return fmt.Errorf("failed to check uniqueness: %w", err)
	}

	if len(docs) == 0 {
		return nil
	}

	duplicates := make(map[string]interface{})
	for _, doc := range docs {
		dataMap := doc.Data()

		for _, f := range filters {
			if val, ok := dataMap[f.Path]; ok && val == f.Value {
				duplicates[f.Path] = f.Value
			}
		}
	}

	if len(duplicates) > 0 {
		var parts []string
		for field, value := range duplicates {
			parts = append(parts, fmt.Sprintf("%s=%v", field, value))
		}
		return fmt.Errorf("duplicate fields: %s", strings.Join(parts, ", "))
	}

	return nil
}

// StoreCreate inserts a new document into the given collection with an auto-generated ID.
//
// Parameters:
//   - ctx: request context for timeout/cancellation control
//   - collection: Firestore collection name
//   - data: struct or map to be stored
//
// Returns:
//   - error if the write operation fails
//
// Example:
//
//	err := app.StoreCreate(ctx, "users", userData)
func (a *App) StoreCreate(ctx context.Context, collection string, data interface{}) error {
	if err := a.CheckForDuplicate(ctx, collection, data); err != nil {
		return err
	}

	docRef := a.StoreDoc(collection).NewDoc()

	entity, ok := data.(StoreEntity)
	if !ok {
		return fmt.Errorf("data must implement Identifiable")
	}
	entity.SetID(docRef.ID)
	_, err := docRef.Create(ctx, data)
	return err
}

// StoreCreateWithId creates a new document with a specific ID in the given collection.
// It fails if a document with the same ID already exists.
//
// Parameters:
//   - ctx: request context
//   - collection: Firestore collection name
//   - id: custom document ID
//   - data: struct or map to be stored
//
// Returns:
//   - error if document already exists or write fails
//
// Example:
//
//	err := app.StoreCreateWithId(ctx, "users", "user123", userData)
func (a *App) StoreCreateWithId(ctx context.Context, collection string, id string, data interface{}) error {
	if err := a.CheckForDuplicate(ctx, collection, data); err != nil {
		return err
	}
	entity, ok := data.(StoreEntity)
	if !ok {
		return fmt.Errorf("data must implement Identifiable")
	}
	entity.SetID(id)
	docRef := a.StoreDoc(collection).Doc(id)
	_, err := docRef.Create(ctx, data)
	return err
}

// StoreUpdate updates or overwrites a document with the given ID.
// If the document does not exist, it will be created.
//
// Parameters:
//   - ctx: request context
//   - collection: Firestore collection name
//   - id: document ID
//   - data: updated data (replaces existing document)
//
// Returns:
//   - error if update fails
//
// Example:
//
//	err := app.StoreUpdate(ctx, "users", "user123", updatedData)
func (a *App) StoreUpdate(ctx context.Context, collection string, id string, data interface{}) error {
	docRef := a.StoreDoc(collection).Doc(id)
	_, err := docRef.Set(ctx, data)
	return err
}

// StoreDelete deletes a document from the given collection by ID.
//
// Parameters:
//   - ctx: request context
//   - collection: Firestore collection name
//   - id: document ID
//
// Returns:
//   - error if deletion fails
//
// Example:
//
//	err := app.StoreDelete(ctx, "users", "user123")
func (a *App) StoreDelete(ctx context.Context, collection string, id string) error {
	_, err := a.StoreDoc(collection).Doc(id).Delete(ctx)
	return err
}
