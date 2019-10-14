// Copyright 2019 The vogo Authors. All rights reserved.
// author: wongoo
// since: 2018/12/27

// zkclient is a client of zookeeper
//
// ## Features:
// - auto reconnect
// - set/get/delete value
// - support string/json codec, and you can implement your own
// - real-time synchronize data from zookeeper to memory
//
//## Modules
//
//- `Codec`: value encode/decode
//- `Watcher`: loop watch control
//- `Handler`: include `ValueHandler` and `MapHandler`, set/get value, handle event, synchronize value, trigger listener
//- `Listener`:  include `ValueListener` and `ChildListener`,  listen value change
//
//## API
//
//- `Sync*`:  synchronize data into memory
//- `SyncWatch*`:  synchronize data into memory, and listen value change
//- `Decode*`: set value into zookeeper
//- `Encode*`: get value from zookeeper
package zkclient
