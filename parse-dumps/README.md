# geonames_parse
--
    import "github.com/go-geo/geonames/parse-dumps"

Iterates over records in raw `download.geonames.org/export/dump` files (fetched
via `fetch-dumps` package).

## Usage

#### type AdminRec

```go
type AdminRec struct {
	Code      string
	Name      string
	NameAscii string
	Id        int64
}
```

admin1CodesASCII.txt and admin2Codes.txt

#### type CountryRec

```go
type CountryRec struct {
	Code struct {
		Iso2   string
		Iso3   string
		IsoNum string
		Fips   string
	}
	Name       string
	Capital    string
	AreaSqKm   int64
	Population int64
	Continent  string
	Tld        string
	Currency   struct {
		Code string
		Name string
	}
	CallingCode string
	PostalCode  struct {
		Format string
		Regex  string
	}
	Languages []string
	Id        int64
	Neighbors []string
}
```

countryInfo.txt

#### type FeatureRec

```go
type FeatureRec struct {
	Code string
	Name string
	Desc string
}
```

featureCodes_en.txt

#### type HierarchyRec

```go
type HierarchyRec struct {
	ParentId int64
	ChildId  int64
	Type     string
}
```

hierarchy.txt

#### type Iterator

```go
type Iterator struct {
	//	Directory containing raw `download.geonames.org/export/dump` files
	DirPath string

	FileNames struct {
		Admin1, Admin2, Countries, Features, Hierarchy, Languages, Places, Postal, Timezones string
	}
}
```

Provides file parsing and record iteration

#### func  NewIterator

```go
func NewIterator(dirPath string) (me *Iterator)
```
Initializes `me.DirPath` and all `me.FileNames`.

#### func (*Iterator) Admin1

```go
func (me *Iterator) Admin1(onRec func(index int, rec *AdminRec)) (err error)
```
Calls `onRec` for each `AdminRec` found in `me.FileNames.Admin1`.

The `rec` pointer always points to the exact same `struct`, but its field values
differ for each `onRec` invocation.

#### func (*Iterator) Admin2

```go
func (me *Iterator) Admin2(onRec func(index int, rec *AdminRec)) (err error)
```
Calls `onRec` for each `AdminRec` found in `me.FileNames.Admin2`.

The `rec` pointer always points to the exact same `struct`, but its field values
differ for each `onRec` invocation.

#### func (*Iterator) AdminAll

```go
func (me *Iterator) AdminAll(onRec func(index int, rec *AdminRec)) (err error)
```
Calls `onRec` for each `AdminRec` found in both `me.FileNames.Admin1` and
`me.FileNames.Admin2`.

The `rec` pointer always points to the exact same `struct`, but its field values
differ for each `onRec` invocation.

#### func (*Iterator) Countries

```go
func (me *Iterator) Countries(onRec func(index int, rec *CountryRec)) (err error)
```
Calls `onRec` for each `CountryRec` found in `me.FileNames.Countries`.

The `rec` pointer always points to the exact same `struct`, but its field values
differ for each `onRec` invocation.

#### func (*Iterator) Features

```go
func (me *Iterator) Features(onRec func(index int, rec *FeatureRec)) (err error)
```
Calls `onRec` for each `FeatureRec` found in `me.FileNames.Features`.

The `rec` pointer always points to the exact same `struct`, but its field values
differ for each `onRec` invocation.

#### func (*Iterator) Hierarchy

```go
func (me *Iterator) Hierarchy(onRec func(index int, rec *HierarchyRec)) (err error)
```
Calls `onRec` for each `HierarchyRec` found in `me.FileNames.Hierarchy`.

The `rec` pointer always points to the exact same `struct`, but its field values
differ for each `onRec` invocation.

#### func (*Iterator) Languages

```go
func (me *Iterator) Languages(onRec func(index int, rec *LanguageRec)) (err error)
```
Calls `onRec` for each `LanguageRec` found in `me.FileNames.Languages`.

The `rec` pointer always points to the exact same `struct`, but its field values
differ for each `onRec` invocation.

#### func (*Iterator) Places

```go
func (me *Iterator) Places(onRec func(index int, rec *PlaceRec)) (err error)
```
Calls `onRec` for each `PlaceRec` found in `me.FileNames.Places`.

The `rec` pointer always points to the exact same `struct`, but its field values
differ for each `onRec` invocation.

#### func (*Iterator) PostalCodes

```go
func (me *Iterator) PostalCodes(onRec func(index int, rec *PostalRec)) (err error)
```
Calls `onRec` for each `PostalRec` found in `me.FileNames.Postal`.

The `rec` pointer always points to the exact same `struct`, but its field values
differ for each `onRec` invocation.

#### func (*Iterator) Timezones

```go
func (me *Iterator) Timezones(onRec func(index int, rec *TimezoneRec)) (err error)
```
Calls `onRec` for each `TimezoneRec` found in `me.FileNames.Timezones`.

The `rec` pointer always points to the exact same `struct`, but its field values
differ for each `onRec` invocation.

#### type LanguageRec

```go
type LanguageRec struct {
	Iso_639_3 string
	Iso_639_2 string
	Iso_639_1 string
	Name      string
}
```

iso-languagecodes.txt

#### type PlaceRec

```go
type PlaceRec struct {
	Id        int64
	Name      string
	NameAscii string
	NamesAlt  []string
	LonLat    []float64
	Feature   struct {
		Class string
		Code  string
	}
	Country struct {
		Code     string
		CodesAlt []string
	}
	Admin struct {
		Code1, Code2, Code3, Code4 string
	}
	Population   int64
	Elevation    int64
	TimezoneName string
}
```

allCountries.txt and null.txt

#### type PostalRec

```go
type PostalRec struct {
	CountryCode string
	PostalCode  string
	PlaceName   string
	Admin       struct {
		Name1, Code1, Name2, Code2, Name3, Code3 string
	}
	LonLat   []float64
	Accuracy int64
}
```

zip_allCountries.txt

#### type TimezoneRec

```go
type TimezoneRec struct {
	CountryCode  string
	TimezoneName string
	OffsetGmt    float64
	OffsetDst    float64
	OffsetRaw    float64
}
```

timeZones.txt

--
**godocdown** http://github.com/robertkrimen/godocdown
