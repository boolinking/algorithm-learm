package _5_array

import "errors"

type Silce struct {
	data []interface{}
	size uint
}

func NewSilce() *Silce {
	return &Silce{
		data:make([]interface{},10), //默认大小
		size: 0,
	}
}

//当前silce的容量
func (s *Silce) Cap() uint  {
	return uint(cap(s.data))
}

//当前Silce包含元素的个数
func (s *Silce) Size() uint  {
	return s.size
}

func (s *Silce) IsEmpty() bool  {
	if s.Size() == 0 {
		return true
	}
	return false
}

func (s *Silce) Set(index uint, data interface{}) error {
	err := s.checkIndex(index)
	if err != nil {
		return err
	}
	s.data[index] = data
	return nil
}

func (s *Silce) Add(data interface{})  {
	if s.Size() == s.Cap() {
		s = reSize(s.Size() * 2,s)
	}
	 i := s.size + 1
	s.data[i] = data
}

func (s *Silce) Remove(index uint) (interface{},error) {
	err := s.checkIndex(index)
	if err != nil {
		return nil , err
	}
	v := s.data[index]
	for i:= index; i < s.Size()-1; i++  {
		s.data[i] = s.data[i + 1]
	}
	s.data[s.size] = nil
	s.size--
	return v,nil
}



func reSize(capacity uint, s *Silce) *Silce {
	newSilce := Silce{
		data:make([]interface{},capacity),
		size:0,
	}
	for i:= uint(0); i < s.Size() ; i++ {
		newSilce.data[i] = s.data[i]
	}
	return &newSilce
}


func (s Silce) checkIndex(index uint) error {
	if index > s.Size() {
		errors.New("index out of range")
	}
	return nil
}




