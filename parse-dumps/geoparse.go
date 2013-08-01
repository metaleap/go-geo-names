//	Iterates over records in raw `download.geonames.org/export/dump` files (fetched via `fetch-dumps` package).
package geonames_parse

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/go-geo/ugeo"
	"github.com/go-utils/ufs"
	"github.com/go-utils/uslice"
	"github.com/go-utils/ustr"
)

func checkLonLat(lonLat []float64) []float64 {
	if len(lonLat) == 2 && lonLat[0] >= ugeo.LonMin && lonLat[0] < ugeo.LonMax && lonLat[1] >= ugeo.LatMin && lonLat[1] < ugeo.LatMax {
		return lonLat
	}
	return nil
}

func replaceDoubleSpaces(s string) string {
	for strings.Index(s, "  ") >= 0 {
		s = strings.Replace(s, "  ", " ", -1)
	}
	return s
}

//	Provides file parsing and record iteration
type Iterator struct {
	//	Directory containing raw `download.geonames.org/export/dump` files
	DirPath string

	FileNames struct {
		Admin1, Admin2, Countries, Features, Hierarchy, Languages, Places, Postal, Timezones string
	}
}

//	Initializes `me.DirPath` and all `me.FileNames`.
func NewIterator(dirPath string) (me *Iterator) {
	me = &Iterator{DirPath: dirPath}
	fn := &me.FileNames
	fn.Admin1 = "admin1CodesASCII.txt"
	fn.Admin2 = "admin2Codes.txt"
	fn.Countries = "countryInfo.txt"
	fn.Features = "featureCodes_en.txt"
	fn.Hierarchy = "hierarchy.txt"
	fn.Languages = "iso-languagecodes.txt"
	fn.Places = "allCountries.txt"
	fn.Postal = "zip_allCountries.txt"
	fn.Timezones = "timeZones.txt"
	return
}

func (me *Iterator) iterate(fileName string, skipFirst bool, i int, onRec func(int, []string)) (int, error) {
	file, err := os.Open(filepath.Join(me.DirPath, fileName))
	if file != nil {
		defer file.Close()
		if err == nil {
			err = ufs.ReadLines(file, skipFirst, func(ln string) {
				if !strings.HasPrefix(ln, "#") {
					onRec(i, uslice.StrEach(ustr.Split(ln, "\t"), strings.TrimSpace, replaceDoubleSpaces))
					i++
				}
			})
		}
	}
	return i, err
}

func (me *Iterator) admin(fileName string, i int, onRec func(int, *AdminRec)) (int, error) {
	var r AdminRec
	return me.iterate(fileName, false, i, func(index int, rec []string) {
		r.Code = rec[0]
		r.Name = rec[1]
		r.NameAscii = rec[2]
		r.Id = ustr.ParseInt(rec[3])

		if r.Name == r.NameAscii {
			r.NameAscii = ""
		}
		if len(r.Name) == 0 {
			r.Name, r.NameAscii = r.NameAscii, ""
		}
		onRec(index, &r)
	})
}

//	Calls `onRec` for each `AdminRec` found in `me.FileNames.Admin1`.
//
//	The `rec` pointer always points to the exact same `struct`, but its field values differ for each `onRec` invocation.
func (me *Iterator) Admin1(onRec func(index int, rec *AdminRec)) (err error) {
	_, err = me.admin(me.FileNames.Admin1, 0, onRec)
	return
}

//	Calls `onRec` for each `AdminRec` found in `me.FileNames.Admin2`.
//
//	The `rec` pointer always points to the exact same `struct`, but its field values differ for each `onRec` invocation.
func (me *Iterator) Admin2(onRec func(index int, rec *AdminRec)) (err error) {
	_, err = me.admin(me.FileNames.Admin2, 0, onRec)
	return
}

//	Calls `onRec` for each `AdminRec` found in both `me.FileNames.Admin1` and `me.FileNames.Admin2`.
//
//	The `rec` pointer always points to the exact same `struct`, but its field values differ for each `onRec` invocation.
func (me *Iterator) AdminAll(onRec func(index int, rec *AdminRec)) (err error) {
	var i int
	if i, err = me.admin(me.FileNames.Admin1, i, onRec); err == nil {
		_, err = me.admin(me.FileNames.Admin2, i, onRec)
	}
	return
}

//	Calls `onRec` for each `CountryRec` found in `me.FileNames.Countries`.
//
//	The `rec` pointer always points to the exact same `struct`, but its field values differ for each `onRec` invocation.
func (me *Iterator) Countries(onRec func(index int, rec *CountryRec)) (err error) {
	var r CountryRec
	_, err = me.iterate(me.FileNames.Countries, false, 0, func(index int, rec []string) {
		r.Code.Iso2 = rec[0]
		r.Code.Iso3 = rec[1]
		r.Code.IsoNum = rec[2]
		r.Code.Fips = rec[3]
		r.Name = rec[4]
		r.Capital = rec[5]
		r.AreaSqKm = ustr.ParseInt(rec[6])
		r.Population = ustr.ParseInt(rec[7])
		r.Continent = rec[8]
		r.Tld = rec[9]
		r.Currency.Code = rec[10]
		r.Currency.Name = rec[11]
		r.CallingCode = rec[12]
		r.PostalCode.Format = rec[13]
		r.PostalCode.Regex = rec[14]
		r.Languages = uslice.StrEach(ustr.Split(rec[15], ","), strings.TrimSpace)
		r.Id = ustr.ParseInt(rec[16])
		r.Neighbors = uslice.StrEach(ustr.Split(rec[17], ","), strings.TrimSpace)
		onRec(index, &r)
	})
	return
}

//	Calls `onRec` for each `FeatureRec` found in `me.FileNames.Features`.
//
//	The `rec` pointer always points to the exact same `struct`, but its field values differ for each `onRec` invocation.
func (me *Iterator) Features(onRec func(index int, rec *FeatureRec)) (err error) {
	var r FeatureRec
	_, err = me.iterate(me.FileNames.Features, false, 0, func(index int, rec []string) {
		r.Code = rec[0]
		r.Name = rec[1]
		r.Desc = rec[2]
		onRec(index, &r)
	})
	return
}

//	Calls `onRec` for each `HierarchyRec` found in `me.FileNames.Hierarchy`.
//
//	The `rec` pointer always points to the exact same `struct`, but its field values differ for each `onRec` invocation.
func (me *Iterator) Hierarchy(onRec func(index int, rec *HierarchyRec)) (err error) {
	var r HierarchyRec
	_, err = me.iterate(me.FileNames.Hierarchy, false, 0, func(index int, rec []string) {
		r.ParentId = ustr.ParseInt(rec[0])
		r.ChildId = ustr.ParseInt(rec[1])
		r.Type = rec[2]
		onRec(index, &r)
	})
	return
}

//	Calls `onRec` for each `LanguageRec` found in `me.FileNames.Languages`.
//
//	The `rec` pointer always points to the exact same `struct`, but its field values differ for each `onRec` invocation.
func (me *Iterator) Languages(onRec func(index int, rec *LanguageRec)) (err error) {
	var r LanguageRec
	_, err = me.iterate(me.FileNames.Languages, true, 0, func(index int, rec []string) {
		r.Iso_639_3 = rec[0]
		r.Iso_639_2 = rec[1]
		r.Iso_639_1 = rec[2]
		r.Name = rec[3]
		onRec(index, &r)
	})
	return
}

//	Calls `onRec` for each `PlaceRec` found in `me.FileNames.Places`.
//
//	The `rec` pointer always points to the exact same `struct`, but its field values differ for each `onRec` invocation.
func (me *Iterator) Places(onRec func(index int, rec *PlaceRec)) (err error) {
	var r PlaceRec
	_, err = me.iterate(me.FileNames.Places, false, 0, func(index int, rec []string) {
		r.Id = ustr.ParseInt(rec[0])
		r.Name = rec[1]
		r.NameAscii = rec[2]
		r.NamesAlt = uslice.StrEach(ustr.Split(rec[3], ","), strings.TrimSpace)
		r.LonLat = checkLonLat(ustr.ParseFloats(rec[4], rec[5]))
		r.Feature.Class = rec[6]
		r.Feature.Code = rec[7]
		r.Country.Code = rec[8]
		r.Country.CodesAlt = uslice.StrEach(ustr.Split(rec[9], ","), strings.TrimSpace)
		r.Admin.Code1 = rec[10]
		r.Admin.Code2 = rec[11]
		r.Admin.Code3 = rec[12]
		r.Admin.Code4 = rec[13]
		r.Population = ustr.ParseInt(rec[14])
		r.Elevation = ustr.ParseInt(rec[len(rec)-4])
		if el2 := ustr.ParseInt(rec[len(rec)-3]); el2 != 0 && r.Elevation == 0 {
			r.Elevation = el2
		}
		r.TimezoneName = rec[len(rec)-2]

		r.NamesAlt = uslice.StrWithout(r.NamesAlt, true, r.Name, r.NameAscii)
		if r.Name == r.NameAscii {
			r.NameAscii = ""
		}
		if len(r.Name) == 0 {
			r.Name, r.NameAscii = r.NameAscii, ""
		}

		onRec(index, &r)
	})
	return
}

//	Calls `onRec` for each `PostalRec` found in `me.FileNames.Postal`.
//
//	The `rec` pointer always points to the exact same `struct`, but its field values differ for each `onRec` invocation.
func (me *Iterator) PostalCodes(onRec func(index int, rec *PostalRec)) (err error) {
	var r PostalRec
	_, err = me.iterate(me.FileNames.Postal, false, 0, func(index int, rec []string) {
		r.CountryCode = rec[0]
		r.PostalCode = rec[1]
		r.PlaceName = rec[2]
		r.Admin.Name1 = rec[3]
		r.Admin.Code1 = rec[4]
		r.Admin.Name2 = rec[5]
		r.Admin.Code2 = rec[6]
		r.Admin.Name3 = rec[7]
		r.Admin.Code3 = rec[8]
		r.LonLat = checkLonLat(ustr.ParseFloats(rec[len(rec)-3], rec[len(rec)-2]))
		r.Accuracy = ustr.ParseInt(rec[len(rec)-1])
		onRec(index, &r)
	})
	return
}

//	Calls `onRec` for each `TimezoneRec` found in `me.FileNames.Timezones`.
//
//	The `rec` pointer always points to the exact same `struct`, but its field values differ for each `onRec` invocation.
func (me *Iterator) Timezones(onRec func(index int, rec *TimezoneRec)) (err error) {
	var r TimezoneRec
	_, err = me.iterate(me.FileNames.Timezones, true, 0, func(index int, rec []string) {
		r.CountryCode = rec[0]
		r.TimezoneName = rec[1]
		r.OffsetGmt = ustr.ParseFloat(rec[2])
		r.OffsetDst = ustr.ParseFloat(rec[3])
		r.OffsetRaw = ustr.ParseFloat(rec[4])
		onRec(index, &r)
	})
	return
}
