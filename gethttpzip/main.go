package main

// 从http端读取zip文件并解析得到压缩包里的内容
import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://192.168.0.119:8007/static/map/map.pbf")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		log.Fatal(err)
	}

	// Read all the files from zip archive
	for _, zipFile := range zipReader.File {
		fmt.Println("Reading file:", zipFile.Name)
		unzippedFileBytes, err := readZipFile(zipFile)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(string(unzippedFileBytes))
		_ = unzippedFileBytes // this is unzipped file bytes

	}
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
