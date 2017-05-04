//Package mtio is a library to interface with magnetic tape drive devices
//using the mt driver ioctls.
package mtio

import (
	"fmt"
	"os"
	"unsafe"
)

//NewMtOp returns a pointer to a new MtOp
func NewMtOp(opts ...Option) *MtOp {
	m := &MtOp{
		//default op is MTRESET (0)
		count: 1,
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

//Option sets various options for NewMtOp
type Option func(*MtOp)

//WithOperation sets the operation of the MtOp
func WithOperation(op int16) Option {
	return func(m *MtOp) { m.op = op }
}

//WithCount sets the operation of the MtOp
func WithCount(count int32) Option {
	return func(m *MtOp) { m.count = count }
}

//DoOp executes the MTIOCTOP ioctl for the given MtOp
//The first argument is the open file handle on the
//approrpriate tape device
func DoOp(f *os.File, m *MtOp) error {
	if m.op != MTMKPART && m.count < 0 {
		return fmt.Errorf("negative repeat count")
	}
	return ioctl(f.Fd(), MTIOCTOP, uintptr(unsafe.Pointer(m)))
}

//GetStatus takes the open file handle of the appropriate
//tape device and returns a pointer to the MtGet Status
func GetStatus(f *os.File) (*MtGet, error) {
	m := &MtGet{}
	err := ioctl(f.Fd(), MTIOCGET, uintptr(unsafe.Pointer(m)))
	if err != nil {
		return &MtGet{}, err
	}
	return m, nil
}

//GetPos takes the open file handle of the appropriate
//tape device and returns a pointer to the MtPos position
//information
// NB: this seems to get EIO on many devices, better to use status results
func GetPos(f *os.File) (*MtPos, error) {
	m := &MtPos{}
	err := ioctl(f.Fd(), MTIOCPOS, uintptr(unsafe.Pointer(m)))
	if err != nil {
		return &MtPos{}, err
	}
	return m, nil
}
