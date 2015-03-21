package cbcluster

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _data_couchbase_node_service_template = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x93\x5f\xcb\xd3\x30\x14\xc6\xef\xfb\x29\x72\x21\x0c\x84\xbc\xf1\x42\x6f\x5e\xe9\x45\x9d\x7d\xa5\x37\xeb\x68\xeb\x10\xc6\x28\x59\x7a\xe6\xc2\xd2\x24\xe6\x4f\x37\x19\xfb\xee\x66\xeb\x58\x75\x45\x27\xe2\x5d\xf8\xf1\x3c\x4f\xce\x73\xe0\x2c\x3f\x4b\xee\x56\xd1\x47\xb0\xcc\x70\xed\xb8\x92\x31\x53\x9e\x6d\xd7\xd4\x42\x2d\x55\x03\x51\xb2\x71\x60\xe2\x46\xb1\x1d\x98\x27\x0b\xa6\xe3\x0c\xa2\x02\xbe\x79\x6e\xc0\xde\xf3\x5e\x0c\x8e\x35\x63\xe9\x2f\x34\x5a\x96\xfd\x6b\x15\x55\xbc\x05\xe5\x5d\xe9\xa8\x71\x25\xb0\xf8\xcd\x40\x94\xee\x41\x2a\x3b\x6e\x94\x6c\x41\xba\x17\x2e\x20\x26\x21\x8b\xc0\x00\xa3\xf4\x00\xec\x12\x30\x37\x10\x63\xe2\xad\x21\x6b\x2e\x49\x3f\x1d\xda\x71\x21\xd0\xad\xd6\x03\xb1\x69\x7f\x27\xbd\x57\x6a\x1f\x62\x9d\x80\xef\x0d\xc8\x77\x7c\x7f\x20\x37\x1f\x3e\xd7\x04\x83\x8f\x47\xf4\x34\xfd\x50\x2f\xd2\xa2\xcc\xf2\x19\x3a\x9d\x9e\x2f\x24\x9f\x55\x49\x36\x4b\x8b\xba\x4a\x3e\x05\xf8\xcf\xbf\x30\xe1\x6d\xd8\x37\xfe\xaa\x1e\xe4\xc6\x97\xc0\xe0\xd9\x22\xcc\xd0\x64\x54\xd9\x4b\x84\xb1\xa4\x2d\x0c\xd5\x11\xee\x10\x51\xda\x0d\xdf\x91\x8e\x9a\xe7\x31\x3a\x3b\xc1\xc5\x5b\x65\xdd\xff\x58\x06\xfa\xc9\x77\x1e\x7d\x72\x6d\xa1\xf4\xdf\x95\xf8\xf3\x28\x0f\x36\x86\xbc\x6e\xa8\x03\xbc\x37\x54\xeb\x90\x39\x32\x22\x03\xad\xea\x00\x53\xd9\x60\x03\x6b\x2a\xa8\x64\x61\x55\x58\x28\x46\x05\xe6\x1a\xbd\x9a\xe6\x45\x9a\x97\xf5\xbc\xc8\x16\x49\x95\xd6\xd9\x7c\xf1\xf6\x3d\xb2\xbe\x51\xe8\x3a\xa7\x0d\x55\x86\xe0\x49\xb8\x84\x2f\xf8\x45\x00\x84\x2b\x9c\x2a\xb9\x11\x9c\x39\x7b\x77\x83\xaf\x6f\x67\xf3\x23\x00\x00\xff\xff\x6f\x86\x56\xbc\xaf\x03\x00\x00")

func data_couchbase_node_service_template_bytes() ([]byte, error) {
	return bindata_read(
		_data_couchbase_node_service_template,
		"data/couchbase_node@.service.template",
	)
}

func data_couchbase_node_service_template() (*asset, error) {
	bytes, err := data_couchbase_node_service_template_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/couchbase_node@.service.template", size: 943, mode: os.FileMode(420), modTime: time.Unix(1426775948, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _data_couchbase_sidekick_service_template = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x52\x4d\x6b\xdb\x40\x14\xbc\xef\xaf\xd8\x43\x21\xa7\x8d\x7a\x68\x2f\x85\x85\x2a\xa9\x52\x74\x88\x64\x24\x39\x14\x8c\x11\xca\xea\xb9\x7a\x68\xb5\xab\xee\x47\x92\x62\xfc\xdf\x2b\x59\xc5\xae\x2d\xd3\x9a\xde\x96\xd9\x99\x79\xf3\x86\xb7\x5a\x2a\x74\x6b\xf2\x05\xac\x30\xd8\x3b\xd4\x8a\x0b\xed\x45\xf3\x5c\x59\x28\x2d\xd6\xd0\xa2\x68\x49\xb8\x71\x60\x78\xad\x45\x0b\xe6\xd6\x82\x79\x41\x01\x24\x83\x1f\x1e\x0d\xd8\x73\x7c\x22\x83\x13\xf5\x9c\x7a\x82\x4e\xc4\x8d\x04\x70\x73\xe6\x29\x7c\x87\xaa\xb6\x85\xfe\x23\x9b\xd2\x35\x7c\xde\x6e\xe9\xed\x32\x89\x8b\x32\x59\x3e\xde\x45\x19\xdd\xed\xce\xcc\xaf\xe7\x93\x55\x3e\xbd\xd6\xa4\xc0\x0e\xb4\x77\xb9\xab\x8c\xcb\x41\xf0\xf7\x24\x52\x2f\x68\xb4\xea\x40\xb9\x07\x94\xc0\x83\x61\x8f\x00\x8e\x20\x89\xde\x40\xec\xf9\x0b\x03\x9c\x05\xde\x9a\xe0\x19\x55\x30\x35\x43\x5b\x94\x92\x1e\xa2\xb0\x43\xad\x7f\x57\x99\xee\x9f\x9a\x73\x49\xef\x87\x41\x4e\xc2\xcf\x1a\xd4\x47\x7c\x7d\x0b\x8e\x06\x42\x7a\x3b\x34\xc2\xbe\xeb\x4f\x63\x0b\xf7\x69\x52\x84\x71\x12\x65\x65\x11\x7e\x1d\x7a\x38\xfa\xf2\xbd\xe1\xa0\x69\x28\x13\xf4\x66\x96\xca\x2b\xca\x98\xaa\x3a\xb8\x90\x6e\xfc\x01\xc7\x1b\x6d\xdd\xff\xc5\xa0\xbe\xaf\x2b\x07\xec\xd5\x54\x7d\x3f\x4c\x9b\x09\xa9\x1d\x33\xb2\x8b\xa3\xa5\x16\x95\x64\xd8\xf3\x77\xf7\x69\x16\xa5\x79\xb9\xc8\xe2\xa7\xb0\x88\xca\x78\xf1\xf4\xe1\xe6\xf7\x86\xba\x9f\xb5\x66\x07\xf0\x52\xd5\x64\xf5\x8d\x3d\x8c\x67\xb8\x26\x8f\x95\x68\x50\x41\xba\xb9\xfe\xa2\x7e\x05\x00\x00\xff\xff\x28\xcb\x1c\xaa\x5a\x03\x00\x00")

func data_couchbase_sidekick_service_template_bytes() ([]byte, error) {
	return bindata_read(
		_data_couchbase_sidekick_service_template,
		"data/couchbase_sidekick@.service.template",
	)
}

func data_couchbase_sidekick_service_template() (*asset, error) {
	bytes, err := data_couchbase_sidekick_service_template_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/couchbase_sidekick@.service.template", size: 858, mode: os.FileMode(420), modTime: time.Unix(1426775948, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _data_sync_gw_node_service_template = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x54\xc1\x4e\xdb\x40\x10\xbd\xef\x57\xcc\x01\x89\xb6\xd2\x7a\x7b\xe9\x85\xca\x87\x28\x0d\xc8\x52\x31\x55\xe2\x88\x56\x51\x64\x99\xf5\xc4\xd9\x62\xef\xba\xbb\xe3\x98\x08\xf1\xef\x5d\x63\x20\x35\xa1\x42\x41\xdc\xac\xf1\x9b\x99\xf7\xde\xcc\xec\x62\xae\x15\x2d\xd9\x37\x74\xd2\xaa\x9a\x94\xd1\xa1\xdb\x6a\x99\x16\x6d\xaa\x4d\x8e\x6c\xb4\x22\xb4\x61\x6e\xe4\x35\xda\xc0\xa1\xdd\x28\x89\x6c\x8a\x7f\x1a\x65\xd1\x3d\x8f\xf7\x60\x24\x99\xef\x43\x07\xd1\x1e\xb8\x2a\x11\x69\x1f\x39\x0c\xb3\xc5\xac\xff\x5a\xb2\x44\x55\x68\x1a\x9a\x51\x66\x69\x86\x32\xfc\xcc\x26\x7a\xa3\xac\xd1\x15\x6a\x3a\x55\x25\x86\xc2\x77\x11\xb8\x0b\xb2\xc9\x0d\xca\x7b\xfc\x0f\x8b\x21\x17\x8d\xb3\xe2\x4a\x69\xd1\xf3\x86\x6b\x55\x96\xf0\x20\xf7\x15\xa8\xad\x5e\x06\x3e\xc7\xd5\x8d\x2f\x49\x25\x6e\x73\xd4\x5f\x54\x7b\x23\xba\x2c\x5e\x64\x84\x6d\xb6\xe5\xd2\x58\x34\xee\xe4\xf6\x16\x82\xf1\x45\x9c\x8c\xa2\x78\x32\x4d\x93\xd1\x19\xdc\xdd\x1d\x58\x56\x9a\x46\xae\xaf\x32\x87\x5c\x96\x8d\xf3\x6e\xf2\xc2\xbc\xa1\xae\x6d\x34\x70\xae\x91\xc2\xb5\x71\xf4\xb6\x0e\xd0\xd4\xb9\xd7\xc7\x5b\x9b\xd5\xb5\xaf\xb9\x97\x08\x5d\xd6\xe5\x28\x4a\xd2\x79\x9c\x44\xdf\xd3\xe9\x3c\x8e\xa3\xf8\x60\x72\x7c\x03\x62\x6d\x2a\x14\x9d\x8b\x27\xbb\xcf\xf7\x61\xdd\xcf\xa9\xf5\x23\xd2\x2b\x55\x80\xc5\xd6\x2a\x42\xdf\x3f\x47\x47\x4a\x67\xdd\x69\xfc\xd3\x5f\x04\xc3\x84\xe0\xb7\x33\x7a\x27\x27\xbc\xd7\xe1\x79\xac\x81\x4b\x38\x9e\xfd\x8a\xc7\xe9\xd9\x65\x3a\xbe\x38\x3f\x8f\x92\xf0\xe8\x43\x77\x0d\x92\x4a\x28\x90\x60\xc7\x39\x90\xa6\x1a\x2c\x8c\xff\x55\x55\x8a\x3e\x7e\x05\xf8\x8f\x39\x59\x85\x8f\x8b\x79\xb0\x57\x2f\xac\xe6\x93\x0d\xae\x53\xd1\x71\x3f\x1a\x72\x07\x5e\xc0\x6b\x36\x1c\x3f\xf8\x60\xea\xbd\x91\x3a\x1f\x7c\x3a\x24\xb6\xf8\xc9\x4f\xbb\x63\x5f\xb2\xb1\xcf\x2e\x95\x24\x37\x78\x7d\x3e\x3d\xbe\x02\x7f\x03\x00\x00\xff\xff\x6d\xf0\x49\xd8\xa6\x04\x00\x00")

func data_sync_gw_node_service_template_bytes() ([]byte, error) {
	return bindata_read(
		_data_sync_gw_node_service_template,
		"data/sync_gw_node@.service.template",
	)
}

func data_sync_gw_node_service_template() (*asset, error) {
	bytes, err := data_sync_gw_node_service_template_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/sync_gw_node@.service.template", size: 1190, mode: os.FileMode(420), modTime: time.Unix(1426981301, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _data_sync_gw_node_service_template_ = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x54\xc1\x6e\x13\x31\x10\xbd\xfb\x2b\x7c\xa8\x54\x40\x72\xcc\x85\x4b\xd1\x1e\x50\x68\x11\x87\x02\x22\x45\x80\xa2\x28\xda\x7a\x27\x9b\xa1\xf6\x78\xb1\xc7\xd9\xe6\xef\xf1\x76\xdb\x86\xcd\xb6\x6a\x72\x9b\xbc\x3c\xcf\x7b\x6f\x3c\xde\xf9\x0f\x42\x5e\x88\x8f\x10\x4d\xc0\x86\xd1\x53\x11\xb7\x64\x96\x75\xbb\x24\x5f\x81\xf8\xb0\x62\x08\x45\xe5\xcd\x0d\x84\x49\x84\xb0\x41\x03\xe2\x3b\xfc\x4d\x18\x20\xee\xe3\x3d\x19\xd8\x54\x63\xea\x00\xed\x89\x2b\x0b\xc0\x63\xe6\x10\x16\xf3\x59\x5f\x2d\xc4\x15\x3a\xf0\x89\x67\x5c\x06\x9e\x81\x29\xde\x8a\x73\xda\x60\xf0\xe4\x80\xf8\x02\x2d\x14\x3a\xab\x68\xd8\x81\xe2\xfc\x16\xcc\x1d\xff\x5b\x80\x42\xe9\x14\x83\xbe\x46\xd2\xbd\x6f\x79\x83\xd6\xca\xfb\xb8\x2f\x50\x83\x7b\x9a\xb8\xcf\x6b\x52\x6e\xc9\x16\xb6\x15\xd0\x3b\x6c\x6f\x75\x77\x4a\xd5\x25\x43\x5b\x6e\x95\xf1\x01\x7c\x3c\xb3\xf9\x67\xe4\x23\x1b\x19\x9f\xcc\xfa\xba\x8c\xa0\x8c\x4d\x31\xcf\x4f\xd5\xfe\xa0\x4e\x21\x91\x54\x8a\x80\x8b\xb5\x8f\x7c\xa0\x39\x99\x9a\x2a\x17\xaa\x0d\x65\xd3\xe4\x26\x23\x75\xd9\x96\xc8\x2a\x11\xa3\x55\x59\x81\x90\xea\xa3\x6c\xa8\x8d\xd4\x6b\xef\x40\x77\xb2\x67\xbb\xf2\xd0\xcc\xfb\x06\xfb\x28\x6d\x4e\x41\x2b\xac\x65\x80\x36\x20\x43\x56\xac\x32\x19\xa9\xec\x56\xfb\x3f\x45\xdd\xf3\x26\x7f\xa2\xa7\x9d\xef\xe2\xce\x70\x16\x5c\x4b\x65\xe4\xe9\xec\xf7\x97\xe9\xf2\xd3\xcf\xe5\xf4\xeb\xe5\xe5\xe7\xab\xe2\xe4\x55\xb7\xc4\x86\xad\xac\x81\xe5\xce\xdc\xc4\x78\x37\x18\x65\xfe\xcb\x39\xe4\xd7\xef\xa5\x7c\x66\x0a\xa5\x83\x87\x7d\x3a\x7a\x28\x4f\x5c\xda\x63\xfa\xd8\xa5\xe8\xbc\x9f\x0c\xbd\x4b\x55\xcb\x67\xd2\x9f\xde\xc7\xf7\xcd\xe8\xca\x62\x06\x1f\xd7\x5e\xcc\x7f\xa9\x8b\xee\x69\x2e\xc4\x34\x9f\xb6\x68\x38\x0e\xbe\x15\x6f\x1e\xde\xec\xbf\x00\x00\x00\xff\xff\x27\x60\x83\xe6\x54\x04\x00\x00")

func data_sync_gw_node_service_template__bytes() ([]byte, error) {
	return bindata_read(
		_data_sync_gw_node_service_template_,
		"data/sync_gw_node@.service.template~",
	)
}

func data_sync_gw_node_service_template_() (*asset, error) {
	bytes, err := data_sync_gw_node_service_template__bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/sync_gw_node@.service.template~", size: 1108, mode: os.FileMode(420), modTime: time.Unix(1425612410, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _data_sync_gw_sidekick_service_template = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x52\xcb\x6a\xdc\x30\x14\xdd\xeb\x2b\xb4\x28\x64\xa5\xb8\x8b\x76\x53\x10\xd4\x49\x9d\xe2\x45\xec\xc1\xf6\x84\xc2\x10\x8c\x22\xdf\x8c\x2f\x96\x25\x57\x8f\x38\x21\xe4\xdf\xab\xc9\xa4\x49\x3b\x86\x76\xc8\x4e\x1c\xce\x4b\x87\xbb\x59\x6b\xf4\xd7\xe4\x1b\x38\x69\x71\xf2\x68\x34\x77\x0f\x5a\xb6\xdb\xb9\x75\xd8\xc1\x80\x72\x20\xe9\xad\x07\xcb\x3b\x23\x07\xb0\xa7\x0e\xec\x1d\x4a\x20\x15\xfc\x0c\x68\xc1\x1d\xe2\x7b\x32\x78\xd9\x2d\xa9\x7f\xa1\x67\xa8\x3b\xd7\x98\xd7\x38\x6d\x3a\xf8\xfa\xf8\x48\x4f\xd7\x45\xde\xb4\xc5\xfa\xf2\x2c\xab\xe8\xd3\xd3\x81\xf1\xb1\x6c\xb2\xa9\xf7\xaf\x6b\xd2\xe0\x08\x26\xf8\xda\x0b\xeb\x6b\x90\xfc\x23\xc9\xf4\x1d\x5a\xa3\x47\xd0\xfe\x02\x15\xf0\x24\x16\x4b\xe0\x0d\x24\xd9\x3d\xc8\x67\xfe\xca\x02\x67\x49\x70\x36\xb9\x41\x9d\xec\xbf\x4a\x07\x54\x8a\xee\x8a\xb0\xed\xcc\x5e\x57\xfa\xb7\xc6\x8e\xff\x51\x1c\x0a\xa6\x10\x43\xbc\x82\x87\x0e\xf4\x67\x9c\xef\x13\x69\x82\xec\x6f\x84\x03\x26\x55\x70\x71\x0b\xb6\x35\x5f\x76\x0b\x9c\x97\x45\x93\xe6\x45\x56\xb5\x4d\xfa\x3d\x6e\xf0\xe6\xcb\x9f\x0d\xa3\xa6\xa7\x4c\xd2\x93\x45\xa7\xa0\x29\x63\x5a\x8c\xb0\xe8\xb6\xc3\xc1\xf3\xde\x38\xff\xbe\x12\x34\x4c\x9d\xf0\xc0\x66\x2b\xa6\x29\x66\xfd\x0e\x78\x91\x51\x25\x82\x96\xfd\x9f\x79\xca\x48\xa1\x18\x4e\xfc\xc3\x79\x59\x65\x65\xdd\xae\xaa\xfc\x2a\x6d\xb2\x36\x5f\x5d\x7d\x3a\x79\xf9\x94\x99\x16\x43\xb9\x08\x2e\xb7\x25\x9b\x1f\xec\x42\x01\xc4\xdb\xbe\x14\xb2\x47\x0d\xe5\xed\xd1\xc7\xf3\x2b\x00\x00\xff\xff\x90\x04\x86\x17\x15\x03\x00\x00")

func data_sync_gw_sidekick_service_template_bytes() ([]byte, error) {
	return bindata_read(
		_data_sync_gw_sidekick_service_template,
		"data/sync_gw_sidekick@.service.template",
	)
}

func data_sync_gw_sidekick_service_template() (*asset, error) {
	bytes, err := data_sync_gw_sidekick_service_template_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/sync_gw_sidekick@.service.template", size: 789, mode: os.FileMode(420), modTime: time.Unix(1426775948, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _data_sync_gw_sidekick_service_template_ = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x4f\x4f\xc2\x40\x10\xc5\xef\xfb\x29\x7a\x30\xe1\xb4\x56\x13\xbd\x98\x6c\x22\x6a\x49\x38\x18\x08\x45\x62\x42\x48\xb3\x6c\x07\x3a\xe9\x76\xb7\xee\x1f\x0a\xdf\xde\xe5\x8f\x88\x34\xa8\xb7\xcd\xdb\x37\xef\x37\x33\x99\xe9\x9b\x42\x37\x23\x2f\x60\x85\xc1\xda\xa1\x56\xcc\x6e\x94\xc8\x96\x4d\x66\x31\x87\x12\x45\x49\xba\x0b\x07\x86\xe5\x5a\x94\x60\xae\x2d\x98\x15\x0a\x20\x23\xf8\xf0\x68\xc0\x9e\xeb\x7b\x33\x38\x91\xb7\xad\x3f\xd4\x27\x54\xb9\x1d\xeb\x23\x4e\xe9\x1c\x1e\x6f\xcf\x72\x2e\x7c\x92\x69\xba\x7f\xcd\xc8\x18\x2b\xd0\xde\xa5\x8e\x1b\x97\x82\x60\x37\x24\x51\x2b\x34\x5a\x55\xa0\x5c\x0f\x25\xb0\x38\x60\x63\xf8\x16\x49\xb2\x06\xb1\xf3\x0f\x0d\x30\x1a\x7b\x6b\xe2\x39\xaa\x78\x3f\x48\x54\xa2\x94\xd1\x96\x4b\x97\x0d\x3d\xee\xe0\xf7\x1a\x53\xfd\x51\x71\x5e\x50\xfb\x00\x71\x12\x36\x39\xa8\x7b\x6c\xd6\xb1\xd0\x5e\x14\x73\x6e\x81\x0a\xe9\x6d\x18\x9d\x2e\xf5\x83\xe4\x0e\xec\x49\xbf\x6c\x17\x11\x5c\x45\x44\x45\xd4\x69\x75\xe1\x55\x44\xa9\xe2\x15\xb4\xba\xd9\xea\xe0\x58\xa1\xad\xfb\x2f\x36\xf2\x75\x1e\x1e\xb4\x31\xbc\xae\x43\xfa\x57\xe4\xc1\x18\x49\xee\x95\x28\x4e\x09\x52\x0b\x2e\x29\xd6\xec\xea\x79\x30\x4a\x06\x69\x36\x1c\xf5\x27\xdd\x71\x92\xf5\x87\x93\xbb\xce\x61\x0c\x5d\xb7\x96\x61\x83\xd8\xde\x1f\x99\xbe\xd3\x9e\x04\x08\xd7\xf9\xca\x45\x81\x0a\x06\x8b\x4b\xf7\xf0\x19\x00\x00\xff\xff\xf5\x65\x76\x7c\xc6\x02\x00\x00")

func data_sync_gw_sidekick_service_template__bytes() ([]byte, error) {
	return bindata_read(
		_data_sync_gw_sidekick_service_template_,
		"data/sync_gw_sidekick@.service.template~",
	)
}

func data_sync_gw_sidekick_service_template_() (*asset, error) {
	bytes, err := data_sync_gw_sidekick_service_template__bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/sync_gw_sidekick@.service.template~", size: 710, mode: os.FileMode(420), modTime: time.Unix(1425612431, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _data_test_fleet_api_units_json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9d\xeb\x53\xdb\xb8\x1a\xc6\xbf\xf3\x57\x68\x32\x74\x68\xcf\x1c\xe3\xdc\x48\x48\x98\xcc\x94\x52\x7a\xca\x87\x1e\x18\x42\xbb\xbb\xb3\xd9\x61\x1c\x5b\x4e\xb4\x38\x76\xd6\x97\xa4\x0c\xc3\xff\xbe\x92\x43\x42\x42\x6e\xd4\x97\x56\x82\x87\x0f\x94\x24\x8a\xfd\xbc\xaf\x94\x57\xcf\xaf\x96\x9c\xbb\x1d\xc2\x7f\x0a\x91\xcb\xc2\xa0\xd0\xfc\x33\x7e\x24\x7e\xee\x66\x7f\xc5\xaf\x9b\x91\xef\x53\x37\x6c\x87\x46\x48\x0b\xcd\x02\x73\x0d\x33\x64\x23\x5a\xf8\xef\x62\x33\x8b\x06\xcc\xa7\xd6\xb6\x66\xae\x31\x10\x2f\x9b\x5d\x3b\xb8\x36\x5c\xd7\x8b\x5c\x93\xbe\xdf\x0f\xa8\x3f\x62\xe6\x52\x63\x6f\x18\x32\xcf\x5d\x10\xb7\x5a\xe4\xd3\xc3\x7f\xa4\x81\xe9\xb3\xf8\xdd\x4f\x8e\x39\x6b\x19\x50\x33\x7e\xb9\x59\xf8\xca\x13\xb0\xae\xd5\xc8\x70\x22\x71\xc0\xe3\x07\xa9\xe4\xe4\xc3\xa7\x76\x61\xa9\xed\xfd\xf2\xdb\x37\x0b\xfc\xc0\x5c\x2b\xb8\xf2\xb2\x11\x17\x27\xd3\xf5\x2c\xfa\xfe\x0d\x9b\xa5\x32\xbd\xc6\x63\x3b\xa4\xbe\xd4\x0a\x4f\xbf\x53\x93\x0f\x38\x7f\xed\xf9\x1f\x55\xb6\x57\x8e\xb0\x25\xa1\x7a\x97\xb9\x7a\xd0\x27\x9a\x49\x3a\x85\x71\x9f\x39\x94\x84\x7e\x44\x8f\x88\xe5\x11\x1a\x9a\x96\x19\x3a\x24\xa0\x21\xd1\x4d\x2f\x32\xfb\x5d\x23\xa0\xfb\xa6\x37\xd0\x45\x7c\xfa\x7c\x90\x64\xef\x8e\x74\x3a\x9d\x42\xdf\x0b\x42\xf1\x6f\x33\x7e\xf4\xe6\xb3\xf8\x7d\xbf\x47\x34\x2d\xe4\x07\xaa\x15\x8f\x02\x87\xd2\x21\xa9\x1e\x1c\x59\x9e\x4b\x3b\x85\xcc\x72\xe2\x0d\xb3\x4b\x49\x14\xf8\x71\x5a\xa6\xf1\xfb\x83\xad\xe1\x67\x10\xc7\x17\xc3\xec\x33\x97\x9e\xdb\xdb\x03\xf9\x5d\xfb\xc4\xd3\x98\xd5\x20\x5c\x78\xe6\xaf\x9d\x15\x01\x6c\x2e\x90\x8e\xc1\x0b\x45\x9f\x5a\x5b\x0a\xe4\xba\x66\x83\x49\xdc\x67\x1f\x79\x9b\x62\xa9\x4a\x6b\xa5\x8a\x61\x1f\xd6\x0f\xaa\xf5\x6a\xc5\xa8\x15\xeb\x8d\x46\xa5\x6e\x35\xea\x76\xd9\xac\xd4\x9e\x55\x5c\x4b\xa8\xae\xa9\xc4\xa1\xba\xa2\xba\xa2\xba\xbe\xc0\xea\x5a\xad\x96\x4a\x8d\x7a\xa3\x5a\xaa\xd4\xba\xd5\x7a\xc3\x6e\x1c\x18\x5d\xbb\x4c\xcb\xb4\x58\xb2\x8b\xbc\xf0\x3e\xab\xba\x96\x51\x5d\x53\x89\x43\x75\x45\x75\x45\x75\x7d\x81\xd5\xd5\xae\x51\xd3\x3e\xac\x1e\x76\xeb\x86\x51\xb5\x6a\xb5\xc3\x72\xb5\x5e\x36\x0e\x69\xa9\x61\x94\xea\x66\xad\xf4\xac\xea\x5a\x41\x75\x4d\x25\x0e\xd5\x15\xd5\x15\xd5\x55\x96\xea\x9a\xe1\x7f\x9d\xc6\x62\xe5\x2f\x8e\x33\xad\x19\x74\xfb\x15\x1b\x50\x2f\x0a\xe3\x4f\x75\x9b\x9a\x99\x8d\xe2\x62\x16\x1f\x2d\x77\xc4\x7c\xcf\x1d\xf0\xce\xfe\xc4\x2b\x43\x76\x9f\x30\xfe\xc9\xd2\xe9\xe3\xc1\xb3\xac\x8c\x17\x7e\x76\x3a\xb5\x59\x29\xb0\x3c\xf3\x86\xfa\xe4\x86\x39\x0e\x11\xbd\xaf\x8c\x62\x5e\xbb\x94\xd2\x3b\x8c\x78\x86\x43\x87\xde\x5a\xd4\x3d\x60\xe3\xef\xba\xcc\xea\x63\xe1\x7c\x52\x88\xa7\xcf\x3d\x6a\xf6\x3d\x12\x9f\x83\xb9\xbd\x38\xeb\x44\xd4\x08\x66\x35\xc9\xee\xc9\xf9\xe5\xe9\x79\xfb\xfa\xe2\xf2\xec\xdb\xf1\xd5\xe9\xf5\xd9\xc5\xb7\xea\x9e\xbc\x7e\x60\x16\xd2\xc9\xf9\xd7\x93\xcf\x1f\x8e\xdb\xa7\xd7\xed\xd3\xcb\x6f\xa7\x97\x5c\x78\x6b\xf7\xad\xce\x2b\xf1\xfc\x04\xa9\xf5\x68\xa8\xd9\xcc\x0f\x96\x6d\xc2\xf4\x91\x26\x32\xa1\x05\x62\x2a\x78\x47\x3a\x51\xb1\x58\xae\x4d\x7e\x93\xe9\x38\x8d\x5c\x6e\x12\x44\x58\x93\xcc\x69\x1a\xaf\x0f\xfc\xa0\xdc\x8f\x3c\xad\x16\x44\x1b\x11\x7d\x64\xf8\xba\xc3\xba\x93\x99\xd8\x32\x42\xa3\xb9\xfc\x94\x38\x20\x0d\x5b\xc2\x92\x2c\x0d\x29\x32\x8d\xb1\x53\xf0\xa9\xed\xd3\xa0\xaf\xc5\x4f\xdb\xd4\x08\x23\x9f\xea\x0f\xed\x8f\x1e\xd4\x08\xf9\x67\x1f\x5b\xbb\x77\x2b\x3a\xf2\x9e\x68\xdd\x88\x07\x11\xb6\x26\x6d\x67\x41\xb7\xfa\x61\x38\x6c\xea\xba\x78\xd7\x52\x1e\xef\x9b\x87\xc5\x46\x49\x27\x9a\xef\x79\x61\x6b\x95\xf8\x11\xa3\xe3\x0b\xdf\xfb\x7e\xdb\x29\x64\x37\x56\xf2\xf0\x49\x0f\x5d\x18\xf0\x83\x67\x55\x6c\x4e\x3c\xd7\x76\x98\x19\x06\xd9\x9b\xa1\xff\xa8\x01\x9a\x69\x2e\x92\xc4\x36\x4a\x81\x0b\x24\xf0\x51\xe9\xb4\xc1\x47\xc1\x47\xc1\x47\xc1\x47\xc1\x47\xc1\x47\xad\x16\x0b\x1f\x95\xe6\x72\x68\xec\xa3\x14\xb8\x14\x0a\x1f\x95\x4e\x1b\x7c\x14\x7c\x14\x7c\x14\x7c\x14\x7c\x14\x7c\xd4\x6a\xb1\xf0\x51\x69\x16\x3e\xc4\x3e\x4a\x81\x45\x0f\xf0\x51\xe9\xb4\xc1\x47\xc1\x47\xc1\x47\xc1\x47\xc1\x47\xc1\x47\xad\x16\x0b\x1f\x95\xf4\xba\xde\x74\xf0\xa9\x73\x71\x6f\x41\xb1\x5c\xcb\x33\x27\x23\x3b\xc3\x85\x99\x97\xf4\x9f\x88\x8f\x84\x67\x0c\xeb\x5f\x22\x2f\xc3\xd4\x89\xb9\x41\xda\xc4\x65\x2c\x4e\x66\x0f\x3f\xd3\xe6\x0d\x65\x93\x06\xbc\x58\x89\x17\xd3\x7a\xa8\x8c\x6c\xc1\x18\xb2\x8b\xde\xce\x19\x33\xab\x2a\xea\x02\xf5\xb5\xca\x7e\x71\xbf\xd4\x74\xf8\x8c\x1d\x48\x3b\x82\x7e\x20\x2a\xd3\x89\x02\x5e\xdc\xb5\x9e\x97\x43\x4c\x39\xe1\xc7\xd2\x40\x9b\x83\x84\x69\x60\x31\x0d\x08\x20\x99\x3d\x23\xbc\x74\x73\xf9\xa9\xb5\x34\xb0\xa9\xe3\xc9\xdc\xab\x22\x50\x39\x8d\xf8\x73\x92\xb6\x39\xf4\xa5\xd1\x41\xa2\x21\xa7\x10\xaa\x8d\x7d\x63\x38\xe4\x47\x59\x6a\x4a\x7c\x3a\xf0\x46\x54\x33\x5c\x4b\xf3\x69\xd7\x70\x0c\xb1\x0b\x48\xd3\x1c\xcf\x34\x1c\x8d\x0d\x57\x32\xef\x11\x09\x22\xcb\x23\x0b\xc0\x30\x3d\x70\x16\xa9\xcd\x01\x1b\x16\xdc\xa9\x22\xec\x90\xf4\x5a\xf6\x22\x3b\xa8\x70\x41\x1b\xec\x20\x8d\x3c\xb0\x43\x22\x71\x60\x87\x44\xd2\xc0\x0e\x60\x87\x9f\x25\x1a\xec\x00\x76\x00\x3b\x80\x1d\x5e\x05\x3b\x24\x5d\xbf\xb1\xc8\x0e\x2a\x2c\xe2\x00\x3b\x48\x23\x0f\xec\x90\x48\x1c\xd8\x21\x91\x34\xb0\x03\xd8\xe1\x67\x89\x06\x3b\x80\x1d\xc0\x0e\x60\x87\x57\xc1\x0e\xa9\xd7\x2c\x05\xcc\xa2\x37\xcc\xbc\x51\x6b\xdd\xd2\x54\xb5\x5c\x46\x18\x0c\x91\x58\xdb\x6b\x62\x88\x0c\xd3\x66\x8b\xc2\x26\x6d\xde\xb2\x56\x97\xed\x9d\x2a\xd7\x2d\xdb\x94\xaa\x83\x73\x54\x29\x33\xca\x82\x17\x37\xf2\xa2\x96\xe1\xfc\xf7\xf3\xc1\x51\x7a\xf5\x60\xad\xc7\x3e\xca\x05\x28\x62\xc0\xd2\x56\x9e\x6c\x0a\x15\xad\x7c\x37\xf0\xe4\xbe\x29\x23\x8f\xc1\x9e\xc3\x7d\x68\x9f\x3f\xbd\x48\x04\x3d\xa9\x17\x5b\xcd\xa0\x47\xa9\x05\x57\x80\x1e\x19\xe4\x01\x7a\x12\x89\x03\xf4\x24\x54\x97\x27\xf4\x94\xe5\xec\xe0\x1c\x55\x02\x7a\x00\x3d\x80\x9e\x55\xea\x01\x3d\x80\x9e\x57\x09\x3d\x9b\xa6\x17\x89\xa0\x27\xf5\x2a\xb1\x19\xf4\x28\xb5\x52\x0c\xd0\x23\x83\x3c\x40\x4f\x22\x71\x80\x9e\x84\xea\xf2\x84\x9e\x8a\x9c\x1d\x9c\xa3\x4a\x40\x0f\xa0\x07\xd0\xb3\x4a\x3d\xa0\x07\xd0\xf3\x2a\xa1\x67\xd3\xf4\xa2\x3e\xf4\x04\xb7\xae\x79\xdd\x1b\x2b\x73\x43\xae\x79\xbd\x72\xb9\x12\x80\x4e\x62\x6d\x00\x9d\x44\xca\x5e\x17\xe8\xc0\x9a\xab\x66\xcd\x1f\x6a\xb5\x32\xa2\xb9\x1f\x97\x5c\xf2\x56\x13\x2e\xf4\x6b\x3d\xee\x1a\xc6\xc6\x2d\xf7\x92\x3e\xf5\x02\xec\xd6\xf9\xe5\x31\x6d\xd8\x6d\xb2\xbe\xc3\xb6\x13\xc3\xd8\x60\xa1\x16\xb9\x21\x73\x34\x7e\x06\x97\xb9\x3d\x45\x13\x22\x76\x2b\xf5\xbd\x01\xd5\x45\x02\x9a\x8f\x7f\x26\x85\xab\x49\x52\xc7\x3c\x9f\xae\xcd\x7a\xc4\xa7\x63\x9f\x85\x62\x1f\x0e\xf7\xd5\x21\x73\x0d\x21\x7e\xee\x8c\xfa\xfe\xe2\x1b\xf6\xff\x0e\x78\x70\xf2\x33\x6b\xfb\x8f\xff\x9f\x5c\xff\xef\xb7\xeb\x93\xf3\x2f\x5f\xce\xae\x5a\xbb\x6f\xa7\xdf\xf5\xde\x5b\xfe\xae\xfb\xf9\x71\xc6\x5f\x1a\x0c\x58\xf8\xee\x88\x90\x0d\xd8\xfb\x50\x0a\x7f\xb8\x9f\x56\x8c\xe8\x59\x87\xc4\x8c\x2b\xb4\xef\x2e\x6a\x27\x5a\x8f\x6c\xeb\x10\x75\x38\x37\xbb\x49\x24\xfb\xad\x5b\xf3\xfc\xf4\xb2\x37\x6e\x2d\x90\xad\x02\xab\x17\x41\xb6\xb2\xc8\x03\xd9\x26\x12\x07\xb2\x4d\xa8\x0e\x64\x0b\xb2\x05\xd9\x82\x6c\x41\xb6\x20\x5b\x90\x2d\xc8\x36\x79\x8f\x80\x6c\x95\x22\xdb\x84\xbb\xf3\x16\xc8\x56\x81\x25\xaa\x20\x5b\x59\xe4\x81\x6c\x13\x89\x03\xd9\x26\x54\x07\xb2\x05\xd9\x82\x6c\x41\xb6\x20\x5b\x90\x2d\xc8\x16\x64\x9b\xbc\x47\x40\xb6\x4a\x91\x6d\xca\xd5\xc8\x2a\xdd\x6a\xf3\xa9\x66\xb9\x78\x03\x84\x9b\x58\xdb\x6b\x22\xdc\x4c\x77\x11\xae\xd9\x53\x20\x55\xe7\xe6\xa6\x11\xc8\xab\x22\xf2\xc6\x46\x47\xf2\x1d\x78\xeb\xd8\x57\x05\xed\xb2\x10\x63\x4e\xae\x7e\x8b\x25\x9f\xef\xa1\x0c\xf6\x0e\xce\x8c\xf6\x03\x53\x4e\x0c\xd9\x0b\xdc\x30\x98\xc3\xe8\xce\x7e\xbb\xe0\xb3\x67\x12\x89\xec\x79\xca\x25\x95\x2a\xdd\x14\x12\xf6\x5c\x26\x79\xb0\xe7\x89\xc4\xe5\x67\xcf\x25\xbd\xaf\x61\x6e\x1a\x61\xcf\x61\xcf\x61\xcf\x9f\x6a\x87\x3d\x87\x3d\xdf\x9c\xcf\x97\x61\xcf\x15\xb9\x81\x61\xca\x75\x61\x2a\xdd\xbe\x10\xf6\x5c\x26\x79\xb0\xe7\x89\xc4\xe5\x67\xcf\x25\xbd\x03\x5f\x6e\x1a\x61\xcf\x61\xcf\x61\xcf\x9f\x6a\x87\x3d\x87\x3d\xdf\x9c\xcf\x97\x61\xcf\x93\xdc\x6a\x6f\x67\xf2\xf8\x7e\x67\xe7\xdf\x00\x00\x00\xff\xff\xcf\x53\xc7\xf0\x34\xd5\x00\x00")

func data_test_fleet_api_units_json_bytes() ([]byte, error) {
	return bindata_read(
		_data_test_fleet_api_units_json,
		"data-test/fleet_api_units.json",
	)
}

func data_test_fleet_api_units_json() (*asset, error) {
	bytes, err := data_test_fleet_api_units_json_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data-test/fleet_api_units.json", size: 54580, mode: os.FileMode(420), modTime: time.Unix(1426775948, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _data_test_fleet_api_units_json_ = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\x56\x2a\xcd\xcb\x2c\x29\x56\xb2\x8a\x8e\xad\xe5\x02\x04\x00\x00\xff\xff\x6c\xb7\x9b\x40\x0d\x00\x00\x00")

func data_test_fleet_api_units_json__bytes() ([]byte, error) {
	return bindata_read(
		_data_test_fleet_api_units_json_,
		"data-test/fleet_api_units.json~",
	)
}

func data_test_fleet_api_units_json_() (*asset, error) {
	bytes, err := data_test_fleet_api_units_json__bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data-test/fleet_api_units.json~", size: 13, mode: os.FileMode(420), modTime: time.Unix(1425751310, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"data/couchbase_node@.service.template": data_couchbase_node_service_template,
	"data/couchbase_sidekick@.service.template": data_couchbase_sidekick_service_template,
	"data/sync_gw_node@.service.template": data_sync_gw_node_service_template,
	"data/sync_gw_node@.service.template~": data_sync_gw_node_service_template_,
	"data/sync_gw_sidekick@.service.template": data_sync_gw_sidekick_service_template,
	"data/sync_gw_sidekick@.service.template~": data_sync_gw_sidekick_service_template_,
	"data-test/fleet_api_units.json": data_test_fleet_api_units_json,
	"data-test/fleet_api_units.json~": data_test_fleet_api_units_json_,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"data": &_bintree_t{nil, map[string]*_bintree_t{
		"couchbase_node@.service.template": &_bintree_t{data_couchbase_node_service_template, map[string]*_bintree_t{
		}},
		"couchbase_sidekick@.service.template": &_bintree_t{data_couchbase_sidekick_service_template, map[string]*_bintree_t{
		}},
		"sync_gw_node@.service.template": &_bintree_t{data_sync_gw_node_service_template, map[string]*_bintree_t{
		}},
		"sync_gw_node@.service.template~": &_bintree_t{data_sync_gw_node_service_template_, map[string]*_bintree_t{
		}},
		"sync_gw_sidekick@.service.template": &_bintree_t{data_sync_gw_sidekick_service_template, map[string]*_bintree_t{
		}},
		"sync_gw_sidekick@.service.template~": &_bintree_t{data_sync_gw_sidekick_service_template_, map[string]*_bintree_t{
		}},
	}},
	"data-test": &_bintree_t{nil, map[string]*_bintree_t{
		"fleet_api_units.json": &_bintree_t{data_test_fleet_api_units_json, map[string]*_bintree_t{
		}},
		"fleet_api_units.json~": &_bintree_t{data_test_fleet_api_units_json_, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

