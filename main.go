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

	if parent.freq > current.freq {
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
	lIdx := (index * 2) + 1
	rIdx := (index * 2) + 2

	if len(mh.arr) == 2{
		c := (*mh).arr[index]
		l := (*mh).arr[lIdx]
		if l.freq < c.freq {
			mh.swap(lIdx, index)
		}
	}
	if index > len(mh.arr) - 1 || lIdx >=len((*mh).arr) - 1  {
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
		// fmt.Printf("LN -> %d %c\n", left_node.freq, left_node.char)
		right_node, _ := minheap.remove_by_index(0)
		// fmt.Printf("RN -> %d %c\n", right_node.freq, right_node.char)

		new_node := Node{freq: left_node.freq + right_node.freq, left: &left_node, right: &right_node}
		minheap.insert(new_node)
	}
}

func generate_huffman_table(
	node *Node,
	arr[]int,
	byte_count int,
	b byte,
	huffman_table *map[byte]Huffman_table_item,
){

	// if we are traversing the tree and byte_count > 8
	// we need to add full byte to arr
	// then create a new byte 
	// fill it and add to arr
	// repeat until char is found

	// TODO: take care of byte count
	if node.left == nil && node.right == nil {
		_, ok := (*huffman_table)[node.char]
		if ok {
			// append completed byte to byte arr
			entry := (*huffman_table)[node.char]
			entry.b_arr = append(entry.b_arr, b)
			entry.bits_to_read += byte_count
			(*huffman_table)[node.char] = entry
		} else {
			(*huffman_table)[node.char] = Huffman_table_item{b_arr: []byte{b}, bits_to_read: byte_count}
		}
		// fmt.Printf("Bits: %08b for %c\n",b,node.char)
		b = 0
		byte_count = 0

		return 
	}

	if byte_count > 8{
		// append completed byte to byte arr
		entry := (*huffman_table)[node.char]
		entry.b_arr = append(entry.b_arr, b)
		entry.bits_to_read += byte_count
		(*huffman_table)[node.char] = entry
		b = 0
		byte_count = 0
	}

	byte_count += 1
	// shift left
	b <<= 1

	if node.left != nil {
		new_arr := arr
		new_arr = append(arr, 0)
		generate_huffman_table(node.left, new_arr, byte_count, b, huffman_table)
	}

	if node.right != nil {
		new_arr := arr
		new_arr = append(arr, 1)
		// if right leave then use OR operator and set 1
		b |= 1
		generate_huffman_table(node.right, new_arr, byte_count, b, huffman_table)
	}
}

func show_minheap(arr *[]Node){
	for i:= range (*arr){
		fmt.Printf("%c %d", (*arr)[i].char, (*arr)[i].freq)
	}
	fmt.Print("\n")
}


type Huffman_table_item struct{
	b_arr []byte
	bits_to_read int
}

func main(){
	hashmap := make(map[byte]int)
	test := "big boobas"
	// test := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam luctus varius imperdiet. Phasellus a dui enim. Proin consequat ante eget euismod semper. Maecenas non pellentesque felis."
	for i := range test{ 
		hashmap[test[i]] += 1
	}

	minheap := MinHeap{}
	for key, value := range hashmap{
		minheap.insert(Node{freq: value, char: key})
	}
	huffman_encoding_shenanigans(&minheap)

	// later I need to allow to store byte arr
	// huffman_table := make(map[byte]byte)
	huffman_table := make(map[byte]Huffman_table_item)
	test_arr := []int{}
	generate_huffman_table(&minheap.arr[0], test_arr, 0, byte(0), &huffman_table)

	// for i := range huffman_table{
	// 	fmt.Printf("%c %d and %v\n", i, huffman_table[i].bits_to_read, huffman_table[i].b_arr)
	// }
	
	// fmt.Printf("From: %08b\n", from)



	// arr_to_move := []int{0b11111001, 0b11111111}
	// index_to_move := 0
	// total_to_move := 13
	//
	// arr_to_fill := []int{0b11000000}
	// index_to_fill := 0
	// to_fill := 4

	byte_arr_to_move := []int{0b01101000}
	byte_index_to_move := 0
	total_bits_to_move := 5

	byte_arr_to_fill := []int{0b11111000}
	byte_index_to_fill := 0
	bits_to_fill := 3

	fmt.Print("To move: ")
	for i:= range byte_arr_to_move{
		fmt.Printf("%08b ", byte_arr_to_move[i])
	}

	fmt.Printf("\nTo fill: ")
	for i:= range byte_arr_to_fill{
		fmt.Printf("%08b ", byte_arr_to_fill[i])
	}
	fmt.Println()

	bits_per_byte_to_move := 0
	for total_bits_to_move > 0{
		fmt.Printf("Total to move: %d\n", total_bits_to_move)
		fmt.Printf("Left %08b\n", byte_arr_to_move[byte_index_to_move])

		if bits_per_byte_to_move == 0{
			// do this only when the whole byte is transfered
			if total_bits_to_move >= 8{
				bits_per_byte_to_move = 8
			} else {
				bits_per_byte_to_move = total_bits_to_move
			}
		}

		if bits_to_fill <= 0{
			// byte was filled
			// time to create a new one
			fmt.Println("new byte created")
			byte_arr_to_fill = append(byte_arr_to_fill, 0b00000000)
			bits_to_fill = 8
			byte_index_to_fill += 1
		}

		if bits_to_fill == 8 && bits_to_fill == bits_per_byte_to_move{
			fmt.Println("to_fill == 8 && to_fill == to_move")
			// 0000.0000 <- 1111.1111
			// to fill 8, to move 8
			byte_arr_to_fill[byte_index_to_fill] |= byte_arr_to_move[byte_index_to_move]
			bits_to_fill = 0

			// important
			total_bits_to_move -= bits_per_byte_to_move
			continue
		}

		if bits_to_fill == 8 && bits_to_fill > bits_per_byte_to_move{
			fmt.Println("to_fill == 8 && to_fill > to_move")
			// 0000.0000 <- 0000.1011
			// to fill 8, to move 4
			byte_arr_to_fill[byte_index_to_fill] |= byte_arr_to_move[byte_index_to_move]
			// shift to the right for new bits
			byte_arr_to_fill[byte_index_to_fill] <<= (bits_to_fill - bits_per_byte_to_move)
			bits_to_fill -= bits_per_byte_to_move

			// important
			total_bits_to_move -= bits_per_byte_to_move
			bits_per_byte_to_move = 0
			// we used all bits, move the next byte

			if byte_index_to_fill + 1 <= len(byte_arr_to_move) - 1{
				fmt.Println("Moving to the next byte")
				byte_index_to_move+=1
			}

			continue
		}

		if bits_to_fill >= bits_per_byte_to_move {
			fmt.Println("to_fill >= to_move")
			// 1111.0000 <- 0000.1011
			// to fill 4, to move 4
			fmt.Printf("max to move: %d, %d, max to fill: %d, %d\n", len(byte_arr_to_move) - 1, byte_index_to_move, len(byte_arr_to_fill) - 1, byte_index_to_fill)
			byte_arr_to_fill[byte_index_to_fill] |= byte_arr_to_move[byte_index_to_move]
			bits_to_fill =- bits_per_byte_to_move
			total_bits_to_move -= bits_per_byte_to_move
			continue
		}

		if bits_to_fill != 0 && bits_to_fill < bits_per_byte_to_move{
			fmt.Println("to_fill != 0 && to_fill < to_move")
			// 1111.0000 <- 1111.1011
			// to fill 4, to move 8
			// 1111.1011 >> 4 = 0000.1111
			selected_bits := byte_arr_to_move[byte_index_to_move] >> (8 - bits_to_fill)
			fmt.Printf("Selected bits: %08b\n", selected_bits)
			byte_arr_to_fill[byte_index_to_fill] |= selected_bits
			// hide used bits by using mask
			// 1111.1011 & 0000.1111 (mask) = 0000.1011
			byte_arr_to_move[byte_index_to_move] &= get_bit_mask(bits_to_fill)
			bits_per_byte_to_move -= bits_to_fill
			total_bits_to_move -= bits_per_byte_to_move

			// if we used all bits, then move the next byte
			if bits_per_byte_to_move <= 0 && len(byte_arr_to_move) > 1{
				fmt.Println("Moving to the next byte")
				byte_index_to_move+=1
			}

			bits_to_fill = 0
			continue

		}
	}

	fmt.Printf("Result: ")
	for i:= range byte_arr_to_fill{
		fmt.Printf("%08b ", byte_arr_to_fill[i])
	}
	fmt.Printf("\n")


	// create a huffman tree from the string
	// create a huffman table 
	// use table for encding a string
	
	// send chars in descending order
	// send freqs in descending order
	// send encoded value

	// generate a tree
	// decode value by using the tree
}

func get_bit_mask(number_of_used_bits int) int{
	// returns mask in styles 0000.1111
	// where 0000 is a number of used bits
	return 255 >> (8 - number_of_used_bits)
}
