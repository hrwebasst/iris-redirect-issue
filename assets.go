// Code generated by go-bindata.
// sources:
// assets/static/css/pure-min.css
// assets/static/css/style.css
// assets/static/javascript/alert.js
// assets/templates/dude.html
// assets/templates/index.html
// assets/templates/start.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsStaticCssPureMinCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x3b\x6b\x8f\xe4\x36\x72\xdf\xf7\x57\xc8\xbb\x58\xdc\xcc\x40\xd2\x48\xea\xe7\x48\x58\x23\x3e\x7b\xef\xe2\xe0\xce\x08\x60\x7f\x48\xb0\x37\x01\xd8\x52\x75\x37\x33\x92\x28\x90\xec\x79\x6c\x5f\xff\xf7\x80\xa4\x1e\x24\x45\x75\xcf\xd8\x71\xe0\x0f\x71\xc3\x3b\x22\xab\x58\x2c\xd6\x8b\x2c\x3e\x6e\x6f\xbe\x79\xf7\xef\x07\x0a\xde\x63\x14\x2e\xc3\xe8\xdd\xf7\xa4\x79\xa1\x78\xb7\xe7\x5e\x12\xc5\x73\xef\x3f\xd1\x9e\x90\x6f\xbc\x1f\xeb\x3c\xf4\xbe\x2b\x4b\x4f\x82\x98\x47\x81\x01\x7d\x84\x22\x7c\xf7\x37\x9c\x43\xcd\xa0\xf0\x0e\x75\x01\xd4\xe3\x7b\xf0\xfe\xfc\xf3\x0f\x5e\x5b\x1d\xbe\xdb\x73\xde\xb0\xf4\xf6\x76\x87\xf9\xfe\xb0\x09\x73\x52\xdd\xbe\x08\x9a\xb7\xcd\x81\xc2\xed\xa6\x24\x9b\xdb\x0a\x31\x0e\xf4\xf6\x6f\x3f\x7e\xff\xf9\xa7\x9f\x3f\x87\x55\xf1\xee\xe6\xf6\x9d\xe0\xab\x26\xb4\x42\x25\xfe\x0a\x61\xce\x98\xf7\xf8\x5f\xb3\x30\xf2\xfe\xe9\xfd\xfd\xc7\x5f\x3a\xf2\xde\x3f\xbd\x1d\xe6\x21\x26\xb7\x3d\xaa\x36\x80\xab\xfc\xda\xfb\x09\xe7\xa4\x44\xcc\xfb\x2b\x2a\x4b\xb4\xdb\x03\xf5\x50\x5d\x78\xff\x46\x6a\xc4\xf7\xa8\xf6\x7e\x02\x54\xb6\xbd\x79\x56\x6f\xb3\x30\x0a\x93\x8b\xdd\x79\x37\xb7\x7b\x5e\x95\xc7\x2d\xa9\x79\xb0\x45\x15\x2e\x5f\x52\x86\x6a\x16\x30\xa0\x78\x9b\x05\x15\x0b\x38\x3c\xf3\x80\xe1\xaf\x10\xa0\xe2\xbf\x0f\x8c\xa7\x71\x14\x7d\xcc\x82\x27\xd8\x3c\x60\xee\x86\x9e\x36\xa4\x78\x39\x56\x88\xee\x70\x9d\x46\x27\x44\x39\xce\x4b\xf0\x11\xc3\x05\xf8\x05\x70\x84\x4b\xe6\x6f\xf1\x2e\x47\x0d\xc7\xa4\x16\x9f\x07\x0a\xfe\x96\x10\x0e\xd4\xdf\x03\x2a\xc4\x9f\x1d\x25\x87\xc6\xaf\x10\xae\xfd\x0a\xea\x83\x5f\xa3\x47\x9f\x41\x2e\x5b\xb0\x43\x55\x21\xfa\x72\x2c\x30\x6b\x4a\xf4\x92\x6e\x4a\x92\x3f\x9c\xd0\xa1\xc0\xc4\xcf\x51\xfd\x88\x98\xdf\x50\xb2\xa3\xc0\x98\xff\x88\x0b\x20\x3d\x26\xae\x4b\x5c\x43\x20\x1b\x64\x8f\x20\x58\x43\x65\x80\x4a\xbc\xab\xd3\x0d\x62\x20\xa0\x8a\x50\x5a\x13\x7e\xf5\x25\x27\x35\xa7\xa4\x64\xf7\xd7\x3d\x89\x9a\xd4\x90\xed\x41\x28\x29\x8d\x4e\x5f\xf6\xb8\x28\xa0\xbe\xf7\x39\x54\x4d\x89\x38\x18\x78\x27\x74\xdc\xa0\xfc\x41\x8c\xa5\x2e\x82\x9c\x94\x84\xa6\x9c\xa2\x9a\x35\x88\x42\xcd\x4f\x28\x45\x39\xc7\x8f\xe0\xa3\x74\x4f\x1e\x81\x1e\xc9\x81\x0b\x16\x84\xd8\x36\x1b\xfa\x85\x63\x5e\xc2\xfd\x71\x43\x68\x01\x34\xd8\x10\xce\x49\x95\xc6\xcd\xb3\x57\x10\xce\xa1\x38\x6d\x7c\xc6\x29\xa9\x77\x4a\x83\x4f\x8a\xa9\x55\x14\x9d\x8a\x6d\xad\xea\x18\x7f\x29\x21\xc5\x1c\x95\x38\x3f\xed\xe3\xb6\x12\x7f\x85\x34\x81\x2a\x6b\x95\x14\x2e\x57\x50\x79\xd1\xa9\x42\xf4\x41\x63\x38\xfd\xb0\xdd\x46\x99\xe2\xfa\x43\x14\x45\x27\x56\xa1\xb2\xd4\x48\xac\xa3\x8f\x27\x76\xd8\xf8\xec\xd0\x68\xb5\xab\xc5\xc7\x4c\x4a\xb9\x13\x52\xd6\x10\x86\x85\xe2\x52\x0a\x25\x12\xe3\x9d\x14\xbd\xa0\xc4\x49\x93\x06\xe1\x02\x2a\x41\xfb\xd8\x0e\x3a\x08\x13\x51\x83\xab\x5d\x2b\x8d\x34\x3a\xb1\xc7\x9d\xd4\x52\x4a\x09\xe1\xd7\x47\x21\xc0\x6d\x49\x9e\x52\xa5\x92\x93\xb2\xab\xce\x10\x63\xa8\xbc\x79\xd4\x3c\x9f\xf6\xf4\x18\x54\xe4\x6b\xb0\x21\xcf\x82\x5f\x5c\xef\x52\xa1\x65\xa8\xb9\xa8\xca\x26\xaa\x7b\x85\x37\x14\x86\x9e\xd0\x81\x93\x53\x4e\x0a\xf0\x1f\x36\x85\xdf\x50\xf0\x19\xaa\x1a\xc3\x9f\x2a\x52\x13\xd6\xa0\x1c\xfc\xfe\x2b\x1b\x64\x15\x43\x75\xda\x1c\x38\x27\xb5\x8f\xeb\xe6\xc0\x7d\xd2\x70\x65\xf9\x0c\x4a\xc8\xb9\x2f\x3c\x0c\x51\x40\x47\xa5\x06\x5c\xef\x81\x62\x2e\x29\xf4\x85\xde\xd5\x14\xa5\x81\xbd\x47\xcc\xf0\xa6\x84\xae\x07\x45\xf2\x28\x9d\x56\x5a\xe1\x96\xd0\x4a\xd9\x69\x8b\x21\xa2\x81\x27\x19\xf9\xc2\x5f\x1a\xf8\xa4\xaa\xef\x7d\xad\x4a\x04\x4f\x6e\xd4\xb0\xc3\xa6\xc2\xfc\xfe\xd8\x05\x05\xd4\x34\x80\x28\xaa\x73\x48\x55\xfb\x2c\x3f\x50\x46\x68\xda\x10\x5c\x73\xa0\x6d\x67\x5f\x0a\xcc\xd0\xa6\x84\xe2\x5e\xef\xb6\xaf\x3c\xb6\x8d\x0a\xd8\xa2\x43\xc9\xdb\x46\x69\x2a\x75\xb7\x25\xf9\x81\x05\xb8\xae\x81\x2a\x4e\xc6\xf5\xbd\x99\x64\x0d\x2a\x0a\xa1\xce\xe8\x24\x51\x8f\xba\x6d\xaa\x50\x78\xd2\x46\x93\xef\x21\x7f\xd8\x90\x67\x73\xd0\xa8\xc0\x44\xf8\x61\x6f\x1b\xbd\x4b\x3e\xdb\xf4\x55\x8b\xfa\x50\x6d\x80\xde\xa7\x69\x27\x15\xc9\x54\xc0\x1a\x5c\x07\xba\xc2\x27\xb0\xc9\x81\x9b\xd8\xc7\x96\x61\x69\x71\xba\xf0\x01\xd1\x7c\xef\x14\xbe\xd0\xf3\x16\x43\x59\x64\xe7\xec\xbd\x6b\xf8\x26\x77\x70\x70\x30\xf0\xae\x2a\x82\x5c\x30\x51\x3a\x06\x3b\xd5\xa0\x80\x9c\x50\x24\xe2\x84\x6b\x34\xd2\x4c\xe5\x70\x18\xf0\x4e\xb9\x22\x14\x32\x52\xe2\xc2\x63\xb8\x7c\x04\xda\xbb\x82\x97\x34\x83\x62\xc2\xd9\x02\x2a\x2f\x5c\x26\xf2\xcf\x4a\xc4\x91\x12\x76\x50\x17\x2e\x1b\xe9\x1d\xce\x74\xf2\xce\x2f\x47\x91\x96\x0b\x73\xed\x22\x74\x4e\xca\x12\x35\x0c\xd2\xee\x23\x6b\x01\xc2\xef\x5b\xfa\x85\xcf\xf7\xc7\xa1\xbf\x50\x45\x2b\xbf\x9b\x48\x8c\xf9\xe3\x1b\x5c\x35\x84\x72\x54\xf3\x53\x28\x96\x1c\x81\x88\x7f\x15\x7a\x0e\x9e\x70\xc1\xf7\x6a\x2e\xd6\x0c\x23\x33\x27\x44\xd5\x64\x77\x2c\x81\x73\x8d\x87\x20\x9c\xc5\x50\x65\x37\x56\xb5\x72\x85\xec\xe6\x89\xd0\x42\xc3\x9d\xcf\xa0\xca\x64\xc8\xa0\x20\x16\x49\xa2\x96\x34\x1c\x57\xf8\x2b\xb0\x06\xa0\xc8\xf4\x68\xf7\x17\x0a\xf0\x33\xaa\x99\xff\x1d\xc5\x15\xf1\xdf\xff\x40\x09\x2e\x3c\x51\xf3\xde\xff\x57\x28\x1f\x41\x44\x7d\x01\x44\xa5\xaf\xad\x34\x3a\xbe\x3b\xbd\x6f\x4b\x18\x2c\x53\x14\x02\xa9\x09\x4a\x9e\xbc\x27\x8a\x9a\x01\xbf\x62\x12\x2c\x0d\xb9\xfd\xb6\x50\x45\xb5\x9c\x64\x82\xd6\x7e\x53\x89\xc4\x38\xa2\xbc\xef\x62\x12\x61\x0a\x70\x0a\x49\x03\x14\x05\xa4\x2e\x5f\xbc\x34\x20\x41\x43\x41\x06\x1f\xbf\x13\xba\x43\x8c\xad\x42\x0e\xee\xd5\xc8\x8d\x59\x9b\x7d\x25\x62\xae\xcf\xdc\x5a\x32\xa8\xb7\x75\xd6\x9c\xca\x49\x63\xeb\x4d\x5a\x72\xcb\xa0\xf7\x25\x2f\x11\x63\xde\xcd\xa7\xf7\x8a\xab\xf7\xf7\x13\xeb\xc0\x8e\xed\x20\xf6\xfb\x2f\xe3\x3b\xd1\xbe\x67\xfd\x77\xa2\x7d\xc7\xc1\xbc\xff\x9e\x69\xdf\x71\xb0\xd0\xf0\x17\x1a\xce\xf0\x3d\xd7\xbe\x17\xda\x77\x1c\x2c\xb5\xfa\xa5\x56\xbf\xd6\xe8\xac\x35\x9c\xe1\x7b\xa5\x7d\xc7\x41\x9c\x68\x48\x5a\x61\xa5\x17\x62\x03\x2f\x0e\x92\xb9\xc6\x79\xa2\x0f\x4f\x2b\xcc\xf5\xc2\x42\x2f\x2c\xf5\xc2\x4a\x2f\xac\xf5\xc2\x9d\x5e\x88\x23\xa3\x64\xf0\x10\x1b\x4c\xc4\x06\x17\xb1\xc1\x46\x6c\xf0\x11\x1b\x8c\xc4\x06\x27\xb1\xc1\x4a\x6c\xf0\x92\x18\xbc\x24\xa6\x3c\x0c\x5e\x12\x83\x97\x44\xf0\xf2\x47\x70\x00\xa5\xc3\xa3\x0a\xa4\xf3\x30\x5e\x2e\x57\x1f\xb3\x9b\xbe\x38\x5b\xac\x3e\x9e\x9c\x26\x92\x0c\xcd\xd6\xe1\x6c\x36\x9b\xf5\xcd\xd6\xe1\x2c\x4a\x66\x7a\xb3\xb5\x61\x16\x6d\xab\x38\x09\x17\x7d\x9b\x38\x09\xe7\xcb\x3b\xbd\xcd\xd2\xb0\x9e\xae\xcd\x32\x5c\xea\x1c\x8a\xb2\xc5\xe2\xa2\x45\x4d\xa2\x01\xeb\x2e\xbc\xd3\x89\x2f\x06\x82\x49\x14\xae\x75\xde\x45\xd9\x62\xde\x34\xd6\xae\xdd\xc0\x7a\x32\x37\xa9\xaf\x34\xac\x3b\x53\xa0\xa2\x6c\xb1\x3b\x33\x0c\xbe\x6d\x37\x9b\x99\x12\x15\x65\x83\x2b\xdd\xa1\xef\xb4\x76\x2b\x4d\xa4\xb3\x95\x29\xd2\xa4\x17\xcd\x7c\x10\xcd\x6c\x24\x1a\xdd\xbb\x23\xcd\x34\x62\x53\xf2\xa2\x6c\x0e\x45\xb7\xa3\x85\x29\x55\x51\xb6\xa4\x9a\x98\x2e\xdb\x36\x5c\x0c\x9c\xcd\x2d\xce\x62\xcd\x74\x16\x96\xa1\x2e\x6c\x4b\x35\x43\x96\x66\x40\x0b\xcb\x56\x17\xb6\xb1\xce\x7a\x29\x2d\x07\x5e\x16\x23\x29\xad\xcd\x38\xd2\x35\xd1\x4d\x7a\x99\xd8\xf2\x9f\x99\xe1\xa6\x6b\x64\xd9\xf4\x72\x64\xd3\x9a\x45\xad\x2c\x7b\x5d\xd9\xf6\x6a\x4c\x2b\x9a\x49\xad\x06\xc6\x56\x96\xc1\xc6\x9a\x05\xad\x2c\x8b\x5d\xd9\x16\x3b\xef\xe5\xb3\x8e\x74\x2c\x4b\x3e\x4b\x33\x42\x76\x4d\x2c\xbb\x5e\xdb\x76\xad\x4f\x48\x89\x66\x50\x6b\xdd\xb0\xd7\x96\x61\x9b\x53\x52\xa2\x99\xd3\x9d\x65\xb4\x77\xb6\xd1\x26\x9a\x51\xdd\x59\x46\x7b\x37\x32\x5a\xe7\xbc\xaf\xcf\xc5\x89\x1e\xaa\xa2\xa8\x6b\xda\xa6\x2e\xce\x68\xdf\x45\xf7\x51\x32\x96\x3d\xed\x31\x07\x19\xdb\xc5\xaa\x5f\x2e\xe2\xac\xc8\x5e\xe1\xa2\x28\x41\x05\x77\x55\x93\x83\xc8\x2c\xad\x44\xb3\x5f\xde\x1d\x18\xd0\xa0\xa0\x68\xa7\x36\x6f\x8c\x6a\x95\x14\xb7\x00\x91\x27\x39\x6a\xd9\xb8\x72\x8c\x35\xce\xa3\xb4\x1c\xd1\xce\xc0\x34\x90\xb3\xd6\x10\xa0\x23\xb9\xed\xf3\x87\xac\xdf\x0d\x31\x44\xae\xaf\xe4\xf4\x3d\x83\x76\xd7\x41\xe4\x0d\x7d\x76\x24\xb2\x22\x91\x10\xb4\x1b\x3d\xf3\xf9\xbc\xfd\xa4\xbb\x0d\xba\x8a\x7c\xf1\x0b\xd7\xd7\xd9\x28\xe7\xfa\x70\x77\x77\xd7\xf7\xef\x69\xd8\xd1\x75\x36\xda\xf6\xfa\xf0\x79\x29\x7e\x4a\x69\x43\xaa\xa7\x84\xd7\x0e\x5c\xa4\xd9\x07\x96\x26\x8d\x39\xfe\x40\xee\x8d\xf9\x86\x48\x1c\x55\x52\x3e\xc7\x2d\x2e\x39\xd0\xb4\xa1\x64\x87\x8b\xf4\x87\xff\xf8\xb1\x42\x3b\xf8\xa5\xdb\xed\x08\xff\x8e\x73\x4a\x18\xd9\xf2\x70\x27\x7a\x83\x9a\x5f\xc9\x05\xfd\xf7\x82\x49\xc6\xe9\xa7\x3f\x7d\x88\xda\xff\xfe\xe4\x7b\x50\x17\x1a\x20\x46\x3d\xe0\xaf\x6d\xe3\x5f\x44\x42\x6b\x0e\x17\x8b\x0e\xfb\x4c\xa6\xef\x45\x18\x3a\xa2\x7e\xe4\x45\x7e\xe4\x09\x05\xf8\x5b\x4a\xaa\x2b\x6d\x37\xf0\xda\x97\x92\x0a\x18\x27\xcd\xd5\x3c\xfa\xe8\xeb\xf2\x8f\x16\xd7\xd7\x3e\x27\x57\x7a\x5d\x7c\x7d\x7d\xa6\x67\xd5\xe1\xc0\x80\xd6\x93\x4d\xd9\xb3\x7b\x8b\x9d\x84\x85\x11\x8e\xa8\x92\x66\x44\x2d\x7a\x05\x2d\xf2\xbf\xcf\xdf\x6f\x24\x78\x72\x18\xd3\xb0\x1d\x6b\xd8\x63\xbb\x69\x6b\x34\x50\x75\x6a\x8b\x68\x8f\x0a\xf2\x94\x46\x9e\xf8\x09\x77\x31\x7a\x5a\x5c\x7b\xb8\x66\xc0\x85\x29\x78\x4b\x0b\x9a\xb4\xc0\x6c\xd8\x50\x68\x37\x5f\xff\x71\x67\xf0\xa0\xed\x9e\x19\xac\x75\xd5\xee\x5a\x87\xdb\x0c\x30\x3d\x7f\xb5\x61\xfd\xe8\xda\xfd\x92\x91\xf0\xa5\x1b\xbf\xdd\xf7\xa0\x96\xf4\x3f\x6d\x51\xc9\xe0\xba\x23\x80\xca\x66\x8f\xae\x88\x58\xdd\xf3\x97\x4f\xf3\xe8\x3a\x0b\x1e\xf6\xbc\x2a\x83\xb6\x2a\x0d\xe7\x2a\xa2\x6a\x65\xed\xb3\x9d\x05\x6a\x22\x66\x87\x92\x3c\x41\x91\x69\x6a\x91\x1b\x49\x66\x78\x91\x3b\x2f\xe6\xc6\xfd\x6f\x8b\xbf\x41\x43\x71\x85\xe8\x8b\x29\x4d\x35\x5d\x40\xe1\x23\x27\x2e\x72\x22\x8f\x0f\x10\x3e\x44\xd1\x6a\x0d\xab\x2e\x5e\x6f\xb7\x5d\x86\x2e\x84\xac\xef\xde\x8a\x50\xdb\x99\x87\x0d\x6b\x10\x63\x22\x87\x9a\x82\x43\x85\x70\x39\x05\x3c\xd0\x49\x50\x81\x38\x4c\xc1\x2a\x52\xf3\xfd\x14\x90\xe3\x6a\xb2\xa1\x20\xfa\x1a\x78\x50\x92\x1c\x4d\xb2\xf6\x04\xf0\x30\x05\x6b\x37\x60\x27\xa0\xed\x8e\xe5\x14\xe7\x30\xd9\xa5\x54\x91\x01\x6c\xb7\xf5\xb5\x9a\x7e\xc3\xd1\x98\x90\xc3\x25\x54\x99\x73\xe1\x34\x9e\x86\xf3\x3c\xd7\x0d\x5c\x06\x90\x36\xf2\xcc\x9a\x67\xef\x43\x51\x14\xd6\x0c\x3b\x6f\x9e\x27\xd6\x54\xbf\xcb\x32\x66\x10\x8b\x3a\x55\x13\xb2\xb9\xbf\xfe\xbf\x1e\xf0\xef\x3c\x34\x5d\xe3\xc3\xd0\x12\x31\xb4\x45\xbf\xf7\xe7\x74\x51\x23\xf2\x4e\x3a\xea\x59\x2c\xe5\xae\x67\x51\x84\xd3\x9e\x45\x90\xae\x7b\x16\x43\x39\xf0\x59\x14\xe9\xa6\x17\xbb\x79\x3d\x56\xeb\xd2\x67\x71\xa5\x63\x9f\xc5\xe8\xce\x57\xce\xe1\x74\xc7\x12\x67\x47\x07\x17\x58\x51\xea\x1f\xa3\xb4\x99\xc2\xa8\xbe\x73\x7e\x7b\xbd\x61\xcd\xff\x71\x72\xf7\x97\xcf\xdf\x9d\xf5\xa6\x5f\x47\x41\x71\xbd\xc5\xe5\x05\x75\xa8\x03\xb0\xf3\x43\xef\x0e\xce\x2c\x4e\xf8\x1e\xd7\x9d\xe3\x2a\x2e\xb2\x0e\x24\x1c\x16\x1d\x38\xf1\x1c\xec\xa9\xcf\x8e\xa8\x3f\x82\x48\x8e\xba\xc3\x5d\x19\x41\x22\xe7\xf1\x87\xd3\xe3\x46\xab\xa7\x49\xaf\xbb\x88\xa9\x3c\xef\x22\x9a\xf0\xbe\x8b\x48\xd2\x03\x2f\x62\x29\x2f\xbc\x88\x26\x7d\xec\x55\x5d\xbe\x0d\xb3\xf5\xc8\x8b\xf8\xd2\x2b\x2f\x62\xb5\x9e\x79\x11\xaf\xf5\xce\xcb\xa3\x86\x57\xb0\xa6\xbc\xd4\x8d\xa6\x3c\xd5\x0d\xeb\xbc\x75\x7c\x4a\x6d\xac\x35\x47\x0b\x36\x40\x50\x40\xd1\x2d\xd8\x72\x54\x24\xc5\xec\xac\x33\xff\x2e\x1d\x7c\xa1\x80\x0a\x52\x97\x2f\xae\xf1\x3a\x61\xfd\x78\x7b\xa8\x63\x35\x0a\x00\x5d\xc7\xab\xd5\xca\x0a\x3b\x79\x9e\x8f\x07\x2a\xe3\x43\x8a\xeb\x47\x54\xe2\x62\x3a\x1a\xba\x30\xf4\x38\xda\xc1\xdb\xab\x11\x1f\x36\x77\x73\x34\x5f\x5b\x0c\xc0\xdd\x2c\x49\x8a\xcb\x71\xaf\x23\xf6\xea\x28\xf8\xaa\x06\x56\x4c\x34\xdb\x74\x11\x72\x9a\xd5\xf6\xc2\x46\xbb\x3d\x96\xc8\x2b\x30\x53\xcb\xa1\x91\x5e\xac\xbc\xa0\xd5\x73\x75\x28\x39\x6e\x4a\xb8\x37\x2e\x14\x68\x78\x25\xda\x40\x69\xc6\x55\x4f\xac\x62\x74\x9c\xfe\x18\xbe\x3b\x6e\xb7\x8e\xda\x23\x75\xca\x6e\xe7\x46\x8a\xbe\x3a\x79\x37\x42\x75\xa6\x1d\x67\x0f\x94\x64\x48\x57\x3d\x74\x57\x9f\x44\x65\x67\x6c\xb3\xd9\x2c\x1b\x5f\x8c\x6a\x45\x02\x0b\xf1\xd3\xfa\x0d\x18\x47\xf9\x03\x14\xe7\xf2\x23\x17\x8e\x2b\x4f\x72\xe1\x8d\xf2\x25\x17\x92\x95\x37\xb9\x50\xec\xfc\xc9\x85\x33\xca\xa3\x9c\x83\xb3\xf2\xa5\xa9\xce\xde\x82\x37\xce\xaf\x5c\xd8\x76\x9e\xe5\xc2\x19\xe7\x5b\x2e\xac\x71\xde\xe5\x56\xe3\x45\x96\x46\x79\x98\x0b\x49\x46\x03\x17\xce\x28\x67\xeb\x21\xd2\x59\x5c\x80\x3e\xa9\x33\xed\xbc\x73\x2b\x79\x13\x25\x9a\x34\x4f\x23\x4b\x7a\x2d\x05\x99\xc6\x75\x14\x7c\x07\xa0\xe3\xc9\x05\x1b\x0f\xb1\x83\xa8\xaa\x3d\x94\x4d\xa0\xb2\x32\x1d\xa9\x02\xc6\xd0\x0e\x5a\xc8\xeb\x8e\x75\x6f\xda\x9d\x7f\x67\xfe\xe9\x1a\x50\x2f\xcb\xf1\xf9\xae\x0b\xbd\x5d\x38\xaa\x8b\x9b\x81\xba\xb4\x63\x45\x11\x33\x1f\x3b\xd7\xb2\x0d\x87\xda\x49\x83\xbc\x9d\xeb\x4e\x54\xdd\x19\x75\x17\xdb\x86\xeb\x96\x91\x17\xcb\x20\xe9\x54\x9f\xc1\x04\xeb\x6f\x2f\xb6\x91\x38\xf2\xe2\xd8\x91\x4c\xaa\xb2\xfc\x0c\xe4\x24\x00\xc5\x78\xc5\x6c\x80\x8f\xd6\x96\x3b\x54\xa3\x03\x81\xf1\x6a\x5c\x89\xc4\x8a\xfe\x7d\xf4\x8d\x9a\xe7\xa9\x26\xb6\x4d\x1a\xc0\x09\x5f\xe9\xd8\x11\x74\x07\xc9\x45\x5e\x10\x37\xcf\x56\x76\xef\xba\x61\x2a\x6f\x91\xc6\x17\x58\x1a\x4f\xde\x2e\xc6\xda\xe9\xfa\x6b\x80\xeb\x02\x9e\xd3\xd9\x05\x92\x98\x32\x1e\xe4\x7b\x5c\x3a\x74\x60\x13\x1e\x70\xe5\xb5\xd7\xf1\xd0\xe6\xcd\xb3\x27\xfe\x8f\xfa\xb9\xd0\x9c\x4e\xcf\xf6\x9f\x96\xe8\xd7\xb0\xa2\x35\x9b\xe6\xea\x95\xdc\xbc\x81\x03\xab\xd7\x20\x19\xeb\xd9\x8b\xbc\x56\x20\x17\xbb\x6f\x0f\xc2\xba\x78\x39\xb3\xe3\xa5\xe1\x14\xf1\xf8\xec\x72\x84\x93\x04\xb3\xfe\xe0\x7a\x12\x29\x0e\x92\xe1\x44\x7f\x1a\x69\xd6\xdf\x77\x38\x83\xa4\x5d\xb9\x18\x23\xfd\x86\x80\xdc\x3a\x56\x50\xc2\x96\x1b\x2b\xaa\xe5\x72\x39\x11\xc3\x86\xa3\xc3\x70\xbd\xb2\xa2\x67\xdb\xa5\xe5\xbd\x1a\xc9\x51\xe3\x7f\xa9\xa0\xc0\xc8\x93\x77\xe7\x58\x4e\x01\x6a\xf9\xae\xe1\xaa\xbf\xd8\xe8\xa5\xf3\x75\xd4\x3c\x5f\x1f\xb5\x61\xb7\xc7\x1e\xc6\x8d\xe3\x4e\xb9\xab\x51\x38\x1d\xcf\xa3\x93\xf9\xe2\xff\xef\x91\xff\xd1\xf6\xc8\xf5\x14\xc4\x58\xf9\x3b\xaf\xb7\x6a\xc1\x66\xac\x6e\x0d\x3a\xd6\xf7\x08\x68\x2b\x7c\x84\x60\x68\x7c\x04\xd5\x54\x3e\x82\xe9\x3a\x1f\x01\x0d\xa5\x8f\x59\xd6\xb4\xea\xa4\xfb\x2a\x04\x53\xef\x23\x34\x5d\xf1\x23\xa0\xa9\xf9\x11\xd8\x54\xbd\x43\xe4\xd3\xdd\xb6\xdb\xe5\xa6\xae\x2f\xaf\x8b\x8c\xc5\x99\xc3\x50\xb4\xf5\x9a\x88\x72\xd9\x54\xb6\xf9\x2b\x16\x60\xbf\x32\x14\x3b\x20\x56\xc4\xd4\xc2\xa4\x4c\x9e\x8d\x23\x84\xc8\x0b\xd7\x50\x9d\xda\xce\x2b\xa8\x0f\xc7\xdf\xe5\x38\x43\x50\x0e\xb6\xf8\x19\x8a\x63\xbf\xa4\x92\xc5\x4c\xce\x16\x91\x5c\x56\x45\x99\xbd\x14\x92\xcd\x4a\xcc\xba\x35\x9e\x2c\x63\x0e\xd5\x71\xb4\x30\xb3\x5b\x1c\xc5\x3f\xed\x33\x25\x79\x88\x3c\xda\x53\x88\x4e\x23\xa2\xfd\xe9\x6b\x8f\xdc\xee\x66\x68\x4a\x6d\x7b\xa8\x1f\x74\x9e\xf6\x80\x44\x4b\x4b\xf4\xce\x3b\x29\xe3\x6b\x48\x3a\xe1\x3d\xa1\xf8\x2b\xa9\x39\x2a\xb5\xa5\xc3\xab\xdb\x78\xb6\x10\x5c\x13\xf5\xe5\xa6\x42\x18\xfe\x45\xac\x76\xcc\x97\x11\x19\x34\x88\x22\x4e\xe8\x9b\xee\xe7\x9e\xcb\xe3\x7a\x2e\x6d\xae\x8f\xae\x50\x2e\xa1\x72\x01\x48\xad\xe3\xf8\x61\x7d\x8f\x36\x8c\x94\x07\x0e\xca\x1e\xa5\xd0\x95\x49\x8e\xcd\xc6\x6d\xa5\xee\xa1\xf7\xbd\x6a\x66\x2e\x5f\x54\x28\xe5\xb6\x57\xa8\x74\x3a\x72\x33\x56\xdd\x4c\x52\xb7\x2a\xbe\x75\x90\xd3\x65\xae\xae\x50\xb8\xb0\xec\xc4\xc7\x1e\xab\xc1\x3e\x62\x7d\xbb\x6f\x2d\x43\x4f\xd1\x96\x0f\x77\x13\xda\xf5\xdd\x42\xae\xef\xd4\xfb\x85\xf7\xff\x48\x16\x7f\x5e\xbf\xd7\x42\x8d\x7c\xd1\x77\x59\x3c\xaf\xe8\xd5\xe8\xe3\xf3\x7b\x9d\x26\xcb\x29\x29\x4b\xf9\x4e\xa6\x7b\x54\x13\xbc\xa4\xaa\x36\xeb\x6b\x9e\xbb\x77\x7b\xce\x96\x93\x1e\x73\xc6\x55\xde\x46\xe8\xb2\xeb\xb9\x87\xe4\xb8\xae\xa8\x8d\x52\x8d\x49\x1f\xa5\x34\xab\xa0\x62\x41\x5f\xa7\x05\xbf\xfe\xfd\x57\x0f\x93\x1d\x09\x7b\xe6\xe4\x90\xef\xcd\x24\x3d\x7a\x03\x9f\xda\x73\x2b\x59\xb7\x41\xd4\x75\xe3\xc5\x0a\x04\xe3\x6d\xe5\x3c\xcf\xfb\x60\x3b\xa4\x61\x6a\xb7\xf6\xb2\x21\x0d\x94\xdb\xa8\xd9\xf4\x4f\x2b\x63\x39\x73\xf7\x89\x7e\x38\x3c\x98\x31\x22\xb7\xf5\x7a\xf1\xd0\x34\x40\x73\xc4\xfa\x33\x88\xc5\x72\x51\x2c\xe7\xf6\x1c\x70\x1c\x4e\x28\x9c\xc1\xe6\xdc\xf6\xb9\x73\x22\xb1\xee\x59\x19\x2c\x4e\x6c\xa4\x18\xed\x8e\xfd\x7d\xa5\x85\x0b\xec\xd9\x4e\xa6\x9e\x07\x9f\x7d\x4c\x7c\x36\xd6\xd8\xfc\x0f\x44\x47\xb5\x6a\xb7\xc3\x79\xd0\x63\x1a\x89\xba\xa0\x64\xb3\xea\xbf\x02\x47\xbe\x4c\xe5\xd0\x1f\xdd\x44\x51\x67\x3b\x6f\x7c\x4e\x97\x41\xd5\xf0\x97\x20\x87\xb2\x64\x29\xdb\x93\x27\xd7\xf1\xc8\x46\xfc\x74\xfa\x5e\xfb\xf4\x5c\xeb\xbe\x7d\x51\x2b\xdf\x49\x7b\xeb\xc5\xc7\xdb\xd8\x43\xf6\x2b\xb5\x7e\x47\x4a\x9e\x4c\x8c\x6e\x1f\x1b\x3d\xf0\xce\x32\xda\xe2\xbe\x1b\x93\x9a\xb7\x6c\xee\xba\x81\x29\xaf\xe8\x6f\x0e\x6a\x91\xda\x7e\xec\x9b\xd9\xaf\x7c\xa7\xf6\xef\x3a\x7e\x1c\x1b\x53\x1d\x6f\xc6\x3e\x94\xc6\x67\xc7\x8e\x49\x49\x98\xb9\xcb\x3c\x22\xf1\xd3\x9e\x8b\x8f\x56\xe1\xf6\xa3\x6f\xb9\x60\xb7\xb8\x7c\x8d\x89\x4b\xdc\x80\x14\x85\x13\xff\xc3\x36\x11\x3f\x03\x97\x71\x8a\x1b\x28\x3c\x4e\xd3\x9a\xef\xd5\x40\xaf\x92\x3a\x88\xaf\x5f\x4d\x42\x89\x05\x54\x9f\x53\xc7\x4e\x63\x53\xd3\xda\x6d\x48\xf1\xf2\x2d\xa7\xda\x3e\xd7\xb7\x36\x2d\x97\xc0\xf5\x48\x6a\x5a\x95\x01\xe9\x0d\x6c\x30\x21\x6d\xdf\xee\x55\xac\xea\xe4\xde\xc2\xec\xff\x04\x00\x00\xff\xff\x6f\x05\x93\x99\x86\x43\x00\x00")

func assetsStaticCssPureMinCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsStaticCssPureMinCss,
		"assets/static/css/pure-min.css",
	)
}

func assetsStaticCssPureMinCss() (*asset, error) {
	bytes, err := assetsStaticCssPureMinCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/static/css/pure-min.css", size: 17286, mode: os.FileMode(420), modTime: time.Unix(1424722911, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsStaticCssStyleCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xca\x4f\xa9\x54\xa8\xe6\x52\x50\x50\x50\x48\xce\xcf\xc9\x2f\xb2\x52\x48\xca\x49\x4c\xce\xb6\xe6\xaa\xe5\xe2\x4a\xcb\x4c\xcd\x49\x29\x4e\x2d\x81\xca\xe7\x26\x16\xa5\x67\xe6\xe9\x96\xe4\x17\x58\x29\x18\x19\x14\x54\x58\x73\xd5\x02\x02\x00\x00\xff\xff\xb9\x17\x4d\x4a\x3e\x00\x00\x00")

func assetsStaticCssStyleCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsStaticCssStyleCss,
		"assets/static/css/style.css",
	)
}

func assetsStaticCssStyleCss() (*asset, error) {
	bytes, err := assetsStaticCssStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/static/css/style.css", size: 62, mode: os.FileMode(420), modTime: time.Unix(1476135025, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsStaticJavascriptAlertJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4f\xcc\x49\x2d\x2a\xd1\x50\xaa\xcc\x2f\x2e\x2e\x56\xd2\xb4\x06\x04\x00\x00\xff\xff\x78\x6a\x06\xca\x11\x00\x00\x00")

func assetsStaticJavascriptAlertJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsStaticJavascriptAlertJs,
		"assets/static/javascript/alert.js",
	)
}

func assetsStaticJavascriptAlertJs() (*asset, error) {
	bytes, err := assetsStaticJavascriptAlertJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/static/javascript/alert.js", size: 17, mode: os.FileMode(420), modTime: time.Unix(1475869220, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesDudeHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x91\x41\x6b\x84\x30\x10\x85\xef\xfe\x8a\x59\x61\x61\x17\xda\x8d\xd0\x9b\x66\x3d\xb7\xf7\xde\x4b\x34\x93\x3a\x90\x64\x24\x89\x88\x88\xff\xbd\xa8\xbb\xd8\xbd\x25\x8f\xcc\xfb\x5e\xde\xc8\x2e\x39\x5b\x67\x00\x00\xb2\x43\xa5\xf7\xe3\x76\x4d\x94\x2c\xd6\x9f\x68\x2d\xc3\x3c\xc3\x10\x31\x78\xe5\x10\x96\x05\x4c\x60\x07\x5f\x81\xa2\x14\xfb\xab\x6c\x77\x10\x87\x85\x6c\x58\x4f\x87\xdb\x7c\x06\x32\x70\x02\x8a\x3f\x4a\x3b\xf2\x70\x5e\xb2\x03\xd5\x43\x4c\x93\xc5\x7b\xde\xb2\xe5\x50\x36\x76\xc0\xca\xb0\x4f\xef\x23\xd2\x6f\x97\x4a\x68\xd8\xea\x0a\x36\xc9\x28\x47\x76\x2a\x21\x62\x20\x53\xe5\x07\x62\xc3\xbc\xe6\xa4\x08\xca\xc3\x06\x3c\x49\xd1\xbf\xc4\x41\xaf\xc9\xac\x31\x9e\xda\xb4\x7e\xf3\x39\xbd\x1c\xba\x14\x83\xfd\x57\x4b\x6c\x03\xf5\xa9\x8e\x98\xbe\xc9\x21\x0f\xe9\x62\x06\xdf\x26\x62\x7f\xb9\xce\x23\x79\xcd\xe3\xcd\x72\xab\x56\xe5\xd6\x05\x34\xf7\x5c\xe4\xcb\xdb\x47\x51\x14\xd7\x4a\x8a\xc7\xf8\xa3\xaf\xbd\x24\x29\xb6\x2d\xfc\x05\x00\x00\xff\xff\xc0\xe1\x56\x7f\x8c\x01\x00\x00")

func assetsTemplatesDudeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesDudeHtml,
		"assets/templates/dude.html",
	)
}

func assetsTemplatesDudeHtml() (*asset, error) {
	bytes, err := assetsTemplatesDudeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/dude.html", size: 396, mode: os.FileMode(420), modTime: time.Unix(1476204409, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\xc1\x6e\xc3\x20\x10\x44\xcf\xe1\x2b\xb6\x48\x39\xa6\xdc\x13\xec\x73\xff\xa2\x22\x66\x89\x57\x5d\xc0\x62\xb1\x2a\xcb\xf2\xbf\x57\xc2\x69\x95\x54\xaa\xd4\x39\x0e\xf3\x46\x0c\xd8\xb1\x46\xee\x95\x1d\xd1\xf9\x5e\xd9\x4a\x95\xb1\x7f\x43\xe6\x0c\xeb\x0a\xb3\x60\x49\x2e\x22\x6c\x9b\x35\xfb\x99\xb2\x4c\xe9\x03\xc6\x82\xa1\xd3\x46\xaa\xab\x34\x98\x41\xc4\x4c\x73\xc1\x53\xa4\xf4\x3a\x88\x68\x88\xe8\xc9\x75\xda\x31\x6b\x28\xc8\x9d\x96\xba\x30\xca\x88\x58\x35\x98\xbf\x5b\x5a\xec\x9f\x15\xe6\x7e\xe9\x6b\xf6\x4b\xaf\x0e\xeb\x11\x28\x00\xc9\xbb\xf3\x91\x12\x74\x1d\xe8\x65\x9e\x34\x1c\x37\xa5\x0e\x76\x82\x86\x77\x7a\xc8\x9c\xcb\x59\xc1\x5d\x3b\xd5\xcc\x86\xdc\x0a\x62\x6a\x10\x3c\xe8\xca\x33\x3e\x22\xc8\x82\xbf\x33\x05\xfd\x53\x24\x79\x0a\x70\xdc\x2e\x3f\x66\xc8\xa9\x9e\x3e\x91\x6e\x63\x3d\xc3\x35\xb3\xbf\xec\x56\x70\x91\x78\x39\x83\x60\xa1\x70\xd1\xfd\xf3\xc3\x03\x09\xb8\x04\x6d\xd4\x8b\x35\xd3\xbe\xf4\xbb\x5d\x59\xb3\xcf\xb7\xa6\xfd\xe4\x57\x00\x00\x00\xff\xff\xf0\xfd\x58\x0c\xd0\x01\x00\x00")

func assetsTemplatesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesIndexHtml,
		"assets/templates/index.html",
	)
}

func assetsTemplatesIndexHtml() (*asset, error) {
	bytes, err := assetsTemplatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/index.html", size: 464, mode: os.FileMode(420), modTime: time.Unix(1476204518, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesStartHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x91\xc1\x6a\xc3\x30\x10\x44\xef\xf9\x8a\x8d\x21\x90\x40\x1a\x05\x7a\xb3\x95\x9c\xfb\x01\x85\x1e\x8b\x6c\xad\xea\x2d\xb2\x36\x48\x2b\x82\x31\xfe\xf7\x62\x3b\xc1\xcd\x4d\x1a\x31\x9a\xb7\xb3\xba\x95\xce\x5f\x37\x00\x00\xba\x45\x63\x97\xe3\x7c\x15\x12\x8f\xd7\x0f\xf4\x9e\x61\x18\x20\x27\x8c\xc1\x74\x08\xe3\xa8\xd5\xf2\xb6\x59\x7c\x6a\x35\xea\x9a\x6d\xbf\xfe\x31\xec\x80\x1c\x6c\x81\xd2\xb7\xb1\x1d\x05\xd8\x8d\x9b\x35\xe0\x06\x49\x7a\x8f\x97\xa2\x61\xcf\xb1\xac\x7d\xc6\xca\x71\x90\xb7\x3b\xd2\x4f\x2b\x25\xd4\xec\x6d\x05\xb3\xe4\x4c\x47\xbe\x2f\x21\x61\x24\x57\x15\x6b\xc4\x1c\xf3\x42\x07\x94\xc0\x04\x98\x03\xb7\x5a\xdd\x5e\x70\x30\x58\x72\x13\xc6\x53\xeb\xa7\xe1\x9e\xee\x71\x3c\xae\x7c\x5f\x08\xbf\x39\x09\x24\x31\x51\xd0\x82\x81\x48\x96\x9a\xec\x39\x27\xdf\x43\xc3\xec\xc1\x62\xc7\xa5\xae\x23\xa8\x7f\xc5\xa5\x26\xd2\x4d\xae\x09\xe5\x93\x3a\xe4\x2c\x7b\x97\x43\x23\xc4\x61\x7f\x18\xee\x14\x2c\xdf\x4f\x9e\x1b\x33\x29\xa7\x36\xa2\xbb\x14\xaa\x18\x8f\xef\xe7\xf3\xf9\x50\x69\xf5\xb0\x3f\xba\x5d\x0a\xd5\x6a\xde\xd3\x5f\x00\x00\x00\xff\xff\x05\x60\xdc\x64\xae\x01\x00\x00")

func assetsTemplatesStartHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesStartHtml,
		"assets/templates/start.html",
	)
}

func assetsTemplatesStartHtml() (*asset, error) {
	bytes, err := assetsTemplatesStartHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/start.html", size: 430, mode: os.FileMode(420), modTime: time.Unix(1476204458, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/static/css/pure-min.css": assetsStaticCssPureMinCss,
	"assets/static/css/style.css": assetsStaticCssStyleCss,
	"assets/static/javascript/alert.js": assetsStaticJavascriptAlertJs,
	"assets/templates/dude.html": assetsTemplatesDudeHtml,
	"assets/templates/index.html": assetsTemplatesIndexHtml,
	"assets/templates/start.html": assetsTemplatesStartHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"static": &bintree{nil, map[string]*bintree{
			"css": &bintree{nil, map[string]*bintree{
				"pure-min.css": &bintree{assetsStaticCssPureMinCss, map[string]*bintree{}},
				"style.css": &bintree{assetsStaticCssStyleCss, map[string]*bintree{}},
			}},
			"javascript": &bintree{nil, map[string]*bintree{
				"alert.js": &bintree{assetsStaticJavascriptAlertJs, map[string]*bintree{}},
			}},
		}},
		"templates": &bintree{nil, map[string]*bintree{
			"dude.html": &bintree{assetsTemplatesDudeHtml, map[string]*bintree{}},
			"index.html": &bintree{assetsTemplatesIndexHtml, map[string]*bintree{}},
			"start.html": &bintree{assetsTemplatesStartHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
