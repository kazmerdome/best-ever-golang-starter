package mongodb

import (
	"github.com/kazmerdome/best-ever-golang-starter/internal/util/filter"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUuidFilter(filter *filter.UuidFilter) bson.M {
	uuidFilter := bson.M{}
	if filter != nil {
		if filter.Eq != nil {
			uuidFilter["$eq"] = *filter.Eq
		}
		if len(filter.In) > 0 {
			uuidFilter["$in"] = filter.In
		}
		if filter.Ne != nil {
			uuidFilter["$ne"] = *filter.Ne
		}
	}
	return uuidFilter
}

func GetIntFilter(filter *filter.IntFilter) bson.M {
	intFilter := bson.M{}
	if filter != nil {
		if filter.Eq != nil {
			intFilter["$eq"] = *filter.Eq
		}
		if filter.Gt != nil {
			intFilter["$gt"] = *filter.Gt
		}
		if filter.Gte != nil {
			intFilter["$gte"] = *filter.Gte
		}
		if filter.Lt != nil {
			intFilter["$lt"] = *filter.Lt
		}
		if filter.Lte != nil {
			intFilter["$lte"] = *filter.Lte
		}
	}
	return intFilter
}

func GetFloat64Filter(filter *filter.Float64Filter) bson.M {
	intFilter := bson.M{}
	if filter != nil {
		if filter.Eq != nil {
			intFilter["$eq"] = *filter.Eq
		}
		if filter.Gt != nil {
			intFilter["$gt"] = *filter.Gt
		}
		if filter.Gte != nil {
			intFilter["$gte"] = *filter.Gte
		}
		if filter.Lt != nil {
			intFilter["$lt"] = *filter.Lt
		}
		if filter.Lte != nil {
			intFilter["$lte"] = *filter.Lte
		}
	}
	return intFilter
}

func GetRegexFilter(value string) bson.M {
	regex := bson.M{"$regex": value, "$options": "i"}
	return regex
}

func GetStringFilter(filter *filter.StringFilter) bson.M {
	stringFilter := bson.M{}
	if filter != nil {
		if filter.Eq != nil {
			stringFilter["$eq"] = *filter.Eq
		}
		if filter.Regex != nil {
			stringFilter["$regex"] = *filter.Regex
			stringFilter["$options"] = "i"
		}
	}
	return stringFilter
}

func GetDateFilter(filter *filter.DateFilter) bson.M {
	dateFilter := bson.M{}
	if filter.Gt != nil && filter.Lt != nil {
		dateFilter["$gt"] = *filter.Gt
		dateFilter["$lt"] = *filter.Lt
	} else if filter.Gt != nil && filter.Lte != nil {
		dateFilter["$gt"] = *filter.Gt
		dateFilter["$lte"] = *filter.Lte
	} else if filter.Gte != nil && filter.Lt != nil {
		dateFilter["$gte"] = *filter.Gte
		dateFilter["$lt"] = *filter.Lt
	} else if filter.Gte != nil && filter.Lte != nil {
		dateFilter["$gte"] = *filter.Gte
		dateFilter["$lte"] = *filter.Lte
	} else {
		if filter.Eq != nil {
			dateFilter["$eq"] = *filter.Eq
		}
		if filter.Gt != nil {
			dateFilter["$gt"] = *filter.Gt
		}
		if filter.Gte != nil {
			dateFilter["$gte"] = *filter.Gte
		}
		if filter.Lt != nil {
			dateFilter["$lt"] = *filter.Lt
		}
		if filter.Lte != nil {
			dateFilter["$lte"] = *filter.Lte
		}
	}
	return dateFilter
}

func GetSortFilter(sort *filter.SortFilter) bson.D {
	sortBy := bson.D{}
	if sort != nil {
		if sort.SortBy != "" {
			order := 1
			if sort.SortOrder != nil && *sort.SortOrder == filter.SortOrderDesc {
				order = -1
			}
			sortBy = bson.D{{Key: sort.SortBy, Value: order}}
		}
	}
	return sortBy
}
