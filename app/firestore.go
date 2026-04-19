package app

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"cloud.google.com/go/firestore"
)

// 1. Change return type to []firestore.BooleanExpression
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
	filters := GetUniqueFields(data)
	if len(filters) > 0 {
		entityFilters := make([]firestore.EntityFilter, 0, len(filters))
		for _, f := range filters {
			entityFilters = append(entityFilters, f)
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
		if len(docs) > 0 {
			return fmt.Errorf("duplicate value found in one of the unique fields")
		}
	}
	_, _, err := a.StoreDoc(collection).Add(ctx, data)
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
