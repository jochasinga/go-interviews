Go Interviews
=============
This is a collection of concepts, algorithms, theories and questions common in technical interviews, all implemented and explained in Go.

This will eventually be a collection of wrapped-up key takeaways from various sources, mostly from [Interview Cake](https://www.interviewcake.com) and [Programming Interviews](http://www.programmerinterview.com/)

Chapters
--------
+ [The Art of Cramming](#the-art-of-cramming)
+ [Theories and Algorithms](#theories-and-algorithms)
    + [Time and Space Complexity](#time-and-space-complexity)
    + [Big O Notation](#big-o-notation)
        + [O(1)](#o(1))
        + [O(n)](#o(n))
        + [O(n<sup>2</sup>)](#o(n<sup>2</sup>))
        + [N could be the *actual* input, or the *size* of the input](#n-could-be-the-actual-input-or-the-size-of-the-input)
        + [Drop the constants](#drop-the-constants)
        + [Drop the less significant terms](#drop-the-less-significant-terms)
        + [It's about the "worse case"](#its-about-the-worst-case)
    + [Logarithms](#logarithms)
+ [Data structures](#data-structures)

The Art of Cramming
-------------------
[![young boy cramming](https://upload.wikimedia.org/wikipedia/commons/2/21/Cramming.jpg)](https://upload.wikimedia.org/wikipedia/commons/2/21/Cramming.jpg)

> **cramming** (n.): To study with determination. Often in a parrot-learning fashion and in a short amount of time for a predefined mission such as a test or an exam.   

Are you one of the people who, despite being a good (if not expert) developer, having a hard time getting a job as one? You see many of more novice developers get the job while you sit around with great projects on Github but unemployed? I'm one of them.

Like it or not, many technical interviews, as well as any professional tests in any field for that matters, are designed to filter unqualified candidates out. "Unqualified" isn't equivalent to "no talent". It just means not complying to a certain standard. Yes, **passing a test or an interview is all about complying to a standard and taking good amount of time and motivation to study it**.

Theories and Algorithms
-----------------------
Most interviews will include one or more of these concepts that getting yourself acquianted with is a MUST. Remember that these interviews don't just measure how experienced you're as a software developer or engineer, but also how book-smart you are and how you will likely be able to solve real-world problems by applying them.

### Time and Space Complexity
The study of algorithms and how they grow in terms of time and memory consumptions under increasing arbitrary number of input data. It is the same concept as the [Big O Notation](#big-o-notation).

### Big O Notation
![Big O symbol](https://upload.wikimedia.org/wikipedia/commons/thumb/a/a5/Big-o-approx-logo.png/80px-Big-o-approx-logo.png)    
+ **Remember O as in "Order of"**
+ Used to classify algorithms by how they respond to changes in input size, such as how the processing time of an algorithm changes as the problem size becomes extremely large
+ Saying "time increases (or decreases) on the order of _n_ to the amount of input.

Note that it is mostly all about **functions** and **for loops** and how more complex a function gets with more input.

#### O(1)

```go
func printFirstInt(integers []int) {
        fmt.Print(integers[0])
}
```
**This function runs in O(1) time (or "constant time") relative to its input**. The input list could be 1 item or 1,000 items, but this function would still just require one "step."

#### O(n)

```go
func printAllInt(integers []int) {
        for _, num := range integers {        
                fmt.Println(num)
        }
}
```
**This function runs in O(n) time (or "linear time"), where n is the number of items in the list**. If the list has 10 items, we have to print 10 times. If it has 1,000 items, we have to print 1,000 times.

#### O(n<sup>2</sup>)

```go
func printAllPossibleOrderedPairs(integers []int) {
        for _, first := range integers {        
                for _, second := range integers {
                        fmt.Println(first, second)
                }
        }
}
```

Here we're nesting two loops. If our list has _n_ items, our outer loop runs _n_ times and our inner loop runs nnn times for each iteration of the outer loop, giving us _n_^2 total prints. **Thus this function runs in O(n<sup>2</sup>) time (or "quadratic time")**. If the list has 10 items, we have to print 100 times. If it has 1,000 items, we have to print 1,000,000 times.

![graph showing various types of big O's](http://static1.squarespace.com/static/514fbc34e4b0123f55d5a024/t/54ff5ebce4b0dee69e7d0a5f/1426022077515/)

#### N could be the *actual* input, or the *size* of the input
Both of these functions have O(n) runtime, even though one takes an integer as its input and the other takes a list:

```go
func sayHiNTimes(time int) {
        for i := 0; i <= time; i++ {
                fmt.Println("hi")
        }
}

func printAllItems(items []string) {
        for item := range items {
                fmt.Println(item)
        }
}
```

So sometimes n is an actual number that's an input to our function, and other times nnn is the number of items in an input list (or an input map, or an input object, etc.).

#### Drop the constants
The cool thing about Big O's are they care more about what happens as _n_ **gets arbitrarily large**. As _n_ gets really big, adding 100 or dividing by 2 has a decreasingly significant effect.

```go
func printTwice(nums []int) {
        for num := range nums {
                fmt.Println(num)
        }   
        for num := range nums {
                fmt.Println(num)
        }
}
```
This is O(2*n*), but we just call O(*n*).

```go
func printFirstItemThenFirstHalfThenSayHi100Times(nums []int) {
        fmt.Println(nums[0])

        midIndex := len(nums) / 2
        index := 0

        for index < midIndex {
                fmt.Println(nums[index])
                index++
        }

        for i := 0; i < 100; i++ {
                fmt.Println("hi")
        }
}
```

This is O(1+*n*/2+100), but we just call O(*n*).

#### Drop the less significant terms

```go
func printAllNumsThenAllPairSums(nums []int) {
        fmt.Println("These are the numbers:")
        for num := range nums {
                fmt.Println(num)
        }

        fmt.Println("and these are their sums:")
        for firstNum := range nums {
                for secondNum := range nums {
                        fmt.Println(firstNum + secondNum)
                }
        }
}
```
Here our runtime is O(*n*+*n*<sup>2</sup>), which we just call O(*n*<sup>2</sup>). Even if it was O(*n*<sup>2</sup>/2+100*n*), it would still be O(*n*<sup>2</sup>).

Similarly:
+ O(*n*<sup>3</sup> + 50*n*<sup>2</sup> + 10000) is O(*n*<sup>3</sup>)
+ O((*n*+30) * (*n*+5)) is O(*n*<sup>2</sup>)



#### It's about the "worse case"
When we say this algorithm takes O(*n*) time, we're actually talking about a worst-case scenerio.

```go
func printHello(words []string) {
        for word := range words {
                if strings.Contains(word, "hello") {
                        fmt.Println(word)
                }
        }
}
```

This function could have taken O(1) time if a string containing "hello" is actually the first (or tenth, for what it's worth. See [Drop the constants](#drop-the-constants)) member in the slice. **Worst case is, the culprit is the last member, making this case an O(_n_)**.

#### Space complexity
Apart from time, some time we want to optimize for using less memory instead of (or in addition to) using less time. This is very similar to talking about time complexity. We simply look at the total size (relative to the size of the input) of any new variables we're allocating.

This function takes O(1) space (no new variable allocated)

```go
func sayHiNTimes(n int) {
        for time := range n {
            fmt.Println("hi")
        }
}
```

This function takes O(_n_) space (the size of `hiList` scales with the size of the input):

```go
func listOfHiNTimes(n int) []string {
        hiList := []string
        for time := range n {
                hiList = append(hiList, "hi")
        }
        return hiList
}
```

### Logarithms
**"What power must we raise this base to, to get this answer?"** That's all a logarithm is asking. So if we say:     
> log<sub>10</sub>100          

The 10 is called the *base*. Think of the 100 as the "answer." We are taking the log of 100. This pronounced "log base 10 of 100" OR

**What power do we need to raise the base (10) to to get the answer (100)?**

10<sup>*x*</sup> = 100

What *x* gets us our result of 100? 2!

10<sup>2</sup> = 100

Thus,

log<sub>10</sub>100 = 2

#### What logs are *used* for
**Solving for _x_ when _x_ is in an exponent.**    

If we wanted to solve this:    

10<sup>*x*</sup> = 100

We take the log<sub>10</sub> of both sides.

log<sub>10</sub>10<sup>*x*</sup> = log<sub>10</sub>100

What power to raise base (10) to get 10<sup>*x*</sup>? Duh, it's power of *x*! Thus

*x* = log<sub>10</sub>100    
*x* = 2

#### Where logs come up in algorithms and interviews
**"How many times must we double 1 before we get to _n_?"** OR    
**"How many times must we divide _n_ in halfn order to get back down to 1?"**    

The answer to both questions is log<sub>2</sub>*n*!

#### Logs in [binary search](#binary-search)
This comes up in the time cost of **binary search**, which is an algorithm for finding a target number in a sorted array. The process goes like this:
1. **Start with the middle number: is it bigger or smaller than the target?** Since the list is sorted, this tells us if the target would be in the _left_ half or the _right_ half.
2. **We've effectively divided the problem in half.** We can "rule out" the whole half of the list.
3. **Repeat the same approach iteratively on the new half-size problem.** Keep doing this until we find the number that matches the target or "rule out" the whole data set.

```go
func binarySearch(target int, nums []int) bool {
        // see if target appears in nums

        // we think of floorIndex and ceilingIndex as "walls" around
        // the possible positions of our target so by -1 below we mean
        // to start our wall "to the left" of the 0th index
        // (we /don't/ mean "the last index")
        floorIndex := -1
        ceilingIndex := len(nums)

        // if there isn't at least 1 index between floor and ceiling,
        // we've run out of guesses and the number must not be present
        // this reads "while there is at least a member in the list"
        for floorIndex + 1 < ceilingIndex {
                // find the index halfway between the floor and ceiling
                // we use integer division, so we'll never get a "half index"
                distance := ceilingIndex - floorIndex
                halfDist int := distance / 2
                guessIndex := floorIndex + halfDist

                guessValue := nums[guessIndex]

                switch {
                case guessValue == target:
                return true
                case guessValue > target:

                // target is to the left
                // move ceiling to the left half
                ceilingIndex = guessIndex
                default:
                // target is to the right
                // move floor to the right half
                floorIndex = guessIndex
                }
        }
        return false
}
```

So what's the time cost of binary search? The only non-constant part of our time cost is the number of times our while loop runs. Each step of our while loop cuts the range (dictated by `floorIndex` and `ceilingIndex`) in half, until our range has just one element left.

**So the question is, "how many times must we divide our original list size (nnn) in half until we get down to 1?**

> n * 1/2 * 1/2 * 1/2 * 1/2 * ... = 1

How many _1/2_'s are there? We don't know yet, but let's call it _x_:

n * (1/2)<sup>x</sup> = 1

Now we solve for _x_:

n * 1<sup>x</sup>/2<sup>x</sup> = 1
n * 1/2<sup>x</sup> = 1
n/2<sup>x</sup> = 1
n = 2<sup>x</sup>

Now to get the _x_ out of the exponent, take the log<sub>2</sub> of both sides:

log<sub>2</sub>n = log<sub>2</sub>2<sup>x</sup>

The right hand side asks, "what power must we raise 2 to, to get 2<sup>x</sup>?" Well that's _x_!

log<sub>2</sub>n = x

There it is. The total time cost of binary search is O(log<sub>2</sub>*n*).

### Binary Search
![Binary search from Wikipedia](https://upload.wikimedia.org/wikipedia/commons/f/f7/Binary_search_into_array.png)

Binary search always results in O(log *n*) time.

See [Logs in binary search](#logs-in-binary-search)

Data structures
---------------
Some data structures we should know of.

### Linked Lists
A linked list node contains a value or data and a pointer (or pointers) to the next node(s).
![singly linked list](http://www.cs.usfca.edu/~srollins/courses/cs112-f08/web/notes/linkedlists/ll2.gif)
##### Singly linked list
Often, our only connection to the list itself is a variable pointing to the head. From there we can walk down the list to all the other nodes.

```go
type LinkedListNode struct {
	    Value int
	    Next  *LinkedListNode
}
```

How you create a linked list:

```go
a := &LinkedListNode{Value:5}
b := &LinkedListNode{Value:9}
c := &LinkedListNode{Value:1}
a.Next = b
b.Next = c
```
**Advantages**
1. Linked lists have **constant-time O(1) insertions and deletions in any position** (you just change some pointers). Arrays require O(*n*) time to do the same thing, because you'd have to "shift" all the subsequent items over 1 index.
2. Linked lists **can continue to expand as long as there is space on the machine**. Arrays (in low-level languages) must have their size specified ahead of time. Even in languages with "dynamic arrays" that automatically resize themselves when they run out of space (like Python, Ruby and JavaScript), the operation to resize a dynamic array has a large cost which can make a single insertion unexpectedly expensive.

**Disadvantages**
1. To access or edit an item in a linked list, you have to take O(*i*) time to walk from the head of the list to the *i*th item (unless of course you already have a pointer directly to that item). Arrays have constant-time lookups and edits to the *i*th item.

##### Doubly linked list
![Doubly linked list](http://www.cs.usfca.edu/~srollins/courses/cs112-f08/web/notes/linkedlists/ll5.gif)

```go
type LinkedListNode struct {
        Value    int
        Next     *LinkedListNode
        Previous *LinkedListNode
}
```
And how to create one

```go
a := &LinkedListNode{Value: 5}
b := &LinkedListNode{Value: 1}
c := &LinkedListNode{Value: 9}

a.Next = b
b.Previous = a
b.Next = c
c.Previous = b
```

Doubly linked lists allow us to traverse our list backwards. In a singly linked list, if you just had a pointer to a node in the middle of a list, there would be no way to know what its previous node was. Not a problem in a doubly linked list.
