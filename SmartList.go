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
func (SList *SList) Put(elements ...interface{}) {
    SList.elements = append(SList.elements, elements...)
}

// returns element from a list
func (SList *SList) Get(index int) (interface{}){
    return SList.elements[index]
}

// replaces element in a list
func (SList *SList) Sit(index int, element interface{}) {
    SList.elements[index] = element
}

// clears a list out
func (SList *SList) Clr() {
    SList.elements = []interface{}{}
    SList.to_remove = []int{}
}

// returns list's length
func (SList *SList) Len() (int) {
    return len(SList.elements)
}

// returns all the list's elements
func (SList *SList) ToList() ([]interface{}) {
    return SList.elements
}

// adds given indexes to removal queue from the list during execution
func (SList *SList) Remove(indexes ...int) {
    SList.to_remove = append(SList.to_remove, indexes...)
}

// executes a removal of indexes specified using Remove()
func (SList *SList) Execute() {

    // if list's removal queue len is equal to 0 - do nothing
    if ( len(SList.to_remove) == 0 ) {
        return
    }

    // turning index removal queue into a set of indexes
    SList.to_remove = Normalize(SList.to_remove)

    // remove elements specified
    for x := 0; x < len(SList.to_remove); x++ {
        rm_index := SList.to_remove[x]

        // checking if out of bounds
        if ( rm_index > len(SList.elements)-1 ) {
            fmt.Println("Index", rm_index, "out of bounds, omitting")
            continue
        }

        // removing element
        SList.elements = append(SList.elements[:rm_index], SList.elements[rm_index+1:]...)

        // decreasing values of the remaining indexes
        // to fit SList.elements' indexes after an item removal
        for y := x; y < len(SList.to_remove); y++ {
            if ( SList.to_remove[y] > rm_index ) {
                SList.to_remove[y]--
            }
        }
    }

    // clearing out remove list
    SList.to_remove = []int{}
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
