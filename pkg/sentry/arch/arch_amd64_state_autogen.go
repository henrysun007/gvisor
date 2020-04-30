// automatically generated by stateify.

// +build amd64
// +build amd64
// +build amd64

package arch

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *context64) beforeSave() {}
func (x *context64) save(m state.Map) {
	x.beforeSave()
	m.Save("State", &x.State)
	m.Save("sigFPState", &x.sigFPState)
}

func (x *context64) afterLoad() {}
func (x *context64) load(m state.Map) {
	m.Load("State", &x.State)
	m.Load("sigFPState", &x.sigFPState)
}

func init() {
	state.Register("pkg/sentry/arch.context64", (*context64)(nil), state.Fns{Save: (*context64).save, Load: (*context64).load})
}