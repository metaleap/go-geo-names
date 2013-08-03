package geonames_makedb

import (
	"fmt"
	"strings"

	"github.com/go-forks/mgo/bson"
	"github.com/go-geo/geonames/parse-dumps"
	"github.com/go-utils/udb/umgo"
	"github.com/go-utils/uslice"
	"github.com/go-utils/ustr"
)

var (
	mAdmins    = map[string]int64{}
	mCountries = map[string]int{}
	mFeatures  = map[string]int{}
	mTimezones = map[string]int{}
)

var (
	//	Administrative divisions
	CollAdminsName            = "admins"
	CollAdminsCap             = 40000
	CollAdminsField_Country   = "c"
	CollAdminsField_Code      = "d"
	CollAdminsField_Name      = "n"
	CollAdminsField_NameAscii = "" // as of now, all utf-8 names are 'western-readable' (no exotic scripts), so ascii would be redundant
)

func onAdmin(i int, r *geonames_parse.AdminRec) {
	if concat := ustr.Split(r.Code, "."); len(concat) > 1 {
		mAdmins[r.Code] = r.Id
		m := umgo.Sparse(bson.M{
			"_id": r.Id,
			CollAdminsField_Country:   mCountries[concat[0]],
			CollAdminsField_Code:      strings.Join(concat[1:], "."),
			CollAdminsField_Name:      r.Name,
			CollAdminsField_NameAscii: r.NameAscii,
		})
		recs = append(recs, m)
	}
}

var (
	//	Countries
	CollCountriesName                    = "countries"
	CollCountriesCap                     = 300
	CollCountriesField_Name              = "n"
	CollCountriesField_GeoId             = "i"
	CollCountriesField_AreaSqKm          = "q"
	CollCountriesField_PhoneCode         = "a"
	CollCountriesField_Capital           = "m"
	CollCountriesField_CodeFips          = "f"
	CollCountriesField_CodeIso2          = "i2"
	CollCountriesField_CodeIso3          = "i3"
	CollCountriesField_CodeIsoNum        = "i1"
	CollCountriesField_Tld               = "t"
	CollCountriesField_Continent         = "w"
	CollCountriesField_CurrencyCode      = "c"
	CollCountriesField_CurrencyName      = "v"
	CollCountriesField_Languages         = "l"
	CollCountriesField_NeighborCountries = "s"
	CollCountriesField_Population        = "p"
	CollCountriesField_PostalFormat      = "z"
	CollCountriesField_PostalRegex       = "r"
)

func onCountry(i int, r *geonames_parse.CountryRec) {
	mCountries[r.Code.Iso2] = i + 1
	recs = append(recs, umgo.Sparse(bson.M{
		"_id": i + 1, CollCountriesField_Name: r.Name, CollCountriesField_GeoId: r.Id,
		CollCountriesField_AreaSqKm: r.AreaSqKm, CollCountriesField_PhoneCode: r.CallingCode,
		CollCountriesField_Capital: r.Capital, CollCountriesField_CodeFips: r.Code.Fips,
		CollCountriesField_CodeIso2: r.Code.Iso2, CollCountriesField_CodeIso3: r.Code.Iso3,
		CollCountriesField_CodeIsoNum: r.Code.IsoNum, CollCountriesField_Tld: r.Tld,
		CollCountriesField_Continent: r.Continent, CollCountriesField_CurrencyCode: r.Currency.Code,
		CollCountriesField_CurrencyName: r.Currency.Name, CollCountriesField_Languages: r.Languages,
		CollCountriesField_NeighborCountries: r.Neighbors, CollCountriesField_Population: r.Population,
		CollCountriesField_PostalFormat: r.PostalCode.Format, CollCountriesField_PostalRegex: r.PostalCode.Regex,
	}))
}

var (
	//	Features classes & codes
	CollFeaturesName       = "features"
	CollFeaturesCap        = 700
	CollFeaturesField_Name = "n"
	CollFeaturesField_Code = "c"
	CollFeaturesField_Desc = "d"
)

func onFeature(i int, r *geonames_parse.FeatureRec) {
	mFeatures[r.Code] = i + 1
	recs = append(recs, umgo.Sparse(bson.M{
		"_id": i + 1, CollFeaturesField_Name: r.Name, CollFeaturesField_Code: r.Code, CollFeaturesField_Desc: r.Desc,
	}))
}

var (
	//	The actual "geo-names"
	CollPlacesName             = "places"
	CollPlacesCap              = 8520000
	CollPlacesField_Country    = "c"
	CollPlacesField_Elevation  = "e"
	CollPlacesField_LonLat     = "l"
	CollPlacesField_Name       = "n"
	CollPlacesField_NameAscii  = "a"
	CollPlacesField_NamesAlt   = "m"
	CollPlacesField_Population = "p"
	CollPlacesField_Timezone   = "t"
	CollPlacesField_Feature    = "f"
	CollPlacesField_Admin12    = "d"
)

func onPlace(_ int, r *geonames_parse.PlaceRec) {
	if r.Name, r.NameAscii = placeName(r.Name), placeName(r.NameAscii); len(r.Name) == 0 {
		r.Name = r.NameAscii
	}
	if r.Name == r.NameAscii {
		r.NameAscii = ""
	}
	if r.NamesAlt = uslice.StrEach(r.NamesAlt, placeName); len(r.Name) == 0 {
		r.Name = ustr.FirstNonEmpty(r.NamesAlt...)
	}
	r.NamesAlt = uslice.StrWithout(r.NamesAlt, true, r.Name, r.NameAscii)

	if len(r.Name) == 0 {
		return
	}

	m := bson.M{
		"_id": r.Id, CollPlacesField_Country: mCountries[r.Country.Code], CollPlacesField_Elevation: r.Elevation,
		CollPlacesField_LonLat: r.LonLat, CollPlacesField_Name: r.Name,
		CollPlacesField_NameAscii: r.NameAscii, CollPlacesField_NamesAlt: r.NamesAlt,
		CollPlacesField_Population: r.Population, CollPlacesField_Timezone: mTimezones[r.TimezoneName],
		CollPlacesField_Feature: mFeatures[fmt.Sprintf("%s.%s", r.Feature.Class, r.Feature.Code)],
	}
	ad := mAdmins[fmt.Sprintf("%s.%s.%s", r.Country.Code, r.Admin.Code1, r.Admin.Code2)]
	if ad == 0 {
		ad = mAdmins[fmt.Sprintf("%s.%s", r.Country.Code, r.Admin.Code1)] // more-general only if more-specific wasnt found
	}
	m[CollPlacesField_Admin12] = ad
	recs = append(recs, umgo.Sparse(m))
}

var (
	//	Postal codes
	CollPostalsName             = "zips"
	CollPostalsCap              = 900000
	CollPostalsField_PlaceName  = "n"
	CollPostalsField_PostalCode = "z"
	CollPostalsField_Country    = "c"
	CollPostalsField_Accuracy   = "a"
	CollPostalsField_LonLat     = "l"
	CollPostalsField_Admins     = "d"
)

func onPostal(i int, r *geonames_parse.PostalRec) {
	m := umgo.Sparse(bson.M{
		"_id": i + 1, CollPostalsField_PlaceName: title(r.PlaceName), CollPostalsField_PostalCode: r.PostalCode,
		CollPostalsField_Country: mCountries[r.CountryCode], CollPostalsField_Accuracy: r.Accuracy,
		CollPostalsField_LonLat: r.LonLat,
	})
	ad := map[string]string{r.Admin.Code1: r.Admin.Name1, r.Admin.Code2: r.Admin.Name2, r.Admin.Code3: r.Admin.Name3}
	for k, v := range ad {
		if len(k) == 0 || len(v) == 0 {
			delete(ad, k)
		} else {
			ad[k] = title(v)
		}
	}
	m[CollPostalsField_Admins] = ad
	recs = append(recs, m)
}

var (
	//	Timezones
	CollTimezonesName            = "timezones"
	CollTimezonesCap             = 420
	CollTimezonesField_Name      = "n"
	CollTimezonesField_OffsetGmt = "g"
	CollTimezonesField_OffsetDst = "d"
	CollTimezonesField_OffsetRaw = "r"
)

func onTimezone(i int, r *geonames_parse.TimezoneRec) {
	mTimezones[r.TimezoneName] = i + 1
	recs = append(recs, umgo.Sparse(bson.M{
		"_id": i + 1, CollTimezonesField_Name: strings.Replace(r.TimezoneName, "_", " ", -1),
		CollTimezonesField_OffsetGmt: r.OffsetGmt, CollTimezonesField_OffsetDst: r.OffsetDst, CollTimezonesField_OffsetRaw: r.OffsetRaw,
	}))
}
