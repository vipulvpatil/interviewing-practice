package sortable

type Sortable interface {
	Swap(i, j int)
	Less(i, j int) bool
	Len() int
}

type Sorter interface {
	Sort(arr Sortable)
}
