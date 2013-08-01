# geonames_fetch
--
    import "github.com/go-geo/geonames/fetch-dumps"

Fetches files from `download.geonames.org/export/dump` --- see
http://download.geonames.org/export/dump/readme.txt for details.

## Usage

```go
var (
	//	Default base URL to be combined with relative file URLs in `GeoFiles`.
	BaseUrl = "http://download.geonames.org/export/"

	//	Maps relative file URLs to local destination file names.
	GeoFiles = map[string]string{
		"admin1CodesASCII.txt":  "dump/admin1CodesASCII.txt",
		"admin2Codes.txt":       "dump/admin2Codes.txt",
		"allCountries.txt":      "dump/allCountries.zip",
		"alternateNames.txt":    "dump/alternateNames.zip",
		"countryInfo.txt":       "dump/countryInfo.txt",
		"featureCodes_en.txt":   "dump/featureCodes_en.txt",
		"hierarchy.txt":         "dump/hierarchy.zip",
		"iso-languagecodes.txt": "dump/iso-languagecodes.txt",
		"timeZones.txt":         "dump/timeZones.txt",
		"zip_allCountries.txt":  "zip/allCountries.zip",
	}
)
```

#### func  FetchAllFiles

```go
func FetchAllFiles(outDir string) (errs []error)
```
Fetches all files in `GeoFiles` in parallel using `FetchFile`.

#### func  FetchFile

```go
func FetchFile(outDir, fileName, relUrl string) (err error)
```
Downloads the file at `BaseUrl + relUrl` to `outDir + fileName`.

If it is a ZIP archive file, it is extracted in place and deleted.

--
**godocdown** http://github.com/robertkrimen/godocdown
