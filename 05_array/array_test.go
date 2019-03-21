package _5_array

import "testing"

func TestArray_Insert(t *testing.T) {
	capacity := 10
	arr := New(uint(capacity))
	for i:=0; i < capacity -1 ;i++  {
		err := arr.Insert(uint(i),i+1)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	arr.Print()
	arr.Insert(uint(6),3)
	arr.Print()

}

func TestArray_Delete(t *testing.T) {
	capacity := 10
	arr := New(10)
	for i:=0; i < capacity -3 ;i++  {
		err := arr.Insert(uint(i),i+1)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	arr.Print()
	_, err := arr.Delete(5)
	if err != nil {
		t.Fatal(err.Error())
	}
	arr.Print()
}

func TestArray_Find(t *testing.T) {
	capacity := 10
	arr := New(10)
	for i:=0; i < capacity -3 ;i++  {
		err := arr.Insert(uint(i),i+1)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	arr.Print()
	t.Log(arr.Find(0))
	t.Log(arr.Find(1))
	t.Log(arr.Find(2))
}