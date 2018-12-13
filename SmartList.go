// SmartList provides an easy way to manage slices of vars of any type.
package SmartList

import (
    "fmt"
)

type SList struct {
    elements    []interface{}
    to_remove   []int
}

// appends a number of elements to a list
func (sl *SList) Put(elements ...interface{}) {
    sl.elements = append(sl.elements, elements...)
}

// returns element from a list
func (sl *SList) Get(index int) (interface{}){
    return sl.elements[index]
}

// replaces element in a list
func (sl *SList) Sit(index int, element interface{}) {
    sl.elements[index] = element
}

// clears a list out
func (sl *SList) Clr() {
    sl.elements = []interface{}{}
    sl.to_remove = []int{}
}

// returns list's length
func (sl *SList) Len() (int) {
    return len(sl.elements)
}

// returns all the list's elements
func (sl *SList) ToList() ([]interface{}) {
    return sl.elements
}

// adds given indexes to removal queue from the list during execution
func (sl *SList) Remove(indexes ...int) {
    sl.to_remove = append(sl.to_remove, indexes...)
}

// executes a removal of indexes specified using Remove()
func (sl *SList) Execute() {

    // if list's removal queue len is equal to 0 - do nothing
    if ( len(sl.to_remove) == 0 ) {
        return
    }

    // turning index removal queue into a set of indexes
    sl.to_remove = Normalize(sl.to_remove)

    // remove elements specified
    for x := 0; x < len(sl.to_remove); x++ {
        rm_index := sl.to_remove[x]

        // checking if out of bounds
        if ( rm_index > len(sl.elements)-1 ) {
            fmt.Println("Index", rm_index, "out of bounds, omitting")
            continue
        }

        // removing element
        sl.elements = append(sl.elements[:rm_index], sl.elements[rm_index+1:]...)

        // decreasing values of the remaining indexes
        // to fit sl.elements' indexes after an item removal
        for y := x; y < len(sl.to_remove); y++ {
            if ( sl.to_remove[y] > rm_index ) {
                sl.to_remove[y]--
            }
        }
    }

    // clearing out remove list
    sl.to_remove = []int{}
}

// checks if a list of ints contains element specified
func Includes(list []int, search int) (bool) {
    for x := 0; x < len(list); x++ {
        if ( list[x] == search ) {
            return true
        }
    }

    return false
}

// removes duplicates in a list of ints
func Normalize(list []int) ([]int) {
    normalized := []int{}

    for x := 0; x < len(list); x++ {
        value := list[x]
        if ( !Includes(normalized, value) ) {
            normalized = append(normalized, value)
        }
    }

    return normalized
}
