package main

import (
	"fmt"
)

type Node struct {
	freq int
	char byte
	left *Node
	right *Node
}

type MinHeap struct{
	arr []Node
}

func (mh *MinHeap) insert(node Node){
	if len(mh.arr) == 0{
		mh.arr = append(mh.arr, node)
		return
	}
	mh.arr = append(mh.arr, node)
	heapify_up(&mh.arr, len(mh.arr) - 1)
}

// left child (i * 2) + 1
// right child (i * 2) + 2
// parent (i - 1) / 2

func heapify_up(arr *[]Node, index int){
	if index == 0 || (*arr)[index].freq >= (*arr)[(index - 1) / 2].freq{
		return
	}

	current := (*arr)[index]

	parentIdx := (index - 1) / 2
	parentValue := (*arr)[parentIdx]

	(*arr)[index] = parentValue

	(*arr)[parentIdx] = current

	heapify_up(arr, parentIdx)
}

// replace root or element which needs to be deleted with the last element
func (mh *MinHeap) remove_by_index(index int) (Node, bool){
	if index > len(mh.arr) || index < 0{
		return Node{}, false
	}

	result := mh.arr[index]

	// set last item to top
	mh.arr[index] = mh.arr[len(mh.arr) - 1]
	if len(mh.arr) > 0{
		mh.arr = mh.arr[:len(mh.arr) - 1]
		heapify_down(&mh.arr, index)
	}

	return result, true
}

func heapify_down(arr *[]Node, index int){
	left_child_index := (index * 2) + 1
	right_child_index := (index * 2) + 2
	lenght := len(*arr)


	index_for_swapping := index

	if left_child_index < lenght && (*arr)[left_child_index].freq < (*arr)[index].freq{
		index_for_swapping = left_child_index
	}

	if right_child_index < lenght && (*arr)[right_child_index].freq < (*arr)[index].freq{
		index_for_swapping = right_child_index
	}
	
	if index_for_swapping != index{
		current_node := (*arr)[index]
		(*arr)[index] = (*arr)[index_for_swapping]
		(*arr)[index_for_swapping] = current_node
		heapify_down(arr, index_for_swapping)
	}
}

func huffman_encoding_shenanigans(minheap *MinHeap){
	for len(minheap.arr) > 1{
		left_node, _ := minheap.remove_by_index(0)
		right_node, _ := minheap.remove_by_index(0)
		if left_node.freq + right_node.freq == 6 || left_node.freq + right_node.freq == 5{
			fmt.Printf("%v\n", left_node)
			fmt.Printf("%v\n", right_node)
		} 
		new_node := Node{freq: left_node.freq + right_node.freq, left: &left_node, right: &right_node}
		minheap.insert(new_node)
	}
}

func traverse_huffman_tree(node *Node){
	arr := []*Node{node}

	for len(arr) > 0 {
		current_node := arr[0]
		fmt.Printf("%c %d\n", current_node.char, current_node.freq)
		arr = arr[1:]
		if current_node.left != nil{
			arr = append(arr, current_node.left)
		}
		if current_node.right != nil{
			arr = append(arr, current_node.right)
		}
	}
}

func main(){
	hashmap := make(map[byte]int)
	test := "balls"
	for i := range test{
		hashmap[test[i]] += 1
	}

	minHeap := MinHeap{}
	for key, value := range hashmap{
		minHeap.insert(Node{freq: value, char: key})
	}

	huffman_encoding_shenanigans(&minHeap)
	fmt.Printf("Top of huffman tree: %v\n", minHeap.arr[0])
	traverse_huffman_tree(&minHeap.arr[0])
}


