package geonames_makedb

import (
	"log"
	"strings"

	"github.com/go-forks/mgo"
	"github.com/go-forks/mgo/bson"
	"github.com/go-geo/geonames/parse-dumps"
	"github.com/go-utils/ustr"
)

var (
	//	How many records are at most passed per `mgo.Collection.Insert` call at once
	BatchSize = 125000

	//	Whether to `log.Printf` progress
	Log = true

	//	All-upper-case strings (where applicable) longer than this value are `Title`d (`FOO BAR` becomes `Foo Bar`).
	//	Set to `0` to disable this.
	TitleAllUpper = 1

	recs []interface{}
)

//	Inserts all records from `geo` into the specified `db`, in the following order:
//	Time zones, features, countries, administrative divisions, postal codes, places (geo-names).
func Insert(geo *geonames_parse.Iterator, db *mgo.Database) (err error) {
	recs = make([]interface{}, 0, CollTimezonesCap)
	if err = insert(geo, db, CollTimezonesName, CollFeaturesCap, geo.Timezones(onTimezone)); err == nil {
		if err = insert(geo, db, CollFeaturesName, CollCountriesCap, geo.Features(onFeature)); err == nil {
			if err = insert(geo, db, CollCountriesName, CollAdminsCap, geo.Countries(onCountry)); err == nil {
				if err = insert(geo, db, CollAdminsName, CollPostalsCap, geo.AdminAll(onAdmin)); err == nil {
					if err = insert(geo, db, CollPostalsName, CollPlacesCap, geo.PostalCodes(onPostal)); err == nil {
						if err = insert(geo, db, CollPlacesName, 0, geo.Places(onPlace)); err == nil {
						}
					}
				}
			}
		}
	}
	return
}

func insert(geo *geonames_parse.Iterator, db *mgo.Database, collName string, nextRecsCap int, err error) error {
	switch collName {
	case CollCountriesName:
		prepCountries()
	}
	if err == nil && len(recs) > 0 {
		if Log {
			log.Printf("Insert %v %#v..", len(recs), collName)
		}
		var max int
		for i := 0; i < len(recs); i += BatchSize {
			if max = i + BatchSize; max >= len(recs) {
				max = len(recs)
			}
			if Log && i > 0 {
				log.Printf("\t%v done..", i)
			}
			if err = db.C(collName).Insert(recs[i:max]...); err != nil {
				break
			}
		}
		if Log && err == nil {
			log.Print("\tall done.")
		}
		if nextRecsCap > 0 {
			if Log {
				log.Printf("Loading <%v records..", nextRecsCap)
			}
			recs = make([]interface{}, 0, nextRecsCap)
		}
	}
	return err
}

func prepCountries() {
	var (
		m    bson.M
		cn   []string
		ok   bool
		cnc  string
		refs []int
	)
	for _, r := range recs {
		m = r.(bson.M)
		if cn, ok = m[CollCountriesField_NeighborCountries].([]string); ok {
			refs = make([]int, 0, len(cn))
			for _, cnc = range cn {
				refs = append(refs, mCountries[cnc])
			}
			m[CollCountriesField_NeighborCountries] = refs
		}
	}
}

func title(str string) string {
	if TitleAllUpper > 0 && len(str) > TitleAllUpper {
		if ustr.IsUpper(str) {
			str = strings.ToLower(str)
		}
		str = strings.Title(str)
	}
	return str
}
