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


func TriRapideParallele (liste []int, wg *sync.WaitGroup){
	defer wg.Done()
	if len(liste)<2{
		return
	}
	index:=partition(liste)
	if len(liste)>1000{
		subWg := &sync.WaitGroup{}
		subWg.Add(2)
		go TriRapideParallele(liste[:index],subWg)
		go TriRapideParallele(liste[index+1:],subWg)
		subWg.Wait()
	}
	else {
		QuickSortSeq(liste[:index])
		QuickSortSeq(liste[index+1:])
	}

}
