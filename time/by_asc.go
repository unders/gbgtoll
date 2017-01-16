package time

import "time"

// ByAsc implements sort.Interface
type ByAsc []time.Time

func (t ByAsc) Len() int           { return len(t) }
func (t ByAsc) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t ByAsc) Less(i, j int) bool { return t[j].After(t[i]) }
