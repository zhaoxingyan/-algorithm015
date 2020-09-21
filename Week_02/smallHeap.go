package Week_02

type Heap struct {
	a        []int
	capacity int
	count    int
}

func NewHeap(capacity int) *Heap{
	a := make([]int, capacity+1)
	return &Heap{a, capacity, 0}
}

func (this *Heap) Insert(data int)  {
	if this.capacity == this.count {
		return
	}
	this.count++
	this.a[this.count] = data
	i := this.count
	for i/2 > 0 && this.a[i] < this.a[i/2] {
		swap(this.a, i, i/2)
		i = i / 2
	}
}

func (this *Heap) removeSmall() {
	if this.count <= 0 {
		return
	}
	this.a[1] = this.a[this.count]
	this.count--
	//向下做堆化
	i := 1
	for {
		minPos := i
		if i*2 <= this.count && this.a[i] > this.a[i*2] { minPos = i*2}
		if i*2+1 <= this.count && this.a[minPos] > this.a[i*2+1] {minPos = i*2+1}
		if minPos == i {
			break
		}
		swap(this.a, i, minPos)
		i = minPos
	}
}

func swap(a []int, i int, j int)  {
	a[i], a[j] = a[j], a[i]
}
