package unionfind

// IDSlice represent quick find
type IDSlice []int

// NewQuickFind initialize new
func NewQuickFind(n int) IDSlice {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}
	return data
}

// Connected check connection between p and q.
func (id IDSlice) Connected(p, q int) bool { return id[p] == id[q] }

// Union merge components p and q by id.
func (id IDSlice) Union(p, q int) {
	var pid = id[p]
	var qid = id[q]
	for i := 0; i < len(id); i++ {
		if id[i] == pid {
			id[i] = qid
		}
	}
}
