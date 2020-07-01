// automatically generated by stateify.

// +build go1.12
// +build !go1.16

package vfs

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *mountTable) StateTypeName() string {
	return "pkg/sentry/vfs.mountTable"
}

func (x *mountTable) StateFields() []string {
	return []string{
		"seed",
		"size",
	}
}

func (x *mountTable) beforeSave() {}

func (x *mountTable) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.seed)
	m.Save(1, &x.size)
}

func (x *mountTable) afterLoad() {}

func (x *mountTable) StateLoad(m state.Source) {
	m.Load(0, &x.seed)
	m.Load(1, &x.size)
}

func init() {
	state.Register((*mountTable)(nil))
}