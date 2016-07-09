// Copyright 2014 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!windows

package tcp

var soOptions = [soMax]soOption{
	soBuffered:  {0, -1},
	soAvailable: {0, -1},
	soCork:      {0, -1},
}

func buffered(s uintptr) int  { return -1 }
func available(s uintptr) int { return -1 }
