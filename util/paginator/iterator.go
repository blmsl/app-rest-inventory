package paginator

import (
	"sync"
)

//Iterator Simple object to iteratie over a pager
type Iterator struct {
	paginator *Paginator
	pageCount int
	mux       sync.Mutex
}

//GetIterator Generates an Iterator over a Pager
func (p *Paginator) GetIterator() (*Iterator, error) {
	return &Iterator{paginator: p}, nil
}

//GetCursorPosition Gets the current page number
func (i *Iterator) GetCursorPosition() int {
	return i.pageCount
}

//Next Gets the Next Page. Also moves the cursor position position ahead.
func (i *Iterator) Next() (interface{}, error) {
	i.mux.Lock()
	defer i.mux.Unlock()

	i.pageCount++
	return i.paginator.Page(i.pageCount)
}

//Prev Gets the previous page. Also takes the cursor back
func (i *Iterator) Prev() (interface{}, error) {
	i.mux.Lock()
	defer i.mux.Unlock()

	i.pageCount--
	return i.paginator.Page(i.pageCount)
}

//Peek Gets the next page without moving the cursor position ahead
func (i *Iterator) Peek() (interface{}, error) {
	return i.paginator.Page(i.pageCount)
}
