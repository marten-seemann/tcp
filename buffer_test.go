// Copyright 2014 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcp_test

import (
	"net"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/mikioh/tcp"
)

func TestCorkAndUncork(t *testing.T) {
	switch runtime.GOOS {
	case "darwin", "freebsd", "linux", "openbsd", "solaris":
	case "dragonfly":
		t.Log("you may need to adjust the net.inet.tcp.disable_nopush kernel state")
	default:
		t.Skipf("%s/%s", runtime.GOOS, runtime.GOARCH)
	}

	const N = 1280
	const M = N / 10

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := ln.(*net.TCPListener).SetDeadline(time.Now().Add(200 * time.Millisecond)); err != nil {
			t.Error(err)
			return
		}
		c, err := ln.Accept()
		if err != nil {
			t.Error(err)
			return
		}
		defer c.Close()
		b := make([]byte, N)
		n, err := c.Read(b)
		if err != nil {
			t.Error(err)
			return
		}
		if n != N {
			t.Errorf("got %d; want %d", n, N)
			return
		}
	}()

	c, err := net.Dial(ln.Addr().Network(), ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	tc, err := tcp.NewConn(c)
	if err != nil {
		t.Fatal(err)
	}
	defer tc.Close()
	if err := tc.Cork(); err != nil {
		t.Fatal(err)
	}
	b := make([]byte, N)
	for i := 0; i+M <= N; i += M {
		if _, err := tc.Write(b[i : i+M]); err != nil {
			t.Fatal(err)
		}
		time.Sleep(5 * time.Millisecond)
	}
	if err := tc.Uncork(); err != nil {
		t.Fatal(err)
	}

	wg.Wait()
}

func TestBufferOptions(t *testing.T) {
	switch runtime.GOOS {
	case "darwin", "linux":
	default:
		t.Skipf("%s/%s", runtime.GOOS, runtime.GOARCH)
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				break
			}
			defer c.Close()
		}
	}()

	c, err := net.Dial(ln.Addr().Network(), ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	tc, err := tcp.NewConn(c)
	if err != nil {
		t.Fatal(err)
	}
	defer tc.Close()

	opt := tcp.BufferOptions{
		UnsentThreshold: 1024,
	}
	if err := tc.SetBufferOptions(&opt); err != nil {
		t.Error(err)
	}
}

func TestBuffered(t *testing.T) {
	switch runtime.GOOS {
	case "darwin", "dragonfly", "freebsd", "linux", "netbsd", "openbsd":
	default:
		t.Skipf("%s/%s", runtime.GOOS, runtime.GOARCH)
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	m := []byte("HELLO-R-U-THERE")
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := ln.(*net.TCPListener).SetDeadline(time.Now().Add(200 * time.Millisecond)); err != nil {
			t.Error(err)
			return
		}
		c, err := ln.Accept()
		if err != nil {
			t.Error(err)
			return
		}
		defer c.Close()
		if err := c.(*net.TCPConn).SetReadBuffer(65535); err != nil {
			t.Error(err)
			return
		}
		tc, err := tcp.NewConn(c)
		if err != nil {
			t.Error(err)
			return
		}
		time.Sleep(100 * time.Millisecond)
		n := tc.Buffered()
		if n != len(m) {
			t.Errorf("got %d; want %d", n, len(m))
			return
		}
		t.Logf("%v bytes buffered to be read", n)
	}()

	c, err := net.Dial(ln.Addr().Network(), ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	if _, err := c.Write(m); err != nil {
		t.Fatal(err)
	}

	wg.Wait()
}

func TestAvailable(t *testing.T) {
	switch runtime.GOOS {
	case "freebsd", "linux", "netbsd":
	default:
		t.Skipf("%s/%s", runtime.GOOS, runtime.GOARCH)
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	m := []byte("HELLO-R-U-THERE")
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		d := net.Dialer{Timeout: 200 * time.Millisecond}
		c, err := d.Dial(ln.Addr().Network(), ln.Addr().String())
		if err != nil {
			t.Error(err)
			return
		}
		defer c.Close()
		if err := c.(*net.TCPConn).SetWriteBuffer(65535); err != nil {
			t.Error(err)
			return
		}
		tc, err := tcp.NewConn(c)
		if err != nil {
			t.Error(err)
			return
		}
		defer tc.Close()
		if _, err := c.Write(m); err != nil {
			t.Error(err)
			return
		}
		n := tc.Available()
		if n <= 0 {
			t.Errorf("got %d; want >0", n)
			return
		}
		t.Logf("%d bytes write available", n)
	}()

	c, err := ln.Accept()
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	wg.Wait()
}