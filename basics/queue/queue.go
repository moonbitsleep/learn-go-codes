package queue

//用切片定义一个队列
type Queue []int

// q发生改变，采用指针接收者
func (q *Queue) Push(v int)  {
	*q = append(*q, v)
}

func (q *Queue) IsEmpty() bool{
	return len(*q) == 0
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

