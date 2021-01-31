package consistent

import (
	"hash/fnv"
	"sort"
	"strconv"
)

type Consistent struct {
	KeyList []int
}

func NewConsistent(keyListString []string) Consistent {
	keyList := make([]int, len(keyListString))

	for i := 0; i < len(keyListString); i++ {
		keyNum, _ := strconv.Atoi(keyListString[i])
		keyList[i] = keyNum
	}

	sort.Sort(sort.IntSlice(keyList))

	return Consistent{
		KeyList: keyList,
	}
}

func (c Consistent) GetNearestKey(
	rawKey string,
) string {
	h := fnv.New32a()
	h.Write([]byte(rawKey))
	realKey := h.Sum32()

	idx := sort.Search(len(c.KeyList), func(i int) bool {
		return c.KeyList[i] >= int(realKey)
	})

	if idx >= len(c.KeyList) {
		idx = 0
	}

	return strconv.Itoa(c.KeyList[idx])
}
