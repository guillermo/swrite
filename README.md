# swrite
Go library that implements a SlowWriter
PACKAGE DOCUMENTATION

```
PACKAGE DOCUMENTATION

package swriter
    import "."

    Package swriter implements a buffered writer that dumps Writes with a
    maxium frequence. SlowWriter will save the content of the Writes and
    dump after it timesout.

	w := swriter.New(os.Stdout, time.Second / 60)
	w.Write("hello")
	w.Write("world")
	// After 1/60 seconds it will dump to os.Stdout all the previous writes.

TYPES

type SlowWriter struct {
    // contains filtered or unexported fields
}

func New(w io.Writer, d time.Duration) *SlowWriter
    New returns a new SlowWriter

func (w *SlowWriter) Close() error
    Close stops the internal gorutines and return any error that could
    happen during Writes

func (w *SlowWriter) Flush()
    Flush will inmediatly flush all the data

func (w *SlowWriter) Write(data []byte) (int, error)
    Write will queue the data to be send and wait a max of the time specify
    during creation

```
