# geonames_makedb
--
    import "github.com/go-geo/geonames/make-mongodb"


## Usage

```go
var (
	//	How many records are at most passed per `mgo.Collection.Insert` call at once
	BatchSize = 125000

	//	Whether to `log.Printf` progress
	Log = true

	//	All-upper-case strings (where applicable) longer than this value are `Title`d (`FOO BAR` becomes `Foo Bar`).
	//	Set to `0` to disable this.
	TitleAllUpper = 1
)
```

```go
var (
	//	Administrative divisions
	CollAdminsName            = "admins"
	CollAdminsCap             = 40000
	CollAdminsField_Country   = "c"
	CollAdminsField_Code      = "d"
	CollAdminsField_Name      = "n"
	CollAdminsField_NameAscii = "" // as of now, all utf-8 names are 'western-readable' (no exotic scripts), so ascii would be redundant
)
```

```go
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
```

```go
var (
	//	Features classes & codes
	CollFeaturesName       = "features"
	CollFeaturesCap        = 700
	CollFeaturesField_Name = "n"
	CollFeaturesField_Code = "c"
	CollFeaturesField_Desc = "d"
)
```

```go
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
```

```go
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
```

```go
var (
	//	Timezones
	CollTimezonesName            = "timezones"
	CollTimezonesCap             = 420
	CollTimezonesField_Name      = "n"
	CollTimezonesField_OffsetGmt = "g"
	CollTimezonesField_OffsetDst = "d"
	CollTimezonesField_OffsetRaw = "r"
)
```

#### func  Insert

```go
func Insert(geo *geonames_parse.Iterator, db *mgo.Database) (err error)
```
Inserts all records from `geo` into the specified `db`, in the following order:
Time zones, features, countries, administrative divisions, postal codes, places
(geo-names).

--
**godocdown** http://github.com/robertkrimen/godocdown
