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



func TriParallele(liste[]) []int{
	NbCoeurs := runtime.NumCPU()
	if len(liste)<100*NbCoeurs{
		QuickSortSeq(liste)
		return liste
	}
	var wg sync.Waitgroup{}
	wg.Add(NbCoeurs)
	taille := len(liste)/NbCoeurs
	segments := make([][]int,NbCoeurs)
	for i:=0;i<NbCoeurs;i++{
		deb := i*taille
		fin := (i+1)*taille
		if i == NbCoeurs-1{
			fin = len(liste)
		}
		segments[i]=liste[deb:fin]
		go func(partie []int){
			defer wg.Done()
			QuickSortSeq(partie)
		}(segment[i])
	}
	wg.Wait()
	return FusionnerSegments(segments)
}
