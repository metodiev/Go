package main

type Interface interface{
    //Find number of elements in collection
    Len() int

    //Less method is used for identify which elements amont index i and j are lesser and is used for sorting
    Less(i, j int)bool

    //Swap method is used for swapping elemnts with indexes i and j
    Swap(i, j int)
}

type Human struct {
    name string
    age int
}
