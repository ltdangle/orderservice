package helpers

import "database/sql"

// Helper functions to convert sql.Null* types.
func NullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	// Return some default value or an empty string if it's null
	return ""
}
