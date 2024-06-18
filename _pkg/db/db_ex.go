package pkg

import (
	"log"

	"gorm.io/gorm"
)

type DBEx struct{ DB *gorm.DB }

func SelectById[T interface{}](db *gorm.DB, orderId string, object *T) error {
	err := db.Model(new(T)).Where("id = ?", orderId).First(&object).Error
	if err != nil && err.Error() != "record not found" {
		return err
	}
	return nil
}

func SelectAllByField[T interface{}](db *gorm.DB, objects *[]T, orderId string, field string) error {
	if field == "" {
		field = "id"
	}
	err := db.Model(new(T)).Where(field+" = ?", orderId).Scan(&objects).Error
	if err != nil {
		return err
	}
	return nil
}

func Raw[T interface{}](db *gorm.DB, objects *[]T, raw string, variables ...interface{}) {
	// db.Raw("SELECT * FROM products WHERE price > ? ORDER BY price DESC LIMIT ? OFFSET ?", 100, 10, 20).Scan(&products)
	db.Raw(raw, variables...).Scan(&objects)
}

func Query[T interface{}](db *gorm.DB, object *[]T, strSelect string, where []map[string]string, orders []string, page Page) {
	resultOrm := db.Model(new(T))
	if strSelect != "" {
		resultOrm.Select(strSelect)
	}
	if len(where) > 0 {
		for k, v := range where {
			if err := resultOrm.Where(k, v).Error; err != nil {
				log.Println("failed to query ", where, err)
			}
		}
	}

	if len(orders) > 0 {
		for _, v := range orders {
			resultOrm.Order(v)
		}
	}

	if page.PageSize != 0 {
		resultOrm.Offset(int(page.PageNum))
		resultOrm.Limit(int(page.PageSize))
	}

	resultOrm.Find(&object)
}

func QueryAll[T interface{}](db *gorm.DB, object *[]T, strSelect string, where []map[string]string, strOrder []string) {
	Query(db, object, strSelect, where, strOrder, Page{})
}

func UpdatesWhere[T comparable](db *gorm.DB, o T, where map[string]string, vars ...[]string) *gorm.DB {
	// vars: col1,col2
	model := db.Model(&o)
	if len(where) > 0 {
		for k, v := range where {
			model = model.Where(k, v)
		}
	}
	if len(vars) > 0 {
		model = model.Select(vars[0])
	}

	r := model.Updates(o)
	if r.Error != nil {
		log.Println("DB UpdatesWhere", r.Error, o)
	}
	return r
}

func SaveBatch[T interface{}](db *gorm.DB, objs []T, vars ...string) *gorm.DB {
	r := db.Save(objs)
	if r.Error != nil {
		log.Println("DB Inserts", r.Error, objs)
	}
	return r
}
