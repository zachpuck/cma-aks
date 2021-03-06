// Code generated by go-bindata.
// sources:
// api/api.proto
// DO NOT EDIT!

package protobuf

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _apiProto = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x3a\x4d\x73\xdb\x38\xb2\x77\xfd\x8a\x2e\x5d\x9e\xf3\x2a\x91\x12\x27\xf3\xde\x94\xbd\xde\x5a\xad\xe4\x49\x54\x4e\x64\x97\xe9\x64\x6a\x4e\x2a\x88\x6c\x51\x58\x91\x00\x07\x00\xa5\x28\x29\xff\xf7\x2d\x7c\x90\x02\x48\x4a\x51\x32\xde\xda\x19\x1d\x6c\x91\xe8\x6e\xf4\x37\xba\x1b\x1a\x0e\x61\xcc\x8b\x9d\xa0\xe9\x4a\xc1\xf9\xcb\x57\x3f\x43\x44\x72\x59\xb2\x14\xa2\x49\x04\xe3\x8c\x97\x09\xcc\x88\xa2\x1b\x84\x31\xcf\x8b\x52\x51\x96\xc2\x03\x92\x1c\x48\xa9\x56\x5c\xc8\x41\x6f\x38\xec\x0d\x87\xf0\x9e\xc6\xc8\x24\x26\x50\xb2\x04\x05\xa8\x15\xc2\xa8\x20\xf1\x0a\xab\x95\xe7\xf0\x09\x85\xa4\x9c\xc1\xf9\xe0\x25\x9c\x69\x80\xbe\x5b\xea\x3f\xbb\xd4\x24\x76\xbc\x84\x9c\xec\x80\x71\x05\xa5\x44\x50\x2b\x2a\x61\x49\x33\x04\xfc\x1c\x63\xa1\x80\x32\x88\x79\x5e\x64\x94\xb0\x18\x61\x4b\xd5\xca\xec\xe3\xa8\x68\x4e\xe0\x37\x47\x83\x2f\x14\xa1\x0c\x08\xc4\xbc\xd8\x01\x5f\xfa\x80\x40\x94\x63\x5a\x7f\x56\x4a\x15\x17\xc3\xe1\x76\xbb\x1d\x10\xc3\xf0\x80\x8b\x74\x98\x59\x50\x39\x7c\x3f\x1d\x5f\xcf\xa2\xeb\x17\xe7\x83\x97\x0e\xe9\x23\xcb\x50\x4a\x10\xf8\x7b\x49\x05\x26\xb0\xd8\x01\x29\x8a\x8c\xc6\x64\x91\x21\x64\x64\x0b\x5c\x00\x49\x05\x62\x02\x8a\x6b\xa6\xb7\x82\x6a\xbd\x3d\x07\xc9\x97\x6a\x4b\x04\x6a\x32\x09\x95\x4a\xd0\x45\xa9\x02\x9d\x55\x2c\x52\x19\x00\x70\x06\x84\x41\x7f\x14\xc1\x34\xea\xc3\x3f\x47\xd1\x34\x7a\xae\x89\xfc\x3a\x7d\x78\x77\xfb\xf1\x01\x7e\x1d\xdd\xdf\x8f\x66\x0f\xd3\xeb\x08\x6e\xef\x61\x7c\x3b\x9b\x4c\x1f\xa6\xb7\xb3\x08\x6e\x7f\x81\xd1\xec\x37\xb8\x99\xce\x26\xcf\x01\xa9\x5a\xa1\x00\xfc\x5c\x08\x2d\x01\x17\x40\xb5\x36\x31\x31\xaa\x8b\x10\x03\x16\x96\xdc\xb2\x24\x0b\x8c\xe9\x92\xc6\x90\x11\x96\x96\x24\x45\x48\xf9\x06\x05\xd3\x9e\x50\xa0\xc8\xa9\xd4\x56\x95\x40\x58\xa2\xc9\x64\x34\xa7\x8a\x28\xf3\xaa\x25\xd7\xa0\xa7\x41\x3e\xd0\x78\x45\x30\x83\x4f\xc8\xf0\x0b\x25\xf0\xb7\x7c\x63\xbf\xfd\x23\xcd\x09\xcd\x06\x31\xcf\xff\xde\xeb\xc9\x1d\x53\xe4\x33\x5c\x41\xbf\x10\x5c\xf1\xd7\xfd\xcb\x5e\xaf\x20\xf1\x5a\x73\x10\xe7\x84\xac\xe5\x65\xaf\x47\xf3\x82\x0b\x05\xfd\x94\xf3\x34\xc3\x21\x29\xe8\x90\x30\xc6\x1d\x03\x03\x83\xd9\xbf\xac\xc1\xcc\x73\xfc\x22\x45\xf6\x42\x6e\x49\x9a\xa2\x18\xf2\xc2\x80\x76\xa2\xf5\x7a\x76\x15\xce\x52\x51\xc4\x83\x94\x28\xdc\x92\x9d\x5d\x8e\xe7\x29\xb2\xb9\xa3\x32\x70\x54\x06\xbc\x40\x46\x0a\xba\x39\xaf\x56\x9e\xc1\x15\x7c\xed\x01\x50\xb6\xe4\x17\xe6\x1b\x80\xa2\x2a\xc3\x0b\xe8\x8f\xb3\x52\x2a\x14\xf0\x81\x30\x92\xa2\x80\xd1\x4d\x04\xef\x30\x2b\xf4\xd7\xbb\x69\xff\xd2\x00\x6f\x6c\xd4\x5c\x40\x7f\xf3\x72\xf0\x6a\xf0\xd2\xbd\x8e\x39\x53\x24\x56\x15\x49\xfd\x61\x24\xd7\x54\x1b\xda\x75\xf0\xfa\x53\x8a\xec\x02\xfa\xda\xe1\xe5\xc5\x70\x98\x52\xb5\x2a\x17\x5a\xd9\x43\x69\x63\xfe\x45\xcc\x62\x35\x8c\x73\xf2\x82\xac\xa5\x87\x87\xda\x2a\x17\xd0\x6f\x9b\xc9\x01\x3d\xea\x7f\xe6\x0f\x7e\x56\x28\x18\xc9\xe6\x09\x8f\x65\xc5\xdc\x8f\xec\x9b\xa0\x8c\x05\x35\x5a\xd5\x32\x71\x81\x40\x16\xbc\x54\x70\x82\xd2\x1e\x7b\x00\x32\x5e\x61\x8e\xf2\x02\xde\x3d\x3c\xdc\x45\x97\xcd\x37\xfa\x45\xcc\x99\x2c\xcd\x9b\xbe\x0b\x60\xbd\xdb\xf0\x5f\x92\x33\x43\xa6\x10\x3c\x29\xe3\x43\xeb\x8f\x97\xbd\x9e\x44\xb1\xa1\x31\xd6\x3c\x59\x71\x75\x5c\xd2\x2c\xd3\xf8\x1b\x6a\x32\x1e\x81\xd8\x42\x98\x75\x51\xc4\x30\x16\x48\x14\x56\x78\x67\xc1\xe3\x07\x99\x3e\x03\x81\xaa\x14\x4c\x36\x96\xee\xb1\xc8\x76\xcf\x3c\x9b\xd7\xfe\x69\xfc\x7f\x40\x0a\x3a\xd0\x7a\xae\xbc\x6e\xff\x29\x4a\x05\x17\xd0\x37\x11\xb2\x79\x35\x74\xfc\xf4\x03\x98\x05\x4f\x76\x1a\xe8\x7f\xf7\xaf\x1f\x9d\x81\x03\xc1\x04\x2a\x41\x71\x63\xd3\x85\x54\x44\x95\x52\xa7\xd8\x5a\x4a\x9d\x0a\x80\x2a\x09\xeb\x72\x81\x31\x67\x4b\x9a\x9a\x6c\x12\x73\xc6\x30\x56\x74\x43\xd5\xae\xd6\xc4\x5b\x54\xb5\x1a\xf6\xdf\x43\x1d\xec\xdf\xff\xb8\x02\x52\x3c\xae\x80\x4e\x49\x13\xcc\x50\x61\x87\xfd\x26\x66\xa1\x66\x3c\x78\x0c\x79\x0f\x96\x7e\x9c\x7d\xc7\xc9\x77\x4b\x50\xdb\x8a\x40\x46\xa5\xd2\x76\x72\x88\xb2\xc3\x04\xef\x35\xc8\x59\xf8\x7c\xc8\x14\x7a\xed\xa9\xcd\x31\xd4\x3c\x7e\x5b\xa2\x52\xb0\x2a\x2f\x9a\xc4\x2a\x72\x13\x9a\x2e\x43\x90\x82\x82\x8e\x4c\xcf\x5c\x6f\x51\xb9\xea\x63\xea\x81\x9f\xed\x5f\xb7\x84\x74\xef\x9f\x4c\x40\xc7\xee\xc9\xd6\xda\x10\x9a\x99\x82\xa2\x2c\x52\x41\x12\x0c\x23\xac\xc3\x74\x1f\x2b\xb8\xb3\xf6\xbb\x43\x26\xac\xd6\x9f\xdc\x8c\x8e\xe9\x6f\x48\xeb\xa0\x3a\xc4\x72\x8c\xd5\xf1\x15\x3e\x87\xe2\x84\x6b\x4f\x9d\x21\xdb\xa2\xc0\x8f\x64\x4a\xc6\x13\x84\x98\x97\x4c\x75\x98\x6e\xc6\x13\x1c\xeb\x35\xdf\x36\xf5\xcb\x43\xc6\xab\x01\x9e\xdc\x7a\x32\x26\xd9\xb7\x6c\x67\x60\xf6\x96\x93\x5d\x12\x46\x1a\xa6\xb6\xa1\xff\x14\xca\xe4\xaf\x3c\xb9\xfd\x1a\xc2\x54\xd6\x3b\xc1\x78\xc8\x4c\x00\xd6\x87\x5a\xa9\xb8\x26\x46\x59\x5a\x4b\x78\x6d\x40\x1c\xeb\xa3\x3d\x00\x9c\x1d\x5a\x09\x25\x3f\x04\xf5\xe4\x5a\xf0\x98\x3f\xdd\x93\x1f\x7b\x3d\x64\x65\x5e\x95\x37\x91\x3d\xe9\xeb\x22\x67\xc6\x15\x48\xb4\xe6\x8e\x1e\x46\x0f\x1f\xa3\xf9\xc7\x59\x74\x77\x3d\x9e\xfe\x32\xbd\x9e\xc0\x15\xbc\xbc\xac\x40\x1f\x56\x08\x77\xf7\xb7\x9f\xa6\xd1\xf4\x76\x36\x9d\xbd\x35\x55\x03\x02\x65\x89\xae\xaa\x50\x9a\x4a\xa2\x52\x34\x95\xb0\x40\xad\xc4\xd8\x94\x3e\xc9\xc0\x50\x09\xd0\xaf\xe0\x55\x40\xfb\xfe\xe3\xec\x9b\x64\x57\x44\xd3\x45\x56\x91\xb5\x55\x8a\x84\x65\x99\x65\x3b\x28\xa5\xb6\x84\xdd\xaa\xa2\x76\x05\xe7\xe1\x2e\xd7\xe3\xdb\xd9\x78\xfa\xbe\x7b\x27\xa2\x40\xf2\x1c\x61\xcb\xc5\x5a\xd3\x25\xba\xd2\xc1\x6c\xe7\x84\x49\x38\x43\xdd\xc7\x79\x2c\x3d\x07\x59\xc6\x2b\x20\xd2\x65\x42\x0d\xa6\x97\x73\x62\x18\xe6\xc2\x06\x55\xd5\x35\x3a\xe6\x3c\x26\xae\xe0\x75\xc0\x60\xf4\x70\x7b\x77\x77\xb2\x7a\x6d\x49\x91\x38\xfb\x39\xcc\x2b\x78\x13\x90\xbc\xbe\xbf\xbf\xbd\x3f\x4a\x4f\xb7\xdb\x0b\x84\x92\x59\x15\x1a\x64\x8b\x75\x05\x3f\x05\xb4\x26\xd7\x6f\xef\x47\x93\xeb\xc9\x51\x72\xae\xaf\x96\x50\x4a\x1d\x74\xb1\xf1\x7b\xc5\x41\xa0\x54\xba\x0b\xd0\xe6\x82\x65\xc9\xcc\x02\xc9\xaa\x52\xb2\xa6\x7d\x05\xff\x77\xa9\x3d\x37\x47\x29\x75\xab\xd8\x2c\xad\x3d\xff\x25\x39\x56\xa3\x81\x6a\x77\xc5\xb5\x2c\x75\xe5\xee\xb4\xa3\x1b\x71\x96\x9a\x06\xab\xe5\x7a\x06\x56\xb7\xb9\x55\x9f\x6c\x6a\x0b\x03\x11\x6c\x7d\xe7\xe0\xa2\x02\xe3\x3d\x92\x71\xb1\x43\xdc\x9a\x2c\xe0\x35\x15\x2b\x34\xfd\xbb\xf1\x0b\x15\xb0\xbd\x25\xd2\x67\x1a\x16\x3b\x3b\x3d\xd1\xca\x44\x69\x43\x74\xc1\x79\x06\x7c\xdd\x12\x20\x41\x45\x68\x26\x9b\x9a\x70\xa8\x5a\xef\x05\x67\xd2\xda\xd5\x31\x36\x55\x98\xd7\x80\x4d\x11\x82\x22\xfe\x14\x6d\x67\x9c\xaf\x31\x81\xb2\x38\xaa\xeb\xb1\xc0\x04\x99\xa2\x24\x93\x1a\xef\xf7\x12\xc5\xae\x9e\x51\xf8\xf5\xc3\xe8\x4b\x29\xd0\x87\x8e\xbd\xef\x87\x99\x6d\x28\x7b\xda\x08\x1a\x1b\xb8\x72\x27\x15\xe6\x6d\x75\xfa\xca\x99\x18\x7d\x1e\x55\x51\xb3\x5d\xf0\x6d\x4c\x94\x8e\x50\x6f\xef\xff\x91\x56\x19\x8a\xeb\xa6\x58\x09\xbe\xfb\x1e\x3d\xb9\xbe\xe1\x0f\x29\xa9\xdd\xc2\xec\x19\x1e\xf3\x32\x4b\x02\x55\x2d\xb0\xe2\xd3\x45\x4f\x97\xe3\x8d\x92\x84\xda\x00\xee\x28\xe0\xc3\xb6\xb2\xc9\xba\x13\xdc\x02\xcc\x2b\x1e\xfd\x5c\x1d\x05\xb8\x7e\x0c\x84\x27\x99\xdb\xc3\x64\xd1\x4e\x9f\x70\xed\x8f\x27\x6d\xa8\x5c\x89\x44\xc4\x2b\x90\xe5\xa2\x1e\x56\xd8\x46\xd7\xef\xb0\x7e\xd0\x1f\xeb\xf6\xea\x47\x7d\xd2\x21\xbd\xef\xec\xfa\xb0\xb0\x67\x60\x47\x40\xb7\x99\xf2\x81\xf6\xcc\x4c\x8e\x18\x87\x26\x01\x0f\x1d\xb1\xdf\xe1\xc3\xe7\xff\x49\xef\x08\x0f\xca\xc3\xf8\x07\x3c\xe4\x4d\x97\x42\xbc\x48\xff\xeb\xab\xc5\xcf\x3c\xde\xdc\x46\xf1\x6a\x6c\xa3\xbf\x1e\x20\xeb\xc1\x37\xeb\x87\xef\xd6\xf4\x4f\x87\xcf\xc3\xe0\x08\xed\x4c\x99\xf5\xb9\xfa\x02\x32\xba\x46\x20\x6b\xf9\xcd\x03\xbc\x1a\x24\xf0\x25\xdc\x94\x0b\x14\x0c\x15\x06\x58\xeb\x9f\xe5\xbc\x02\x32\xe6\xf0\xb1\x47\x37\xd1\xb7\x4e\xfe\xd1\x4d\x64\x38\x26\x3a\x11\x58\x9d\x9f\x70\xa4\x53\x09\xef\x46\xfb\xa8\x5e\xd1\x74\x35\x77\x53\x01\xaa\xab\x9e\x96\xa6\x97\x64\x21\x68\xec\xce\xd4\x52\x36\x4a\x17\x54\xba\x2e\x9d\x3b\x20\xa3\x66\xad\x67\x87\x1b\x87\x79\xad\x74\xe3\x7f\x53\x29\x57\x55\x69\x65\xbb\xca\x34\xad\xbc\xf6\x35\xd0\x4b\x51\x4c\x13\x43\x64\x74\x37\x85\x51\x1c\xa3\x0c\x94\x4a\x8a\x62\xde\x08\x07\x8d\xf6\x80\x8c\x30\x55\xe3\x91\x16\x9e\xb2\x00\xcd\xd2\xfc\x8e\x48\xb9\xe5\x22\x39\x82\x59\x54\x20\xad\xaa\xb9\x99\xbe\xbb\xf1\xfd\x2c\x6f\x59\xb7\x49\x61\x38\x34\xea\x21\xb1\x69\x7e\x8d\xe6\x6d\xfd\x15\xd4\x39\xb6\xe9\x70\x3e\x20\x50\xf2\x52\xc4\x28\xe1\x0c\x3f\x5f\x40\xc6\x49\x02\x0b\x92\x11\x16\xa3\x78\xd6\x50\xb0\x8b\x12\x3b\x69\x1e\xb9\x4d\x02\x55\x8f\x33\x8a\x4c\x4d\x13\x38\x23\x6b\x72\x61\x34\x3f\x79\xe6\x73\x1e\x1b\x80\x2e\x75\x5b\xd4\x08\x63\x81\xca\xa1\x57\x6a\xea\xa2\x20\x2d\xa0\x3b\x1f\x3a\x63\xb4\x72\xf6\xd0\x19\x8c\xd4\x13\xa2\x08\x8c\x91\x35\x92\x47\xc6\x6d\xe0\x1c\xab\x62\x16\x25\x0d\xab\x8c\x13\x4f\xd6\x8a\x9a\xb3\x43\x35\xaf\x0f\x6c\xa5\xd3\x1a\xc9\xd6\xfa\xbf\xb5\x8e\xb1\x49\xed\xfa\x0d\xe3\xec\xf7\xed\x34\x8c\x63\xaf\x7a\xf4\x5d\x6d\xca\xa4\x32\x17\x97\xa9\xe0\x65\xd1\x38\x87\x47\x37\x51\xb5\xfe\x56\x2f\x03\x75\x4f\x73\x0b\x1d\x46\x3b\x49\x1b\xe8\xfa\x0d\x28\xfd\xc7\x46\x76\xf7\x96\x90\xe0\x92\x32\x04\x02\x6a\x57\xa0\xe9\x7f\x59\x99\x2f\x74\xfe\x59\xd6\x1b\x5a\xca\xb5\x0f\x36\xf9\xda\x4f\x1b\x9c\x6d\x99\x77\x88\x99\x6d\x6a\x80\xce\xa4\xdb\xe4\xcc\x70\x72\x16\x29\xc2\x12\x22\x92\xf9\xe4\x7c\xbe\x39\x7f\x0e\xa8\xe2\xc1\xb3\x26\x21\x03\x5a\x9b\xd5\x11\xfa\x40\x19\xcd\xcb\xbc\x4b\x10\x38\x4b\x70\x49\xca\x4c\x19\x17\xfa\x82\x82\xef\x49\x52\xa6\x5e\x9f\x43\x4e\xd9\xfc\xf7\x92\x30\x65\xf3\xe9\xeb\x6a\xf0\xe1\x2b\x1a\x88\xb0\x42\x0e\x37\x24\x2b\x11\x0a\x42\x85\x6b\xf7\xdd\x80\x68\xc7\x4b\x13\xe1\x44\x61\xca\x05\xfd\xe2\xc7\xb7\xd6\xf1\x86\xe2\xd6\x5c\x36\xf1\x8c\x26\xc6\x5c\x0b\x9a\xd5\x63\xa4\x4a\xd3\x66\xaf\x50\xbb\x24\x85\x35\xee\x9a\x6a\x58\xe3\xae\xa5\x4e\x0d\x6a\xd8\x6b\x02\x5b\x9e\x6b\xa5\x3d\xba\x84\xf5\x16\x95\x7f\xee\xe9\x84\x17\xd9\x89\xb9\x57\x88\xee\x47\xe3\xf0\xd5\xe1\xd9\x82\x54\xe7\xc8\x0a\xbb\xaa\xae\xdb\x78\xcd\xe2\x75\x09\xbc\x40\x61\x43\x5d\xf7\xab\xb7\x37\x07\xfa\xa8\x8a\x54\xc7\xc4\xbe\xe5\x7d\x8a\xa4\xd5\x50\x25\xa5\xba\x59\x2d\xb8\xa4\x8a\x8b\x96\xd6\x52\xaa\xbc\x43\xbc\xa1\xbd\x15\xc2\x8a\xc8\x55\xed\xc6\x54\x41\xcc\xf3\x9c\xaa\x2e\x2a\x76\xa5\xe5\x88\x1d\x47\xb9\x12\x88\x46\xd4\x38\x43\xc2\x60\xbb\x42\x66\x32\x59\x27\x59\x0d\x3c\xb7\x83\x91\xda\x13\x1d\xe9\x89\x7e\xc9\x97\x36\x0b\x36\x71\xcd\xcb\x79\x62\xf1\xde\x04\x78\x9f\xf6\x16\x4e\x79\x9d\xec\x62\x9e\x17\x34\x6b\x79\x4a\xca\x3d\xfd\xfc\x14\xd0\x19\x5b\x0c\xb1\xaf\x2a\x3c\xbc\xb8\x5a\x34\x83\x17\x0f\xeb\x2e\x23\x4a\x5b\x0e\xa8\xb2\x4a\xb0\x80\xf6\x88\x1e\x82\x28\x99\xf9\x09\x81\xab\x98\x3c\x8a\x45\x85\x78\x05\xff\xdf\x8c\xc8\x4a\x24\xcf\x29\xcc\x52\x87\xaf\x38\x69\xe6\x7e\xa9\x5c\x75\x36\x2e\x04\xf6\xb7\x2c\xd5\x49\x5d\xdd\xb6\x74\x34\x64\xde\x65\xca\x9f\x75\xaa\x51\x47\xa8\xf9\x6d\x4c\xfb\x0e\xe9\x88\x54\x7f\x7c\xfc\xe1\x02\xa9\x63\x57\xf0\x8f\x2a\xb7\xe1\xfe\x5e\xcb\x63\xbe\x85\x5b\x19\xb1\x66\xbc\xc2\xfe\xea\x6b\x35\x88\x6b\xaf\x79\x68\xdd\x1a\x9d\x62\x37\xb7\xf3\x1f\x1a\xfc\x85\x1b\x9f\x36\xf9\xeb\xb8\xc6\x3a\x79\xf4\x57\x69\xeb\xbf\x39\xf6\x3b\x22\xf3\x9f\xba\x55\x0b\xf9\xee\xe8\xd5\x1e\x0f\x88\xd8\xaa\x74\x9f\xa2\x66\x7d\xd5\x9d\x9e\xbc\xcb\xb5\x76\x08\xfb\x37\x85\x7f\xb5\xcc\x74\x9a\x60\x4f\x31\x07\x53\x41\x6e\x32\xfb\x16\x1a\x4c\x0b\x7e\x74\xea\xf2\xd0\x46\xdc\x5f\x73\xda\x3a\xd2\x2b\xf6\xbd\x88\x68\x5c\x78\xb6\x4d\x73\x78\xe6\x73\xd0\x0c\x4f\x34\x71\x74\xc4\xab\xda\xbd\xd6\x46\xc0\x0a\x4f\x70\x6e\x54\xe4\x77\x31\x8a\x2b\x92\x41\x82\xd2\xfc\xf6\x51\xc3\xc8\x0e\x45\xbc\x39\xa8\x88\x86\x29\xb7\x2b\x24\xc7\x32\x9b\xb9\xbd\x3d\x3c\x46\xfe\xee\x49\x6f\x98\xb3\x8e\x5c\xd2\x1e\x0d\xa4\xa3\x16\x7b\x6b\x7b\x2d\xdd\x4e\x26\x89\x7f\x6b\x5c\x95\x3a\xda\xaa\xa5\x10\xc8\x54\xb6\x03\xce\xcc\x1f\xe7\x58\xa9\xed\xfd\x24\x90\x2c\xe3\x5b\x4c\xb4\x6f\x8f\x6e\xa2\xe7\xb0\xa5\x59\x66\x6f\x16\x4d\x00\xe5\x5c\xa0\xee\x41\x18\xbc\x02\x64\x4a\xec\x20\x23\xf5\xef\x27\xea\x9e\x52\xbb\x81\x27\x95\x0e\x27\xdb\xc4\xe9\xbd\xea\xbe\xf2\xfc\x90\xaf\xd9\x11\xc8\xf7\xc6\xbb\x2b\x41\x83\x21\x46\x27\x13\x41\x2d\x3f\xf3\x3d\xf1\xf4\x36\xb2\xdd\xfd\x19\x8f\x34\x21\xd1\xbc\xaf\x87\x43\x6d\x5f\xa3\xa1\x24\x9f\x7f\x84\x24\xf9\xdc\xd9\x49\x9e\xe0\x6c\xc7\x8f\xfa\x8e\x5f\x1f\x98\xc0\xb0\xbd\xe7\x09\x37\x7d\x8f\xbd\x7f\x07\x00\x00\xff\xff\xf8\xd4\x88\xbe\xe4\x2d\x00\x00")

func apiProtoBytes() ([]byte, error) {
	return bindataRead(
		_apiProto,
		"api.proto",
	)
}

func apiProto() (*asset, error) {
	bytes, err := apiProtoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.proto", size: 11748, mode: os.FileMode(420), modTime: time.Unix(1539384363, 0)}
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
	"api.proto": apiProto,
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
	"api.proto": &bintree{apiProto, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
