// Code generated by go-bindata. DO NOT EDIT.
// sources:
// iatemplates/mongodb_connections_memory_usage.yml (788B)
// iatemplates/mongodb_high_memory_usage.yml (722B)
// iatemplates/mongodb_restarted.yml (586B)
// iatemplates/mysql_down.yml (344B)
// iatemplates/mysql_restarted.yml (573B)
// iatemplates/mysql_too_many_connections.yml (740B)
// iatemplates/node_high_cpu_load.yml (599B)
// iatemplates/node_low_free_memory.yml (605B)
// iatemplates/node_swap_filled_up.yml (611B)
// iatemplates/postgresql_down.yml (361B)
// iatemplates/postgresql_restarted.yml (595B)
// iatemplates/postgresql_too_many_connections.yml (720B)

package data

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _iatemplatesMongodb_connections_memory_usageYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x52\xc1\x8e\x9b\x30\x10\xbd\xf3\x15\x4f\x51\x23\x25\x55\x49\xb3\xab\xee\xc5\x52\x2b\xb5\xea\x35\xb7\xde\xa2\x08\x19\x18\xc0\x12\xb6\xd1\xd8\xd0\x45\x29\xff\x5e\xd9\x0b\x88\xb4\x5a\x2e\xd8\x8f\x99\xc7\x9b\xf7\x26\x4d\xd3\xc4\x93\xee\x5a\xe9\xc9\x89\x04\x48\x61\xa4\x26\x81\x4e\xeb\x4c\x5b\x53\xdb\x32\xcf\x0a\x6b\x0c\x15\x5e\x59\xe3\x32\x4d\xda\xf2\x98\xf5\x4e\xd6\x94\x00\xc0\x40\xec\x94\x35\x02\x4f\xf1\xea\x7a\xad\x25\x8f\x02\x97\x58\x88\xde\x51\x89\x7c\xc4\x25\x70\xfd\xfc\x81\x0d\x57\xac\xa7\xd7\x8e\x05\xfe\xa4\xf1\x12\xdb\x43\xf5\xc1\xd8\x92\xb2\xa0\xe4\x88\xc3\x22\xc3\xb9\xad\x92\x7b\x38\x67\x7e\xec\xe8\xeb\xae\xe8\x99\xc9\xf8\xdd\x74\xc4\x47\x3c\x9d\x9f\xbf\xcc\xaf\x99\xf4\x33\xac\x79\xa4\x8c\xe7\x79\x94\x0b\xe9\x5f\xd6\xcb\x36\xcb\x47\x4f\xee\x38\xf7\x04\x82\xf3\x7c\xfe\x86\xeb\x15\x27\xdf\x30\xb9\xc6\xb6\x25\x6e\xb7\xf8\xa1\x93\x2c\x75\x34\x2d\x3c\x8b\x71\x6b\xd9\x8c\x6f\x2c\xf9\x8e\x8e\xb8\x20\xe3\x65\x4d\xa8\xd8\xea\xe0\x46\xa5\xea\x9e\xa9\x84\x96\xaf\x4a\xf7\x7a\xed\xea\x8d\xf2\x02\xbb\xfd\x6e\x45\xc2\xac\x02\x55\x6b\xa5\x5f\x31\x96\xa6\x26\x81\xeb\xf9\x53\xd0\x7b\x5b\xf1\x41\xb6\x3d\x09\x3c\xbf\x44\xa4\xb2\x2c\xf0\xf2\xc6\xed\x68\x20\x56\x7e\x14\xf8\x2d\xd9\x28\x53\x47\x54\x1a\x63\xbd\x8c\xc6\x8a\xe4\x1f\xd9\x4b\x74\x8d\xaa\x1b\xe8\x25\xd6\x30\x43\x3e\x6e\xf3\xc4\xe1\x7e\xc7\x87\x56\xe6\xd4\xba\x93\x23\x1e\x54\xf1\x66\x38\xa6\x69\xb1\xb5\x24\x57\xb0\xea\x7c\xdc\x98\x35\x75\x20\x74\x46\xcd\x98\xa6\x3d\x6c\xb5\xfc\xe7\xa0\x2d\x13\x7c\x23\xcd\x7f\x19\xec\x8f\x50\x2e\xee\xd7\xca\x92\x8f\x78\x5f\xc2\x83\x54\x6b\xb6\x95\xeb\x6a\x60\x9a\x4e\xc9\xdf\x00\x00\x00\xff\xff\x47\x74\x6f\x37\x14\x03\x00\x00")

func iatemplatesMongodb_connections_memory_usageYmlBytes() ([]byte, error) {
	return bindataRead(
		_iatemplatesMongodb_connections_memory_usageYml,
		"iatemplates/mongodb_connections_memory_usage.yml",
	)
}

func iatemplatesMongodb_connections_memory_usageYml() (*asset, error) {
	bytes, err := iatemplatesMongodb_connections_memory_usageYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "iatemplates/mongodb_connections_memory_usage.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc5, 0x23, 0x17, 0x9d, 0x17, 0x6d, 0xc0, 0xdf, 0xf9, 0x9d, 0xa4, 0x3f, 0x27, 0xa4, 0x7b, 0x1e, 0x4e, 0x17, 0x32, 0x7b, 0xfd, 0x41, 0x77, 0x48, 0xa7, 0x77, 0xd5, 0x2c, 0x1f, 0x87, 0xd0, 0xb3}}
	return a, nil
}

var _iatemplatesMongodb_high_memory_usageYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x52\xc1\xaa\xdb\x30\x10\xbc\xe7\x2b\x86\x47\x03\x49\xa9\xd3\xbc\xd2\x42\xd1\xa1\xd0\xd2\x6b\x6e\xbd\x85\x60\xe4\x78\x6d\x0b\xbc\x92\x59\xc9\xe9\x33\xaf\xfe\xf7\x22\xc5\x56\x53\x4a\x7d\xb1\x34\xda\x19\xed\xce\xa8\x28\x8a\x4d\x20\x1e\x7a\x1d\xc8\xab\x0d\x50\xc0\x6a\x26\x85\x81\xb9\x64\x67\x5b\x57\x57\x65\x67\xda\xae\x64\x62\x27\x53\x39\x7a\xdd\xd2\x06\x00\x6e\x24\xde\x38\xab\xf0\x9c\xb6\x7e\x64\xd6\x32\x29\x9c\x52\x21\x46\x4f\x35\xaa\x09\xa7\x28\xf2\xfd\x5b\xaa\xa1\x97\x41\x14\x7e\x15\x69\x93\x28\xb1\x62\x67\x5d\x4d\x65\xbc\x76\x8f\xdd\x7a\xa7\xf7\xf1\xc6\x52\xc8\x9b\x9a\x6c\xc0\x5b\x3c\x1f\x3f\x7c\x5c\x7e\xfb\x45\xe0\x3d\x9c\xfd\x9b\x9e\xd6\x4b\xab\x27\xe2\x1f\x2e\xe8\xbe\xac\xa6\x40\x7e\xe5\x44\x85\xe3\xb2\xfe\x82\xf3\x19\x87\xd0\x09\xf9\xce\xf5\x35\x2e\x97\x74\x30\x68\xd1\x9c\xdc\x88\xdf\xea\x48\x2e\x5b\xf0\x87\x91\xbf\x62\x20\xb9\x92\x0d\xba\x25\x34\xe2\x18\x57\x67\x1b\xd3\x8e\x42\x35\x58\xbf\x18\x1e\x39\xb3\x46\x6b\x82\xc2\xd3\xf6\x29\x23\x61\x1a\x48\xa1\xe9\x9d\x0e\x19\x13\x6d\x5b\x52\x38\x1f\xdf\xc5\x7e\x2f\x19\xbf\xe9\x7e\x24\x85\xcf\xf7\x11\x1a\x27\x0a\x9f\xee\xda\x9e\x6e\x24\x26\x4c\x0a\x3f\xb5\x58\x63\xdb\x84\x6a\x6b\x5d\xd0\xc1\x38\x9b\x07\xfa\x93\xd4\x3d\x1a\xc4\x7c\xc1\x6b\x6c\x71\x86\xdd\xeb\x2b\xde\xf4\xba\xa2\xde\x1f\x3c\xc9\xcd\x5c\xef\x0e\x63\x9e\x57\x1f\x6b\xf2\x57\x31\x43\x48\x4f\x20\x47\x0a\x44\x66\x6a\x12\xf3\xbc\x85\x6b\x56\xe1\x1d\x3b\x21\x84\x4e\xdb\x7f\x4c\xdf\xee\x61\x7c\x7a\x30\x59\xa5\x9a\xf0\xff\x16\x62\xea\x0f\xa7\x39\x7f\xcc\xf3\x61\xf3\x3b\x00\x00\xff\xff\x11\xd2\x6c\x1e\xd2\x02\x00\x00")

func iatemplatesMongodb_high_memory_usageYmlBytes() ([]byte, error) {
	return bindataRead(
		_iatemplatesMongodb_high_memory_usageYml,
		"iatemplates/mongodb_high_memory_usage.yml",
	)
}

func iatemplatesMongodb_high_memory_usageYml() (*asset, error) {
	bytes, err := iatemplatesMongodb_high_memory_usageYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "iatemplates/mongodb_high_memory_usage.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x12, 0xf7, 0xd5, 0x2, 0xe4, 0xb4, 0x9b, 0xa9, 0x8f, 0xdd, 0xea, 0xe6, 0x77, 0xf8, 0x64, 0x36, 0xad, 0xb6, 0x52, 0xf2, 0xee, 0xd6, 0xac, 0x59, 0xfe, 0x1f, 0x8a, 0x50, 0x84, 0x1, 0x90, 0x17}}
	return a, nil
}

var _iatemplatesMongodb_restartedYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\x31\x6f\xe3\x30\x0c\x85\x77\xff\x8a\x37\xdc\x70\x07\x9c\x03\x07\x5d\x0a\xa1\x53\xd1\xb5\x53\xc7\x20\x30\x18\x9b\x71\x04\x58\x92\x21\xd2\x4e\x83\x34\xff\xbd\xb0\xe2\xa8\xe9\x10\x6d\xe4\xfb\xf4\xf0\x48\x96\x65\x59\x28\xbb\xa1\x27\x65\x31\x05\x50\xc2\x93\x63\x83\xc1\xb9\xda\x05\xdf\x85\x76\x57\x47\x16\xa5\xa8\xdc\x16\x00\x30\x71\x14\x1b\xbc\xc1\x3a\x95\x32\x3a\x47\xf1\x64\xf0\x3e\xd3\x6f\xaf\xf8\x4d\xf3\xe7\x10\x0d\xbe\xca\x54\x00\x37\x4b\xeb\x45\xc9\x37\x5c\x8f\x83\x5a\xc7\xb5\x70\x13\x7c\x2b\x0b\xf5\x82\xcd\x06\x2b\x3d\x44\x96\x43\xe8\x5b\x6c\xb7\x49\x18\x28\x92\x4b\x29\xe7\x77\x4b\x9a\xb1\xa5\x7f\x17\xe9\xe3\xea\x8a\x7d\x0c\x0e\x7a\x60\xf4\x24\x7a\x0b\x98\xf1\xd1\x5b\x35\x90\x5c\xeb\x69\x60\x83\x7d\x1f\xe8\x87\x89\xe4\x3b\x36\xd8\x54\xff\xb1\x7e\xae\xaa\x6d\x16\x26\xea\x47\x36\x78\xaa\xaa\xd4\xda\x87\x68\xb0\xae\xae\x66\xc2\x13\x47\xab\x27\x83\x23\x45\x6f\x7d\x97\xba\xe4\x7d\x50\x52\x1b\x7c\x1e\xe5\xf1\x0e\xf1\xf7\x7c\xc6\x9f\x9e\x76\xdc\xcb\x4a\x38\x4e\xb6\xe1\x7a\x9e\x1b\x97\xcb\xbf\xe5\x77\xcb\xd2\x44\x3b\x68\x3a\x4a\xde\x34\xb2\xd9\x63\x07\x04\x7f\xaf\xfa\xd0\x66\x29\xbb\x1c\x49\xee\xe2\xcc\x74\x1a\x79\xfe\xbd\x1c\x0d\xd4\x85\x55\xf1\x1d\x00\x00\xff\xff\x53\x44\xe2\x7e\x4a\x02\x00\x00")

func iatemplatesMongodb_restartedYmlBytes() ([]byte, error) {
	return bindataRead(
		_iatemplatesMongodb_restartedYml,
		"iatemplates/mongodb_restarted.yml",
	)
}

func iatemplatesMongodb_restartedYml() (*asset, error) {
	bytes, err := iatemplatesMongodb_restartedYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "iatemplates/mongodb_restarted.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x35, 0x6e, 0xdf, 0xbe, 0x87, 0x93, 0x9f, 0xe6, 0x59, 0x9c, 0x31, 0x1e, 0x48, 0xb1, 0xa, 0xe5, 0xf, 0x9a, 0xa9, 0x2e, 0xc1, 0x20, 0xd2, 0xd0, 0xce, 0xe, 0x75, 0x6c, 0xd7, 0xf0, 0x68, 0xb5}}
	return a, nil
}

var _iatemplatesMysql_downYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\x4d\x6a\xc3\x30\x14\x84\xf7\x3e\xc5\x2c\xba\xb0\xa1\x0a\xed\xa2\x1b\x41\x6e\xd0\x2e\x4a\x0f\x60\x14\xf9\x15\x04\xfa\xeb\x7b\x8a\x5b\x91\xe6\xee\xc1\xb2\x09\xd9\x78\x27\xcd\x68\x3e\xcd\x28\xa5\xba\x42\x21\x7b\x53\x48\x74\x07\x28\x44\x13\x48\x23\x87\x30\x86\x2a\x3f\x7e\x9c\xd2\x6f\xec\x00\x60\x26\x16\x97\xa2\xc6\x6b\xbb\xca\x39\x04\xc3\x55\xe3\xa3\x7e\x7d\xbe\xe3\xfe\x8c\xfe\x32\xeb\xc5\xc5\xa9\xa2\x17\xe2\xd9\x59\x1a\x17\xea\x33\x62\x9a\xd6\xe3\x80\x7e\xa5\x9f\xf3\x80\xe3\x11\x2f\x2d\xfa\x9d\x58\xe3\x4d\x56\x3c\xcd\xc4\xae\x54\x0d\xcb\xae\x38\x6b\x7c\x93\x4d\x8c\xa9\x98\xe2\x52\x6c\x75\x77\x8a\xa0\xbf\x5c\xf0\xe4\xcd\x89\xbc\x1c\x1e\x2b\xe0\x7a\x1d\xb6\xd8\x44\x62\xd9\xe5\xd2\x26\xfd\xab\x4d\xc5\x46\xd9\xcf\x23\xc5\x47\xf7\x3e\x69\xb1\x9c\xb4\xef\x0f\xdd\x2d\x00\x00\xff\xff\x11\xfb\x5c\x89\x58\x01\x00\x00")

func iatemplatesMysql_downYmlBytes() ([]byte, error) {
	return bindataRead(
		_iatemplatesMysql_downYml,
		"iatemplates/mysql_down.yml",
	)
}

func iatemplatesMysql_downYml() (*asset, error) {
	bytes, err := iatemplatesMysql_downYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "iatemplates/mysql_down.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x37, 0xb6, 0xea, 0x41, 0xdb, 0xb3, 0xf1, 0x78, 0x1f, 0x78, 0x5f, 0xb6, 0xfb, 0x3e, 0x1a, 0x95, 0x6, 0xc3, 0xa0, 0xd8, 0xa9, 0x75, 0x67, 0xf0, 0x44, 0xcf, 0xff, 0x87, 0xd5, 0xee, 0x84, 0x9}}
	return a, nil
}

var _iatemplatesMysql_restartedYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\xcd\x0e\xd3\x40\x0c\x84\xef\x79\x8a\x39\x70\x00\x89\x54\xa9\xb8\xa0\x15\x8f\x00\x07\xd4\x63\x55\x45\x6e\xe3\xa6\x2b\xed\x4f\x58\x3b\x29\x51\xe9\xbb\xa3\x6c\xd3\xa5\x1c\xba\x37\x7b\x3e\x8f\xc6\xde\xba\xae\x2b\x65\x3f\x38\x52\x16\x53\x01\x35\x02\x79\x36\x18\xbc\x6f\xfd\x2c\xbf\x5c\x9b\x58\x94\x92\x72\x57\x01\xc0\xc4\x49\x6c\x0c\x06\xdb\x5c\xca\xe8\x3d\xa5\xd9\xe0\xc7\xbc\xfb\xf9\x1d\xff\xb3\xfc\x7b\x48\x06\x7f\xea\x5c\x00\x0f\xbb\xde\xc5\x23\xb9\x56\x94\x74\x94\x76\x1c\xd4\x7a\x5e\x81\x6f\xd8\xef\xb1\xd1\x4b\x62\xb9\x44\xd7\xe1\x70\xc8\xc2\x40\x89\x7c\x0e\xb7\xbc\x67\xc0\x82\xad\xfd\x97\x2c\x3b\x3e\xc5\xd0\x09\xce\x29\x7a\xe8\x85\xe1\x48\xf4\x99\xad\xe0\x63\xb0\x6a\x20\xa5\xd6\x79\x60\x83\xb3\x8b\xf4\x8f\x49\x14\x7a\x36\xd8\x37\x9f\xb1\xfd\xda\x34\x87\x22\x4c\xe4\x46\x36\xf8\xd2\x34\xb9\x75\x8e\xc9\x60\xdb\x3c\xcc\x84\x27\x4e\x56\x67\x83\x2b\xa5\x60\x43\x9f\xbb\x14\x42\x54\x52\x1b\x43\x59\xe5\xdd\xf1\xf0\xf1\x76\xc3\x07\x47\x47\x76\xb2\x11\x4e\x93\x3d\x71\xbb\x6c\x8d\xfb\xfd\xd3\x3a\xdb\xb1\x9c\x92\x1d\x34\xff\x45\x39\x31\x56\xab\xf7\xf3\x88\xe1\x55\x0d\xb1\x2b\x52\xf1\xb8\x92\xbc\x84\x59\xe8\xbc\xee\x32\x2d\xeb\x69\xa9\x8f\x9b\xea\x6f\x00\x00\x00\xff\xff\xc9\x01\x74\x0c\x3d\x02\x00\x00")

func iatemplatesMysql_restartedYmlBytes() ([]byte, error) {
	return bindataRead(
		_iatemplatesMysql_restartedYml,
		"iatemplates/mysql_restarted.yml",
	)
}

func iatemplatesMysql_restartedYml() (*asset, error) {
	bytes, err := iatemplatesMysql_restartedYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "iatemplates/mysql_restarted.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5, 0x27, 0x78, 0x63, 0xd0, 0x52, 0x66, 0x65, 0xc, 0xd4, 0xbc, 0xe8, 0x1f, 0xc4, 0xb5, 0xed, 0x90, 0x66, 0x46, 0xe6, 0x90, 0x98, 0xa2, 0x2b, 0x9e, 0x45, 0x97, 0xb5, 0x20, 0xce, 0xa6, 0x5}}
	return a, nil
}

var _iatemplatesMysql_too_many_connectionsYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\x4f\x8f\xd3\x30\x10\xc5\xef\xfd\x14\x4f\x2b\x2a\xb5\x88\x96\xee\x61\x25\xe4\x03\x12\x77\x38\x20\x8e\x55\x65\x4d\x9b\x49\x6a\x64\x7b\xc2\xd8\x09\x1b\x2d\xfd\xee\xc8\xe9\x26\x74\x17\x6d\x4e\xf1\x9b\xf1\xfc\xf9\x3d\x6f\x36\x9b\x45\xe6\xd0\x7a\xca\x9c\xcc\x02\xd8\x20\x52\x60\x83\x36\x04\x1b\x86\xf4\xcb\xdb\x2c\x62\x03\xc5\xc1\x9e\x24\x46\x3e\x65\x27\x31\x2d\x00\xa0\x67\x4d\x4e\xa2\xc1\xfd\x78\x4c\x5d\x08\xa4\x83\xc1\xb7\xe1\xc7\xf7\xaf\xb8\xc9\x86\x8b\xe8\x12\x8f\x59\xfc\xd8\xaa\xc1\x9f\xcd\x78\x00\x02\x3d\x5a\xe9\x59\x6d\x76\x81\x57\xd7\x86\x8d\x97\x23\x79\x9b\x32\xe5\x2e\xd9\x7c\x56\xa6\x2a\x4d\xdd\xb9\xda\x3f\x84\xc3\x1a\x1f\xe1\x9a\x28\xea\x62\x83\xd5\x4f\x39\xae\xa7\x7a\xb7\x15\x7a\x52\x47\x47\xcf\xc9\x96\x36\xaf\xc7\x07\xde\xe3\x7e\xb7\x7b\xfe\xff\x8c\xfd\x1e\xdb\xd2\x2c\x9d\xc5\x57\x38\x1c\xc6\x40\x4b\x4a\x61\x04\x53\xbe\x09\xce\x9c\xf6\xac\xdf\x2c\xff\x05\x2d\xeb\x89\x63\xa6\x86\x51\xab\x84\x02\xa2\x76\x4d\xa7\x5c\x95\x6d\x5d\xe8\xc2\x7c\xab\x8b\x2e\x1b\xdc\x2d\xef\x66\x25\x0f\x2d\x1b\xd4\x5e\x28\xcf\x9a\x52\x6c\xd8\x60\xbf\xfb\x50\xe6\x3d\xcc\x7a\x4f\xbe\x63\x83\x4f\xd7\x15\x6a\x51\x83\x87\x6b\xed\xc4\x3d\xab\xcb\x83\xc1\x6f\xd2\xe8\x62\x33\xaa\x14\xa3\x64\x1a\x01\x4c\x0b\xbd\xf2\x2c\x8b\xa0\x58\xfd\xc2\xbc\xd5\xd3\x13\xde\x79\x3a\xb2\x4f\xdb\xc4\xda\xbb\x13\xdb\x82\x01\x97\xcb\x44\xbd\xe2\x74\x52\xd7\xe6\xf1\x35\xcc\xde\x02\xe5\xe6\x38\x25\x2e\x97\x25\xa4\x7e\x59\x37\x88\x32\xf2\x99\xe2\x7f\xe8\x97\x6b\x90\x72\x79\x33\xff\x08\x1f\x07\xbc\x3d\x08\x24\xde\x46\xa3\x54\x73\x68\xbb\xf8\x1b\x00\x00\xff\xff\x15\xf2\x1e\x02\xe4\x02\x00\x00")

func iatemplatesMysql_too_many_connectionsYmlBytes() ([]byte, error) {
	return bindataRead(
		_iatemplatesMysql_too_many_connectionsYml,
		"iatemplates/mysql_too_many_connections.yml",
	)
}

func iatemplatesMysql_too_many_connectionsYml() (*asset, error) {
	bytes, err := iatemplatesMysql_too_many_connectionsYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "iatemplates/mysql_too_many_connections.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5f, 0xb6, 0x7e, 0x51, 0x42, 0x86, 0x51, 0x99, 0x1b, 0x87, 0x2f, 0x6b, 0xa3, 0xc5, 0x72, 0x59, 0x7, 0x65, 0x74, 0xb1, 0x59, 0x1d, 0xbf, 0xfb, 0x3a, 0xce, 0xac, 0x9c, 0x4, 0xe8, 0x53, 0x77}}
	return a, nil
}

var _iatemplatesNode_high_cpu_loadYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x4f\x8b\xdb\x30\x10\xc5\xef\xf9\x14\x0f\xd3\x05\xbb\xd4\x21\x39\x2c\x14\x41\x0b\xa5\xf7\xd2\x4b\x4f\xc6\x98\x59\x6b\x6c\x0b\xf4\x0f\x49\x76\xd7\xa4\xf9\xee\xc5\x4e\xa2\x16\x4a\x57\x27\xcd\x4f\xf3\x47\xef\x4d\x5d\xd7\x87\xc4\xc6\x6b\x4a\x1c\xc5\x01\xa8\x61\xc9\xb0\x80\x37\xa6\xb3\x4e\x72\x37\xa9\x71\xea\x7a\x3f\x77\xda\x91\x3c\x00\xc0\xc2\x21\x2a\x67\x05\xce\x7b\x18\x67\x63\x28\xac\x02\xdf\x9c\x64\x6c\xe9\xf8\xfa\xfd\x07\x72\x3a\xbf\xfa\x20\xf0\xab\xde\x03\xa0\x3c\xa3\x06\x2d\x23\x5e\xd6\x72\x1f\xb0\xcd\xab\x50\x06\x4a\x7c\x03\xdb\xb0\xc8\xbd\xb3\x32\x76\xc9\x25\xd2\x17\xe3\x24\x7f\x2a\x94\xd4\x5c\x5c\x9b\x67\xd3\x56\x55\x75\xef\xf6\x1e\xe7\xd3\xe9\x7e\xff\x8c\xa6\xc1\x31\x4d\x81\xe3\xe4\xb4\x44\xdb\xee\x0f\x9e\x02\x99\x5d\xdb\x76\x1e\xfa\x72\xda\x9d\xff\xa5\xe3\x0b\x3c\x87\x9e\x6d\xa2\x91\x31\x04\x67\xd0\x3b\x3b\xa8\x71\x0e\x2c\x61\xe8\x55\x99\xd9\xe4\xaa\xd9\xaa\x24\x50\x3c\x15\x99\xa4\xd5\xb3\xc0\xa0\x1d\xa5\xcc\x02\xd9\x91\x05\x9a\xd3\x87\xed\xbf\x6d\xe6\x0b\xe9\x99\x05\x3e\xde\x24\x0c\x2e\x08\x3c\xdf\x7a\x47\x5e\x38\xa8\xb4\x0a\xfc\xa4\x60\x95\x1d\x77\x4a\xd6\xba\x44\x49\x39\x9b\x05\xbd\x61\x3f\xca\xcb\x05\xef\x34\xbd\xb0\x8e\xc7\xec\x35\xae\xd7\x87\x7b\x92\x63\x1f\x94\x4f\xfb\x36\xf3\x86\x80\xff\x94\xfd\x69\xac\x22\x8c\x0b\x8c\x34\x91\xfd\xc7\xf5\xa7\xe3\xe1\x77\x00\x00\x00\xff\xff\x2d\x86\x10\xdc\x57\x02\x00\x00")

func iatemplatesNode_high_cpu_loadYmlBytes() ([]byte, error) {
	return bindataRead(
		_iatemplatesNode_high_cpu_loadYml,
		"iatemplates/node_high_cpu_load.yml",
	)
}

func iatemplatesNode_high_cpu_loadYml() (*asset, error) {
	bytes, err := iatemplatesNode_high_cpu_loadYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "iatemplates/node_high_cpu_load.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x64, 0x8f, 0x37, 0xdd, 0x39, 0xd6, 0x4a, 0xbc, 0x1f, 0xe1, 0x72, 0x63, 0xe9, 0x32, 0xa, 0xec, 0x30, 0x1c, 0x2f, 0x18, 0x70, 0xb7, 0xaf, 0xef, 0x22, 0x1c, 0x8d, 0xf9, 0xb1, 0x24, 0x5, 0xc6}}
	return a, nil
}

var _iatemplatesNode_low_free_memoryYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x51\x4d\xab\xdb\x30\x10\xbc\xfb\x57\x0c\x8f\x3e\x78\x29\x75\x9a\x14\x7a\x11\xbd\xe4\x07\xb4\xa7\xde\x42\x30\x9b\x78\xe5\x08\xf4\x85\xb4\x76\x6a\xd2\xfc\xf7\x12\xd9\x31\xa1\xa5\x4f\x27\xed\xec\xce\xb0\x33\x5b\xd7\x75\x25\xec\xa2\x25\xe1\xac\x2a\xa0\x86\x27\xc7\x0a\xd1\xb9\xc6\x87\x96\x1b\x1b\x2e\x8d\x4e\xcc\x8d\x63\x17\xd2\x58\x01\xc0\xc0\x29\x9b\xe0\x15\xb6\xa5\xcc\xbd\x73\x94\x46\x85\x1f\xa1\x65\x84\x5e\x10\x34\x9e\xc6\xf9\x57\x4c\x0a\xbf\xeb\x52\x00\x45\x76\x6a\x37\xdf\xd9\xed\x06\x32\x96\x8e\x96\x9b\xe3\x28\x9c\xf1\xf9\xef\x81\x9f\x41\xc8\x4e\xcd\x59\xe1\x23\xb6\x9b\xcd\xfc\xff\x86\xfd\x1e\x6b\x39\x27\xce\xe7\x60\x5b\x1c\x0e\xa5\x11\x29\x91\x2b\x8e\xee\xef\xe1\x6a\x19\x9b\xf1\xa7\xdd\x77\x88\x9c\x4e\xec\x85\x3a\x86\x4e\xc1\xe1\x14\xbc\x36\x5d\x9f\xb8\x85\x33\xde\xb8\xde\x2d\xac\xde\x1b\x51\x78\x79\x7d\x59\x10\x19\x23\x2b\x68\x1b\x48\x16\x2c\x91\xef\x58\x61\xbf\xf9\x74\xdf\xf7\xb0\xe0\x03\xd9\x9e\x15\xbe\x4c\x16\x74\x48\x0a\x5f\x27\xed\xcc\x03\x27\x23\xa3\xc2\x85\x92\x37\xbe\x2b\x28\x79\x1f\x84\xc4\x04\xbf\x18\x7a\x27\x72\xbc\x5d\xaf\xf8\x60\xe9\xc8\x36\xaf\x4b\x94\x77\xef\xb8\xdd\x56\x33\xb7\xe5\x7c\x4a\x26\x4a\xb9\xe0\x72\x15\xe0\x3f\xb4\x87\xac\xc9\xd0\xc6\x5a\xe3\x3b\xf4\x11\x6f\x96\x73\x86\x9c\xc9\xff\x93\xff\x2b\x2c\x6b\x59\xad\xab\x3f\x01\x00\x00\xff\xff\x56\x49\x79\xff\x5d\x02\x00\x00")

func iatemplatesNode_low_free_memoryYmlBytes() ([]byte, error) {
	return bindataRead(
		_iatemplatesNode_low_free_memoryYml,
		"iatemplates/node_low_free_memory.yml",
	)
}

func iatemplatesNode_low_free_memoryYml() (*asset, error) {
	bytes, err := iatemplatesNode_low_free_memoryYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "iatemplates/node_low_free_memory.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb6, 0x52, 0x79, 0xc3, 0x6c, 0x9a, 0x12, 0x66, 0x3c, 0xab, 0xee, 0x7, 0x7e, 0x2e, 0xf6, 0xee, 0x2d, 0xda, 0x97, 0x86, 0x17, 0x73, 0x92, 0x96, 0x5c, 0x9e, 0x3c, 0xe4, 0xc9, 0xf, 0x9a, 0x6f}}
	return a, nil
}

var _iatemplatesNode_swap_filled_upYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x51\xcd\xca\xdb\x30\x10\xbc\xfb\x29\x86\x8f\x06\x92\x52\xa7\xc9\xa1\x50\x74\x28\xf4\xd2\x63\x2f\xed\x2d\x04\xb3\x89\xd7\xb6\x40\x7f\xac\xe4\x24\x26\xcd\xbb\x17\x2b\x8e\x49\x5a\x3e\x9d\xb4\xb3\x9a\x5d\xcd\x4c\x59\x96\x45\x62\x1b\x0c\x25\x8e\xaa\x00\x4a\x38\xb2\xac\x10\xac\xad\x9c\xaf\xb9\x8a\x67\x0a\x55\xa3\x8d\xe1\xba\xea\x43\x01\x00\x27\x96\xa8\xbd\x53\xd8\xe6\x32\xf6\xd6\x92\x0c\x0a\x3f\x7d\xcd\xe8\x74\xdb\x61\x24\x61\x24\x69\xd7\x62\x62\xf1\x25\x88\xc2\x9f\x32\x17\xc0\x72\x8b\x12\xcb\xbc\xc2\xb2\xf5\x32\x54\xbf\xce\x14\x7e\x08\x73\x75\x18\x12\x47\x7c\xc6\xbf\xcd\xdf\x3e\x91\xb9\x77\x57\xab\x69\xcc\x47\x6c\x37\x9b\xe9\xfe\x0d\xbb\x1d\xd6\xa9\x13\x8e\x9d\x37\x35\xf6\xfb\xdc\x08\x24\x64\xb3\xb8\xf1\x3c\x04\xce\xcf\x26\xfc\x49\xc7\x77\x04\x96\x23\xbb\x44\x2d\xa3\x11\x6f\x71\xf4\xae\xd1\x6d\x2f\x5c\xc3\xd2\x45\xdb\xde\xce\xac\xde\xe9\xa4\xf0\xb6\x78\x9b\x91\x34\x04\x56\x68\x8c\xa7\x34\x63\x42\xae\x65\x85\xdd\xe6\xd3\xf8\xdf\xfd\x8c\x9f\xc8\xf4\xac\xf0\xf5\x2e\xa1\xf1\xa2\xf0\xe5\x3e\x3b\xf2\x89\x45\xa7\x41\xe1\x4c\xe2\xb4\x6b\x33\x4a\xce\xf9\x44\x49\x7b\x37\x0b\x7a\xb5\x3f\x3b\xaf\xe3\x93\xf9\x58\x5e\xaf\xf8\x60\xe8\xc0\x26\xae\xb3\xa5\xa3\x01\xb8\xdd\x1e\x16\xd6\x1c\x8f\xa2\x43\xca\x91\xce\xf9\x00\xef\xd0\x5e\x56\x70\x3d\x6e\xb0\x5e\x18\xa9\x23\xf7\x5f\x00\x8b\x75\xf1\x37\x00\x00\xff\xff\xb6\x7a\x5a\xdf\x63\x02\x00\x00")

func iatemplatesNode_swap_filled_upYmlBytes() ([]byte, error) {
	return bindataRead(
		_iatemplatesNode_swap_filled_upYml,
		"iatemplates/node_swap_filled_up.yml",
	)
}

func iatemplatesNode_swap_filled_upYml() (*asset, error) {
	bytes, err := iatemplatesNode_swap_filled_upYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "iatemplates/node_swap_filled_up.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd4, 0x19, 0x1f, 0xac, 0x4f, 0x61, 0x99, 0xd3, 0xc4, 0xea, 0xf0, 0x11, 0x5a, 0xdf, 0x50, 0x49, 0x6b, 0xe8, 0x2b, 0xe4, 0x5f, 0xda, 0xa5, 0x6e, 0xbb, 0xe5, 0x84, 0x6b, 0xaf, 0xbd, 0xaa, 0xd3}}
	return a, nil
}

var _iatemplatesPostgresql_downYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xbd\x4e\xc3\x40\x10\x84\x7b\x3f\xc5\x14\x14\xb6\xc4\x45\x50\xd0\x9c\x94\x37\xa0\x00\xf1\x00\xd6\xc5\x5e\xa2\x93\xee\x8f\xdd\x8d\xc1\x0a\x79\x77\xe4\x73\x14\xb9\x21\xdd\xdd\xee\xcc\xec\x37\xc6\x98\x46\x29\x96\xe0\x94\xc4\x36\x80\x41\x72\x91\x2c\x4a\x8c\x7d\xc9\xa2\x47\x26\xf9\x0a\xfd\x98\xbf\x53\x03\x00\x13\xb1\xf8\x9c\x2c\x9e\xeb\x57\x4e\x31\x3a\x9e\x2d\xde\x56\xed\xc7\xfb\x2b\x6e\x5a\xfa\x29\x6c\x17\x09\x0e\x33\x5a\x21\x9e\xfc\x40\xfd\x92\xff\x88\x94\xc7\xf5\xd9\xa1\x2d\xc7\xfe\x54\x3a\xec\xf7\x78\xaa\xbe\xcf\xcc\x16\x2f\xb2\x1e\xa0\x89\xd8\xeb\x6c\x31\xb0\x57\x3f\xb8\x50\xc7\x2e\xa5\xac\x4e\x7d\x4e\x95\xfa\x1e\x0a\xda\xf3\x19\x0f\xc1\x1d\x28\xc8\x6e\x0b\x81\xcb\xa5\xbb\x7a\x47\x92\x81\x7d\xd1\xda\xec\xd7\x5c\xa7\xd8\x46\xfd\x1f\x82\x9c\xb6\xdb\x5b\xb3\x65\xe5\xa5\x32\xec\x9a\xbf\x00\x00\x00\xff\xff\x00\x94\x6f\x28\x69\x01\x00\x00")

func iatemplatesPostgresql_downYmlBytes() ([]byte, error) {
	return bindataRead(
		_iatemplatesPostgresql_downYml,
		"iatemplates/postgresql_down.yml",
	)
}

func iatemplatesPostgresql_downYml() (*asset, error) {
	bytes, err := iatemplatesPostgresql_downYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "iatemplates/postgresql_down.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcb, 0xba, 0x70, 0x4c, 0xaa, 0x5e, 0xa6, 0xcb, 0x70, 0xb0, 0x2d, 0x4a, 0x6a, 0x6b, 0xc, 0xc4, 0xd6, 0x65, 0x38, 0x97, 0x0, 0x98, 0x55, 0xf0, 0xcb, 0x87, 0x97, 0x37, 0x2a, 0x9e, 0x92, 0xcb}}
	return a, nil
}

var _iatemplatesPostgresql_restartedYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\xcd\x8e\x13\x41\x0c\x84\xef\x79\x8a\x3a\x70\x00\x89\x89\x26\xe2\x82\x5a\x3c\x02\x07\xd0\x1e\xa3\x68\x64\x32\xce\xa4\xa5\xfe\xc3\x76\xb2\x44\xcb\xbe\x3b\x9a\xce\xa4\x19\x0e\x6c\xdf\xec\x2a\x7f\x2a\xb7\xbb\xae\xdb\x18\xc7\x12\xc8\x58\xdd\x06\xe8\x90\x28\xb2\x43\x89\x71\x28\x59\x6d\x12\xd6\x9f\x61\x10\x56\x23\x31\x1e\x37\x00\x70\x65\x51\x9f\x93\xc3\xae\x96\x7a\x89\x91\xe4\xe6\xf0\xed\x3e\xf0\xf4\xfd\x2b\xfe\x1d\xe0\x5f\x45\x1c\x7e\x77\xb5\x00\xca\x54\xd9\x91\xd4\x58\x86\x4b\x31\x1f\x79\x50\x3e\xe6\x34\xea\x62\xf9\x82\xfd\x1e\x5b\x3b\x0b\xeb\x39\x87\x11\x87\x43\x15\x0a\x09\xc5\x1a\x74\x7e\x8f\xb0\xcd\xb6\xf4\x57\x91\x9e\xee\x54\x9c\x24\x47\xd8\x99\x11\x48\xed\x91\xae\xd9\x2f\xc9\x9b\x83\xb6\xda\x6e\x85\x1d\x4e\x21\xd3\x5f\x8f\x50\x9a\xd8\x61\xdf\x7f\xc4\xee\x73\xdf\x1f\x9a\x70\xa5\x70\x61\x87\x4f\x7d\x5f\x5b\xa7\x2c\x0e\xbb\xfe\x0e\x53\xbe\xb2\x78\xbb\x39\x3c\x93\x24\x9f\xa6\xda\xa5\x94\xb2\x91\xf9\x9c\xda\x2a\x6f\xfe\x21\xde\xbf\xbc\xe0\x5d\xa0\x1f\x1c\x74\xab\x2c\x57\x7f\xe4\x61\x5e\x1d\xaf\xaf\x1f\x16\xc0\xc8\x7a\x14\x5f\xac\xde\xa5\xfd\x34\xd6\xbc\xff\x43\x90\xd3\x5a\x4d\x79\x6c\x52\x03\x3d\x93\xae\x12\xcd\xee\xba\xf8\x3c\xbd\x9c\x0e\x34\xe5\xed\xe6\x4f\x00\x00\x00\xff\xff\x4a\x83\x13\xda\x53\x02\x00\x00")

func iatemplatesPostgresql_restartedYmlBytes() ([]byte, error) {
	return bindataRead(
		_iatemplatesPostgresql_restartedYml,
		"iatemplates/postgresql_restarted.yml",
	)
}

func iatemplatesPostgresql_restartedYml() (*asset, error) {
	bytes, err := iatemplatesPostgresql_restartedYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "iatemplates/postgresql_restarted.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x69, 0x14, 0xd1, 0xe6, 0x42, 0xee, 0x9a, 0xc3, 0xb3, 0x2e, 0x48, 0x4, 0x9, 0xe1, 0xb3, 0x37, 0xaa, 0xe8, 0x87, 0x92, 0xc2, 0xfd, 0xf8, 0xf2, 0x95, 0x44, 0x2, 0x32, 0x35, 0x50, 0xb5, 0x94}}
	return a, nil
}

var _iatemplatesPostgresql_too_many_connectionsYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\x41\x8f\xd3\x30\x10\x85\xef\xfd\x15\x8f\x8a\x4a\xbb\x2b\x52\xba\x07\x24\xe4\x03\x12\x77\x0e\x20\x8e\x55\x15\xcd\x26\x93\xd4\x52\x3c\x0e\xf6\x24\x34\xea\x96\xdf\x8e\xec\x6e\x43\x16\x44\x6e\x7e\x9e\x37\xf9\x66\x9e\x8b\xa2\x58\x29\xbb\xbe\x23\xe5\x68\x56\x40\x01\x21\xc7\x06\xbd\x73\x65\xef\xa3\xb6\x81\xe3\x8f\xae\x54\xef\x4b\x47\x32\x95\x95\x17\xe1\x4a\xad\x97\xb8\x02\x80\x91\x43\xb4\x5e\x0c\x1e\xf3\x31\x0e\xce\x51\x98\x0c\xbe\x5e\xbd\xdf\xbf\x7d\xc1\xc2\x02\x2b\x18\x22\xe7\x52\x3e\xf5\xc1\xe0\xb9\xc8\x87\xec\xbc\xeb\xdb\x32\x2a\x69\x49\x95\xda\xd1\x6a\xfa\xdb\x20\x7a\xae\x49\x13\xd4\x9b\x5f\xeb\x1b\xea\xf6\xe1\xf9\x06\xb7\xbe\xdc\xbf\x74\xf8\x84\xe4\x67\x55\x2b\x6d\x2c\x1d\x9d\x96\xb0\x78\xc0\x7e\x8f\xad\x1e\x03\xc7\xa3\xef\x6a\x1c\x0e\x78\x8f\xc7\xdd\x2e\x9b\x7b\x0a\xe4\xf2\xfc\xe9\xbb\xed\x60\x2e\x7e\xd1\x17\xe3\x7d\x46\xcf\xa1\x62\x51\x6a\x19\x4d\xf0\x2e\x4d\xd9\xd8\x76\x08\x5c\xc3\xd1\xc9\xba\xc1\xcd\xae\x41\xac\x1a\xac\x37\xeb\x59\xd1\xa9\x67\x83\xa6\xf3\xa4\xb3\x16\x48\x5a\x36\xd8\xef\xde\x25\xac\xc3\xac\x8f\xd4\x0d\x6c\xf0\xf1\x4a\xda\xf8\x60\xf0\xe1\xda\x3b\xf2\xc8\xc1\xea\x64\xf0\x93\x82\x58\x69\xb3\x4a\x22\x5e\x29\x0f\x6d\x56\x7f\x61\x2f\x52\x51\xef\x91\x12\x7d\x15\xcf\xdd\xf9\x8c\xb7\x1d\x3d\x71\x17\xb7\x91\xc3\x68\x2b\x2e\xd3\x2e\x70\x99\xb7\x5c\x73\xac\x82\xed\x35\x87\x3e\xa7\x07\x24\x67\x46\xc5\xe5\xb2\x81\x6f\x5e\xf7\x75\x3e\x30\xf4\x48\xf2\x4f\x0a\x9b\x7b\x50\xe0\xf4\x2a\xfe\xac\xf9\x69\xc2\xff\x41\xe0\x65\x79\x2b\xbe\x9e\xaf\xb6\xab\xdf\x01\x00\x00\xff\xff\xc8\x6f\x55\x57\xd0\x02\x00\x00")

func iatemplatesPostgresql_too_many_connectionsYmlBytes() ([]byte, error) {
	return bindataRead(
		_iatemplatesPostgresql_too_many_connectionsYml,
		"iatemplates/postgresql_too_many_connections.yml",
	)
}

func iatemplatesPostgresql_too_many_connectionsYml() (*asset, error) {
	bytes, err := iatemplatesPostgresql_too_many_connectionsYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "iatemplates/postgresql_too_many_connections.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbd, 0xee, 0xa5, 0xff, 0x2a, 0x4d, 0xe1, 0x3b, 0x24, 0x70, 0x4a, 0xd6, 0x52, 0xe8, 0xd2, 0x5d, 0x1f, 0x9c, 0x7f, 0x79, 0x4c, 0x6c, 0x5e, 0x20, 0x8e, 0x71, 0x59, 0xb7, 0xad, 0xc4, 0x1e, 0x0}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"iatemplates/mongodb_connections_memory_usage.yml": iatemplatesMongodb_connections_memory_usageYml,
	"iatemplates/mongodb_high_memory_usage.yml":        iatemplatesMongodb_high_memory_usageYml,
	"iatemplates/mongodb_restarted.yml":                iatemplatesMongodb_restartedYml,
	"iatemplates/mysql_down.yml":                       iatemplatesMysql_downYml,
	"iatemplates/mysql_restarted.yml":                  iatemplatesMysql_restartedYml,
	"iatemplates/mysql_too_many_connections.yml":       iatemplatesMysql_too_many_connectionsYml,
	"iatemplates/node_high_cpu_load.yml":               iatemplatesNode_high_cpu_loadYml,
	"iatemplates/node_low_free_memory.yml":             iatemplatesNode_low_free_memoryYml,
	"iatemplates/node_swap_filled_up.yml":              iatemplatesNode_swap_filled_upYml,
	"iatemplates/postgresql_down.yml":                  iatemplatesPostgresql_downYml,
	"iatemplates/postgresql_restarted.yml":             iatemplatesPostgresql_restartedYml,
	"iatemplates/postgresql_too_many_connections.yml":  iatemplatesPostgresql_too_many_connectionsYml,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"iatemplates": {nil, map[string]*bintree{
		"mongodb_connections_memory_usage.yml": {iatemplatesMongodb_connections_memory_usageYml, map[string]*bintree{}},
		"mongodb_high_memory_usage.yml":        {iatemplatesMongodb_high_memory_usageYml, map[string]*bintree{}},
		"mongodb_restarted.yml":                {iatemplatesMongodb_restartedYml, map[string]*bintree{}},
		"mysql_down.yml":                       {iatemplatesMysql_downYml, map[string]*bintree{}},
		"mysql_restarted.yml":                  {iatemplatesMysql_restartedYml, map[string]*bintree{}},
		"mysql_too_many_connections.yml":       {iatemplatesMysql_too_many_connectionsYml, map[string]*bintree{}},
		"node_high_cpu_load.yml":               {iatemplatesNode_high_cpu_loadYml, map[string]*bintree{}},
		"node_low_free_memory.yml":             {iatemplatesNode_low_free_memoryYml, map[string]*bintree{}},
		"node_swap_filled_up.yml":              {iatemplatesNode_swap_filled_upYml, map[string]*bintree{}},
		"postgresql_down.yml":                  {iatemplatesPostgresql_downYml, map[string]*bintree{}},
		"postgresql_restarted.yml":             {iatemplatesPostgresql_restartedYml, map[string]*bintree{}},
		"postgresql_too_many_connections.yml":  {iatemplatesPostgresql_too_many_connectionsYml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
