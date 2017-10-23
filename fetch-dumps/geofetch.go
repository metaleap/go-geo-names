//	Fetches files from `download.geonames.org/export/dump` --- see http://download.geonames.org/export/dump/readme.txt for details.
package geonames_fetch

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/metaleap/go-util/fs"
	"github.com/metaleap/go-util/net"
)

var (
	//	Default base URL to be combined with relative file URLs in `GeoFiles`.
	BaseUrl = "http://download.geonames.org/export/"

	//	Maps relative file URLs to local destination file names.
	GeoFiles = map[string]string{
		"admin1CodesASCII.txt": "dump/admin1CodesASCII.txt",
		"admin2Codes.txt":      "dump/admin2Codes.txt",
		"allCountries.txt":     "dump/allCountries.zip",
		// "alternateNames.txt":    "dump/alternateNames.zip",
		"countryInfo.txt":     "dump/countryInfo.txt",
		"featureCodes_en.txt": "dump/featureCodes_en.txt",
		"hierarchy.txt":       "dump/hierarchy.zip",
		// "iso-languagecodes.txt": "dump/iso-languagecodes.txt",
		"timeZones.txt":        "dump/timeZones.txt",
		"zip_allCountries.txt": "zip/allCountries.zip",
	}
)

//	Fetches all files in `GeoFiles` in parallel using `FetchFile`.
func FetchAllFiles(outDir string) (errs []error) {
	var wait sync.WaitGroup
	goFetch := func(fileName, relUrl string) {
		defer wait.Done()
		if err := FetchFile(outDir, fileName, relUrl); err != nil {
			errs = append(errs, err)
		}
	}
	for fileName, relUrl := range GeoFiles {
		wait.Add(1)
		go goFetch(fileName, relUrl)
	}
	wait.Wait()
	return
}

//	Downloads the file at `BaseUrl + relUrl` to `outDir + fileName`.
//
//	If it is a ZIP archive file, it is extracted in place and deleted.
func FetchFile(outDir, fileName, relUrl string) (err error) {
	fullUrl := "http://download.geonames.org/export/" + relUrl
	filePath := filepath.Join(outDir, strings.Replace(relUrl, "/", "_", -1))
	err = unet.DownloadFile(fullUrl, filePath)
	if err == nil {
		if strings.HasSuffix(filePath, ".zip") {
			log.Printf("UNZIP: %s from %s\n", fileName, filePath)
			err = ufs.ExtractZipFile(filePath, outDir, true, "zip_", fileName)
		} else {
			err = os.Rename(filePath, filepath.Join(outDir, fileName))
		}
	}
	return
}
