package natsort

import (
	"sort"
	"strconv"
)

type VersionList []string

func (this VersionList) Len() int {
	return len(this)
}

func (this VersionList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this VersionList) Less(i, j int) bool {
	countI, countJ := len(this[i]), len(this[j])
	if (countI == 0 || countI > 0) && countJ == 0 {
		return false
	}

	if countI == 0 && countJ > 0 {
		return true
	}

	partI, partJ := make([]byte, 0, 4), make([]byte, 0, 4)
	kindI, okI, kindJ, okJ := 0, false, 0, false
	indexI, indexJ := 0, 0
	for indexI < countI && indexJ < countJ {
		partI, kindI, indexI, okI = getPart(this[i], indexI, partI[:0])
		partJ, kindJ, indexJ, okJ = getPart(this[j], indexJ, partJ[:0])

		if kindI != kindJ {
			return kindI > kindJ
		}

		ci, cj := len(partI), len(partJ)
		switch kindI {
		case Number:
			pi, _ := strconv.Atoi(string(partI))
			pj, _ := strconv.Atoi(string(partJ))
			if pi != pj {
				return pi < pj
			}
		case Alpha:
			for ii := 0; ii < ci && ii < cj; ii++ {
				if partI[ii] == partJ[ii] {
					continue
				}
				return partI[ii] < partJ[ii]
			}
		case Other:
			if partI[0] != partJ[0] {
				return partI[0] < partJ[0]
			}
		}
	}

	if countI > countJ {
		for indexI < countI {
			partI, kindI, indexI, okI = getPart(this[i], indexI, partI[:0])
			if !okI {
				break
			}
		}
	} else {
		for indexJ < countJ {
			partJ, kindJ, indexJ, okJ = getPart(this[j], indexJ, partJ[:0])
			if !okJ {
				break
			}
		}
	}

	if kindI != kindJ {
		return kindI > kindJ
	}
	return countI < countJ
}

func Sort(versions []string) []string {
	sort.Sort(VersionList(versions))
	return versions
}

func RSort(versions []string) []string {
	sort.Sort(VersionList(versions))
	count := len(versions) - 1
	for i, j := 0, count/2; i <= j; i++ {
		versions[i], versions[count-i] = versions[count-i], versions[i]
	}
	return versions
}
