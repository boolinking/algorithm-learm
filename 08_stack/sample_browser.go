package _8_stack

type SampleBrowser struct {
	backStack *LinkedStack
	forwardStack  *LinkedStack
}

func NewSampleBrowser() *SampleBrowser {
	return &SampleBrowser{
		NewLinkedStack(),
		NewLinkedStack(),
	}
}

func (br *SampleBrowser) open(url string)  {
	br.backStack.Push(url)
	br.forwardStack.Clear()
}

func (br *SampleBrowser) GetBack()  interface{} {
	url := br.backStack.Pop()
	if url != nil {
		br.forwardStack.Push(url)
	}
	return url
}

func (br *SampleBrowser) GetForward()  interface{} {
	url := br.forwardStack.Pop()
	if url != nil{
		br.backStack.Push(url)

	}
	return url
}