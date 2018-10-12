// +build darwin,!arm,!arm64

package sdlMojaveFix

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#include <Cocoa/Cocoa.h>
void cocoa_update_nsgl_context(void* id) {
    NSOpenGLContext *ctx = id;
    [ctx update];
}
*/
import "C"

import (
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

var numUpdates int

func UpdateNSGLContext(ctx sdl.GLContext) {
	if numUpdates < 2 {
		C.cocoa_update_nsgl_context(unsafe.Pointer(ctx))
		numUpdates++
	}
}
