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

func (mh *MinHeap) swap(i, j int){
	(*mh).arr[i], (*mh).arr[j] = (*mh).arr[j], (*mh).arr[i]
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
	if index == 0{
		return
	}

	current := (*arr)[index]

	parentIdx := (index - 1) / 2
	parent := (*arr)[parentIdx]

	if parent.freq > current.freq {;
		(*arr)[parentIdx] = current
		(*arr)[index] = parent
		heapify_up(arr, parentIdx)
	}
}

// replace root or element which needs to be deleted with the last element
func (mh *MinHeap) remove_by_index(index int) (Node, bool){
	lenght := len((*mh).arr)
	if index > len(mh.arr) || index < 0{
		return Node{}, false
	}

	mh.swap(0, lenght - 1)

	result := (*mh).arr[lenght - 1]
	(*mh).arr = (*mh).arr[:lenght - 1]

	mh.heapify_down(0)

	return result, true
}

func (mh *MinHeap) heapify_down(index int) {
	lIdx := index * 2 + 1
	rIdx := index * 2 + 2

	if index >= len((*mh).arr) -1 || lIdx >= len((*mh).arr) - 1 {
		return
	}

	// WE NEED TO REPLACE WITH THE SMALLEST CHILD

	c := (*mh).arr[index]
	l := (*mh).arr[lIdx]
	r := (*mh).arr[rIdx]

	if l.freq > r.freq && c.freq > r.freq {
		mh.swap(rIdx, index)
		mh.heapify_down(rIdx)
	}
	if r.freq > l.freq && c.freq > l.freq {
		mh.swap(lIdx, index)
		mh.heapify_down(lIdx)
	}
	if l.freq == r.freq && c.freq > r.freq{
		mh.swap(lIdx, index)
		mh.heapify_down(lIdx)
	}
}

func huffman_encoding_shenanigans(minheap *MinHeap){
	for len(minheap.arr) > 1{
		left_node, _ := minheap.remove_by_index(0)
		fmt.Printf("LN -> %d %c\n", left_node.freq, left_node.char)
		right_node, _ := minheap.remove_by_index(0)
		fmt.Printf("RN -> %d %c\n", right_node.freq, right_node.char)
		new_node := Node{freq: left_node.freq + right_node.freq, left: &left_node, right: &right_node}
		minheap.insert(new_node)
		fmt.Printf("NEW N -> %d %c\n", new_node.freq, new_node.char)
		fmt.Printf("After insert ")
		show_minheap(&minheap.arr)
		fmt.Println()
	}
}

func traverse_huffman_tree(node *Node, arr []int){
	if node.left != nil {
		arr = append(arr, 0)
		traverse_huffman_tree(node.left, arr)
	}

	if node.right != nil {
		arr = append(arr, 1)
		traverse_huffman_tree(node.right, arr)
	}
	if node.left == nil && node.right == nil {
		fmt.Printf("%v %c\n", arr, node.char)
	}
}

func bfs(node *Node){
	counter := 0
	arr := []*Node{}
	arr = append(arr, node)
	for len(arr) > 0{
		curr := arr[0]
		arr = arr[1:]
		if counter == 3 {
			fmt.Println()
			fmt.Println()
			counter = 0
		}
		char := curr.char
		if char == 0 {
			char = '*'
		}
		fmt.Printf("%d %c ", curr.freq, char)
		counter += 1
		if curr.left != nil{
			arr = append(arr, curr.left)
		}

		if curr.right != nil{
			arr = append(arr, curr.right)
		}
	}
}

func show_minheap(arr *[]Node){
	for i:= range (*arr){
		fmt.Printf("%d ", (*arr)[i].freq)
	}
	fmt.Print("\n")
}

func main(){
	hashmap := make(map[byte]int)
	test := "loremipsumdolorsitamet"
	for i := range test{ 
		hashmap[test[i]] += 1
	}

	minHeap := MinHeap{}
	for key, value := range hashmap{
		minHeap.insert(Node{freq: value, char: key})
	}

	show_minheap(&minHeap.arr)
	huffman_encoding_shenanigans(&minHeap)
	fmt.Printf("Top of huffman tree: %v\n", minHeap.arr[0])
	// bfs(&minHeap.arr[0])
	traverse_huffman_tree(&minHeap.arr[0], []int{})

	// create a huffman tree from the string
	// create an encoded string

	// send tree
	// send encoded value

	// decode value by using the tree
}


