package geonames_parse

//	admin1CodesASCII.txt and admin2Codes.txt
type AdminRec struct {
	Code      string
	Name      string
	NameAscii string
	Id        int64
}

//	countryInfo.txt
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

//	featureCodes_en.txt
type FeatureRec struct {
	Code string
	Name string
	Desc string
}

//	hierarchy.txt
type HierarchyRec struct {
	ParentId int64
	ChildId  int64
	Type     string
}

//	iso-languagecodes.txt
type LanguageRec struct {
	Iso_639_3 string
	Iso_639_2 string
	Iso_639_1 string
	Name      string
}

//	allCountries.txt and null.txt
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

//	zip_allCountries.txt
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

//	timeZones.txt
type TimezoneRec struct {
	CountryCode  string
	TimezoneName string
	OffsetGmt    float64
	OffsetDst    float64
	OffsetRaw    float64
}
