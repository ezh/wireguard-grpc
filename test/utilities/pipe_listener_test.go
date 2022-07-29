package utilities_test

import (
	"context"
	"testing"
	"time"

	"github.com/ezh/wireguard-grpc/test/utilities"
)

func TestPipeListener(t *testing.T) {
	pl := utilities.NewPipeListener()
	recvdBytes := make(chan []byte, 1)
	const want = "hello world"

	go func() {
		c, err := pl.Accept()
		if err != nil {
			t.Error(err)
		}

		read := make([]byte, len(want))
		_, err = c.Read(read)
		if err != nil {
			t.Error(err)
		}
		recvdBytes <- read
	}()

	dl := pl.Dialer()
	conn, err := dl(context.TODO(), "")
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Write([]byte(want))
	if err != nil {
		t.Fatal(err)
	}

	select {
	case gotBytes := <-recvdBytes:
		got := string(gotBytes)
		if got != want {
			t.Fatalf("expected to get %s, got %s", got, want)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timed out waiting for server to receive bytes")
	}
}

func TestUnblocking(t *testing.T) {
	for _, test := range []struct {
		desc                 string
		blockFuncShouldError bool
		blockFunc            func(*utilities.PipeListener, chan struct{}) error
		unblockFunc          func(*utilities.PipeListener) error
	}{
		{
			desc: "Accept unblocks Dial",
			blockFunc: func(pl *utilities.PipeListener, done chan struct{}) error {
				dl := pl.Dialer()
				_, err := dl(context.TODO(), "")
				close(done)
				return err
			},
			unblockFunc: func(pl *utilities.PipeListener) error {
				_, err := pl.Accept()
				return err
			},
		},
		{
			desc:                 "Close unblocks Dial",
			blockFuncShouldError: true, // because pl.Close will be called
			blockFunc: func(pl *utilities.PipeListener, done chan struct{}) error {
				dl := pl.Dialer()
				_, err := dl(context.TODO(), "")
				close(done)
				return err
			},
			unblockFunc: func(pl *utilities.PipeListener) error {
				return pl.Close()
			},
		},
		{
			desc: "Dial unblocks Accept",
			blockFunc: func(pl *utilities.PipeListener, done chan struct{}) error {
				_, err := pl.Accept()
				close(done)
				return err
			},
			unblockFunc: func(pl *utilities.PipeListener) error {
				dl := pl.Dialer()
				_, err := dl(context.TODO(), "")
				return err
			},
		},
		{
			desc:                 "Close unblocks Accept",
			blockFuncShouldError: true, // because pl.Close will be called
			blockFunc: func(pl *utilities.PipeListener, done chan struct{}) error {
				_, err := pl.Accept()
				close(done)
				return err
			},
			unblockFunc: func(pl *utilities.PipeListener) error {
				return pl.Close()
			},
		},
	} {
		t.Log(test.desc)
		testUnblocking(t, test.blockFunc, test.unblockFunc, test.blockFuncShouldError)
	}
}

func testUnblocking(t *testing.T, blockFunc func(*utilities.PipeListener, chan struct{}) error, unblockFunc func(*utilities.PipeListener) error, blockFuncShouldError bool) {
	pl := utilities.NewPipeListener()
	dialFinished := make(chan struct{})

	go func() {
		err := blockFunc(pl, dialFinished)
		if blockFuncShouldError && err == nil {
			t.Error("expected blocking func to return error because pl.Close was called, but got nil")
		}

		if !blockFuncShouldError && err != nil {
			t.Error(err)
		}
	}()

	select {
	case <-dialFinished:
		t.Fatal("expected Dial to block until pl.Close or pl.Accept")
	default:
	}

	if err := unblockFunc(pl); err != nil {
		t.Fatal(err)
	}

	select {
	case <-dialFinished:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("expected Accept to unblock after pl.Accept was called")
	}
}
