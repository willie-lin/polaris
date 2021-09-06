/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package boltdbStore

import (
	"github.com/boltdb/bolt"
	"time"
)

// BoltHandler encapsulate operations around boltdb
type BoltHandler interface {

	// SaveValue save go object into bolt
	SaveValue(typ string, key string, value interface{}) error

	// LoadValues load all objects by keys, return is map[key]value
	LoadValues(typ string, keys []string) (map[string]interface{}, error)

	// LoadValuesByFilter load all objects by filter, return is map[key]value
	LoadValuesByFilter(typ string, filter map[string][]string) (map[string]interface{}, error)

	// Close close boltdb
	Close() error
}

// BoltConfig config to initialize boltdb
type BoltConfig struct {
	// FileName boltdb store file
	FileName string
}

const (
	confPath    = "path"
	defaultPath = "./polaris.bolt"
)

// Parse parse yaml config
func (c *BoltConfig) Parse(opt map[string]interface{}) {
	if value, ok := opt[confPath]; ok {
		c.FileName = value.(string)
	} else {
		c.FileName = defaultPath
	}
}

const (
	defaultTimeoutForFileLock = 5 * time.Second
)

// NewBoltHandler create the boltdb handler
func NewBoltHandler(config *BoltConfig) (BoltHandler, error) {
	db, err := openBoltDB(config.FileName)
	if nil != err {
		return nil, err
	}
	return &boltHandler{db: db}, nil
}

type boltHandler struct {
	db *bolt.DB
}

func openBoltDB(path string) (*bolt.DB, error) {
	return bolt.Open(path, 0600, &bolt.Options{
		Timeout: defaultTimeoutForFileLock,
	})
}

// SaveValue save go object into bolt
func (b *boltHandler) SaveValue(typ string, key string, value interface{}) error {
	//TODO
	return nil
}

// LoadValues load all objects by keys, return is map[key]value
func (b *boltHandler) LoadValues(typ string, keys []string) (map[string]interface{}, error) {
	//TODO
	return nil, nil
}

// LoadValuesByFilter load all objects by filter, return is map[key]value
func (b *boltHandler) LoadValuesByFilter(typ string, filter map[string][]string) (map[string]interface{}, error) {
	//TODO
	return nil, nil
}

func (b *boltHandler) Close() error {
	if nil != b.db {
		return b.db.Close()
	}
	return nil
}