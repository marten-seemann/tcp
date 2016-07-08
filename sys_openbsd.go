// Copyright 2014 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcp

var soOptions = map[int]soOption{
	soBuffered: {0, sysFIONREAD},
	soCork:     {ianaProtocolTCP, sysTCP_NOPUSH},
}
