// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Fri, 02 Sep 2016 17:34:13 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

package pocketsphinx

/*
#cgo pkg-config: pocketsphinx
#include "pocketsphinx.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"sync"
	"unsafe"
)

// Ref returns a reference.
func (x *Decoder) Ref() *C.ps_decoder_t {
	if x == nil {
		return nil
	}
	return (*C.ps_decoder_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *Decoder) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewDecoderRef initialises a new struct.
func NewDecoderRef(ref *C.ps_decoder_t) *Decoder {
	return (*Decoder)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *Decoder) PassRef() *C.ps_decoder_t {
	if x == nil {
		x = new(Decoder)
	}
	return (*C.ps_decoder_t)(unsafe.Pointer(x))
}

// Ref returns a reference.
func (x *Nbest) Ref() *C.ps_nbest_t {
	if x == nil {
		return nil
	}
	return (*C.ps_nbest_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *Nbest) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewNbestRef initialises a new struct.
func NewNbestRef(ref *C.ps_nbest_t) *Nbest {
	return (*Nbest)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *Nbest) PassRef() *C.ps_nbest_t {
	if x == nil {
		x = new(Nbest)
	}
	return (*C.ps_nbest_t)(unsafe.Pointer(x))
}

// Ref returns a reference.
func (x *Seg) Ref() *C.ps_seg_t {
	if x == nil {
		return nil
	}
	return (*C.ps_seg_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *Seg) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewSegRef initialises a new struct.
func NewSegRef(ref *C.ps_seg_t) *Seg {
	return (*Seg)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *Seg) PassRef() *C.ps_seg_t {
	if x == nil {
		x = new(Seg)
	}
	return (*C.ps_seg_t)(unsafe.Pointer(x))
}

// Ref returns a reference.
func (x *File) Ref() *C.FILE {
	if x == nil {
		return nil
	}
	return (*C.FILE)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *File) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFileRef initialises a new struct.
func NewFileRef(ref *C.FILE) *File {
	return (*File)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *File) PassRef() *C.FILE {
	if x == nil {
		x = new(File)
	}
	return (*C.FILE)(unsafe.Pointer(x))
}

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocArgMemory allocates memory for type C.arg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocArgMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfArgValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfArgValue = unsafe.Sizeof([1]C.arg_t{})

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

type stringHeader struct {
	Data uintptr
	Len  int
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = uintptr(unsafe.Pointer(p))
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
	}
	return
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(unsafe.Pointer(h.Data)), C.int(h.Len))
}

// Ref returns a reference.
func (x *Arg) Ref() *C.arg_t {
	if x == nil {
		return nil
	}
	return x.ref8495e8ec
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *Arg) Free() {
	if x != nil && x.allocs8495e8ec != nil {
		x.allocs8495e8ec.(*cgoAllocMap).Free()
		x.ref8495e8ec = nil
	}
}

// NewArgRef initialises a new struct holding the reference to the originaitng C struct.
func NewArgRef(ref unsafe.Pointer) *Arg {
	if ref == nil {
		return nil
	}
	obj := new(Arg)
	obj.ref8495e8ec = (*C.arg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns a reference and creates new C object if no refernce yet.
func (x *Arg) PassRef() (*C.arg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref8495e8ec != nil {
		return x.ref8495e8ec, nil
	}
	mem8495e8ec := allocArgMemory(1)
	ref8495e8ec := (*C.arg_t)(mem8495e8ec)
	allocs8495e8ec := new(cgoAllocMap)
	var cname_allocs *cgoAllocMap
	ref8495e8ec.name, cname_allocs = unpackPCharString(x.Name)
	allocs8495e8ec.Borrow(cname_allocs)

	var c_type_allocs *cgoAllocMap
	ref8495e8ec._type, c_type_allocs = (C.int)(x.Type), cgoAllocsUnknown
	allocs8495e8ec.Borrow(c_type_allocs)

	var cdeflt_allocs *cgoAllocMap
	ref8495e8ec.deflt, cdeflt_allocs = unpackPCharString(x.Deflt)
	allocs8495e8ec.Borrow(cdeflt_allocs)

	var cdoc_allocs *cgoAllocMap
	ref8495e8ec.doc, cdoc_allocs = unpackPCharString(x.Doc)
	allocs8495e8ec.Borrow(cdoc_allocs)

	x.ref8495e8ec = ref8495e8ec
	x.allocs8495e8ec = allocs8495e8ec
	return ref8495e8ec, allocs8495e8ec

}

// PassValue creates a new C object if no refernce yet and returns the dereferenced value.
func (x Arg) PassValue() (C.arg_t, *cgoAllocMap) {
	if x.ref8495e8ec != nil {
		return *x.ref8495e8ec, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref reads the internal fields of struct from its C pointer.
func (x *Arg) Deref() {
	if x.ref8495e8ec == nil {
		return
	}
	x.Name = packPCharString(x.ref8495e8ec.name)
	x.Type = (int32)(x.ref8495e8ec._type)
	x.Deflt = packPCharString(x.ref8495e8ec.deflt)
	x.Doc = packPCharString(x.ref8495e8ec.doc)
}

// Ref returns a reference.
func (x *CommandLn) Ref() *C.cmd_ln_t {
	if x == nil {
		return nil
	}
	return (*C.cmd_ln_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *CommandLn) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewCommandLnRef initialises a new struct.
func NewCommandLnRef(ref *C.cmd_ln_t) *CommandLn {
	return (*CommandLn)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *CommandLn) PassRef() *C.cmd_ln_t {
	if x == nil {
		x = new(CommandLn)
	}
	return (*C.cmd_ln_t)(unsafe.Pointer(x))
}

// Ref returns a reference.
func (x *Logmath) Ref() *C.logmath_t {
	if x == nil {
		return nil
	}
	return (*C.logmath_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *Logmath) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewLogmathRef initialises a new struct.
func NewLogmathRef(ref *C.logmath_t) *Logmath {
	return (*Logmath)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *Logmath) PassRef() *C.logmath_t {
	if x == nil {
		x = new(Logmath)
	}
	return (*C.logmath_t)(unsafe.Pointer(x))
}

// Ref returns a reference.
func (x *Fe) Ref() *C.fe_t {
	if x == nil {
		return nil
	}
	return (*C.fe_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *Fe) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFeRef initialises a new struct.
func NewFeRef(ref *C.fe_t) *Fe {
	return (*Fe)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *Fe) PassRef() *C.fe_t {
	if x == nil {
		x = new(Fe)
	}
	return (*C.fe_t)(unsafe.Pointer(x))
}

// Ref returns a reference.
func (x *NgramModel) Ref() *C.ngram_model_t {
	if x == nil {
		return nil
	}
	return (*C.ngram_model_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *NgramModel) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewNgramModelRef initialises a new struct.
func NewNgramModelRef(ref *C.ngram_model_t) *NgramModel {
	return (*NgramModel)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *NgramModel) PassRef() *C.ngram_model_t {
	if x == nil {
		x = new(NgramModel)
	}
	return (*C.ngram_model_t)(unsafe.Pointer(x))
}

// Ref returns a reference.
func (x *NgramClass) Ref() *C.ngram_class_t {
	if x == nil {
		return nil
	}
	return (*C.ngram_class_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *NgramClass) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewNgramClassRef initialises a new struct.
func NewNgramClassRef(ref *C.ngram_class_t) *NgramClass {
	return (*NgramClass)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *NgramClass) PassRef() *C.ngram_class_t {
	if x == nil {
		x = new(NgramClass)
	}
	return (*C.ngram_class_t)(unsafe.Pointer(x))
}

// Ref returns a reference.
func (x *NgramIter) Ref() *C.ngram_iter_t {
	if x == nil {
		return nil
	}
	return (*C.ngram_iter_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *NgramIter) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewNgramIterRef initialises a new struct.
func NewNgramIterRef(ref *C.ngram_iter_t) *NgramIter {
	return (*NgramIter)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *NgramIter) PassRef() *C.ngram_iter_t {
	if x == nil {
		x = new(NgramIter)
	}
	return (*C.ngram_iter_t)(unsafe.Pointer(x))
}

// Ref returns a reference.
func (x *NgramModelSetIter) Ref() *C.ngram_model_set_iter_t {
	if x == nil {
		return nil
	}
	return (*C.ngram_model_set_iter_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *NgramModelSetIter) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewNgramModelSetIterRef initialises a new struct.
func NewNgramModelSetIterRef(ref *C.ngram_model_set_iter_t) *NgramModelSetIter {
	return (*NgramModelSetIter)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *NgramModelSetIter) PassRef() *C.ngram_model_set_iter_t {
	if x == nil {
		x = new(NgramModelSetIter)
	}
	return (*C.ngram_model_set_iter_t)(unsafe.Pointer(x))
}

// Ref returns a reference.
func (x *Lattice) Ref() *C.ps_lattice_t {
	if x == nil {
		return nil
	}
	return (*C.ps_lattice_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *Lattice) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewLatticeRef initialises a new struct.
func NewLatticeRef(ref *C.ps_lattice_t) *Lattice {
	return (*Lattice)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *Lattice) PassRef() *C.ps_lattice_t {
	if x == nil {
		x = new(Lattice)
	}
	return (*C.ps_lattice_t)(unsafe.Pointer(x))
}

// Ref returns a reference.
func (x *Latnode) Ref() *C.ps_latnode_t {
	if x == nil {
		return nil
	}
	return (*C.ps_latnode_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *Latnode) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewLatnodeRef initialises a new struct.
func NewLatnodeRef(ref *C.ps_latnode_t) *Latnode {
	return (*Latnode)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *Latnode) PassRef() *C.ps_latnode_t {
	if x == nil {
		x = new(Latnode)
	}
	return (*C.ps_latnode_t)(unsafe.Pointer(x))
}

// Ref returns a reference.
func (x *LatnodeIter) Ref() *C.ps_latnode_iter_t {
	if x == nil {
		return nil
	}
	return (*C.ps_latnode_iter_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *LatnodeIter) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewLatnodeIterRef initialises a new struct.
func NewLatnodeIterRef(ref *C.ps_latnode_iter_t) *LatnodeIter {
	return (*LatnodeIter)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *LatnodeIter) PassRef() *C.ps_latnode_iter_t {
	if x == nil {
		x = new(LatnodeIter)
	}
	return (*C.ps_latnode_iter_t)(unsafe.Pointer(x))
}

// Ref returns a reference.
func (x *Latlink) Ref() *C.ps_latlink_t {
	if x == nil {
		return nil
	}
	return (*C.ps_latlink_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *Latlink) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewLatlinkRef initialises a new struct.
func NewLatlinkRef(ref *C.ps_latlink_t) *Latlink {
	return (*Latlink)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *Latlink) PassRef() *C.ps_latlink_t {
	if x == nil {
		x = new(Latlink)
	}
	return (*C.ps_latlink_t)(unsafe.Pointer(x))
}

// Ref returns a reference.
func (x *LatlinkIter) Ref() *C.ps_latlink_iter_t {
	if x == nil {
		return nil
	}
	return (*C.ps_latlink_iter_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *LatlinkIter) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewLatlinkIterRef initialises a new struct.
func NewLatlinkIterRef(ref *C.ps_latlink_iter_t) *LatlinkIter {
	return (*LatlinkIter)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *LatlinkIter) PassRef() *C.ps_latlink_iter_t {
	if x == nil {
		x = new(LatlinkIter)
	}
	return (*C.ps_latlink_iter_t)(unsafe.Pointer(x))
}

// Ref returns a reference.
func (x *Mllr) Ref() *C.ps_mllr_t {
	if x == nil {
		return nil
	}
	return (*C.ps_mllr_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *Mllr) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewMllrRef initialises a new struct.
func NewMllrRef(ref *C.ps_mllr_t) *Mllr {
	return (*Mllr)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *Mllr) PassRef() *C.ps_mllr_t {
	if x == nil {
		x = new(Mllr)
	}
	return (*C.ps_mllr_t)(unsafe.Pointer(x))
}

// Ref returns a reference.
func (x *FsgSet) Ref() *C.fsg_set_t {
	if x == nil {
		return nil
	}
	return (*C.fsg_set_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *FsgSet) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFsgSetRef initialises a new struct.
func NewFsgSetRef(ref *C.fsg_set_t) *FsgSet {
	return (*FsgSet)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *FsgSet) PassRef() *C.fsg_set_t {
	if x == nil {
		x = new(FsgSet)
	}
	return (*C.fsg_set_t)(unsafe.Pointer(x))
}

// allocHashIterMemory allocates memory for type C.hash_iter_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocHashIterMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfHashIterValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfHashIterValue = unsafe.Sizeof([1]C.hash_iter_t{})

// Ref returns a reference.
func (x *FsgSetIter) Ref() *C.hash_iter_t {
	if x == nil {
		return nil
	}
	return x.refb26723b3
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *FsgSetIter) Free() {
	if x != nil && x.allocsb26723b3 != nil {
		x.allocsb26723b3.(*cgoAllocMap).Free()
		x.refb26723b3 = nil
	}
}

// NewFsgSetIterRef initialises a new struct holding the reference to the originaitng C struct.
func NewFsgSetIterRef(ref unsafe.Pointer) *FsgSetIter {
	if ref == nil {
		return nil
	}
	obj := new(FsgSetIter)
	obj.refb26723b3 = (*C.hash_iter_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns a reference and creates new C object if no refernce yet.
func (x *FsgSetIter) PassRef() (*C.hash_iter_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refb26723b3 != nil {
		return x.refb26723b3, nil
	}
	memb26723b3 := allocHashIterMemory(1)
	refb26723b3 := (*C.hash_iter_t)(memb26723b3)
	allocsb26723b3 := new(cgoAllocMap)
	var cidx_allocs *cgoAllocMap
	refb26723b3.idx, cidx_allocs = (C.size_t)(x.Idx), cgoAllocsUnknown
	allocsb26723b3.Borrow(cidx_allocs)

	x.refb26723b3 = refb26723b3
	x.allocsb26723b3 = allocsb26723b3
	return refb26723b3, allocsb26723b3

}

// PassValue creates a new C object if no refernce yet and returns the dereferenced value.
func (x FsgSetIter) PassValue() (C.hash_iter_t, *cgoAllocMap) {
	if x.refb26723b3 != nil {
		return *x.refb26723b3, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref reads the internal fields of struct from its C pointer.
func (x *FsgSetIter) Deref() {
	if x.refb26723b3 == nil {
		return
	}
	x.Idx = (uint)(x.refb26723b3.idx)
}

// allocFsgLinkMemory allocates memory for type C.fsg_link_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFsgLinkMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFsgLinkValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfFsgLinkValue = unsafe.Sizeof([1]C.fsg_link_t{})

// Ref returns a reference.
func (x *FsgLink) Ref() *C.fsg_link_t {
	if x == nil {
		return nil
	}
	return x.refe7ca59ff
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *FsgLink) Free() {
	if x != nil && x.allocse7ca59ff != nil {
		x.allocse7ca59ff.(*cgoAllocMap).Free()
		x.refe7ca59ff = nil
	}
}

// NewFsgLinkRef initialises a new struct holding the reference to the originaitng C struct.
func NewFsgLinkRef(ref unsafe.Pointer) *FsgLink {
	if ref == nil {
		return nil
	}
	obj := new(FsgLink)
	obj.refe7ca59ff = (*C.fsg_link_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns a reference and creates new C object if no refernce yet.
func (x *FsgLink) PassRef() (*C.fsg_link_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refe7ca59ff != nil {
		return x.refe7ca59ff, nil
	}
	meme7ca59ff := allocFsgLinkMemory(1)
	refe7ca59ff := (*C.fsg_link_t)(meme7ca59ff)
	allocse7ca59ff := new(cgoAllocMap)
	var cfrom_state_allocs *cgoAllocMap
	refe7ca59ff.from_state, cfrom_state_allocs = (C.int32)(x.FromState), cgoAllocsUnknown
	allocse7ca59ff.Borrow(cfrom_state_allocs)

	var cto_state_allocs *cgoAllocMap
	refe7ca59ff.to_state, cto_state_allocs = (C.int32)(x.ToState), cgoAllocsUnknown
	allocse7ca59ff.Borrow(cto_state_allocs)

	var clogs2prob_allocs *cgoAllocMap
	refe7ca59ff.logs2prob, clogs2prob_allocs = (C.int32)(x.Logs2prob), cgoAllocsUnknown
	allocse7ca59ff.Borrow(clogs2prob_allocs)

	var cwid_allocs *cgoAllocMap
	refe7ca59ff.wid, cwid_allocs = (C.int32)(x.Wid), cgoAllocsUnknown
	allocse7ca59ff.Borrow(cwid_allocs)

	x.refe7ca59ff = refe7ca59ff
	x.allocse7ca59ff = allocse7ca59ff
	return refe7ca59ff, allocse7ca59ff

}

// PassValue creates a new C object if no refernce yet and returns the dereferenced value.
func (x FsgLink) PassValue() (C.fsg_link_t, *cgoAllocMap) {
	if x.refe7ca59ff != nil {
		return *x.refe7ca59ff, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref reads the internal fields of struct from its C pointer.
func (x *FsgLink) Deref() {
	if x.refe7ca59ff == nil {
		return
	}
	x.FromState = (int32)(x.refe7ca59ff.from_state)
	x.ToState = (int32)(x.refe7ca59ff.to_state)
	x.Logs2prob = (int32)(x.refe7ca59ff.logs2prob)
	x.Wid = (int32)(x.refe7ca59ff.wid)
}

// allocFsgModelMemory allocates memory for type C.fsg_model_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFsgModelMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFsgModelValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfFsgModelValue = unsafe.Sizeof([1]C.fsg_model_t{})

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// allocPCharMemory allocates memory for type *C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPCharValue = unsafe.Sizeof([1]*C.char{})

const sizeOfPtr = unsafe.Sizeof(&struct{}{})

// unpackSSByte transforms a sliced Go data structure into plain C format.
func unpackSSByte(x [][]byte) (unpacked **C.char, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.char) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPCharMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.char)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.char)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.char)(unsafe.Pointer(h.Data))
	return
}

// packSSByte reads sliced Go data structure out from plain C format.
func packSSByte(v [][]byte, ptr0 **C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.char)(unsafe.Pointer(ptr0)))[i0]
		hxf95e7c8 := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxf95e7c8.Data = uintptr(unsafe.Pointer(ptr1))
		hxf95e7c8.Cap = 0x7fffffff
		// hxf95e7c8.Len = ?
	}
}

// Ref returns a reference.
func (x *FsgModel) Ref() *C.fsg_model_t {
	if x == nil {
		return nil
	}
	return x.ref22935ee7
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *FsgModel) Free() {
	if x != nil && x.allocs22935ee7 != nil {
		x.allocs22935ee7.(*cgoAllocMap).Free()
		x.ref22935ee7 = nil
	}
}

// NewFsgModelRef initialises a new struct holding the reference to the originaitng C struct.
func NewFsgModelRef(ref unsafe.Pointer) *FsgModel {
	if ref == nil {
		return nil
	}
	obj := new(FsgModel)
	obj.ref22935ee7 = (*C.fsg_model_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns a reference and creates new C object if no refernce yet.
func (x *FsgModel) PassRef() (*C.fsg_model_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref22935ee7 != nil {
		return x.ref22935ee7, nil
	}
	mem22935ee7 := allocFsgModelMemory(1)
	ref22935ee7 := (*C.fsg_model_t)(mem22935ee7)
	allocs22935ee7 := new(cgoAllocMap)
	var crefcount_allocs *cgoAllocMap
	ref22935ee7.refcount, crefcount_allocs = (C.int)(x.Refcount), cgoAllocsUnknown
	allocs22935ee7.Borrow(crefcount_allocs)

	var cname_allocs *cgoAllocMap
	ref22935ee7.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Name)).Data)), cgoAllocsUnknown
	allocs22935ee7.Borrow(cname_allocs)

	var cn_word_allocs *cgoAllocMap
	ref22935ee7.n_word, cn_word_allocs = (C.int32)(x.NWord), cgoAllocsUnknown
	allocs22935ee7.Borrow(cn_word_allocs)

	var cn_word_alloc_allocs *cgoAllocMap
	ref22935ee7.n_word_alloc, cn_word_alloc_allocs = (C.int32)(x.NWordAlloc), cgoAllocsUnknown
	allocs22935ee7.Borrow(cn_word_alloc_allocs)

	var cvocab_allocs *cgoAllocMap
	ref22935ee7.vocab, cvocab_allocs = unpackSSByte(x.Vocab)
	allocs22935ee7.Borrow(cvocab_allocs)

	var csilwords_allocs *cgoAllocMap
	ref22935ee7.silwords, csilwords_allocs = (*C.bitvec_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Silwords)).Data)), cgoAllocsUnknown
	allocs22935ee7.Borrow(csilwords_allocs)

	var caltwords_allocs *cgoAllocMap
	ref22935ee7.altwords, caltwords_allocs = (*C.bitvec_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Altwords)).Data)), cgoAllocsUnknown
	allocs22935ee7.Borrow(caltwords_allocs)

	var clmath_allocs *cgoAllocMap
	ref22935ee7.lmath, clmath_allocs = (*C.logmath_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Lmath)).Data)), cgoAllocsUnknown
	allocs22935ee7.Borrow(clmath_allocs)

	var cn_state_allocs *cgoAllocMap
	ref22935ee7.n_state, cn_state_allocs = (C.int32)(x.NState), cgoAllocsUnknown
	allocs22935ee7.Borrow(cn_state_allocs)

	var cstart_state_allocs *cgoAllocMap
	ref22935ee7.start_state, cstart_state_allocs = (C.int32)(x.StartState), cgoAllocsUnknown
	allocs22935ee7.Borrow(cstart_state_allocs)

	var cfinal_state_allocs *cgoAllocMap
	ref22935ee7.final_state, cfinal_state_allocs = (C.int32)(x.FinalState), cgoAllocsUnknown
	allocs22935ee7.Borrow(cfinal_state_allocs)

	var clw_allocs *cgoAllocMap
	ref22935ee7.lw, clw_allocs = (C.float32)(x.Lw), cgoAllocsUnknown
	allocs22935ee7.Borrow(clw_allocs)

	x.ref22935ee7 = ref22935ee7
	x.allocs22935ee7 = allocs22935ee7
	return ref22935ee7, allocs22935ee7

}

// PassValue creates a new C object if no refernce yet and returns the dereferenced value.
func (x FsgModel) PassValue() (C.fsg_model_t, *cgoAllocMap) {
	if x.ref22935ee7 != nil {
		return *x.ref22935ee7, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref reads the internal fields of struct from its C pointer.
func (x *FsgModel) Deref() {
	if x.ref22935ee7 == nil {
		return
	}
	x.Refcount = (int32)(x.ref22935ee7.refcount)
	hxfc4425b := (*sliceHeader)(unsafe.Pointer(&x.Name))
	hxfc4425b.Data = uintptr(unsafe.Pointer(x.ref22935ee7.name))
	hxfc4425b.Cap = 0x7fffffff
	// hxfc4425b.Len = ?

	x.NWord = (int32)(x.ref22935ee7.n_word)
	x.NWordAlloc = (int32)(x.ref22935ee7.n_word_alloc)
	packSSByte(x.Vocab, x.ref22935ee7.vocab)
	hxff2234b := (*sliceHeader)(unsafe.Pointer(&x.Silwords))
	hxff2234b.Data = uintptr(unsafe.Pointer(x.ref22935ee7.silwords))
	hxff2234b.Cap = 0x7fffffff
	// hxff2234b.Len = ?

	hxff73280 := (*sliceHeader)(unsafe.Pointer(&x.Altwords))
	hxff73280.Data = uintptr(unsafe.Pointer(x.ref22935ee7.altwords))
	hxff73280.Cap = 0x7fffffff
	// hxff73280.Len = ?

	hxfa9955c := (*sliceHeader)(unsafe.Pointer(&x.Lmath))
	hxfa9955c.Data = uintptr(unsafe.Pointer(x.ref22935ee7.lmath))
	hxfa9955c.Cap = 0x7fffffff
	// hxfa9955c.Len = ?

	x.NState = (int32)(x.ref22935ee7.n_state)
	x.StartState = (int32)(x.ref22935ee7.start_state)
	x.FinalState = (int32)(x.ref22935ee7.final_state)
	x.Lw = (float32)(x.ref22935ee7.lw)
}

// Ref returns a reference.
func (x *FsgArciter) Ref() *C.fsg_arciter_t {
	if x == nil {
		return nil
	}
	return (*C.fsg_arciter_t)(unsafe.Pointer(x))
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *FsgArciter) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFsgArciterRef initialises a new struct.
func NewFsgArciterRef(ref *C.fsg_arciter_t) *FsgArciter {
	return (*FsgArciter)(unsafe.Pointer(ref))
}

// PassRef returns a reference and creates new object if no refernce yet.
func (x *FsgArciter) PassRef() *C.fsg_arciter_t {
	if x == nil {
		x = new(FsgArciter)
	}
	return (*C.fsg_arciter_t)(unsafe.Pointer(x))
}

// allocPMfccMemory allocates memory for type *C.mfcc_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPMfccMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPMfccValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPMfccValue = unsafe.Sizeof([1]*C.mfcc_t{})

// unpackArgSSFloat transforms a sliced Go data structure into plain C format.
func unpackArgSSFloat(x [][]float32) (unpacked **C.mfcc_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.mfcc_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPMfccMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.mfcc_t)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.mfcc_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.mfcc_t)(unsafe.Pointer(h.Data))
	return
}

// packSSFloat reads sliced Go data structure out from plain C format.
func packSSFloat(v [][]float32, ptr0 **C.mfcc_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.mfcc_t)(unsafe.Pointer(ptr0)))[i0]
		hxfa3f05c := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxfa3f05c.Data = uintptr(unsafe.Pointer(ptr1))
		hxfa3f05c.Cap = 0x7fffffff
		// hxfa3f05c.Len = ?
	}
}

// unpackArgSString transforms a sliced Go data structure into plain C format.
func unpackArgSString(x []string) (unpacked **C.char, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.char) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPCharMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.char)(unsafe.Pointer(h0))
	for i0 := range x {
		v0[i0], _ = unpackPCharString(x[i0])
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.char)(unsafe.Pointer(h.Data))
	return
}

// packSString reads sliced Go data structure out from plain C format.
func packSString(v []string, ptr0 **C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.char)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = packPCharString(ptr1)
	}
}

// unpackArgSSByte transforms a sliced Go data structure into plain C format.
func unpackArgSSByte(x [][]byte) (unpacked **C.char, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.char) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPCharMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.char)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.char)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.char)(unsafe.Pointer(h.Data))
	return
}
