package utils

import (
	"reflect"
	"strings"
)

func FilterInputByModel(input map[string]interface{}, item interface{}) map[string]interface{} {
	filtered := make(map[string]interface{})
	validFields := make(map[string]bool)

	// Gunakan reflection untuk baca struct
	t := reflect.TypeOf(item)
	if t.Kind() == reflect.Ptr {
		t = t.Elem() // Dereference pointer
	}

	// Loop melalui semua field di struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue // Skip field tanpa tag json atau yang diignore
		}

		// Handle tag json dengan opsi (e.g., `json:"name,omitempty"`)
		jsonFieldName := strings.Split(jsonTag, ",")[0]
		validFields[jsonFieldName] = true

		// Tambahkan juga nama field (snake_case) untuk kompatibilitas
		gormTag := field.Tag.Get("gorm")
		if gormTag != "" {
			// Contoh: `gorm:"column:subcategory_id"` → ambil "subcategory_id"
			if strings.HasPrefix(gormTag, "column:") {
				columnName := strings.Split(gormTag[len("column:"):], ";")[0]
				validFields[columnName] = true
			}
		}
	}

	// Filter input
	for key, value := range input {
		if validFields[key] {
			filtered[key] = value
		}
	}

	return filtered
}
