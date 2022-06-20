// Code generated for package crd by go-bindata DO NOT EDIT. (@generated)
// sources:
// config/crds/kudo.dev_instances.yaml
// config/crds/kudo.dev_operators.yaml
// config/crds/kudo.dev_operatorversions.yaml
package crd

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configCrdsKudoDev_instancesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x3a\x5b\x73\xdb\x36\xd6\xef\xfe\x15\x67\xd4\x6f\xc6\x76\x2a\x51\x96\xdb\xa4\xa9\xe6\xcb\x66\xb2\x4e\xd2\xf1\x36\x17\x4f\xec\x74\xa7\x6b\xbb\x5b\x88\x3c\x12\x51\x93\x00\x03\x80\x92\xd5\xa6\xff\x7d\xe7\x1c\x80\x14\x65\x51\x97\xb8\x4d\xf7\x61\xf9\x62\x0b\x04\xce\xfd\x0e\xee\xf5\x7a\xbd\x3d\x51\xc8\x1f\xd0\x58\xa9\xd5\x10\x44\x21\xf1\xd6\xa1\xa2\x5f\x36\xba\x79\x6c\x23\xa9\xfb\xd3\xc1\xde\x8d\x54\xc9\x10\x4e\x4a\xeb\x74\xfe\x0e\xad\x2e\x4d\x8c\xcf\x71\x2c\x95\x74\x52\xab\xbd\x1c\x9d\x48\x84\x13\xc3\x3d\x00\xa1\x94\x76\x82\x96\x2d\xfd\x04\x88\xb5\x72\x46\x67\x19\x9a\xde\x04\x55\x74\x53\x8e\x70\x54\xca\x2c\x41\xc3\xc0\x2b\xd4\xd3\xa3\xe8\xeb\xe8\xb8\x77\x14\x1d\x1f\x1d\x0f\x8e\x06\xc7\xdf\x1e\x0f\x1e\x0e\xbe\x7e\xdc\x7b\xf8\xf0\x9b\x44\x1c\x3f\x3c\x1a\x3d\x7e\xf8\x68\x0f\x20\x36\xc8\xc0\x2f\x64\x8e\xd6\x89\xbc\x18\x82\x2a\xb3\x6c\x0f\x40\x89\x1c\x87\x20\x95\x75\x42\xc5\x68\xa3\x9b\x32\xd1\x51\x82\xd3\x3d\x5b\x60\x4c\xa4\x4c\x8c\x2e\x8b\x21\xd4\xeb\xfe\x48\xa0\xd2\x73\x78\x1a\x4e\xf3\x52\x26\xad\xfb\x7e\x69\xf9\x95\xb4\x8e\x5f\x15\x59\x69\x44\xd6\xc0\xc6\xab\x56\xaa\x49\x99\x09\xb3\x58\xdf\x03\xb0\xb1\x2e\x70\x08\x6f\x08\x55\x21\x62\x4c\xf6\x00\x02\xd3\x8c\xba\x17\x08\x9f\x0e\x46\xe8\xc4\xc0\x03\x8a\x53\xcc\x85\x27\x0c\x40\x17\xa8\x9e\x9d\x9d\xfe\xf0\xd5\xf9\xd2\x32\x40\x82\x36\x36\xb2\x70\x2c\xbf\x8a\x46\x90\x16\x5c\x8a\xe0\x37\xc3\x58\x1b\xfe\x59\x53\x0a\xcf\xce\x4e\xa3\x1a\x44\x61\x74\x81\xc6\xc9\x4a\x0c\xfe\x69\x98\x44\x63\xf5\x0e\xc2\x7d\xa2\xc9\xef\x82\x84\x6c\x01\x3d\xe2\xc0\x1c\x26\x81\x0d\xd0\x63\x70\xa9\xb4\x60\xb0\x30\x68\x51\x79\xeb\xa0\x65\xa1\x40\x8f\x7e\xc1\xd8\x45\x70\x8e\x86\x0e\x82\x4d\x75\x99\x25\x64\x34\x53\x34\x0e\x0c\xc6\x7a\xa2\xe4\xaf\x35\x34\x0b\x4e\x33\x9a\x4c\x38\xb4\x0e\xa4\x72\x68\x94\xc8\x60\x2a\xb2\x12\xbb\x20\x54\x02\xb9\x98\x83\x41\x82\x0b\xa5\x6a\x40\xe0\x2d\x36\x82\xd7\xda\x90\x40\xc6\x7a\x08\xa9\x73\x85\x1d\xf6\xfb\x13\xe9\x2a\x73\x8f\x75\x9e\x97\x4a\xba\x79\x9f\x2d\x57\x8e\x4a\xa7\x8d\xed\x27\x38\xc5\xac\x6f\xe5\xa4\x27\x4c\x9c\x4a\x87\xb1\x2b\x0d\xf6\x45\x21\x7b\x4c\xac\x62\x93\x8f\xf2\xe4\x0b\x13\x1c\xc4\xee\x2f\x09\xcf\xcd\xc9\x0e\xac\x33\x52\x4d\x1a\x2f\xd8\xf0\x36\x48\x99\x2c\x90\x74\x2a\xc2\x51\xcf\xc5\x42\x98\xb4\x44\xf2\x78\xf7\xe2\xfc\x02\x2a\xd4\x5e\xe0\x5e\xb6\x8b\xad\x76\x21\x66\x12\x91\x54\x63\x34\x7e\xe7\xd8\xe8\x9c\xa1\xa0\x4a\x0a\x2d\x95\xe3\x1f\x71\x26\x51\x39\xb0\xe5\x28\x97\x8e\xf4\xf7\xa1\x44\xeb\x48\x03\x11\x9c\xb0\x9f\xc3\x08\xa1\x2c\x12\xe1\x30\x89\xe0\x54\xc1\x89\xc8\x31\x3b\x11\x16\x3f\xbb\x90\x49\x9a\xb6\x47\xc2\xdb\x4d\xcc\xcd\x10\x75\x77\xb3\x97\x53\xe3\x45\x15\x31\xd6\xe8\xa4\x72\xb5\xf3\x02\xe3\x25\xd3\x4f\xd0\x4a\x43\xa6\xea\x84\x43\x32\xf0\x6a\x67\xb4\x04\xac\xdd\xe9\x82\xab\x1b\xe1\xb4\x69\xf5\xbe\x15\x3a\xde\x2e\xef\x66\xb2\xe5\x58\x22\x19\x8b\xc1\x31\x1a\xa4\x78\xe0\x34\xd9\x8e\x7f\x15\xaf\x9c\x09\xfe\xb7\x82\x68\x3d\x8d\xb0\x21\x40\xb4\x92\xf9\xec\xec\xb4\x0a\x0a\x3e\x16\x60\x45\x5d\x0b\xde\x0d\x2a\xac\x9e\xb1\xc4\x2c\x39\x13\x2e\xdd\x01\xf7\xfe\xe9\xd8\x23\x63\xd7\x61\x51\x14\x12\x63\x5c\x8a\x3e\x1c\x1c\x51\x24\x61\x91\xac\xcc\x60\x78\xd7\xf5\x0e\x12\x7c\x6f\x11\x9d\x9c\x90\x0a\x04\x39\xa3\x4c\xe0\x1f\xe7\x6f\xdf\xf4\xbf\xd3\x9e\x32\x10\x71\x8c\xd6\x7a\x23\xc8\x51\xb9\x2e\xd8\x32\x4e\x41\xd8\xca\x3e\xce\xe9\x4d\x94\x0b\x25\xc7\x68\x5d\x14\xa0\xa1\xb1\x97\xc7\xd7\x11\xbc\xd4\x06\xf0\x56\xe4\x45\x86\x5d\x90\x5e\x5e\xb5\x27\x57\x4a\x95\xd6\x33\x53\x9f\x85\x99\x74\x29\x93\x54\xe8\x24\x10\x3d\x63\x62\x9d\xb8\x41\xd0\x81\xd8\x12\x21\x93\x37\x38\x84\x0e\x59\x44\x03\xf5\x6f\x94\x85\x7e\xef\xc0\xc1\x2c\x45\x83\xd0\xa1\x9f\x1d\x8f\xb0\x0e\xb9\xb4\x56\x69\x70\x81\xd8\xa5\xc2\x81\x33\x72\x32\x41\xb2\x7d\x8e\x22\xe4\xa9\x87\xa0\x0d\xd1\xaf\x74\x63\x33\x83\x20\x79\x06\x53\x4d\x56\x08\xb9\x3c\xbe\xee\xc0\xc1\x32\x5f\x20\x55\x82\xb7\x70\x0c\x52\x79\xce\x0a\x9d\x1c\x46\x70\xc1\x9a\x99\x2b\x27\x6e\x09\x66\x9c\x6a\x8b\x0a\xb4\xca\xe6\x44\x71\x2a\xa6\x08\x56\xe7\x08\x33\xcc\xb2\x9e\xf7\xd3\x04\x66\x62\x4e\x3c\x54\xa2\x24\xad\x0a\x28\x84\x71\x77\x12\xd2\xc5\xdb\xe7\x6f\x87\x1e\x1b\xa9\x6d\xa2\x08\x05\x85\xbc\xb1\xa4\x74\x43\x79\xc6\x87\x4e\xd6\x39\x11\x52\x7a\x25\x39\x0d\x71\x2a\xd4\x04\x3d\xb5\x08\xe3\x92\x82\x58\xb4\x7f\x1f\x5b\x5f\xcd\x0e\xed\x66\xce\x59\xe2\xae\x73\xfd\xd7\x62\xf0\x8e\xcc\x71\xe1\xb3\x03\x73\x6f\x1a\x76\xb7\x91\x39\xaa\x2d\x8d\x42\x87\xcc\x5f\xa2\x63\x4b\xac\xc5\x58\x38\xdb\xd7\x53\x34\x53\x89\xb3\xfe\x4c\x9b\x1b\xa9\x26\x3d\x32\xac\x9e\xd7\xb6\xed\x73\x25\xd8\xff\x82\xff\xdc\x9b\x17\xae\xef\x76\x65\x88\x37\xff\x15\x5c\x11\x1e\xdb\xbf\x17\x53\x55\x39\xb1\x7b\xac\xdf\x3f\xaf\x12\xcd\x9d\xb3\xe4\x16\xb3\x54\xc6\x69\x55\x0b\x36\x22\x59\x2e\x12\x1f\xea\x84\x9a\x7f\x76\xa3\x25\xd1\x95\x86\x70\xcf\x7b\xa1\x35\xe9\x09\x95\xd0\xff\x56\x5a\x47\xeb\xf7\x92\x55\x29\x77\x72\xd4\xf7\xa7\xcf\xff\x1a\x53\x2e\xe5\xbd\xbc\x72\x4d\x45\x44\x4f\x21\x8c\xc8\xd1\xa1\x69\x29\x09\x44\x92\x70\x2b\x28\xb2\xb3\x8d\x85\xc3\xbd\x71\x67\x42\xbd\xb8\xc5\xb8\x74\xdb\xcb\xa2\xfd\x0b\x4e\x61\xc2\x20\xb8\x99\xa6\x80\x4f\x05\x11\x41\x00\xac\x40\x40\x2c\x14\x15\xaf\x75\xde\x1a\x02\x0c\x0e\x29\xcf\x48\x83\xb1\xa3\x0c\x92\x1a\x5d\x4e\xd2\x50\xde\x72\x72\x80\x58\x1b\x83\xb6\xd0\x2a\xa1\xb4\x51\xcb\xa3\x0a\xf4\xcd\xba\x30\x3a\xab\xa5\x05\xb9\x28\x00\x8e\x0f\x61\x05\xb6\x45\xc7\xf5\x7b\x30\x88\xe5\xf3\x4d\x8e\xf9\x17\x87\x41\x9f\x6e\xfe\x99\xca\x0c\x6b\x6a\xe1\x60\x70\x58\x71\x62\x21\x15\x45\x81\xca\x52\x12\x36\x73\x70\x32\x47\x10\x50\x5a\x34\x21\x2d\x59\x9f\xef\x3c\x71\x5d\x10\x0b\xb2\x0e\x8e\x0f\x1b\x89\x9c\x05\xc6\xae\x6a\xa9\x69\x48\xea\x56\xd2\x4a\x57\xfa\x06\x1f\x66\x29\xaa\x86\x5d\x40\xa2\xd1\xaa\xfd\x7d\x57\x65\x40\x8c\x26\x11\xa1\x43\x23\x75\x22\x63\x18\x89\xf8\xa6\x2c\xb8\x7a\xa9\xf1\x90\x35\x1b\x99\x54\x7d\x0c\xde\x4a\xcb\x42\x09\x7b\xc7\x32\xc3\x08\x9e\xd5\xf6\x95\xcd\x43\x75\xa3\x99\x4b\xa3\x75\xce\x94\xc5\x24\xb9\x8c\xd3\xb9\x5a\x02\xea\xbd\x9d\xf8\x33\xa5\x52\xac\xb8\x4c\x28\x7b\x27\x3b\xc3\x1b\xed\x70\x08\x4b\x52\x0f\xc2\xae\x2a\x7c\x16\x08\x17\x30\x84\x61\x8d\x2d\x58\x5f\x0f\x9d\x9e\xc3\xc9\xfb\x77\xef\x5e\xbc\xb9\x78\xf5\x63\xb0\x3a\x6a\x91\xde\x72\x41\xde\x68\xc7\x1b\xd3\x11\x38\x38\x3d\x39\x24\xd1\x24\x5a\xa1\x2f\x7b\xbc\x3c\x02\x35\xdd\x66\xbd\x31\x93\x59\x46\xf6\x1b\x67\x28\x0c\x41\x7e\x21\xe2\xf4\xae\x8d\xa7\x82\x74\x5d\x2a\xf9\xa1\x44\xa0\xc0\x63\x75\x55\xc1\xb2\x1e\x89\x15\x3e\x32\xa2\x60\xd4\x5b\xa8\x44\x3a\x8f\x80\x4b\x28\x01\x0a\x67\x74\x7c\x35\x9a\x6c\x6e\x12\x8a\x60\xb3\xed\x61\x71\x4b\x38\xa5\xea\xb9\x6c\x05\x7b\xc7\xdb\x6b\x6d\x9d\xf3\x09\x88\x45\x41\x0a\xf5\x0d\x59\xdd\x88\x71\xbc\xd5\x59\xa6\xcb\xfb\xf5\x1c\xbb\x45\x77\x92\x31\xb7\xea\x04\xcd\x1b\x42\xaa\xb3\xc4\x56\x3a\x38\x7d\x1e\x66\x10\x5d\x90\x2a\xce\x4a\x36\x9d\xf7\xef\x4f\x9f\xdb\x08\xe0\xef\x18\x8b\xd2\x52\xb5\x4a\x16\xb0\xef\xe0\xed\x9b\x57\x3f\x92\xe3\xfa\x1d\x41\xfd\x04\x5e\x81\xc8\xa4\x9f\x84\x78\x82\xf9\xb4\xaf\x64\x19\x73\x2d\x03\xa9\x1c\x75\xf0\x64\xaf\x29\x66\x05\x85\xa2\x1b\x04\x5b\x9a\x40\x1d\x01\xe6\xb7\x9c\x34\x20\xd1\x5c\xe1\x4e\xd0\x91\x5d\x8e\x33\xee\xeb\xff\xc4\x1c\xb2\xae\xdd\x6e\xd1\x75\x7b\xc3\xed\x55\xdc\x6c\xb9\xf5\x28\x44\xa7\x95\x9e\x7b\xc7\x96\x3b\x26\x0f\x6e\x8c\x2b\x9b\x8f\x74\x98\xb7\xda\xe0\x12\x75\x9d\x93\x0a\x44\xd5\xe8\x10\x89\x4e\xc8\xcc\x72\x64\x22\x77\x16\xd4\xe5\xb8\xba\x79\xf2\x21\xa9\x69\x9e\x92\x07\x73\x50\x0d\x57\x23\xe8\xf5\x7a\xa1\xc1\x71\xa6\xa4\x1e\x35\x68\x33\x09\x91\x38\x84\x7e\xb2\x0f\xe1\x6d\xc2\x18\x31\x07\xe1\x47\x37\x3e\x4a\x14\xc2\xa5\x10\x79\xf1\x46\x0b\x46\x23\x58\x6e\x32\xd9\x66\x5e\x6a\x1d\xc4\xeb\x11\xfe\xc6\x8c\xf6\xfb\xf0\xae\x1e\x1f\x35\x04\x1e\xe2\x3f\x67\xc5\xb1\xd6\xfb\x76\x99\xa7\xa8\x3a\xfc\xbd\xd2\x33\xd5\x46\x02\xe3\x14\x06\x87\x70\xd5\x79\x36\x15\x32\x13\xa3\x0c\xaf\x3a\x5d\xb8\xea\x9c\x19\x3d\x31\x68\xad\x54\x13\x5a\x20\xf3\xbc\xea\x3c\xc7\x89\x11\x09\x26\x57\x9d\x0a\xf4\x97\x85\x70\x71\xfa\x1a\xcd\x04\xbf\xc7\xf9\x13\x06\xb8\xf4\xea\xdc\x19\xe1\x70\x32\x7f\x92\xd3\x9e\xfa\x5d\x26\xad\xbb\x98\x17\xf8\x84\xd3\x72\x63\xf1\xb5\x28\x96\x00\xd5\x6a\xb5\x70\x79\x9d\xa3\x13\xd3\x41\xb4\x50\xf5\xcf\xbf\x58\xad\x86\x57\x9d\x05\x4f\x5d\x9d\x93\xc1\x14\x6e\x7e\xd5\x81\x25\x0a\x86\x57\x1d\xa6\xa1\x5a\xaf\x88\x1e\x5e\x75\x08\x1b\x2d\x1b\xed\xf4\xa8\x1c\x0f\xaf\x3a\xa3\xb9\x43\xdb\x1d\x74\x0d\x16\x5d\x72\xcb\x27\x0b\x0c\x57\x9d\x9f\xe1\x4a\x55\x44\x6b\x97\xa2\xf1\x9a\xb6\xf0\x7b\xa7\x2d\x02\x6f\x8c\xcf\x00\x99\xb0\xee\xc2\x08\x65\x65\x35\x52\x6f\xdf\x77\xc7\xe0\x57\x8f\x55\x43\x67\x7a\xe3\x4b\x8e\x30\x23\x08\xc2\x72\xf5\x6e\xb2\x5e\xa3\x73\x76\x0a\x6f\x15\x3c\xc8\x50\xcc\x4c\xd5\xd2\xfb\x1c\x35\x42\x5f\x5f\x10\xa8\x52\x25\x68\xb2\x39\xe7\xf3\x85\xb7\x71\x91\x91\x44\x00\xa7\x63\x1f\xcf\x42\x8b\x7e\x43\x56\x47\x01\x13\x15\x94\xb6\xaa\x2c\x98\xae\x1a\x22\x79\x9b\xf7\x92\x00\x86\x43\x6b\x4c\x15\x36\x99\x62\x7b\xd8\x03\x72\xbd\x5c\xb8\x21\x50\x51\xd8\x23\x88\x6b\xf6\x6d\xc9\x27\x00\x39\x5a\x2b\x26\xbb\x09\x3c\xec\xf5\xb9\x25\x2d\x73\xa1\xc0\xa0\x48\x88\xce\xc5\x3b\x95\xc8\x58\x70\xe5\x54\x05\x1f\x31\xd2\xa5\x0f\x07\x0b\xf9\x07\x11\xe7\x62\x4e\xf2\xa5\x72\x81\x0c\xb6\xca\x23\x6b\x88\xc9\xc5\xed\x2b\x54\x13\x97\x0e\xe1\xab\xe3\x6f\x1e\x3d\xbe\x2f\xcf\x55\xb0\xfe\x0e\x15\x1a\xd1\x5e\xca\xb7\xb0\xbf\x7a\xac\x31\xd3\x66\xfe\xa2\x6a\xbc\x1b\x4d\x16\x7b\xfc\x60\x6a\xc9\x0e\x67\xc2\x52\xd1\x0d\x23\x61\x31\x81\xb2\x20\x79\x50\x28\xac\xca\x31\xee\x42\x5b\x81\x49\xdb\x28\x24\x07\xc7\x5d\x18\x05\xd1\xae\xc6\xb6\xcb\xdb\xeb\xa8\x85\x64\x69\xe1\xdb\xee\x1d\x7a\xa8\x8c\x2d\x39\x2d\x70\x93\xc1\x45\x1f\x55\x95\x61\x8e\xb4\x26\x57\x2c\xca\xc7\x6d\x56\x2a\x95\x7b\xf4\xf5\x3a\xa5\x4a\x25\xf3\x32\x1f\xc2\xd1\x46\x75\x52\xd2\x99\xa0\x69\xdd\x63\x50\xd8\x1d\x75\xe8\xb7\x2e\x12\xa4\xa0\xe0\x34\x31\x22\xcf\x85\x93\x31\xc8\x84\x5a\xf4\xb1\xe4\x66\xaa\x36\x64\xdf\x25\xf3\xc1\xaa\xf9\xa8\x65\xb7\x6f\x43\xb4\x69\x98\xf6\x99\xd1\x49\x19\x53\x11\xae\xc7\x8b\x91\x77\x23\x0c\xcd\x0b\xf4\xb6\xef\x4b\x08\xc0\x5b\x12\x75\x7d\x35\xe4\x6f\x8f\x50\x50\x97\x60\x03\xca\xaa\xe4\xf6\x89\x68\x96\x22\x47\x5d\xbe\xe8\x0a\x67\x0c\x53\x65\x65\xc2\xb5\xb3\x80\x49\x29\x8c\x50\x0e\x31\xe1\xbb\x36\xb8\xa8\xf6\x36\x02\x9b\x58\x5c\x95\xd4\x35\xdc\xc5\xa2\xbc\x27\x12\xc3\xf5\x0a\xfb\xe7\x0e\x8e\x39\x38\x3a\xde\xa0\xe9\x7a\xd7\x9a\x2d\x85\x70\x0e\x8d\x1a\xc2\x4f\x97\xcf\x7a\xff\x12\xbd\x5f\xaf\x0f\xc2\x3f\x47\xbd\x6f\xff\xdd\x1d\x5e\x3f\x68\xfc\xbc\x3e\x7c\xfa\x7f\xf7\x0d\x01\x9b\xea\xfa\x3b\x26\x13\xd2\xc3\x62\xe0\xec\xb5\xd8\xe5\xdc\xa1\xc7\x70\x61\x4a\xec\xc2\x4b\x91\x59\xec\xc2\x7b\xc5\x41\x7f\x9d\xa0\x50\x95\xf9\x3a\xa4\x3d\xe8\x10\xa8\xb6\x04\x1a\x5e\x33\x8e\xf5\xef\x03\xee\xfb\x8a\x84\x37\xec\x22\x10\xae\x4c\xf4\xb8\x19\x3f\x1a\x57\x6e\x3c\x6e\x27\x47\xd1\x51\xa8\xec\xa2\x58\xe7\xfd\xc6\x95\x1c\x95\x94\xaf\x85\x9a\xc3\x22\x58\xf9\x3a\xec\xae\x25\x5b\xdf\x35\xc4\x46\x5b\x5b\x0f\xf2\x2c\xdf\x1c\x40\x5d\xac\xf9\x10\x38\x0a\x5d\x8b\x30\x23\xe9\x8c\x30\xf3\x46\x41\x5d\x0d\x59\x4a\x8b\xe3\x32\x83\x03\x8b\x08\x91\xd2\x09\xae\xc6\xcc\x43\x1f\x19\xc5\x48\x66\xd2\xf1\xf8\x3e\x41\x6e\x41\x64\x28\x7d\xf3\x42\x1b\x27\x94\xf3\xee\x64\x70\x82\xb7\xd4\xa6\xe6\x54\x4e\x21\xb7\x45\x07\x89\xb2\x83\xc1\xf1\x57\xe7\xe5\x28\xd1\xb9\x90\xea\x65\xee\xfa\x87\x4f\x0f\x3e\x94\x22\xe3\xab\x06\xea\x45\x5f\xe6\xee\x70\x87\x24\x37\x78\xb4\xd5\x4f\x0e\x2e\xbd\x37\x5c\x1f\x5c\xf6\xc2\x7f\x0f\xaa\xa5\xc3\xa7\x07\x57\xd1\xc6\xf7\x87\x0f\x88\xb4\x86\x8f\x5d\x5f\xf6\x16\x0e\x16\x5d\x3f\x38\x7c\xda\x78\x77\x78\x4f\x77\x33\xf8\xa1\x94\x06\x5b\x5b\xd7\x5e\x4b\x19\xd7\xba\x2d\x14\x18\xad\xef\x7c\x70\x6e\x7d\xe5\x55\xdc\xfa\x8a\xa8\x5e\x3b\x1b\x6c\x1d\xff\x55\x2f\xb9\xc7\x69\x19\x0d\x9e\xaf\x89\x2a\xbb\x8e\x25\x97\xfb\xb9\xb3\x1a\x22\x34\x3f\x5b\xe0\x0b\xc0\x3a\x26\x85\x89\xe2\x95\x22\x83\xb4\x7e\xdc\xe8\x6f\x03\x39\x53\xfb\x5c\xb2\x48\x4e\xb6\x2a\xdc\x5b\x9e\x8f\xf0\x06\xa7\x68\xea\x39\x11\x7c\x5c\xb7\x73\x71\x62\xcb\x33\xf5\x5b\x5e\x18\xa3\x0d\x1f\xf8\xff\x1e\x3f\x7f\xe3\xe5\x33\xf4\x53\xab\x25\x50\x3f\x6d\xc3\xb5\x0e\xe9\x74\xdd\x86\x2f\x3d\xce\x5e\xf5\xb7\xf7\xe5\x1f\xdf\xb9\xb2\x3e\x5d\x4b\xd5\x47\x78\x29\x9c\xc8\x00\x59\x08\xcb\x6c\x9c\x68\x0a\x90\x8e\x1b\xbd\x8f\xf7\x6d\x9d\xde\xfb\x2f\x1f\x16\x9f\x22\xfd\xc1\x9e\x41\x95\x19\x47\xd6\x21\x38\x53\x7e\xae\xc6\x62\xeb\xf9\xf5\xb7\x72\x3b\x1c\x2e\x52\x61\xd7\x89\x6c\xc3\x3c\xc5\x3f\x4b\x5e\x78\x46\x90\x76\xf1\x42\xda\xb7\x16\xe4\x36\x35\xfa\x67\x8b\xc8\x76\xe4\xdd\x3f\x9b\xc4\xf7\x49\x80\x36\x57\x4a\xfe\xf9\x1c\x93\xd0\x4f\x26\x13\x8b\x2d\x54\x6e\x51\x7b\x0b\x2b\xe7\x0e\x8b\x1d\x74\x4f\xb8\xb7\x80\xdd\xcd\x00\xfc\xb3\x93\x19\xf8\x67\x47\xe1\xf8\x67\xbb\x49\xdc\x03\xe8\x2e\xe6\xe1\x9f\xcf\x69\x24\x9f\x4c\xf8\xc6\x24\xbf\xba\xb1\x2d\xe1\x7f\x12\xb0\x6d\x60\x3e\xa1\x1f\xf9\xd3\x44\xb7\x55\x5c\x6b\x6f\x1b\xfe\x67\xef\x1b\xb6\x0a\x6d\xa3\x29\x2c\xb7\x95\x99\x8c\x31\x7c\x1b\x45\xfd\xb5\xe2\xab\x58\xbe\x32\xa6\x86\x66\x46\xec\xf0\xb7\x8d\x24\x0a\xbf\xb9\x6a\x72\xb4\x49\xd0\xf0\x6c\x01\x3f\x94\xfe\xeb\x05\x05\x73\x91\x67\x3c\x1e\x5a\xcc\x00\xac\x9c\x28\x39\x96\xb1\x50\x0e\x66\x7c\x55\xcb\xe0\xa5\xdb\xe7\xf9\xe4\x1f\xbc\x30\x59\x59\xf4\xa3\xa6\x46\xd5\x60\x9d\x36\x14\xc8\x1a\x2b\xe5\xa8\xee\xe3\x2a\xd3\x0a\xa6\x0f\xbf\xfd\xbe\xb7\xf0\x02\x3f\x01\xf5\xcd\xd2\xd2\xa7\xd1\x1d\x5f\x25\x55\x5f\x3e\xf3\xcf\xc6\xdd\x09\x5c\x5e\xef\x79\xc4\x98\xfc\x50\x7d\xde\x4c\x8b\xff\x09\x00\x00\xff\xff\x9a\xc9\xd0\xfd\x64\x2e\x00\x00")

func configCrdsKudoDev_instancesYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_instancesYaml,
		"config/crds/kudo.dev_instances.yaml",
	)
}

func configCrdsKudoDev_instancesYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_instancesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_instances.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configCrdsKudoDev_operatorsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4b\x6f\xdc\x36\x10\xbe\xeb\x57\x0c\xd2\x43\x2e\x95\xf6\xd1\x6c\xeb\xea\x16\xb8\x39\x18\xad\xd3\x20\x0e\x7c\x29\x7a\x18\x89\xb3\xbb\x53\x4b\xa4\x3a\x1c\x2e\xea\xfc\xfa\x82\xa4\xf6\xe9\xdd\xd8\x35\x50\xdd\xf8\x71\x5e\xfc\xf8\xcd\x88\x45\x59\x96\x05\x0e\x7c\x4f\xe2\xd9\xd9\x1a\x70\x60\xfa\x47\xc9\xc6\x95\xaf\x1e\xae\x7c\xc5\x6e\xb2\x99\x15\x0f\x6c\x4d\x0d\xd7\xc1\xab\xeb\x3f\x93\x77\x41\x5a\xfa\x85\x96\x6c\x59\xd9\xd9\xa2\x27\x45\x83\x8a\x75\x01\x80\xd6\x3a\xc5\x08\xfb\xb8\x04\x68\x9d\x55\x71\x5d\x47\x52\xae\xc8\x56\x0f\xa1\xa1\x26\x70\x67\x48\x52\xf0\x6d\xea\xcd\xb4\x7a\x57\xcd\xcb\x69\x35\x9f\xce\x67\xd3\xd9\xfc\xe7\xf9\x6c\x31\x7b\x77\x55\x2e\x16\x3f\x19\x9c\x2f\xa6\xcd\xd5\xe2\xc7\x02\xa0\x15\x4a\xc1\xbf\x70\x4f\x5e\xb1\x1f\x6a\xb0\xa1\xeb\x0a\x00\x8b\x3d\xd5\xe0\x06\x12\x54\x27\xbe\x7a\x08\xc6\x55\x86\x36\x85\x1f\xa8\x8d\xa5\xac\xc4\x85\xa1\x86\x1d\x9e\x5d\xc6\x2a\xf3\x09\x7f\x1f\xbd\x13\xd4\xb1\xd7\x5f\x8f\xe0\xdf\xd8\x6b\xda\x1a\xba\x20\xd8\x1d\x64\x4b\xa8\x67\xbb\x0a\x1d\xca\x1e\x2f\x00\x7c\xeb\x06\xaa\xe1\x63\x4c\x35\x60\x4b\xa6\x00\x18\x0f\x9d\x52\x97\x63\xe1\x9b\x59\x43\x8a\xb3\x1c\xa8\x5d\x53\x8f\xb9\x30\x88\xc1\xec\xfb\x4f\x37\xf7\x3f\xdc\x1d\xc1\x00\x86\x7c\x2b\x3c\x68\xe2\x6f\x5b\x23\xb0\x07\x5d\x13\x64\x63\x58\x3a\x49\xcb\x6d\x45\xf0\xfe\xd3\xcd\x2e\xc0\x20\x11\x56\xde\x92\x90\xbf\x03\x41\x1c\xa0\x27\xe9\xde\xc6\x8a\xb2\x15\x98\xa8\x04\xca\x69\xc7\xa3\x91\x19\x0f\x01\x6e\x09\xba\x66\x0f\x42\x83\x90\x27\x9b\xb5\x11\x61\xb4\xe0\x9a\xbf\xa8\xd5\x0a\xee\x48\xa2\x23\xf8\xb5\x0b\x9d\x89\x92\xd9\x90\x28\x08\xb5\x6e\x65\xf9\xeb\x2e\x9a\x07\x75\x29\x4d\x87\x4a\x5e\x81\xad\x92\x58\xec\x60\x83\x5d\xa0\xef\x01\xad\x81\x1e\x1f\x41\x28\xc6\x85\x60\x0f\x22\x24\x13\x5f\xc1\xad\x13\x02\xb6\x4b\x57\xc3\x5a\x75\xf0\xf5\x64\xb2\x62\xdd\x8a\xbd\x75\x7d\x1f\x2c\xeb\xe3\x24\xe9\x96\x9b\x10\xaf\x77\x62\x68\x43\xdd\xc4\xf3\xaa\x44\x69\xd7\xac\xd4\x6a\x10\x9a\xe0\xc0\x65\x2a\xd6\x26\xc1\x57\xbd\xf9\x4e\xc6\xf6\xf0\x6f\x8f\xc8\xd3\xc7\xa8\x02\xaf\xc2\x76\x75\xb0\x91\x64\xf7\x0d\x96\xa3\xfe\xe2\x8d\xe2\xe8\x9a\x4f\xb1\x27\x33\x42\x91\x8f\xcf\x1f\xee\xbe\xc0\x36\x75\x26\x3c\x73\xbb\x37\xf5\x7b\x9a\x23\x45\x6c\x97\x24\xd9\x72\x29\xae\x4f\x51\xc8\x9a\xc1\xb1\xd5\xb4\x68\x3b\x26\xab\xe0\x43\xd3\xb3\xc6\xfb\xfb\x3b\x90\xd7\x78\x03\x15\x5c\xa7\x2e\x87\x86\x20\x0c\x06\x95\x4c\x05\x37\x16\xae\xb1\xa7\xee\x1a\x3d\xfd\xef\x24\x47\x36\x7d\x19\xc9\x7b\x19\xcd\x87\x03\xea\xd4\x38\xf3\x74\xb0\xb1\x9d\x17\x17\xee\x64\xdb\x68\x77\x03\xb5\x47\xd2\x37\xe4\x59\xa2\x54\x15\x95\xa2\xc0\x8f\xa6\xc9\xb7\x7b\xee\x34\xcb\xc9\xd6\xc5\x83\x25\x0d\x85\x86\xc4\x92\x92\x3f\xdb\xb6\xcf\x7a\x1b\xf7\x1a\xbf\x1e\xd9\x2a\xb2\x25\x79\x72\x10\x00\x56\xea\xcf\xc0\x27\x4c\xde\xee\x42\x8c\x78\x43\x3e\x4e\x85\xdd\x28\xdb\xe7\xa8\xce\xc4\xba\xcc\x64\xfe\xa8\x47\xee\xce\x6f\x9d\x14\xf2\x21\x5a\xa6\x36\xb3\xe0\x12\x86\x5d\x76\x07\x34\x46\xc8\xa7\xb9\x13\xb5\x8a\x6d\x6e\x8e\x38\xb3\xcd\x33\xf5\x3d\x43\x60\xfe\xd2\xf0\x7f\x49\x8d\xf1\x07\x92\x27\x41\xf0\x24\xc9\x0f\x9c\x80\x93\x15\x5a\xfe\x9a\x87\x6a\x04\x5f\x59\xc9\x85\x5e\x38\xdc\x44\x11\x7c\x2c\x9e\x96\x9f\x7e\x6b\xb7\x68\x79\x49\x5e\xff\x93\x86\x82\x9c\xb9\x9f\x8b\xf6\x97\xda\x55\x51\x83\x7f\x49\xc3\x26\xc3\xa3\x96\x75\x8d\x8f\x23\xf1\xb9\x9e\x3d\x9b\xf9\x09\x98\x43\xd5\xa0\x12\x28\x03\xea\x04\x57\x34\x22\xfb\x3a\xb1\x6d\x69\x50\x32\x1f\x4f\x9f\x1f\x6f\xde\x1c\xbd\x2e\xd2\xb2\x75\xd6\x70\x7e\x4e\xc1\x1f\x7f\x16\x39\x2a\x99\xfb\xed\x13\x22\x82\xff\x06\x00\x00\xff\xff\x28\x97\x57\x9d\xc8\x09\x00\x00")

func configCrdsKudoDev_operatorsYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_operatorsYaml,
		"config/crds/kudo.dev_operators.yaml",
	)
}

func configCrdsKudoDev_operatorsYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_operatorsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_operators.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configCrdsKudoDev_operatorversionsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\x6b\x93\xdb\x36\x92\xdf\xe7\x57\x74\x4d\x3e\x38\x49\x49\x94\x67\x12\xdf\xe6\xe6\x5b\xd6\x8e\xaf\xe6\xe2\xb5\x5d\x7e\xec\xd5\x55\x2e\x55\x03\x11\x4d\x11\x3b\x20\xc0\xc3\x43\xb2\x2e\x95\xff\x7e\xd5\x0d\x90\xa2\x34\x7a\xd0\x72\x92\xcb\xd6\x99\x5f\x12\x11\x60\xa3\xdf\x2f\xf4\xf8\x62\x3a\x9d\x5e\x88\x56\xfd\x1d\x9d\x57\xd6\xdc\x80\x68\x15\x7e\x08\x68\xe8\x97\x2f\xee\xbf\xf3\x85\xb2\xb3\xe5\xd5\xc5\xbd\x32\xf2\x06\x9e\x46\x1f\x6c\xf3\x06\xbd\x8d\xae\xc4\x67\x58\x29\xa3\x82\xb2\xe6\xa2\xc1\x20\xa4\x08\xe2\xe6\x02\x40\x18\x63\x83\xa0\xd7\x9e\x7e\x02\x94\xd6\x04\x67\xb5\x46\x37\x5d\xa0\x29\xee\xe3\x1c\xe7\x51\x69\x89\x8e\x81\x77\x47\x2f\x1f\x17\xdf\x16\xd7\xd3\xc7\xc5\xf5\xe3\xeb\xab\xc7\x57\xd7\xff\x7a\x7d\xf5\xe4\xea\xdb\xef\xa6\x4f\x9e\xfc\x45\x8a\xeb\x27\x8f\xe7\xdf\x3d\xf9\x97\x0b\x80\xd2\x21\x03\x7f\xa7\x1a\xf4\x41\x34\xed\x0d\x98\xa8\xf5\x05\x80\x11\x0d\xde\x80\x6d\xd1\x89\x60\x5d\x86\xeb\x8b\xfb\x28\x6d\x21\x71\x79\xe1\x5b\x2c\x09\xa3\x85\xb3\xb1\xbd\x81\xfe\x7d\xfa\x32\x23\x9b\x08\x7d\x95\x81\x64\xbe\xf0\x8a\x56\x3e\xfc\xb8\x6f\xf5\x85\xf2\x81\x77\xb4\x3a\x3a\xa1\x1f\xa2\xc0\x8b\x5e\x99\x45\xd4\xc2\x3d\x58\xbe\x00\xf0\xa5\x6d\xf1\x06\x5e\x12\x1a\xad\x28\x51\x5e\x00\x74\x1f\x13\x5a\xd3\x4c\xdb\xf2\x6a\x8e\x41\x5c\x25\x78\x65\x8d\x8d\x48\x48\x03\xc1\x34\xdf\xbf\xbe\xfd\xfb\x37\x6f\xb7\x5e\x03\x48\xf4\xa5\x53\x6d\x60\x16\xef\x20\x0e\xca\x43\xa8\x11\xd2\x37\x50\x59\xc7\x3f\x77\xd1\x87\xef\x5f\xdf\x16\x3d\xc0\xd6\xd1\x7a\x50\x1d\xc3\xd2\x33\xd0\xa1\xc1\xdb\x9d\xe3\x1f\x11\x86\xf9\x68\x49\xca\x83\xe9\xfc\x7c\x10\xca\x4c\x14\xd8\x0a\x42\xad\x3c\x38\x6c\x1d\x7a\x34\x49\x9d\xe8\xb5\x30\x60\xe7\xff\xc0\x32\x14\xf0\x16\x19\x43\xf0\xb5\x8d\x5a\x92\x96\x2d\xd1\x05\x70\x58\xda\x85\x51\xff\xd3\x43\xf3\x10\x2c\x1f\xa3\x45\x40\x1f\x40\x99\x80\xce\x08\x0d\x4b\xa1\x23\x4e\x40\x18\x09\x8d\x58\x83\x43\x82\x0b\xd1\x0c\x20\xf0\x16\x5f\xc0\xdf\xac\x43\x50\xa6\xb2\x37\x50\x87\xd0\xfa\x9b\xd9\x6c\xa1\x42\x67\x1f\xa5\x6d\x9a\x68\x54\x58\xcf\x58\xd5\xd5\x3c\x06\xeb\xfc\x4c\xe2\x12\xf5\xcc\xab\xc5\x54\xb8\xb2\x56\x01\xcb\x10\x1d\xce\x44\xab\xa6\x8c\xac\x61\x1b\x29\x1a\xf9\x85\xcb\x16\xe5\x1f\x6d\x31\x2f\xac\x49\x2b\x7c\x70\xca\x2c\x06\x0b\xac\xa2\x47\xb8\x4c\x4a\x4a\xa2\x15\xf9\xd3\x44\xc5\x86\x99\xf4\x8a\xf8\xf1\xe6\x87\xb7\xef\xa0\x3b\x3a\x31\x3c\xf1\x76\xb3\xd5\x6f\xd8\x4c\x2c\x52\xa6\x42\x97\x76\x56\xce\x36\x0c\x05\x8d\x6c\xad\x32\x81\x7f\x94\x5a\xa1\x09\xe0\xe3\xbc\x51\x81\xe4\xf7\xdf\x11\x7d\x20\x09\x14\xf0\x94\x1d\x03\xcc\x11\x62\x2b\x45\x40\x59\xc0\xad\x81\xa7\xa2\x41\xfd\x54\x78\xfc\xdd\x99\x4c\xdc\xf4\x53\x62\xde\x38\x36\x0f\x7d\xda\xee\xe6\xc4\xa7\xc1\x42\xe7\x5b\x0e\xc8\x64\xc7\xf0\xde\xb6\x58\x6e\x59\x80\x44\xaf\x1c\x69\x6c\x10\x01\x49\xcf\x77\x3e\x28\xb6\x40\xef\x37\xc1\x64\x86\xed\x5e\x33\x3c\x42\x26\x24\x0f\x6d\xb0\x24\x54\xdf\xf2\xf2\xc3\x8f\xb7\xa8\x79\xba\xb3\xbd\x27\x45\x40\xc0\xa6\x25\x3b\x93\x9d\xee\x85\x5a\x04\x28\x85\x61\xb9\x7b\x94\x64\x8c\xf9\x38\xfa\x5f\x61\x40\x19\x1f\x84\x29\x31\x59\x3d\xf6\xa4\x17\x1f\x43\x41\xe7\xb3\x4e\x60\xfe\xe8\x15\x0b\xee\x0d\x56\xe8\x90\xce\x24\x5d\x12\xca\x78\x40\x63\xe3\xa2\x66\xf5\x73\x4d\x72\x37\xc1\x82\xc6\x00\x6b\x1b\x09\xc7\x96\x30\xb6\x0e\x1a\x2b\x55\xb5\x66\x4c\x1d\x81\x21\xb1\x75\x2e\x69\x3a\x9d\xc2\x4b\x5c\x11\xa1\xbe\x77\x62\x84\x35\x08\x87\x20\x95\x2f\x6d\x74\x62\x81\x12\xe6\x58\x8a\xe8\x99\x66\xa9\xaa\x4a\x95\x51\x87\x75\xc6\x75\x4e\x7c\x23\xf3\x89\x5e\x2c\x10\x56\x35\x1a\xc0\x66\x8e\x52\xa2\x04\x65\xc8\x1d\xfb\x02\xe0\xaa\x80\xdb\x85\xb1\x74\x7e\xa5\x50\x4b\x7a\x77\x4b\xee\xad\xd4\x51\x22\x19\xac\x59\xe7\x15\x58\xd5\xaa\xac\x19\x09\x32\xc1\x05\x1a\x74\x42\xeb\x35\xd4\x96\x01\x14\x00\xcf\xad\xeb\x25\x31\x81\x2e\xc4\x77\xde\x9a\x7c\xe4\x73\x02\xf5\x5a\x84\x04\x67\x6e\x43\x4d\x8e\x7b\x0d\x4e\x38\xd4\x6b\x72\x32\x8a\xd1\x13\x65\x88\x42\x27\xe4\x0b\x80\x6b\x32\xf3\xb4\x98\xe8\xa9\x51\xb7\x19\x55\x0f\xaa\x69\xad\xf7\x6a\xae\x91\xb5\x41\x4a\xb6\x24\x55\xa9\x92\xf7\x71\x4c\x52\x46\xaa\xa5\x92\x43\xa0\xb7\x06\x1a\xeb\xc3\x86\x2d\xbc\xe0\x27\x24\x16\x97\xb8\xdd\x0a\x17\x88\xad\xc2\xb1\x1a\x38\x24\xbd\x61\xa5\xf5\xa0\xd5\x3d\x4e\xe0\xb2\x89\x3e\x24\x21\x82\x35\x7a\xcd\x71\x82\x9c\x04\x7c\xcf\x04\xff\xf5\x92\xe4\x7d\xf9\xfe\xf6\x19\x73\x2d\xf3\x2a\xbd\xa4\x78\x0c\xfc\xfd\x1c\x7b\xd8\x28\x2f\x0b\x3e\xec\x5d\x6d\x3d\x92\xd6\x67\x87\xb7\x42\xad\x3b\xe1\xa2\xdc\x96\x68\x01\xf0\x0d\xb1\xa8\xb4\xc6\x2b\x1f\xc8\x7d\x32\xb7\x58\x07\x0b\x80\xbf\x66\x4d\x21\x85\x4b\x54\x66\x65\xaa\x58\x87\xc3\x24\x85\xd0\xfe\x13\x70\x51\xef\xee\x81\xf9\x3a\x7d\x3b\xc9\x9a\xd0\x88\x7b\xf4\xa0\x02\xd4\xc2\x49\x66\x72\xf4\xe4\xe4\x83\x85\xd6\xa1\x54\x65\x80\x15\x19\xee\x4a\x69\x0d\xb5\x68\x5b\x24\x54\xbe\x2d\xe0\x5d\x8d\x9d\x4e\xf5\x5a\xa0\x9a\xd6\x61\xa9\x3c\x32\xd7\xec\x12\x9d\x5e\x43\x7e\x55\x00\x74\xe1\x88\x78\x21\xba\xf7\xd0\x88\xb6\x65\xff\x60\x41\xc0\xfb\x37\x2f\x08\xb4\xf2\xec\x29\x5a\x67\x65\x2c\x11\x44\x33\x57\x8b\xa8\xc2\x3a\xd9\x71\x64\x7f\xc2\xd1\xbb\x75\x98\x53\x02\x3a\x91\xa2\x8c\x22\xa9\xa7\x88\x96\x21\x0f\xb4\xa4\x14\x3e\xeb\x06\x48\x6c\xd1\x48\x34\xe5\x9a\x50\x22\x23\xaf\x31\x25\x84\x93\x4d\x24\x8c\xad\xc6\xe4\x4e\x8d\x1c\x26\x28\x9d\x87\xca\x1a\xee\x83\x8b\x65\xd2\x62\xe7\x50\xe3\x52\x98\x50\x00\x3c\x29\xe0\x3f\x7a\xe1\xa3\xf0\x4a\xaf\xa1\xac\x85\x59\x20\xa8\xb0\x25\xd0\xce\x39\x28\xbf\x65\xdf\x6c\xb8\xda\x96\x29\x87\x9e\xe4\x70\x99\xd3\x98\xee\x1b\x7a\x58\x3a\xa2\xaa\xc8\x33\x99\xd8\xa0\xb3\xd1\x77\x49\x4f\x01\xf0\xcc\x9a\x47\x8f\x02\xcb\x1a\x0c\xae\xd8\x6f\xa4\x83\xc8\xed\x46\x23\xd1\x65\x63\x43\x49\x8b\x09\x70\xa8\x71\x0d\xd2\xb2\xb8\x72\xe6\x4e\xea\xe9\x03\x0a\x49\x0c\x88\x3e\xb9\xf5\x8c\xc8\x24\x25\xe4\xc4\x7d\x42\x59\xb3\xe8\xed\x52\x49\x3e\x45\x66\x9f\x9f\x00\x0b\x66\x16\x19\xc3\xb4\xb2\x25\xaf\x58\x43\xfe\xd5\x25\x2b\x24\x8f\x5c\xb0\x27\xc2\x0f\xa2\x69\x35\x4e\x38\xfb\x50\x25\xf6\x0e\xdb\xb3\xb2\x0a\xd9\x28\xcf\x12\x71\xb8\x50\x3e\x38\x91\xdc\xfb\x20\x6d\xa8\xe3\xbc\x28\x6d\x33\xa3\x6a\xc3\x19\x0c\xe8\x29\x27\x98\xcd\xb5\x9d\xcf\x48\x58\xc2\xe3\xf4\xaa\xb8\xfa\xcb\xac\x87\x35\x04\x35\x5b\x5e\xcd\xd8\x15\x14\x0b\xfb\xc5\x8b\x27\xdf\x7c\x03\xc5\xa3\x07\x91\xe5\x70\x18\x86\x23\x19\xf1\xde\xb8\x44\xdc\xdf\x51\xb2\xcc\x91\xf0\x30\x0c\x9e\x08\x85\xf4\x54\x9d\xaf\x1e\x71\xf6\xa3\xdb\x2a\x47\xb2\xde\x1e\x5b\x85\x29\x1e\xf7\xe9\x36\xc7\x86\xac\x01\xc2\x00\xa5\x55\x0e\xf3\xda\x24\x69\x43\x0e\xf8\x9b\x74\x9c\x02\x2b\x88\x1c\x18\xfe\xfd\xed\xab\x97\xb3\x7f\xb3\x09\x33\x10\x65\x89\xde\xa7\x74\xa7\x61\x27\xe6\x23\x05\x28\xdf\x65\x42\x6f\x69\xa5\x68\x84\x51\x15\xfa\x50\x64\x68\xe8\xfc\x4f\xd7\x3f\xef\xa8\x88\x4a\xfc\xea\x53\xd7\x2e\xb4\x2b\x9f\x88\xe9\xbf\x85\x95\x0a\x35\xa3\xd4\x5a\x99\x91\x5e\x31\xb2\x81\x4c\xc4\x66\x64\x23\x72\x7c\xb8\x81\x4b\xb2\x8e\xc1\xd1\xbf\x90\xd3\xff\xf5\x12\xbe\x5c\x71\x90\xe1\x18\x70\x99\x0e\xec\x6b\x0c\x8e\x0b\x59\x82\x9b\x83\x59\xf5\x83\x53\x8b\x05\x52\xb8\xe6\xb4\x99\x52\xd3\xaf\x28\x96\xa8\x0a\x8c\x1d\x6c\x66\x10\xc4\xcf\xde\x36\x77\x11\xf9\xe9\xfa\xe7\x4b\xf8\x72\x9b\x2e\x8a\x92\xf8\x01\xae\xc9\x81\x30\x65\xad\x95\x5f\x65\xa7\xea\xd7\x26\x88\x0f\x04\xb3\xa4\xc0\x64\xfa\x68\x57\x8b\x25\x82\xb7\x4d\x8a\x50\xd3\x94\xc6\x49\x58\x89\x35\xd1\xd0\xb1\x92\xa4\x2a\x38\x9e\xee\x54\x60\xef\x5e\x3d\x7b\x75\x93\x4e\x23\xb1\x2d\x4c\xe7\xe6\x2b\x45\xf5\x55\xf2\x9e\x54\x2b\xb0\xcc\x09\x91\x98\x84\x44\x39\x60\xf6\x88\xc9\x03\x57\x91\xb2\xf6\x3d\x36\x36\x42\xd7\x1f\x96\x43\xfb\xd5\x9c\xe3\xd0\xae\x71\xfd\x9f\x15\x1d\x23\x89\xe3\xba\x7f\x04\x71\x2f\x07\x7a\x77\x94\xb8\x8d\x3f\x24\xfa\xa4\x2d\x3d\x91\x56\x62\x1b\xfc\x8c\x42\xf7\x52\xe1\x6a\xb6\xb2\xee\x5e\x99\xc5\x94\x14\x6b\x9a\xa4\xed\x67\xdc\x24\x99\x7d\xc1\xff\x39\x9b\x16\x6e\x6f\x8c\x25\x88\x37\xff\x11\x54\xd1\x39\x7e\x76\x16\x51\x6e\x3b\x53\x1e\x43\xda\xdb\x2e\xc3\xdd\xf9\x96\xcc\x22\xa5\x67\xb9\xf9\x31\xf0\x64\x8d\x90\xc9\xd5\x09\xb3\xfe\xdd\x95\x96\x58\x17\x1d\x9d\xbd\x9e\xe6\x14\x60\x2a\x8c\x9c\xf6\x29\x6a\xb9\x3e\x8b\x57\x51\x8d\x32\x54\x4a\xb8\xff\x10\x55\x8e\xea\x2c\xab\x3c\xd0\x02\xa0\xa7\x15\x4e\x34\x18\xd0\xed\x49\x09\x54\xc0\x66\x6f\xa6\xb0\x45\xfd\xeb\x0e\x02\x94\xa2\x25\x01\xe5\x16\x99\x70\x4a\xcc\x95\xa6\x6c\x38\x39\xe1\xdd\x5e\xde\x1c\x53\x7a\x4c\x25\x5c\x50\x5c\x82\x53\xac\xdb\xd4\xd7\xfb\x12\x89\xe3\x29\x0c\xa1\x56\x89\xa8\xc3\xfe\xc5\x1d\xcc\x9f\xa5\xbd\xa9\xf3\x94\x3f\xcc\xf1\x34\x85\xb8\x9e\x39\xb4\xa5\x4f\x12\xe7\xa9\x96\x3e\x86\xe5\x49\x89\xec\xe2\x32\x0e\xdd\xfe\xc7\x86\xd5\x94\xc4\x9a\x05\xba\xe1\x56\xe2\x77\x6d\x57\x8c\xe5\x86\x04\xce\xbd\x73\x4f\xe3\x7c\x9c\x95\x6f\xb5\x58\xbf\x3c\xe8\xe4\x77\x71\xde\xec\xdf\xea\xa9\xcc\xd7\xf0\xfe\xd6\x9f\x8d\x06\x9a\xd8\x8c\x15\x71\xee\xf3\x68\xe5\x53\x36\xa0\xb5\x5d\x0d\x1a\xa5\xb7\xd5\x50\x0f\x3c\x06\xce\x02\x7e\x30\xb1\xe9\x72\x03\xa3\x74\x5f\xb2\xc6\x4d\x0d\xdd\xa5\x2d\x0c\x58\xa4\x2a\xe1\x00\x4a\x07\x0d\x69\x24\xb9\xdd\x16\xe1\x9c\x58\xef\xdd\xa1\x9a\x26\x06\x31\xd7\xe3\xa4\x92\xfd\x39\x15\xd4\xd5\x8e\x96\x64\x21\xa5\x64\x47\x82\xa8\x02\xba\xac\xee\x2a\x28\xa1\x93\xda\x6b\xdd\xf7\xb7\x87\xfd\xf7\xa3\xc8\xcf\xad\xd5\x28\xcc\xde\x3d\x87\x93\x86\x1d\xcc\x2f\x5f\xe6\x5c\x93\x8e\x1d\x36\xec\x72\x12\xdf\xe9\x57\xce\xd2\xba\xe6\x1e\x54\x4a\x23\x17\x62\xc3\x24\xfc\x2e\xdd\x51\x3c\x7d\xf5\xfe\xe5\xbb\x3b\xda\x6f\xfa\x5a\xb1\xf3\x5f\x9a\xe5\x2c\x38\xb5\xcd\x49\xf6\x7f\x99\xd4\x3b\xe5\x50\xda\x6a\x55\x0a\x7f\x03\xf0\xcb\x2f\x50\xb0\x27\xf4\x05\xc3\x83\x5f\x7f\xbd\x3c\x57\xbb\x73\x7b\xe0\x40\xe8\xd9\xe1\xc8\x9b\xbc\xb9\xcf\xbe\xf7\x08\x55\xf9\x1e\x26\x85\xec\x39\x6e\x39\x33\xa1\x75\xef\xcc\xfc\x84\x12\xfc\x55\x8d\xa1\x46\x37\xf0\x8a\xa4\x16\x3e\x56\x95\x3a\xe5\xef\x8e\x49\x39\xd7\x13\xa3\xc8\x7a\x97\xf6\x82\x92\x14\xe6\x99\x2c\xa6\x49\x0b\x93\x04\xbe\xc0\xe0\x01\x3f\x60\x19\x43\xd7\xa0\x4a\x55\xc4\x46\x95\x59\x87\x7d\xa7\x0b\xb7\x7d\xd7\x36\x17\x03\x03\xb3\xbf\x4b\x1d\x8b\x3b\xce\x57\xd2\x21\x5c\xa2\xf0\x49\x5c\xde\xe0\x07\xe5\x03\x71\x87\x18\xb3\x52\x1e\x41\x85\x47\x1e\xee\x24\xb6\xda\xae\xef\xce\xf6\x64\xec\x53\xa6\xbc\x6d\x14\x5b\xd6\x2d\x0e\x24\xbd\xf1\x4a\x04\xa1\x27\x89\x8b\xbb\xbb\x74\xea\xb9\xa8\x1d\xc9\x19\x8e\xb9\x23\xe2\xdd\x1e\x57\x27\xa4\xe4\x7b\x57\xa1\x5f\x1f\x0d\xe0\xdb\x99\x05\xc9\x61\x43\xac\x00\x8f\x4e\xa5\x3e\xf5\xeb\x5a\x78\xa6\x9f\xe4\x83\xbd\x5a\x97\x96\x8c\x3b\xec\x0f\x70\xa7\x52\x87\x96\x61\x8e\x12\x43\x3e\xbe\x11\x2d\xa1\xc5\x1f\x26\x35\xe1\xfa\x9d\x57\x3b\x35\x3b\x2f\x24\x3c\x3c\x6d\x8b\x11\x5d\x30\xf3\x01\xdb\xcc\x85\xae\x7d\xf1\x63\x9f\x63\x66\x0c\x0e\x46\xd9\xd3\x1c\x49\xcf\x31\xe7\x9c\x9e\x11\x01\x8c\x1e\xc6\xf6\x38\xa4\xed\x30\xc5\xd4\x65\x26\xd3\xc7\x03\x1e\x77\x1c\xd8\x5c\xdf\x3c\x24\x1c\x7c\xe0\xcb\x06\xb1\xb9\x59\x3c\xcc\x0b\x38\x2d\x94\x03\x28\x0e\x2e\x93\xfa\xab\x00\xca\x23\x6c\x95\x03\x23\x77\xf0\x59\x48\xb6\x2c\xe3\x9e\x6b\xa2\xed\x67\x9c\x54\xd2\x73\x5a\x36\xe9\x19\x29\xa1\xbc\x59\xf8\xfb\x11\x67\x8f\xe4\xd7\x19\x08\x9c\xce\x77\x76\x77\x1e\xf0\x53\x1f\x0f\x92\xfb\xa9\xb8\x58\x7f\x84\x9a\xbe\x72\x12\x53\x23\xb1\xb7\xd0\x2e\xfb\xf6\x71\xce\x3c\xda\xf4\xb8\xb4\x30\xb3\xe4\x2f\x36\x19\x0b\x8f\x9a\x48\xb0\xf1\xa0\xb7\x18\x52\x70\x82\x87\x23\xf8\x61\xa2\xd6\x9c\x2b\x42\x70\x11\xcf\xcc\x38\x4f\x31\xea\x8f\x66\xd1\xa7\x84\xb3\x07\x51\x67\xe3\xda\x29\x04\x6d\xbc\x0e\xfd\x7c\x88\xc0\x09\x7e\x1e\x39\xfa\x80\xa5\x6d\xe1\xf3\x62\x53\xb8\xa4\xfd\x20\x96\x42\xe9\x2e\x37\x65\x9e\x1d\xbd\x7c\x87\xd1\xa5\xfc\x3b\xe1\xef\x53\x35\xbc\xd0\x76\x2e\xf4\x04\x5a\xab\xd7\x8d\x75\x6d\xad\x4a\x50\x14\x5b\x9b\xad\xd9\x16\xad\xa1\x8d\x73\xad\x4a\xbd\x1e\x60\xc5\x58\x9e\x11\x80\x0f\xf7\x45\x47\xe8\xfe\x31\x4f\x78\xf2\xe3\x87\xc3\x10\x47\x38\xc4\xb3\x10\x7c\x2b\xe5\x13\x07\xfa\xfb\x60\x62\x1f\x81\xf2\xb9\x8d\xcd\xb5\xa6\x87\x98\xfa\xfa\x4b\xab\x24\xac\x9c\xe2\x71\x96\x92\x07\xd3\x20\x9a\x59\x23\x9c\xaf\x85\xd6\xdc\x93\xe7\x2b\x4c\x56\x7a\xee\x78\xb7\xc2\x79\x84\x12\x1d\x87\xf6\x7c\x8d\x99\x6e\x04\x09\x48\xbe\x0d\xe4\x73\x7f\x54\x46\xa6\xdb\x4e\x69\x57\xc6\x2b\x89\xfd\x7d\xbe\x68\x5b\x67\x45\x59\x83\xe2\x3b\x45\x31\xb8\x85\x4e\xb7\xc7\x94\xdd\xf3\x85\xb1\x58\xf6\x97\xa5\x39\xff\x45\xf0\xa4\xfd\xff\xf0\x36\xd9\x81\xa7\x68\xaa\x3a\x24\xe7\x58\xda\xa6\xbb\xf7\xb4\xd1\xf7\x13\x59\x5d\xdd\xc0\x04\x38\xbe\x5f\x6c\xd4\xa2\x0e\xe0\x70\xa9\xbc\x0a\xbb\x88\x0d\x9b\xea\x9d\xd9\xf3\x96\xee\x04\x03\xca\xfb\x78\xb0\xf8\x18\x13\x33\x8f\x8d\x9a\x1c\x10\xf7\x20\xa2\x8b\xb6\xed\x2f\xbc\x32\xba\x96\xea\x21\xaa\x89\x1d\xb6\x76\xd2\xd1\xdc\xdf\xac\xf0\x4d\xae\xc3\x12\xcd\xb1\xc0\x34\xca\xa9\x4b\x6b\x8e\x86\xf9\xd3\xb5\x17\x3d\x95\x08\x42\x7f\x3a\x98\xae\x56\x3c\xd6\x04\x82\xb1\x94\xd9\x6d\xcf\x75\x8e\x64\x3a\x10\xe7\x8b\x07\xac\x39\x14\x06\x47\x13\xd2\x8a\xf2\x5e\x2c\x8e\x32\x64\x8b\x00\x54\x5c\x62\x13\x6e\xdd\xb7\x6c\x60\x93\x74\x2b\xdd\xbf\xab\xac\x96\xe8\xa8\x2a\x17\x06\xde\xbf\x79\xc1\x13\x0f\x79\x2d\x08\x37\x17\x5a\x17\xdd\xa8\x41\xcf\x89\x61\x93\x66\xc2\x73\xa3\x65\xd0\xa9\x03\xe8\xd0\x5b\xbd\xc4\xdc\x21\x48\x70\xba\x29\x08\x47\x5e\x63\x70\x01\xd6\xfb\x80\xfc\x91\xdc\x9c\x40\xa8\x1e\xcb\x56\x46\xf2\x2c\xd7\xea\x9f\xac\x46\x3d\xa4\xe7\xea\x70\x0f\xec\x81\x0c\x86\x77\x9f\xdb\xad\xa2\x2f\x79\x9a\x61\xd3\x48\xba\xeb\x96\xfd\x5d\x96\xc8\x57\x69\xa0\xb0\xbb\xfe\x40\xf8\xba\x15\x0e\x4d\xf8\x7a\x33\x1e\x96\xa6\x96\x02\x57\x06\x9b\xb6\x04\xc3\xef\x06\xcb\x5a\xdb\x46\x3e\x95\x21\x94\xb5\xd2\xf2\xeb\xbe\x53\x51\x50\xa4\x29\xfa\xfe\xfa\xbe\x90\xfa\x71\x4c\x52\x87\xbb\x0c\x30\x2e\x9d\xdf\xce\x94\x54\x8b\x79\x30\x30\x0d\x0a\xa5\xac\x4e\x24\x12\x3b\xea\x53\x8f\x89\x37\x77\xf9\x45\x57\x97\xf1\x28\x4c\xbe\x33\x92\x47\xcf\x1d\x5f\x13\xa1\x59\x9e\xd2\x81\xf4\x7c\x44\x45\x52\xfd\xd6\x00\xef\xf1\x44\x81\xf1\xb1\xf0\x8e\x64\x4e\x67\x00\x1c\x59\x53\x8d\xaa\x23\x60\x64\xe9\xd5\xda\xa3\x14\x8c\xc2\xbd\x9f\x4a\xfe\x44\x2d\x1f\xc9\xa8\xdf\x94\xfe\x95\x30\xe1\x07\x37\xc2\x13\x1e\x0f\xcf\x27\x44\x77\x66\x57\xaf\xf7\x7e\xe7\x77\xf6\x8e\xf2\x74\x3b\xc7\xee\x0e\x4b\xa5\x48\xd7\xe9\x71\x9b\x89\xaa\x60\xe1\x3f\xbf\xff\xdb\x8b\x0d\x5a\xb0\xe3\xae\x37\x0b\x39\x7c\x92\xab\xa1\x17\x83\xb9\x2d\x99\x27\xc2\xa9\x58\x39\x34\xb9\xbb\x97\x51\xb1\x5d\x38\x21\x49\xf0\xcf\x9d\xdd\x73\x0d\xb5\x45\xcc\xfb\xad\xcd\x4c\x4c\xaa\x1a\x76\x4a\x36\xbf\x99\x3b\x4e\xf0\xb1\x9f\x9e\xfb\x8d\x8a\xbb\xcf\x93\xc4\x9f\x27\x89\x3f\x4f\x12\x7f\x9e\x24\xfe\x3c\x49\xfc\x79\x92\xf8\x93\x27\x89\x4f\x67\xe4\xa7\xa6\x89\x3f\x75\x9e\x78\x44\x96\x76\x62\xa6\xf8\xf3\x54\xf1\xe7\xa9\xe2\x7f\xa6\xa9\xe2\x11\x1a\x7f\xac\x0e\xfc\x67\x98\x2d\xfe\xc4\x3e\xff\x9f\x70\xc2\x78\x24\x45\x47\xa6\x8c\xff\xb4\x73\xc6\xa3\x26\x99\x46\xcc\x1a\xff\xff\x99\x36\x1e\xc1\xb1\x83\x13\xc7\x7f\xc2\x99\xe3\xdf\x6b\x86\x68\xf9\xd1\x7f\x10\x7c\xe8\x6f\x9c\x83\x08\xd1\x7f\xc4\x5f\x39\xf3\xfe\xad\xbf\x73\xb6\x73\x8f\x6e\x39\xfa\x0f\x9d\xf7\x22\xf2\xe0\x65\x02\x39\x68\x1b\xf9\x60\xa9\x36\xce\x6f\x36\x68\x53\x7e\xd0\x06\x94\x2f\x77\xff\xb9\x87\xcb\x34\x5c\xd8\xfd\xfb\x0d\xfc\xb3\xb4\x26\xb5\x62\xfc\x0d\xfc\xf4\xf3\x05\xe4\x36\x6b\xd7\x60\xe0\x97\xff\x1b\x00\x00\xff\xff\x1e\x49\xca\xd4\x3f\x43\x00\x00")

func configCrdsKudoDev_operatorversionsYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_operatorversionsYaml,
		"config/crds/kudo.dev_operatorversions.yaml",
	)
}

func configCrdsKudoDev_operatorversionsYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_operatorversionsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_operatorversions.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"config/crds/kudo.dev_instances.yaml":        configCrdsKudoDev_instancesYaml,
	"config/crds/kudo.dev_operators.yaml":        configCrdsKudoDev_operatorsYaml,
	"config/crds/kudo.dev_operatorversions.yaml": configCrdsKudoDev_operatorversionsYaml,
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
	"config": &bintree{nil, map[string]*bintree{
		"crds": &bintree{nil, map[string]*bintree{
			"kudo.dev_instances.yaml":        &bintree{configCrdsKudoDev_instancesYaml, map[string]*bintree{}},
			"kudo.dev_operators.yaml":        &bintree{configCrdsKudoDev_operatorsYaml, map[string]*bintree{}},
			"kudo.dev_operatorversions.yaml": &bintree{configCrdsKudoDev_operatorversionsYaml, map[string]*bintree{}},
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
