# swrite
Go library that implements a SlowWriter
PACKAGE DOCUMENTATION

```
package swriter
    import "."

    Package swriter implements a buffered writer that dumps Writes with a
    maxium frequence. The first write will wait for more writes and dump at
    a fixed frequence.

TYPES

type SlowWriter struct {
    // contains filtered or unexported fields
}

func New(w io.Writer, d time.Duration) *SlowWriter
    New returns a new SlowWriter

func (w *SlowWriter) Close() error
    Close stops the internal gorutines and return any error that could
    happen during Writes

func (w *SlowWriter) Write(data []byte) (int, error)
    Write will queue the data to be send and wait a max of the time specify
    during creation


```
