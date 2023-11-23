// Copyright (c) 2020-2022 Tigera, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ifstate

import (
	"encoding/binary"
	"fmt"
	"strings"

	"golang.org/x/sys/unix"

	"github.com/projectcalico/calico/felix/bpf/maps"
)

func init() {
	SetMapSize(MapParams.MaxEntries)
}

func SetMapSize(size int) {
	maps.SetSize(MapParams.VersionedName(), size)
}

const (
	KeySize    = 4
	ValueSize  = 4 + 16 + 3*4 + 2*4
	MaxEntries = 1000
)

const (
	FlgWEP   = uint32(0x1)
	FlgReady = uint32(0x2)
	FlgMax   = uint32(0x3)
)

var flagsToStr = map[uint32]string{
	FlgWEP:   "workload",
	FlgReady: "ready",
}

var MapParams = maps.MapParameters{
	Type:         "hash",
	KeySize:      KeySize,
	ValueSize:    ValueSize,
	MaxEntries:   MaxEntries,
	Name:         "cali_iface",
	Flags:        unix.BPF_F_NO_PREALLOC,
	Version:      3,
	UpdatedByBPF: false,
}

func Map() maps.Map {
	return maps.NewPinnedMap(MapParams)
}

type Key [4]byte

func NewKey(ifIndex uint32) Key {
	var k Key

	binary.LittleEndian.PutUint32(k[:], ifIndex)

	return k
}

func (k Key) AsBytes() []byte {
	return k[:]
}

func (k Key) IfIndex() uint32 {
	return binary.LittleEndian.Uint32(k[:])
}

func (k Key) String() string {
	return fmt.Sprintf("{ifIndex: %d}", k.IfIndex())
}

func KeyFromBytes(b []byte) Key {
	var k Key
	copy(k[:], b)
	return k
}

type Value [ValueSize]byte

func NewValue(
	flags uint32,
	name string,
	xdpPol,
	ingressPol,
	egressPol,
	tcIngressFilter,
	tcEgressFilter int,
) Value {
	var v Value

	binary.LittleEndian.PutUint32(v[:], flags)
	copy(v[4:4+15], []byte(name))
	binary.LittleEndian.PutUint32(v[4+16+0:4+16+4], uint32(xdpPol))
	binary.LittleEndian.PutUint32(v[4+16+4:4+16+8], uint32(ingressPol))
	binary.LittleEndian.PutUint32(v[4+16+8:4+16+12], uint32(egressPol))
	binary.LittleEndian.PutUint32(v[4+16+12:4+16+16], uint32(tcIngressFilter))
	binary.LittleEndian.PutUint32(v[4+16+16:4+16+20], uint32(tcEgressFilter))

	return v
}

func (v Value) AsBytes() []byte {
	return v[:]
}

func (v Value) Flags() uint32 {
	return binary.LittleEndian.Uint32(v[:])
}

func (v Value) IfName() string {
	return strings.TrimRight(string(v[4:4+16]), "\x00")
}

func (v Value) XDPPolicy() int {
	return int(int32(binary.LittleEndian.Uint32(v[4+16 : 4+16+4])))
}

func (v Value) IngressPolicy() int {
	return int(int32(binary.LittleEndian.Uint32(v[4+16+4 : 4+16+8])))
}

func (v Value) EgressPolicy() int {
	return int(int32(binary.LittleEndian.Uint32(v[4+16+8 : 4+16+12])))
}

func (v Value) TcIngressFilter() int {
	return int(int32(binary.LittleEndian.Uint32(v[4+16+12 : 4+16+16])))
}

func (v Value) TcEgressFilter() int {
	return int(int32(binary.LittleEndian.Uint32(v[4+16+16 : 4+16+20])))
}

func (v Value) String() string {
	fstr := ""
	f := v.Flags()

	for k := FlgWEP; k < FlgMax; k++ {
		v := flagsToStr[k]
		if f&k != 0 {
			fstr = fstr + v + ","
		}
	}

	if fstr == "" {
		fstr = "host,"
	}

	return fmt.Sprintf(
		"{flags: %s XDPPolicy: %d, IngressPolicy: %d, EgressPolicy: %d, IngressFilter: %d, EgressFilter: %d, name: %s}",
		fstr, v.XDPPolicy(), v.IngressPolicy(), v.EgressPolicy(), v.TcIngressFilter(), v.TcEgressFilter(), v.IfName())
}

func ValueFromBytes(b []byte) Value {
	var v Value
	copy(v[:], b)
	return v
}

type MapMem map[Key]Value

func MapMemIter(m MapMem) func(k, v []byte) {
	ks := len(Key{})
	vs := len(Value{})

	return func(k, v []byte) {
		var key Key
		copy(key[:ks], k[:ks])

		var val Value
		copy(val[:vs], v[:vs])

		m[key] = val
	}
}
