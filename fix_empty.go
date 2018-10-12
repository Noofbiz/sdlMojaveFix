// +build !darwin arm arm64

package sdlMojaveFix

import "github.com/veandco/go-sdl2/sdl"

func UpdateNSGLContext(ctx sdl.GLContext) {}
