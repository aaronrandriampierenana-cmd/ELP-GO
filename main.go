import (
	"runtime"
	"sync"
)
func partition(l []int) int {
	pivot := l[0]
	i := 1
	j := len(l) - 1

	for i <= j {
		for i < len(l) && l[i] <= pivot {
			i++
		}
		for j > 0 && l[j] >= pivot {
			j--
		}
		if i < j {
			l[i], l[j] = l[j], l[i]
		}
	}
	l[0], l[j] = l[j], l[0]
	return j
}
func QuickSortSeq(l []int) {
	if len(l) <= 1 {
		return
	}
	pivotIdx := partition(l)
	QuickSortSeq(l[:pivotIdx])
	QuickSortSeq(l[pivotIdx+1:])
}
func merge(left, right []int) []int {
	taille := len(left) + len(right)
	result := make([]int, 0, taille)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func FusionnerSegments(segments [][]int) []int {
	if len(segments) == 0 {
		return []int{}
	}
	if len(segments) == 1 {
		return segments[0]
	}
	if len(segments) == 2 {
		return merge(segments[0], segments[1])
	}
	milieu := len(segments) / 2
	var gauche []int
	var droite []int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		gauche = FusionnerSegments(segments[:milieu])
	}()
	go func() {
		defer wg.Done()
		droite = FusionnerSegments(segments[milieu:])
	}()

	wg.Wait()
	return merge(gauche, droite)
}

func TriParallele(liste []int) []int {
	NbCoeurs := runtime.NumCPU()
	if len(liste) < 100*NbCoeurs {
		QuickSortSeq(liste)
		return liste
	}
	var wg sync.WaitGroup
	wg.Add(NbCoeurs)
	taille := len(liste) / NbCoeurs
	segments := make([][]int, NbCoeurs)
	for i := 0; i < NbCoeurs; i++ {
		deb := i * taille
		fin := (i + 1) * taille
		if i == NbCoeurs-1 {
			fin = len(liste)
		}
		segments[i] = liste[deb:fin]
		go func(partie []int) {
			defer wg.Done()
			QuickSortSeq(partie)
		}(segments[i])
	}
	wg.Wait()
	return FusionnerSegments(segments)
}
