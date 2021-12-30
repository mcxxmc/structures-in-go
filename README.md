# structures-in-go
 Implementation of some data structures in golang. e.g. Binary search tree.

 Most of the implemented structures support interface{} as the input type,
which means you can use them to store any customized input. However, 
the output type of the structures is also interface{} or []interface{}, 
 which means you need to manually perform type assertion: e.g.,
    
    b = a.(int)

Currently supported and tested structures include:
1. stack
2. queue
3. linked list
4. binary heap
5. binary search tree
6. red black tree
7. fibonacci heap (in process)


Please always use the safe constructor (e.g., `NewStack()`) to initialize any structure.

Some attributes in the structures are protected from outside access for safety reasons; 
please call corresponding methods when necessary.

All trees use the successor for replacement in deletion.

## The "compare" method 
Many structures require a "compare" function as the input for the safe constructor. 
A "compare" function is in the form of `func(a, b interface) int`.
1. "compare" is the class method for comparing different values in the structure (e.g., node values);
2. it should return 1 if a > b , 0 if a == b, -1 if a < b;
3. the first parameter, `a`, should always be an element from the struct other than user input. 
for example, in a tree, `a` should be the same type as the value of the tree node.
4. the second parameter, `b`, may have variant types. A tricky compare method can 
relax the conditions for Search and Delete; please see examples for details. 

One simple example:

    func compareInt (a, b interface{}) int {
        if a.(int) > b.(int) {
            return 1
        } else if a.(int) == b.(int) {
            return 0
        }
        return -1
    }

A tricky example:

    func trickyCompare (a, b interface{}) int {
        a1 := a.(someDataStructure).Val
        var b1 int
        switch b.(type) {
        case SomeDataStructure:
            b1 = b.(someDataStructure).Val
        case int:
            b1 = b.(int)
        }
        if a1 > b1 {return 1} else if a1 == b1 {return 0} else {return -1)
    }
