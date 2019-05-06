// Code generated by go-bindata.
// sources:
// tmpl/css/style.css
// tmpl/index.html
// tmpl/js/mapping.js
// tmpl/js/request_logger.js
// tmpl/js/sorting.js
// tmpl/js/util.js
// tmpl/swagger.json
// DO NOT EDIT!

package console

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

var _tmplCssStyleCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\xdd\x8a\xdb\x30\x10\x85\xaf\xab\xa7\x18\xd8\xdb\xda\x34\xdb\x2c\xdd\x2a\x50\xd8\xfe\xec\x6b\x2c\x23\x69\xec\x08\x2b\x1a\x33\x52\xea\x4d\x43\xde\xbd\xc8\x56\x9a\x16\xea\xbd\x32\x3e\xdf\x9c\x23\x74\x46\x5a\x37\x13\x99\xc1\xe7\x26\x59\xe1\x10\x0c\x0a\x9c\x15\x00\xc0\xe4\x5d\xde\x6b\xd8\xdc\x8f\xaf\x3b\x75\x51\xff\x19\x6c\xb2\xa0\x1d\xea\xb8\x41\x3b\xf4\xc2\xc7\xe8\x1a\xcb\x81\x45\xc3\xdd\xb7\xa7\x1f\xdb\xe7\xa7\xdd\x82\x59\x1c\x49\x13\xa8\xcb\x1a\x36\xe3\x2b\x24\x0e\xde\xc1\x9d\xb5\x76\x35\x7d\x7f\x3c\x98\xf5\xf4\xed\xfd\xa3\xb1\xf8\xa6\x59\xef\xf9\x27\xc9\x7a\xc4\xe6\xfb\xa7\xed\xd7\xe7\x12\xa1\xda\x44\x81\x6c\x26\xf7\x22\x3c\xc1\x19\x3a\x8e\xb9\x99\xc8\xf7\xfb\xac\xc1\x70\x70\xbb\x8b\x6a\x53\x16\x1f\x7b\x38\x43\x0d\xe8\x85\x28\xee\xe0\xa2\xda\x78\x3c\x98\x72\xd2\x95\x38\x94\x81\x05\x63\x4f\x33\x36\xcc\x81\x30\xde\xb8\x09\x47\xaa\xc6\x10\x6e\xf2\x01\x7b\x8a\x19\x67\x32\xd0\xe9\x06\x84\x5c\x11\x55\x46\x13\x08\x5a\xcb\x21\xe0\x98\xa8\xf5\x11\xce\xea\x9d\xf3\x69\x0c\x78\xd2\x33\x6d\x84\xa7\xe5\x4a\xcb\xef\xd2\xc1\x17\xc8\x86\xdd\xa9\x7c\xa5\xde\xb5\x0c\xea\x3f\xd4\xbd\x57\x7f\xe9\xb3\xb2\x5e\x5c\xd7\x75\xb5\x35\x96\xec\x63\xff\xe2\x2d\xc7\x54\xe7\x47\x74\xce\xc7\xbe\xee\xfa\xa1\x3c\x9f\x22\x5f\xbd\x88\xb8\x08\x73\xc5\xc9\xff\x22\xfd\xa1\x7d\xa4\xc3\xbf\x6b\xa8\x59\x57\x93\xfb\xfc\xf0\x71\x5b\xce\xfc\x1d\x00\x00\xff\xff\x65\xd5\x4f\x11\xb0\x02\x00\x00")

func tmplCssStyleCssBytes() ([]byte, error) {
	return bindataRead(
		_tmplCssStyleCss,
		"tmpl/css/style.css",
	)
}

func tmplCssStyleCss() (*asset, error) {
	bytes, err := tmplCssStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/css/style.css", size: 688, mode: os.FileMode(420), modTime: time.Unix(1519645595, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5a\xdb\x52\xe3\x38\xfa\xbf\xe7\x29\xd4\x9e\x99\x3f\xa1\xba\x6d\x13\x68\x1a\x06\x1c\xff\x8b\x81\xcc\x76\x53\xa4\x9b\x69\x18\xd8\x99\xad\xad\x2e\xd9\x52\x62\x05\x59\x32\x92\x9c\xc3\x50\x54\xed\x13\xec\x7b\xec\xc5\xbe\xc4\xee\x9b\xec\x93\x6c\xc9\xb2\x1d\xc7\x39\x60\xe8\x9a\xde\x99\xc9\x45\x62\xcb\xfa\x4e\xfa\x0e\xfa\xe9\x73\xbc\x17\xa7\x1f\x4e\xae\x7e\xba\xe8\x82\xb7\x57\xbd\x73\x7f\xc3\x8b\x54\x4c\xfd\x8d\x0d\x2f\xc2\x10\xf9\x1b\x00\x00\xe0\xc5\x58\x41\x10\x29\x95\xd8\xf8\x2e\x25\xa3\x8e\x75\xc2\x99\xc2\x4c\xd9\x57\xd3\x04\x5b\x20\x34\x77\x1d\x4b\xe1\x89\x72\x35\xfd\x11\x08\x23\x28\x24\x56\x9d\x1f\xaf\xbe\xb7\x0f\x2c\xe0\xe6\x9c\x14\x51\x14\xfb\xbd\x1e\x0f\x6f\xc1\x09\x67\x92\x53\xec\xb9\x66\xd0\x4c\x90\xa1\x20\x89\x02\x52\x84\x1d\x4b\x4b\x94\x87\xae\x1b\x72\x84\x9d\xe1\x5d\x8a\xc5\xd4\x09\x79\xec\x9a\x4b\x7b\xc7\x69\x3b\xaf\x9d\x98\x30\x67\x28\x2d\xdf\x73\x0d\xe9\x3a\x3e\x88\x0d\xa5\x13\x52\x9e\xa2\x3e\x85\x02\x67\xcc\xe0\x10\x4e\x5c\x4a\x02\xe9\x46\x90\x21\x8a\x03\x28\xa4\x33\x94\xee\x8e\xb3\xed\x6c\xcf\x8f\x55\x85\x18\x29\x94\xb0\x5b\x20\x30\xed\x58\x52\x4d\x29\x96\x11\xc6\xca\x02\x91\xc0\xfd\x99\xd4\x18\x4e\x42\xc4\x9c\x80\x73\x25\x95\x80\x89\xbe\xd1\x82\xcb\x01\x77\xd7\xd9\x75\xde\xb8\xa1\x94\xb3\xb1\xcc\xac\x50\x4a\x0b\x10\xa6\xf0\x40\x10\x35\xed\x58\x32\x82\xbb\x07\xaf\xed\xf6\xdd\x41\x7c\x75\xf6\xe1\xf8\x72\x72\x30\x6c\x1f\xa7\x2f\xe1\xde\xcd\xe9\x35\xbb\x20\x3b\xf4\xf6\xfb\xfe\x78\xdc\x3d\x86\x07\xd1\xe9\x29\x1a\xfe\x4c\x93\x73\x3c\x98\x44\xc3\xeb\x5e\xb7\xdd\x1f\x0c\x6f\x2e\xfe\x14\xdf\xfe\x22\xf7\x2d\x10\x0a\x2e\x25\x17\x64\x40\x58\xc7\x82\x8c\xb3\x69\xcc\x53\x69\xf9\xbf\xb2\x51\xb6\x8a\x70\x8c\xd7\x99\xd6\x3f\xbf\xd9\x79\xbf\xdd\xa6\xbd\xbb\x21\xbc\xfd\xee\x76\xb2\x4b\xdd\xde\xb7\x5d\x18\xa5\xe3\xe4\xb2\x8f\xdf\x8f\xae\xdf\xec\x9e\xed\xe1\x5f\xd8\x6e\xfa\xf3\x2f\x30\xb9\xda\x4e\xf7\xbb\x3f\xc9\x3f\xf7\x86\x3f\x5c\xbf\xdc\xee\xb2\x3d\xf1\x98\x69\xcb\xa2\xa2\xa9\x29\xc3\xba\x7b\x86\x4b\x4d\xd8\x8e\x2f\x83\xb3\xd3\xee\x5b\x02\x69\x3f\x4e\xbf\xfb\xee\x87\x8b\x37\xc7\xaf\x7f\x10\x89\xb8\xdb\xfb\x70\xdd\xbf\xd9\xdd\xbf\xf8\xf8\x71\x77\xb8\xd7\x3d\xbf\x9b\x48\xd9\x9e\x5e\xdf\x7d\x50\x0c\x27\xec\xed\xf5\xc5\xb7\xf0\x6c\x7f\x72\xb9\xda\x84\xa5\xc1\x37\xef\x96\x47\x22\xbc\xd4\x7f\xd7\x46\x04\x52\x3e\x70\xdb\xce\xee\x6b\xe7\xdb\x9a\x9b\xcc\xb3\x99\x9f\x16\x62\x41\x4d\x13\x9c\x67\xbb\x9e\xf0\xfc\x94\x5b\xa9\xd0\x70\x85\x3e\xcb\x92\x30\x97\x5b\x51\xea\x0c\x8e\xe0\x65\x36\x6a\x19\x75\x86\xd2\x4d\x15\xa1\x6b\x0a\xc5\x23\xd4\x02\xdf\xa5\x58\xaa\x4f\x94\x0f\x06\x58\x3c\x9f\x4f\x0c\x93\x84\xb0\xc1\xf3\x19\x48\x2e\xd4\x0a\x06\xcb\xf3\xb6\xe6\xab\x3c\x60\xb4\xc3\xb3\x59\x4e\xc5\x81\x59\x99\x67\x30\xc6\x1d\x6b\x44\xf0\x38\xe1\x42\x55\x8a\xfb\x98\x20\x15\x75\x10\x1e\x91\x10\xdb\xd9\xcd\x2b\x40\x18\x51\x04\x52\x5b\x86\x90\xe2\x4e\xdb\x6a\x62\x8a\x99\xa3\x3f\x5f\xb7\x00\xe2\x61\x1a\x63\xa6\xc0\x96\x23\x30\x44\xd3\x56\x3f\x65\xa1\x22\x9c\xb5\xb6\xc0\x3d\x28\x67\xea\xcf\xdc\xcd\x08\x0a\x90\xfb\xe4\x3c\x73\x09\xe8\x00\x86\xc7\xe0\x63\x75\xac\xb5\x75\xb4\x48\x35\x96\x17\x82\x2b\x1e\x72\x0a\x3a\x60\x4c\x18\xe2\x63\x87\xf2\x10\x6a\xa1\x4e\x52\x3e\xea\x74\xc0\xa6\x09\xe2\x4d\xf0\xff\x60\x73\x2c\x75\x34\x6f\x82\x43\x7d\xa9\xaf\x8e\x96\x30\xce\x75\xb8\xc1\xc1\x25\x0f\x6f\xb1\x6a\x55\x64\xbd\x04\xa5\x90\x88\x4b\x05\x5e\x02\xcb\xc5\x61\xc4\xad\xad\x79\x4e\x63\xe9\x70\x16\x63\x29\xe1\x00\x83\x0e\x28\x97\x03\x8f\x30\x53\x0b\x6b\x52\xc8\x9e\x11\x9c\x5d\x7e\x78\xef\x24\x7a\xef\x35\x24\x0e\x82\x0a\xd6\x64\xe8\xcf\xdc\xe2\x39\x94\x0f\xba\x4c\x89\x69\x2b\x67\x54\x23\x78\x38\xda\xd8\x58\xb0\x37\x0f\xe5\xdc\xe8\x9e\xb9\x6b\xcd\x59\x59\x63\x53\x04\x7f\x9f\x50\x7a\x4e\xa4\x6a\x6d\x1d\xcd\xb3\x75\xdd\x9c\x0d\x50\x30\xa0\x18\xe4\xc1\x3e\x37\x87\x62\x55\x30\xba\xd2\x93\xde\x21\xd0\x01\xd6\x57\xf9\x90\x9d\x11\x5a\x47\x0b\x24\x1a\xcd\x60\xd1\xa5\x58\x47\x9b\xf6\xd4\x5f\xfe\x5a\x13\xfe\x75\x6b\x9e\xed\x96\xd3\x27\x0c\xb5\x36\x55\xb4\xb9\xe5\x28\x7e\x2c\x04\x9c\xb6\xb6\x9c\x3e\x17\x5d\x18\x46\xb3\x40\xc5\x86\xe7\x16\xb8\x5f\x58\x64\x2d\x98\x68\xfd\xf2\x39\x0e\x41\x8b\x9e\x20\x7d\xd0\x22\x08\xbc\xd0\x11\xb7\xb9\x05\xee\xe7\x35\x75\x92\x54\x46\xad\xcd\xaf\x36\xc1\x4b\x40\xd0\xd6\xc3\xbc\x5f\xea\x0b\xa8\x05\xe6\x8b\x96\xfb\xe5\xd2\xdc\xd5\x6c\x7b\x55\x5b\x8f\x9a\xa3\x8a\x22\xa3\xd3\xbb\x55\x79\x56\xca\x9b\x55\x1e\xcf\x35\x38\x71\xc3\x0b\x38\x9a\xe6\xf9\x1f\x89\xfc\x02\x91\x11\x08\x29\x94\xb2\x63\xe9\x3a\x02\x09\xc3\xc2\xee\xd3\x94\xa0\x4a\x19\xf0\xa2\x76\x1d\x11\x46\xed\xca\xe3\x0a\x13\xc1\xc7\x15\xc2\xec\x29\x94\x04\xe1\x99\x10\x6a\xc7\xc8\xde\x01\x09\x44\xb6\x20\x83\x48\xd9\xdb\x35\x82\x8c\x28\xa5\x05\x05\x83\x23\xc0\xe0\xc8\x4e\x08\xa5\x32\xbb\x92\x0a\x86\xb7\x18\x2d\x21\x03\xa6\xca\x16\xa4\x30\x54\x64\x84\x2d\xdf\x83\x40\x27\x99\xad\xf8\x60\x40\x75\xc1\x83\x41\x51\x66\xbf\xca\xd3\xcc\xf2\xf3\xc2\x24\x3d\x17\xae\x60\xec\x52\xb2\x52\xe4\x5a\x19\xb9\x6b\x2d\x3f\x4f\x9e\xe7\x8b\x30\x0c\x87\x70\x04\x8d\x7b\x0f\x8f\xac\x79\xb1\x31\x47\x90\x16\x63\x50\x0c\xb0\xd2\xf2\xa7\xbd\x6c\xd8\x3f\x0e\x78\xaa\x9e\x28\xdd\x73\x53\x5a\xf3\xa8\x9b\xb9\xb4\x36\x38\x17\x49\x99\x93\xdb\xda\xb3\x8b\xfc\x2a\x13\x15\x0c\xec\x7c\xfb\x5a\xe5\xcd\x17\xb6\x7d\x05\x03\xd0\xb6\xed\x15\x13\x34\x3b\x82\x3a\x56\xe1\xc8\x2a\xef\x04\x32\x0c\xfa\x10\x61\x40\x18\x28\xa2\x61\x29\x9b\x42\x16\x61\x01\x9f\x00\xc5\x39\x0d\xa0\x58\x25\xb3\x6e\xc6\x62\xd0\xaf\x9b\xad\x57\x67\x22\xed\xf6\xce\x23\x34\x75\xba\x24\xa5\xd4\xa4\x4c\x03\xc2\x8c\x18\x16\xa4\x81\x62\x20\x50\xcc\x46\xb8\x0f\x53\xaa\xb2\xeb\x89\xb4\x40\x76\xae\xeb\x58\x1f\x71\x5f\x60\x19\x59\xd9\x3a\x06\x8a\x9d\x50\x0c\xc5\x39\x1f\x34\x90\x93\x4d\x5d\x6f\xfa\xaa\x70\x9b\x9f\x84\xc8\xe8\x91\x25\x6c\x30\xe5\x85\x6d\xbb\x7a\x79\xd7\xf8\x0d\xac\xf2\x06\x90\x09\x0c\xb1\xd8\xd3\xa8\x6d\xad\xa4\xc7\x1e\x6b\x25\x04\x1f\xaf\x0d\x1e\x3d\xe7\x59\xa1\xa6\x23\x9a\x82\xec\xbb\xf4\xe6\x0c\x29\x58\xd5\x4c\x38\xe7\x83\x0b\x3d\xef\xb1\xd0\x7c\x61\xdb\x05\x3a\xa1\x44\xaa\x46\x8b\x57\x11\x63\xeb\x5a\x53\xcd\x3a\x8a\x6d\x81\x65\xc2\x99\x5c\x9f\x6e\x25\x43\x03\x25\xaa\x2c\x0d\x46\x98\xe3\x69\x00\x87\x1d\xf1\x11\xd6\xc0\x32\x0b\x58\xbb\xdc\xb2\x40\x96\x1c\x88\x8f\x99\x05\x32\xd4\x5c\x68\x42\xe1\x94\xa7\xea\x10\xf4\xc9\x04\xe7\x5b\xfb\x98\x0b\x64\x8f\x05\x4c\x0e\x41\x20\x30\xbc\xb5\xf5\xc0\x51\xd3\x9c\x52\x7a\x37\x2d\x34\x8b\x08\x42\x38\x4b\xa5\x66\xd4\x86\x83\x68\x3e\x39\x17\x59\x0d\x56\x19\xdb\xbb\x96\xef\x49\x25\x38\x1b\xf8\xa7\x50\x61\xcf\xcd\x6f\x9e\xc6\xd6\x55\xd1\xe7\x2a\xd2\x9e\x29\x72\xc2\xd1\x6f\x43\x91\x1e\x56\x11\x47\xff\x3b\x55\xf6\x67\xaa\x5c\x40\x15\x7d\x09\x45\x3c\xb7\x69\x4c\x69\xbe\x65\xd7\xf0\xd1\xc9\x6a\x86\x18\x1b\x30\x6e\x36\xd9\x73\xb3\xc4\xfc\xac\x5a\xdf\xa4\x02\x9b\xea\x6a\x0a\x65\xc3\xda\x3a\xc6\x94\x02\xfd\x65\x4b\xa0\xcf\xc4\xc5\x66\xeb\xe1\xd8\x3f\x87\xb2\x2c\xb4\x20\x4d\x10\x54\x18\x1d\x02\x4f\x26\x90\x65\x95\x8b\x42\xa9\x3e\xe5\xe3\x96\x7f\x03\x49\x06\xcf\x1d\xc7\x73\xf5\x14\xdf\x73\x71\xfc\x2c\x7b\xd6\x3d\x7a\x61\xdb\x40\x63\xa3\x9d\x47\xb1\x51\x01\x40\x97\x62\xa3\x75\x88\xe8\xf7\x0f\x72\xb2\x9f\x50\x60\xa8\xb0\x5d\xe2\xf0\xc6\x99\x75\x92\x11\x36\xcc\x80\x2f\x86\x73\xc0\x97\xc6\x3a\x9f\x07\x76\x6a\x51\xf8\x59\x80\xa1\x02\x16\xe6\x1b\x0a\xcb\xc0\x82\x54\x82\x24\x18\x55\xa1\x43\x23\x3c\xf2\x94\x12\xf9\x84\xcd\x5c\xef\x15\xa6\x47\x67\xb5\xf7\xbe\x31\x58\x2d\x15\xc4\xf2\x7f\x14\xc4\x54\x92\xdc\x86\xfc\x74\xff\x89\x84\x9c\x49\xbd\x9d\x54\x9e\x0d\xe8\x34\x89\xf4\x03\x50\x5e\xd9\x61\x84\x47\x82\x33\x83\x7f\xfc\xa2\xe4\x34\xa5\x4a\x93\x19\x4d\xed\xe7\x49\x9b\xd0\xcc\xba\x9d\x37\xb9\x75\x08\xcb\xd0\xf2\x4f\xb1\x39\xac\x12\xce\x7e\x63\x66\x3e\xd3\xbe\xf6\x76\x6e\x5f\x9c\x61\x0d\xab\xc0\x1c\x7f\x08\xe3\x4a\xe7\x25\x50\x45\x96\xc1\x30\x7f\x08\xc3\x0e\xbe\xb1\x40\x7e\xca\x90\x29\xcd\x3a\x3e\xfa\xf7\xf7\x6d\xdc\xaf\x4a\xd1\x0c\x57\x96\x98\xb2\x41\x65\x6d\x86\x11\x9b\x88\x7c\x9c\xd3\xa3\x48\xf3\x19\xb0\x6b\xc9\x78\x6d\xa8\x72\x5b\xdd\x3c\x33\xb0\x96\x75\xe1\x40\xd6\x84\x2b\x36\xcc\xea\xf6\x9c\xf5\xee\x0c\x24\x33\xd5\x25\x6f\xdb\x01\xc1\xf5\x81\xd6\xbc\x64\xb3\x96\xf7\x5d\x33\xda\xfc\x3d\x1c\x30\x37\x32\xae\x77\x62\x17\xe6\xaf\xee\xbd\x2d\xce\x35\xfd\xe8\x55\x6d\xba\x20\x55\x8a\xb3\xfc\x85\x92\xb9\x29\xb7\xe4\x90\x72\x89\xf3\x9e\x24\x22\x32\x26\x25\x53\xcb\xff\x3f\x45\x62\x2c\x8f\x3c\xd7\xd0\xac\xe0\x1e\xbd\x9e\xd7\x25\xeb\x5b\x59\x79\x4f\xfa\x06\x07\x95\xbe\xf4\xeb\x65\x1d\xcc\xa5\xce\x5c\x34\x51\xc7\xd4\x2a\x03\x13\xbf\x17\x6b\x69\x44\x02\xc8\xc0\xdb\xab\xab\x0b\xa0\xef\x09\x1b\x00\x98\x24\x94\x98\x97\x2a\xa0\xcf\x05\x50\x58\x66\xfd\x7d\xc8\x10\xe8\xeb\x73\x43\xf6\xd2\x4a\x4d\x4d\xeb\x37\x59\x29\xe0\x9c\x84\x98\x85\x18\x9c\xf0\x64\x9a\x41\x60\xf0\xaf\x7f\xfc\xe7\x6f\x7f\x07\x3b\xdb\xed\xfd\x57\xe0\x8c\x0b\x44\x40\x0f\x0a\xf5\xef\x7f\x32\xd0\x2a\x9b\xc1\x91\x52\xc9\xa1\xeb\x0e\xf5\x63\x87\x70\x0b\x14\x6d\xdf\x4f\x01\x85\xec\xd6\xf2\x6b\x13\x34\x4e\xdd\x5a\xa7\xc6\x47\x4c\x31\x94\x18\x81\x94\x21\x2c\x40\xef\xdd\x15\xa0\x5a\x33\x89\x5f\x01\x89\x31\x98\x93\x2c\x0f\x5d\x77\x40\x54\x94\x06\xe6\x8f\x1f\x31\xd4\x05\xf4\x60\xc7\x8d\xf5\xea\xb8\x01\xe5\x81\x1b\x43\xa9\xb0\x70\xcf\xdf\x9d\x74\xdf\x5f\x76\x9d\x18\x2d\xea\x98\x3f\x5b\xd3\x9a\x4e\xfc\x6c\x6d\x11\x56\x90\x50\xe9\x68\x8f\x2b\x41\x82\x54\xe1\xc3\xa7\x68\xb4\x7c\x79\xd6\xd3\x2c\xd5\xaa\x71\x50\xf5\x39\x57\xcf\xcb\x9b\xda\x81\x66\x45\x06\x9d\xe8\xec\x5a\x9d\x3f\x9f\x57\xb1\xae\xba\xbd\x8b\xf3\xe3\xab\xee\x25\x28\x0b\x56\xfe\xea\xb8\xda\xae\xc3\x4c\x89\xe9\xdc\x8b\xec\x89\x3d\xfb\x5f\x8e\xad\x70\x9c\x50\xa8\x70\xf5\x6d\x80\xa7\xc4\xac\x3c\x90\xf0\x36\x43\xed\xf7\xf7\xc5\x5b\xfc\x90\x53\x2e\x1e\x1e\x6a\x2f\x32\xca\x8e\xa3\x09\x02\x7b\x36\x9f\xa5\xf1\xc3\x43\xbd\xdc\x29\xb4\xd8\x37\x9b\x91\xe8\xa3\xfa\xc3\x83\xe7\x2a\xf4\x18\x59\xdb\xf2\xbd\xc0\xaf\x2a\x87\x32\xca\xa0\xbe\xac\x8d\x58\xcd\xf8\x18\xf4\xd8\x4c\x87\xfd\x2a\xa1\x46\x66\x75\xb2\xf9\xbd\x5a\x2f\xef\x5c\x8f\x76\xc5\x8a\x95\xb8\x07\x53\x1c\x2a\x5b\x9f\xf1\x8b\xfe\x29\x22\x32\xa1\x70\x7a\x08\x18\x67\xf5\x53\x59\xa6\x1f\xa7\x1a\xb2\x74\xac\xd7\xab\xb6\x8f\x9c\x51\x02\x11\x22\x6c\x70\x08\xb6\x93\x09\x68\xef\x55\xbf\x56\x65\x45\xb4\xeb\x9f\x1a\x8d\x3d\x37\xda\x5d\x3d\x69\x0e\x69\x51\x18\x60\x0a\xb2\x6f\x3b\x11\x24\x86\x62\x5a\xbe\xd5\x2b\xe1\xd5\x4a\x6e\xeb\x8f\xc1\x89\x28\xdf\x60\x26\x02\xdb\x32\x14\x9c\xd2\xec\xdc\xe9\x7b\x3a\x1c\xfc\xfb\x72\x65\x1f\xb4\x67\xb2\x31\xcf\x4d\xc4\x0a\xf8\xb1\x0e\x7a\xac\xb3\x6b\x31\x43\x34\x8a\xd5\x47\x67\xfc\x65\x4c\x34\xb2\x7e\x55\x1b\x8b\x72\xe7\x9f\xf3\xc1\x17\x31\x8a\xf2\xc1\x53\x0c\x6a\x56\x62\x97\xe7\x66\xe5\x9f\x51\x0b\xd5\xb4\xe8\x67\x34\xaf\xa6\x25\xff\x3a\x4c\xf7\x14\x32\x25\xeb\xc7\x8f\xef\x4c\xa5\x5a\x5a\x63\xfc\xfb\x7b\x34\x3b\x9b\xaf\x28\x44\xeb\x82\xd1\x54\xb0\x59\x2c\x96\xe1\xe9\xcc\x6a\x5b\x71\xae\x59\x2a\xbe\x9c\x6f\x4a\xda\xd3\x35\x90\x0a\xaa\x54\xce\x6b\x60\x62\xd4\x31\x8f\x4e\xf2\x62\xbd\x46\x8d\x46\x2d\xc4\x11\xc1\xe3\xb2\x81\x68\xb6\xa5\x54\x90\x8e\x95\xaf\xb1\xe5\x5f\x13\x3c\xd6\x48\xe1\x33\x64\x60\x44\xd4\x5a\x19\x5d\x44\xd4\xa2\x8c\xe5\xe1\xe5\xb9\xe6\x88\xb4\xe1\xb9\xe6\x0f\xc3\xff\x0d\x00\x00\xff\xff\x14\xa7\x8e\x4a\x48\x2c\x00\x00")

func tmplIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplIndexHtml,
		"tmpl/index.html",
	)
}

func tmplIndexHtml() (*asset, error) {
	bytes, err := tmplIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/index.html", size: 11336, mode: os.FileMode(420), modTime: time.Unix(1534889019, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplJsMappingJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x5f\x6f\xdb\x36\x10\x7f\xcf\xa7\x38\xa8\x06\x24\x21\x8e\x5c\x60\x18\x06\xc8\x51\x1e\x9a\x0d\xe8\x86\x75\x1d\x96\xf6\x29\x0b\x0a\x5a\xba\x58\x6c\x29\x92\x23\x4f\x49\x8c\xc1\xdf\x7d\xa0\xfe\x38\x92\x2c\x3b\x6a\x93\x8e\x4f\xa2\xee\xee\x77\xc7\xbb\xe3\xef\x78\x5b\xca\x94\xb8\x92\xf0\x8e\x69\xcd\xe5\x3a\xc8\x54\xc1\xb8\x0c\xe1\xdf\x93\x13\x00\x00\xca\xb9\x8d\xea\x7f\x90\x40\xfd\x51\x09\x16\x8b\x54\x15\x9a\x0b\x04\xca\x11\x08\x0b\x2d\x18\x61\x25\xba\x63\x06\x8a\x1a\xee\x93\x55\xa5\x49\x11\x12\x98\x05\xde\xab\xe6\xe7\x19\x4a\x32\x1b\x2f\x8c\x72\x2a\x44\x10\x2e\xf7\x8c\x5a\x34\x48\xe0\x2d\x93\x99\xc0\x15\x33\x36\x6a\xfc\x05\x7d\xe8\x70\x79\xb2\xb3\xbf\xe5\x42\xfc\xce\x2d\x41\x02\xed\xb9\x02\x77\x12\x68\xd6\x2c\x5a\x23\xfd\x76\xf5\xfe\x8f\xe0\x9e\xcb\x4c\xdd\x47\x42\xa5\xcc\x69\x45\xda\x28\x52\xa9\x12\x70\x0a\xde\x62\xe1\xc1\x69\x73\xd4\x6a\xcf\x34\x5f\x34\x4e\xbd\xf9\x23\x72\xc6\x88\x75\xd1\xdb\x28\x38\x61\x61\x21\x81\xeb\x9b\x65\x4f\xd6\xcd\x00\xb1\x95\xcb\xdc\x4a\x65\x2e\x0f\x58\x68\xda\xb4\x89\x78\x0c\x16\x59\x9a\x57\x5e\x3a\x4e\xbf\xe0\x66\xde\xe6\x69\xe8\xdc\xad\x46\x74\xed\x5b\x62\x54\xda\x4f\xa9\x12\xca\xf8\x37\x90\x00\xac\x91\x2e\xdd\xee\xcd\xe6\xaa\x92\xb5\x89\x8c\x0c\x5a\xad\xa4\xc5\xa8\xb6\xb9\x54\x19\x0e\x62\xe9\x21\x17\x48\xb9\xca\xc6\x91\xdf\x55\xb2\x0e\xf2\x3f\x25\x5a\x8a\x6a\x93\x11\x50\x97\x2f\xd7\x05\x90\xec\x15\xbf\x05\x19\xb1\x3a\x9c\x49\xa6\x35\xca\x2c\x70\x90\x03\xbb\x6d\x67\xbf\x6d\x9b\x66\xbb\xec\x74\x79\xb7\x7b\x9a\xcf\x46\x3c\x0b\xfc\xd6\x9f\x1f\x46\x4a\x06\x7e\x2a\x78\xfa\xc5\x9f\x83\x1f\xad\x48\x9e\xa5\x06\x19\xe1\x59\xab\x32\x1f\xef\xbe\x37\x4a\x91\x25\xc3\xf4\xcf\x9c\x09\xb5\x8e\x6c\xae\xee\x83\x7e\x05\x89\x93\xc0\x18\xfc\xe6\x32\x42\x0d\xec\xcf\x7b\x4a\x05\x5a\xcb\xd6\x18\xbb\xb0\xce\x09\x1f\x88\x19\x64\xc0\xb3\xc4\x73\x9b\x41\x30\x1e\x58\xda\x08\x4c\xbc\x82\xcb\xb3\x7b\x9e\x51\x1e\xff\xf8\xd3\x6b\xfd\xb0\x74\xfb\x1c\xf9\x3a\xa7\xf8\x87\xd7\xaf\xf5\x83\x77\x71\xbe\x68\xc1\x2e\xfc\xb0\xef\x72\x55\x12\x29\x69\x63\xb8\xde\xef\x38\xc1\x56\x28\x62\xf0\x2f\x99\x4c\x51\x0c\x62\x75\x8b\x55\xa9\x88\x3b\x17\xa7\x3a\xff\x58\xf7\xba\x55\x4b\xa3\x54\x28\x8b\xc3\x3b\x51\xd5\xac\x5f\xd6\xf9\x08\x4c\x1b\xd2\x15\xbb\x1b\x26\xcf\xad\xd4\xda\x4b\xc1\xac\x8d\xc1\x77\xd5\xd3\x86\x17\xcc\x6c\x9e\x1d\xb9\xeb\xe5\x54\x49\x42\x49\x15\xdf\xf9\xaf\x46\xea\xe1\x87\xd1\x1d\x13\x63\xe7\x6a\x10\x4a\xc3\x21\x01\x47\x52\x91\x66\xc6\x62\xd0\x40\x86\xd1\xc7\xbf\x7e\x3d\x68\x85\x32\xd3\x8a\x57\x8e\xbf\x95\xda\x2a\x51\x69\xf8\xa8\x8f\x59\xc4\x3e\xb3\x87\x60\xbc\x62\x6e\xd1\x46\xbb\xc6\xfd\xf3\xfd\xd5\x87\x91\x44\xb6\xab\x34\x22\xde\xc5\x7a\x58\xcd\x71\x5e\xdc\xe6\xf2\xb0\x9a\x2d\xd3\x14\x5d\x19\x87\x94\xbc\xd7\x44\xbb\x0b\x1d\x84\xb0\x3d\x0c\x88\xc6\x28\x33\x02\xc7\x04\x1a\x0a\xbc\x5f\x6a\xb1\x77\x5a\x55\xc7\x92\xe1\x72\xcd\x6f\x37\xb5\x5a\xb8\xdc\x73\x7a\xc4\x53\x73\xb6\x0f\x55\xd6\x3c\xa6\xb5\xe0\x75\xb9\x16\x9f\xad\x92\xde\xf1\xd4\xd4\x56\xbe\xd3\xf4\x47\x15\xb7\x4f\x5f\x9b\x9b\x93\xa1\xf2\x8e\x14\x3b\x6c\x77\xe6\xdc\x8d\x52\xde\x1d\xc7\xfb\xa7\x08\xef\xb1\x9d\x67\x81\x63\xd8\x30\x72\x70\x81\x57\x1a\xee\x75\x22\x7c\xfe\x54\x6e\x5b\x77\xc2\x70\x7e\xbc\xa0\x23\x45\x9c\x83\x2c\x85\x98\x83\xf7\x37\x79\x61\xcf\x74\x02\x77\xc3\x08\x7f\x67\x78\xcb\x25\x77\x01\x8d\x5c\x8a\x1d\x8f\xfb\xe7\xda\xe0\x85\x0f\xa7\x60\x37\x92\xd8\xc3\x5b\xbe\xce\x85\x63\xe6\xdd\xdd\x87\x53\xf0\xcf\x17\x95\xd6\xd1\xa1\xb6\xab\xe3\xf4\x42\x62\xc6\xe9\x85\x0a\xf9\xdd\x88\xe8\xb1\x43\x76\xdc\xf1\x32\x95\xde\x2f\xb4\xb3\x9c\x39\xde\xae\x29\x7c\x64\xbe\x96\x3a\x7b\x81\xf9\x3a\x78\xeb\x39\x49\x35\x18\xda\x8a\xf7\xc4\x2f\xd1\x8a\xae\xce\xc7\x9a\xb0\x0a\x61\x5f\x7e\x6c\xf2\xc3\x94\xe9\x0f\x5f\xff\x02\x80\x09\xaf\x00\xd8\xa3\x34\x38\xf4\x1a\x80\xa7\x5f\x04\xf0\x35\xaf\x82\x6f\x3b\xd1\x81\xd7\x41\xbf\x9b\x8e\xbe\x0e\x60\xca\x1c\x86\xce\x2c\xfe\x78\x6c\x14\xc3\xf4\x71\x0c\xd3\x47\x32\x7c\x8f\xb1\x0c\xff\xeb\x68\x86\xe7\x8c\x67\x98\x3a\xa2\xe1\xc0\x98\x86\xf1\xbe\xbe\x99\x44\xfb\xdb\x93\xff\x02\x00\x00\xff\xff\xc4\x88\x71\x24\xcf\x0f\x00\x00")

func tmplJsMappingJsBytes() ([]byte, error) {
	return bindataRead(
		_tmplJsMappingJs,
		"tmpl/js/mapping.js",
	)
}

func tmplJsMappingJs() (*asset, error) {
	bytes, err := tmplJsMappingJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/js/mapping.js", size: 4047, mode: os.FileMode(420), modTime: time.Unix(1557154603, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplJsRequest_loggerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\x4d\x6f\xe3\x36\x10\xbd\xfb\x57\xcc\x6a\x17\x30\x59\xbb\x4a\xb6\xe8\x49\x86\x80\xa2\x69\x8a\xa0\x48\xb6\xe8\x26\x45\xd1\x53\x40\x4b\x13\x59\x58\x8a\x54\xc9\x61\xbb\xc6\xc2\xff\xbd\x20\xa9\x6f\x38\x1f\x3a\xe4\x30\x7e\xf3\xf8\xf8\xe6\x85\xf3\xe4\x54\x41\xb5\x56\xf0\x19\xff\x71\x68\xe9\x56\x57\x15\x1a\xc6\xe1\xdb\x0a\x00\xe0\xe2\xa2\xd0\x4d\x5b\x4b\x04\x3a\x20\x10\x36\xad\x14\x84\xe1\xa7\x7f\x85\x01\xab\x9d\x29\x10\x72\xf8\xc0\x92\xf7\x26\x32\x7c\x8f\x8a\xcc\x31\xe1\xe9\x81\x1a\xc9\xf8\x6e\x00\xf7\xdd\x90\xc3\x8d\x50\xa5\xc4\xbd\x30\x36\xed\xf8\x59\xa4\xe2\xbb\x55\xc0\xd3\xa1\xb6\xa9\x72\x0d\xe4\x70\x19\x2b\x83\x50\xd7\x96\x82\xf0\xa1\x26\x89\x83\x4c\xff\x7d\x60\xa5\x2e\x5c\x83\x8a\x78\x2a\x88\x0c\x4b\xc8\x63\x92\x2d\x24\x9f\xae\xff\x82\xcf\xd7\x7f\xfc\x79\x7d\xff\xf0\xee\x5d\xd2\x49\xf2\x9f\x45\x7a\xa8\x1b\xd4\x8e\xd8\xc0\xef\x49\x5f\x22\xbb\xbb\xd3\xc5\x17\xb8\xd2\xca\x6a\x89\x09\xdf\xc1\x69\x0b\x3f\x5c\x5e\x5e\x76\xb4\xa7\x85\xdc\x0a\xa9\xb3\xd6\x9f\xc4\xa8\x6e\xd0\x92\x68\xda\xa9\x74\xef\x8e\x19\x41\x90\x83\xc2\xff\xe0\x17\x41\x13\xfc\x77\x1f\xc7\x33\xfa\x1e\x6f\x04\xc5\x86\x49\x7b\x5a\x21\x85\x5e\x0e\x1b\x48\x2e\x12\xd8\x0c\x4d\xfe\x63\x0b\xe8\x9d\x56\x74\x08\xd8\x8f\xe7\x1b\x16\xf8\x5f\x9d\x94\x7f\xa3\x30\x91\x1e\x7e\x82\x57\xf0\x37\xda\x19\x1b\xc1\xd9\x2b\xd0\xbb\x5a\x39\xc2\x37\x82\xef\xb1\xd0\xaa\xb4\x6c\xe2\x89\x41\x72\x46\x0d\xb6\x9c\x1f\x48\xcc\xcf\xad\xb0\xfd\x5c\xe6\x3e\xcf\x23\xb5\x7e\x2f\x85\xa5\xc7\xd8\x53\xae\x79\x4a\xf8\x95\xd8\xb3\x33\x7d\x3e\x03\x57\x5a\x85\x56\xe5\x9a\xad\x17\x28\x96\xf3\xb7\x24\xc8\x59\xc8\xc3\x8f\xa9\x41\xdb\x6a\x65\x31\x8d\xe5\x2b\x5d\xe2\x7c\xf4\x0d\xd2\x41\x97\x23\x3c\xa8\x49\x63\x75\x8e\x6c\x05\x1d\x96\x38\x5f\xdb\x9d\x8b\x1f\xe4\x60\x8f\x8a\xc4\xd7\x9b\xba\x3a\xc8\xba\x3a\x10\xfb\xed\xfe\xf7\x4f\xa9\x25\x53\xab\xaa\x7e\x3a\xb2\x29\xcf\x16\x9c\x2a\xf1\xa9\x56\x58\x6e\xe1\x47\xce\x97\x94\xf1\x0e\x6f\xe6\x8c\xf0\x97\x49\xa5\xae\xde\xce\xe7\xe4\x2b\x12\x0b\x2d\xb5\x81\x3c\x0e\x48\x6a\xf3\xf3\xf1\x3e\xf8\xcd\xa2\xed\x7c\xb5\x8c\xd6\xb7\xde\xa8\x47\xe5\x9a\x0c\xc2\x34\xbb\x4a\x06\x83\x2b\xfd\x55\x32\x18\x2f\x65\xa4\xae\x32\x2f\x7f\x68\x78\xf4\x99\xca\x96\x0f\x44\xd0\xee\x13\xc5\x47\x60\xa1\x4b\xcc\xba\x84\x8c\xd5\x38\xec\xac\x8b\xc2\x58\xf7\xc3\xcd\xc2\xd8\xa7\x0c\x52\x9b\xac\xbb\xef\x69\x96\xd2\xf0\xd4\x4a\x5d\x5d\xfb\x67\x1b\xf2\x31\xb5\x6c\x19\xd3\xfe\x51\xde\x6c\x96\x2e\x86\x6c\xf7\x3e\xc6\xa0\xf7\xe0\x2e\xed\xf3\x0e\xbf\x1a\x20\x1f\x36\x02\xeb\x18\x26\xa8\xe9\x42\x21\xb1\xf7\x0b\x68\xaf\x4b\xbf\x56\x5a\x83\x2d\xaa\x92\x79\x8e\x49\xc3\x6c\x2d\x2c\xcb\xcb\xff\xf6\xd1\xe4\x6e\xdf\x9c\x76\xab\xd3\x6a\x35\x7d\xf5\x0d\x8a\xf2\x38\xdf\x0a\xab\x5e\xd8\x9e\xd4\x95\x44\x61\x6e\x75\x95\xf0\xb4\x90\x75\xf1\xe5\x0c\xf2\xe5\x6b\x60\xd3\xd2\xb1\x57\x7a\xea\x75\xf8\x17\xa7\xc7\x7b\x91\x6b\x9e\x6a\xc5\xd6\xe1\x88\xf5\x16\xd6\x1d\x87\x59\x6f\xe1\xfc\x81\xde\xdd\xba\x0c\x1b\xd9\x4f\x80\xa7\x9e\x85\x25\x24\x4c\x85\x94\x2c\x0c\x4e\x60\x03\x75\xc9\x53\xd2\x55\x35\xda\x16\xc4\xac\xfc\xdf\xff\x03\x00\x00\xff\xff\x02\x4a\x2c\xba\x1d\x08\x00\x00")

func tmplJsRequest_loggerJsBytes() ([]byte, error) {
	return bindataRead(
		_tmplJsRequest_loggerJs,
		"tmpl/js/request_logger.js",
	)
}

func tmplJsRequest_loggerJs() (*asset, error) {
	bytes, err := tmplJsRequest_loggerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/js/request_logger.js", size: 2077, mode: os.FileMode(420), modTime: time.Unix(1519645595, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplJsSortingJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\x41\x6f\xdb\x3c\x0c\xbd\xf7\x57\xf0\x10\xc0\xd2\xf7\x39\x42\x7b\x5d\xd7\x01\x69\xd7\xc3\x80\xa1\x97\xac\xbb\x2b\x32\x13\x6b\x95\x25\x4f\x62\xdc\x16\x45\xff\xfb\x20\xc7\x89\x65\x3b\x29\x72\x18\xa6\x4b\x90\x88\x7c\x8f\xef\x91\x54\xd6\x5b\xab\x48\x3b\x0b\x4b\xe7\x49\xdb\x0d\x23\xb9\x32\xb8\x44\x83\x8a\x9c\xcf\xa1\xfd\x7a\x6f\xb0\x42\x4b\x81\xc3\xdb\xc5\x05\x00\x80\x41\x82\xd0\xc5\x3c\xd6\x70\x03\x99\xd8\x98\xd7\xba\xd4\xca\xd9\xb9\x2a\xb1\xf1\xce\xce\xb7\x75\x76\x3d\x89\xfe\xea\x9e\xed\x89\xf8\xc2\x3d\xdb\x49\x06\x16\x77\x46\x86\xf0\x20\x2b\x8c\x69\xfb\x1f\xb3\xeb\x5d\x21\x87\xf2\x83\xf3\xf4\x23\xd6\xca\x9c\x2f\xd0\xe7\xa0\x09\xab\x70\xe7\xb6\x96\xf8\x5b\x1b\xba\xc7\xf5\xee\x39\xc0\x0d\xcc\x86\x4a\xe1\x7f\xc8\x80\x56\xae\x78\x05\x20\x9f\x71\xb1\x41\x62\xbc\x63\x89\x27\xa6\x89\xc8\xc2\xf6\x9c\x4c\xe6\xb0\x3a\x58\x92\x52\x2c\xe0\x06\x36\x48\x3f\xa5\x61\x92\x5f\x4f\xae\x6f\xfb\xeb\x55\x4a\x11\x8f\x5e\xb3\x05\x7c\x86\xdb\x88\x0b\xa3\xe3\x91\xb6\xde\xc2\xfc\xea\xbf\x56\xe3\x10\xf8\x7d\x0a\xf3\xe5\x43\x98\x33\x50\xba\xc8\xcb\x3e\xe6\x3d\xad\xf7\x60\x7e\x27\x06\x4d\xc5\xdf\x26\x6a\x9b\xd6\xec\x78\x27\x54\xa9\x4d\xe1\xd1\xb2\x8c\x8a\x8c\x0b\xfc\xcd\x92\x2e\x09\xc2\x17\x62\x5c\x90\x7b\xac\x6b\xf4\x77\x32\x20\x1b\x99\xa7\xd7\x6c\x26\x74\x78\xd8\x56\xe8\xb5\x62\x0d\xe7\x53\x75\x91\xae\x96\x3e\xe0\x37\x4b\xac\xc9\xaf\x2e\xf9\x19\x0a\x9b\x44\x61\xaf\x6f\x26\x50\xaa\x92\xc5\xce\xe7\x07\xb1\x4c\xdb\x02\x5f\xf2\x38\x0f\x63\x73\x47\x23\x35\xd0\x1b\x27\x2b\xe3\x42\xd6\x35\xda\x22\x42\xf2\xa1\xa9\x09\xf5\xc1\x56\x49\x24\x55\xf9\x5d\x07\x42\x8b\xbe\xf5\x2a\x87\x8e\x1e\xf7\x2b\xd9\x97\x30\x6b\x23\xb8\x58\x6b\x5b\xb0\x7e\x3b\xb9\x70\x96\x65\xca\x68\xf5\x94\x25\x32\x8e\x74\xaa\x6f\x46\xdb\xb2\x1d\x5a\xed\xb1\x59\x18\xc3\xb8\x30\x68\x37\x54\x0e\xdd\xec\xd7\x6e\x7e\x05\x79\xd2\xcd\x61\x58\x25\xfd\xd3\x22\x2c\xbb\xe5\x65\x33\x46\xa5\x0e\xfc\x88\x05\x27\x75\xc4\x77\xe3\x1f\x29\xf9\x2b\x42\xc6\xbd\x1c\x25\xba\xd5\x2f\x54\x94\x0e\x90\x32\x28\xfd\x32\x7d\xf3\xd2\xe9\xdf\xc5\x0b\x59\x74\x57\x93\xc7\xf1\x04\xed\x31\x54\x48\x47\x26\x9d\x13\x8f\x95\x6b\xf0\x63\x82\x61\xd6\xae\x2b\xe7\xe4\x8d\x0b\xd3\x56\xd3\xa0\x94\xc1\x1f\x8d\x58\x3b\x7f\x1f\x77\x6f\xb8\x02\x43\xac\x6e\x77\x7b\x88\x88\x09\x9f\xda\x8f\x2e\xee\xfd\x4f\x00\x00\x00\xff\xff\x8f\xda\xb9\x25\xdf\x06\x00\x00")

func tmplJsSortingJsBytes() ([]byte, error) {
	return bindataRead(
		_tmplJsSortingJs,
		"tmpl/js/sorting.js",
	)
}

func tmplJsSortingJs() (*asset, error) {
	bytes, err := tmplJsSortingJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/js/sorting.js", size: 1759, mode: os.FileMode(420), modTime: time.Unix(1507114595, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplJsUtilJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\x51\xaf\x93\x30\x14\xc7\xdf\xf9\x14\x27\xc4\x5c\xca\x26\x63\x9a\xbd\xb8\x8d\x11\xbd\x12\x7d\xd0\x68\x72\xaf\x2f\x52\x96\x74\xd0\x31\xb4\x2b\x4b\x5b\xd4\x79\xd9\x77\x37\x65\x64\x63\x5c\x70\x3e\x78\xfb\x40\xd3\xff\x39\xfd\x9d\x3f\xe7\x14\x00\x60\x5d\xf0\x58\x65\x39\x87\x94\xaa\xdb\x9c\xe5\xe2\xcd\xfe\x4e\x11\x55\x48\x24\xab\xed\x36\x4f\xa8\x0d\x0f\x06\xd4\x2b\x5b\x43\x23\x02\x9e\xe7\xc1\xcb\xf1\x18\xca\x12\x1e\xa9\x2f\x9a\xf7\xf4\x12\x54\x15\x82\x83\x29\x8b\x38\xa6\x52\x9a\xb3\x53\xf4\x00\x94\x49\xda\x05\x9f\x8c\x27\x7d\x98\x84\xf0\x94\x8a\xc7\x94\xee\xec\x9f\x44\xf0\x8c\xa7\xcd\x74\xe3\xf8\x35\x7a\xfa\xf0\x91\xaa\x4d\x9e\xa0\x6d\xb5\xb5\x7b\x70\x54\x2b\x8b\xd6\xbb\xe0\xde\xea\x33\x99\xf1\x75\xde\xfd\xa3\x4d\xc2\xe7\x4f\x77\xfd\x88\xbf\xb6\xeb\x82\xf2\xe5\x3f\x40\xde\x06\x1f\x82\xfb\xa0\x97\xf3\x34\x4d\x97\x7b\xae\xc8\xaf\xf7\x59\xba\x61\x59\xba\x51\xe8\x9b\xcc\x79\xd3\x81\x3e\x83\x57\x6d\x23\x41\x77\x8c\xc4\x14\xb9\x37\x6e\xfa\x1c\xac\x1b\xb2\xdd\xcd\x2c\xfb\x2c\xcf\x8f\x32\x53\x17\xea\xe2\xa8\xa6\x5a\x3d\xbb\xa9\x7d\x5e\x72\x91\x89\x30\x2e\x42\xe2\xfc\x7e\xed\x7c\x1d\x3b\xaf\xa2\x87\xc9\xa1\xc4\x38\x5c\x16\x51\x19\x2e\x31\x36\x23\x7b\x60\x22\x2c\x07\x53\xdb\x2f\xf1\x0a\x29\x51\xd0\x72\x4d\x98\xa4\x25\x2f\x18\xb3\xf1\xaa\x74\x7c\x9c\x0c\x91\x3f\xc5\x23\x9c\x0c\x6c\x1f\xf9\xd3\x90\x06\x51\x38\xc4\x4e\xa4\x23\xb6\x6f\x6b\x3b\xa7\xdf\x47\x5b\xa2\xe2\x4d\xbb\xe5\x3f\x88\x80\x98\x49\xf0\xc0\xe2\xc5\x76\x45\x85\x35\xbb\x88\xeb\xc9\xb9\x4b\xd3\x1d\x29\x2a\x55\x8d\x68\x33\x4e\x79\xd3\x67\x57\xf3\xf4\xaa\xeb\x7d\xa7\xfb\x56\xb1\xde\x51\xb7\x6e\x4a\x25\x32\x9e\x76\x5d\x36\x3a\x50\x95\xb5\x73\xff\xae\x5a\xac\x8b\xac\xf2\x9c\x51\xc2\x5b\x55\x9a\x4c\x3d\x88\x7f\xa5\xe9\xdc\x36\xaa\xeb\x35\x5b\x73\xb9\x23\x1c\x62\x46\xa4\xf4\x4c\x0b\x86\x15\x60\x08\x96\xb9\xd0\x87\xaa\x8e\x3e\xce\x5d\x9d\xb7\x68\x20\x0f\xf5\x93\x3b\x18\x7f\x02\x00\x00\xff\xff\x60\x64\x0c\xd1\x73\x05\x00\x00")

func tmplJsUtilJsBytes() ([]byte, error) {
	return bindataRead(
		_tmplJsUtilJs,
		"tmpl/js/util.js",
	)
}

func tmplJsUtilJs() (*asset, error) {
	bytes, err := tmplJsUtilJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/js/util.js", size: 1395, mode: os.FileMode(420), modTime: time.Unix(1519645595, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\xdb\x6e\xdb\x46\x13\xbe\xf7\x53\x0c\xf8\xff\x17\x2d\x60\x98\xae\x7b\x13\xf8\x2e\x27\x14\x46\xed\x22\x95\x93\x02\x45\x10\x04\x23\x72\x24\x6e\xbc\xe4\x32\xbb\x43\xbb\x6c\xa0\x77\x2f\x96\xa4\x78\xd2\x4a\x22\x75\x70\x8c\xc0\x57\xb6\xf6\x30\x3b\x33\xdf\x37\x33\x3b\xcb\x6f\x27\x00\x9e\x79\xc0\xf9\x9c\xb4\x77\x09\xde\xc5\xd9\xb9\x77\x6a\xc7\x44\x32\x53\xde\x25\xd8\x79\x00\xef\x9e\xb4\x11\x2a\xb1\x2b\x7e\xa9\x56\x00\x78\x2c\x58\x92\x1d\xbb\xb9\x51\xc1\xdd\x72\x34\x24\x13\x68\x91\x72\xb5\xfe\x26\x56\xc1\x1d\x08\x03\x08\x4c\x86\x45\x32\x07\x4c\x42\x98\xa1\x61\x48\xb5\x62\xc5\x79\x6a\x07\x59\x29\x09\x33\xa5\x21\xa4\x7b\x92\x2a\x25\x6d\x96\x12\xa5\x08\x28\x31\x54\xab\x03\xe0\x25\x18\x97\x27\x5f\xbd\xaf\x56\x01\x78\x99\x96\x76\x2c\x62\x4e\xcd\xa5\xef\xcf\x05\x47\xd9\xf4\x2c\x50\xb1\xff\x25\x46\xcd\x22\x79\x71\xe1\xc7\x56\x1d\x7f\x2a\xd5\xd4\x8f\xd1\x30\x69\xff\xfa\xea\xf5\xdb\x3f\x6e\xdf\x9e\xc5\xa1\x57\x08\x5a\x9c\x00\x2c\x0a\x1f\x44\xca\xb0\x15\x28\x55\x80\xd2\xfe\xb8\x7c\x71\xfe\xe2\xa2\xf4\x0f\xe3\xdc\x78\x97\xf0\xb1\xd8\xb2\xa2\x96\xa6\xaf\x19\x19\x6e\x54\xeb\x39\xe5\x7d\x44\x50\x3a\x46\x53\xa0\x74\x68\x00\xa5\x84\x6a\x97\x01\xc1\x76\x9c\xc4\x3d\x19\x10\x09\xc4\x14\x2b\x9d\xc3\x4f\xc8\x20\xc9\xfa\x2d\x4b\x58\x48\xbb\x4a\x18\xd0\x64\x88\x7f\x3e\x83\xf7\x91\x30\x10\xe3\x1d\x15\xdb\x53\x65\x8c\x98\x4a\x02\x56\x70\x4f\x5a\xcc\x72\xe0\x08\x19\x70\x79\x08\xc4\xc8\x41\x54\xa0\x01\x26\xa5\x40\xcc\x44\x00\x29\x32\x93\x4e\xe0\x01\xcd\x52\x81\xf0\xb4\x80\x0b\xa5\x51\x56\xd6\x8c\x38\x88\x80\x23\x6a\x94\x0d\x89\x51\x48\x73\x56\x79\xef\xd4\xed\x11\x13\x50\x82\x5a\xa8\xb5\x2e\x79\x09\xcb\x25\xd6\x2a\x32\x86\x12\x16\x28\x65\x6e\x15\x64\x64\x82\x18\xad\xc2\x04\x0f\x91\x32\x54\x8e\x19\x08\x30\x81\x29\x01\xea\xa9\x60\x8d\x5a\xd8\xf5\xc6\x88\x79\x42\xe1\x16\x8d\x62\x4c\x2d\xf1\xd6\x2a\xf4\xb7\xca\x0a\xe9\x31\x26\x38\xb7\xf6\xc6\x8a\x49\xe6\x90\xab\x4c\x83\xe1\x6c\x0a\x95\x04\x03\x0f\x11\x25\x74\x4f\xda\xce\x41\x42\xd4\x62\xd2\xa7\x82\x2d\x53\x34\xf4\x0e\x39\xb2\x72\x7d\x4c\x45\xc9\x21\x13\x44\x14\x53\x43\xa3\x82\xb9\x5e\xbd\x29\x50\x89\xc9\x3a\xf3\x98\xa6\x52\x04\x68\x35\xf4\xbf\x18\x95\x34\x6b\x53\xad\xc2\x2c\x18\xb8\x16\x39\x32\x4d\x6c\xfb\x15\x92\x3e\x4a\xd9\x0e\xb1\x39\x71\xeb\xa7\x55\x37\x8b\x63\xd4\xb9\xb5\xe1\x37\xe2\x0e\x63\x6b\x27\xf6\x23\xa3\x1a\x5b\xc6\x43\x3d\xf6\xa9\xb5\xc1\xc6\x7a\xa1\xe7\x55\xb8\x4d\xf6\x8a\x99\xd5\xb8\xcb\xd8\xd5\x83\x52\xd4\x18\x13\xdb\xcc\x72\x09\x1f\xdb\x33\x9a\x4c\xaa\x12\x43\xa6\x63\x31\x80\x77\x71\x7e\xde\x1b\x72\x30\x57\x6b\xcc\x41\xcd\xca\x90\x6a\xa9\x5b\x7a\xcd\x82\x8c\x2b\x42\x00\xbc\xff\x6b\x9a\xd9\xfd\xff\xf3\x43\x9a\x89\x44\x58\x79\xc6\xb7\xa9\xf4\xc6\x0a\xba\x16\x6d\x87\x2d\xf9\xb4\xfa\xff\xf2\xbf\x45\x9b\xed\x0d\xaa\x85\x52\x14\x8e\x45\xb6\xda\x76\x54\x84\x37\x9d\xf1\x8c\xf4\x68\xa4\x8b\x4a\x30\x14\xe7\x89\x5d\x0c\x86\x95\x3e\x12\xc8\x5b\x0f\x78\x92\x08\xef\x8b\xe9\xcb\xc0\xfe\x9d\x54\x2a\x1c\x0a\xd4\xcf\x25\xdf\xf6\x81\xb6\x2c\xff\x65\xc4\x99\xa6\xe2\x07\x5a\x30\x69\x81\x47\x47\x1e\x1e\x04\x47\xce\xe3\x9e\x79\x30\x84\x07\x59\xb2\x63\x22\xaf\x37\x1e\x35\x95\x6f\x3e\xe5\x49\x42\xfc\xbd\x93\x79\xcb\x0e\x43\x41\xa6\x05\xe7\x85\x7d\x9b\x79\x50\x5e\xe6\xdb\x24\x48\xcb\x2e\xc5\xcd\x82\xbf\xca\xbb\xff\x31\xa0\xdf\x20\xfa\x90\x78\xb7\xf6\xf7\xf0\x5c\xd7\x6c\x55\xd3\xa2\x40\x79\xaa\xc2\xbc\x3f\x63\x37\x08\x5d\x84\x12\xeb\x8c\xf6\x44\x7e\xd2\xf7\x19\xac\xc5\xfc\x87\xe1\xae\x9b\xa3\xcb\x06\xce\x54\x55\x6b\x44\x43\x51\xd6\x0c\x9b\x49\xaa\xba\x51\xcb\xda\x46\xd9\xba\xb1\x1c\x5e\x9a\xb6\x1d\xf3\x24\xd3\xd5\x93\xa9\x48\xcb\xbe\x79\x6c\x43\xa1\x82\xbb\xba\x63\xde\x06\xea\xf2\x8c\x31\x25\x68\xed\x01\x4f\x12\x4e\x1b\x73\xd0\xe0\x03\x52\xac\x66\xb0\x5d\xe2\xf8\x4d\xfd\x7b\xcf\x60\xae\x3c\xe9\x7f\xb3\x7e\xfd\x9c\x22\x47\x8b\x31\x90\xc7\x5d\xfb\x0e\x0d\xf8\x06\xf1\x8f\x5b\x7c\x6a\xef\xb8\xcb\x8f\x6b\xc6\xc5\x04\xd7\x3a\xce\xd3\xf2\xe9\x8c\x75\xfb\x9d\xaa\x9a\xed\x15\xb1\xe3\xd6\x9b\xe3\xb3\x75\x14\x53\xeb\x27\xbb\x34\x5b\xcf\xc2\x0f\x69\x58\xbc\x1a\x1e\x8f\x88\xdb\x4f\x78\xe6\xe2\xe9\x76\xb5\x5f\x39\xae\x68\xbb\x5e\xde\x0e\x5c\x34\xc7\x73\xf4\x70\x57\xbc\xa2\x5e\x83\x26\x93\xc9\xbd\x63\x6d\x8f\xe2\xdf\x7a\x1e\x97\xc4\xb4\x36\xdc\xde\x14\xd3\xc7\x0c\xb7\xed\x27\x3c\x87\xdb\x8f\xcb\xc3\x13\xe8\x7c\x9b\x6b\x49\x6d\x3e\x65\x38\x2e\x40\xad\x2b\x4b\xf3\xb9\xd2\xba\xfd\xba\x5d\xc2\x6a\x87\xa3\x6d\xac\x9a\x61\xc1\x14\x77\xbd\x36\x2a\x51\x74\xaf\x54\xdd\x36\xcb\xa9\x98\x9d\xfd\x4e\x9a\xb5\xd6\xac\xf5\x99\x23\xea\x6a\xfd\xd4\xf4\x0b\x05\x2d\xbd\x53\x6d\xc3\x97\x45\x8f\x75\xde\x87\xc9\x55\x9f\x86\x9b\x92\x76\x3f\x14\x1a\x62\xb4\x08\xde\x15\x70\x58\xd9\xcb\xd7\x85\x9e\xdc\x81\x6f\x01\x0b\x47\x18\x0e\x15\xd5\x0f\x92\xb6\xac\x40\x25\xac\x95\x1c\x26\xea\x75\xb5\x78\x43\x72\x6f\xa5\x94\x8f\x5d\xac\xd6\xb8\xd9\xe1\x21\x97\xa9\x0e\x95\xab\x91\x4f\x1d\x06\x4e\x56\xbc\xdc\x50\x6f\xd2\xff\x98\x3e\x96\x72\x51\xff\x85\xec\x10\xbc\x88\x89\x23\x15\x1e\x5c\x6c\x5a\x7e\x25\x3e\xac\xd0\xaf\x19\xe9\xfc\xb6\x58\xf0\xae\x5d\xe4\x06\x70\xe7\x4f\xe7\x56\xe7\x29\x11\x61\x38\x58\xee\xef\x94\x17\x8f\x58\xeb\xe8\xad\xee\xfa\x30\x6e\xa0\xb7\x5d\xec\x16\x54\xdc\x20\x0f\xe2\xce\x81\x21\x53\xd1\xa2\x8f\xa8\x93\xf4\x6e\xdf\xb6\x42\x60\x88\xeb\x9a\x08\x5a\xc9\x2e\xa3\x03\xc5\x30\x72\x66\x5e\xab\x70\x25\x47\x0d\xf2\x99\x48\x98\xe6\xa4\xbb\x93\x33\xa5\x63\xe4\x6a\xfa\xd7\x8b\x67\xea\xac\xa5\x4e\xcb\xfb\x4e\xba\xd4\x76\xef\x81\xf0\x1d\xed\x66\x53\xf7\xf6\x01\xee\x1b\xc8\x46\x2f\x6c\x7f\x72\xaa\xc0\x18\x60\xdd\x3a\x95\xc7\x98\x3d\x24\x6f\xde\xa3\xcc\x56\x42\x61\x2b\xcc\x3d\xa3\xfa\x85\x7a\x34\x66\xa9\x16\xaa\xfa\x3a\xf4\x98\x31\x19\x92\xc4\xc7\x3e\x33\xd0\xf8\xef\x6e\x67\x4e\x95\x92\x84\x89\x5b\x6c\xfd\x8d\x60\x50\x52\xb8\x5d\xf9\xa2\xb0\xe8\xf6\x98\xff\xe4\xaf\xd0\xd0\x87\xc9\xf5\x61\xd3\x43\x87\x37\xb7\xab\x2a\x8f\x26\x4e\xd5\xa8\x1e\xfe\x3e\x6c\x73\xd7\x2d\x23\xef\x26\xfd\xc0\xb9\xa4\xa5\x5b\x42\x0f\xbb\xab\x35\x0c\x97\x5e\x07\xbb\x07\x3a\x55\x47\x7d\x04\x55\x6d\x93\x7c\xb2\xf8\x2f\x00\x00\xff\xff\x5b\x28\x12\x0f\xdb\x2b\x00\x00")

func tmplSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_tmplSwaggerJson,
		"tmpl/swagger.json",
	)
}

func tmplSwaggerJson() (*asset, error) {
	bytes, err := tmplSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/swagger.json", size: 11227, mode: os.FileMode(420), modTime: time.Unix(1556529826, 0)}
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
	"tmpl/css/style.css": tmplCssStyleCss,
	"tmpl/index.html": tmplIndexHtml,
	"tmpl/js/mapping.js": tmplJsMappingJs,
	"tmpl/js/request_logger.js": tmplJsRequest_loggerJs,
	"tmpl/js/sorting.js": tmplJsSortingJs,
	"tmpl/js/util.js": tmplJsUtilJs,
	"tmpl/swagger.json": tmplSwaggerJson,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"style.css": &bintree{tmplCssStyleCss, map[string]*bintree{}},
		}},
		"index.html": &bintree{tmplIndexHtml, map[string]*bintree{}},
		"js": &bintree{nil, map[string]*bintree{
			"mapping.js": &bintree{tmplJsMappingJs, map[string]*bintree{}},
			"request_logger.js": &bintree{tmplJsRequest_loggerJs, map[string]*bintree{}},
			"sorting.js": &bintree{tmplJsSortingJs, map[string]*bintree{}},
			"util.js": &bintree{tmplJsUtilJs, map[string]*bintree{}},
		}},
		"swagger.json": &bintree{tmplSwaggerJson, map[string]*bintree{}},
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
