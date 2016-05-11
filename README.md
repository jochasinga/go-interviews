Go Interviews
=============
This is a collection of concepts, algorithms, theories and questions common in technical interviews, all implemented and explained in Go.

This will eventually be a collection of wrapped-up key takeaways from various sources. A few to check out:
+ [Interview Cake](https://www.interviewcake.com)
+ [Programming Interviews](http://www.programmerinterview.com/)

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

The Art of Cramming
-------------------
![young boy cramming](https://upload.wikimedia.org/wikipedia/commons/2/21/Cramming.jpg)

Above: A young boy does some last minute cramming while walking to school with his sister. [Wikipedia](https://en.wikipedia.org/wiki/Cramming_%28education%29).
> **cramming** (n.): To study with determination. Often in a parrot-learning fashion and in a short amount of time for a predefined mission such as a test or an exam.   

Are you one of the people who, despite being a good (if not expert) developer, having a hard time getting a job as one? You see many of more novice developers get the job while you sit around with great projects on Github but unemployed? I'm one of them.

Like it or not, many technical interviews, as well as any professional tests in any field for that matters, are designed to filter unqualified candidates out. "Unqualified" isn't equivalent to "no talent". It just means not complying to a certain standard. Yes, **passing a test or an interview is all about complying to a standard and taking good amount of time and motivation to study it**.

Many really skillful, experienced software engineers have failed interviews. Read the crazily popular [blog post](https://medium.com/@evnowandforever/f-you-i-quit-hiring-is-broken-bb8f3a48d324#.yo13elu7f) by [Sahat Yalkabov](https://github.com/sahat) and this [tweet posted on Hacker News](https://news.ycombinator.com/item?id=9695102) by Max Howell. Sahat is an open source elite on Github with 900+ days streak and over 13,500 stars on a [single project](https://github.com/sahat/hackathon-starter), while Max is the creator of [Homebrew](http://brew.sh/).

Theories and Algorithms
-----------------------
Most interviews will include one or more of these concepts that getting yourself acquianted with is a MUST. Remember that these interviews don't just measure how experienced you're as a software developer or engineer, but also how book-smart you are and how you will likely be able to solve real-world problems by applying them.

### Time and Space Complexity
`big o`, `algorithm`, `time`, `resource`, `process`    
This hass nothing with the time and space continuum. It is simply the study of algorithms and how they grow in terms of time and space (computer resource) consumptions under increasing arbitrary number of input data. It is related to the [Big O Notation](#big-o-notation).

### Big O Notation
![Big O symbol](https://upload.wikimedia.org/wikipedia/commons/thumb/a/a5/Big-o-approx-logo.png/80px-Big-o-approx-logo.png)    
`function`, `loop`, `time`, `resource`, `order of`, `scalability`, `worst case`, `best case`
**Remember O as in "Order of"**. Big O is used to classify algorithms by how they respond to changes in input size, such as how the processing time of an algorithm changes as the problem size becomes extremely large, or simply saying "time increases (or decreases) on the order of `n` to the amount of input. Now you see it is not a just boring theory created by some totally [wackly-genius academician](https://en.wikipedia.org/wiki/Edmund_Landau). It can be applied directly to the problems in scalabilities.
Note that it is mostly all about **functions** and **for loops**.

#### O(1)
`once`, `constant time`   
O(1) is simple. How much input you shove into a function, it will always operate just once and take the same amount of time.

```go
func printFirstInt(integers []int) {
        fmt.Print(integers[0])
}
```
**This function runs in O(1) time (or "constant time") relative to its input**. The input list could be 1 item or 1,000 items, but this function would still just require one "step."

#### O(n)
`linear time`  

```go
func printAllInt(integers []int) {
        for _, num := range integers {        
                fmt.Println(num)
        }
}
```
**This function runs in O(n) time (or "linear time"), where n is the number of items in the list**. If the list has 10 items, we have to print 10 times. If it has 1,000 items, we have to print 1,000 times.

#### O(n<sup>2</sup>)
`quadratic time`, `power of 2`  

```go
func printAllPossibleOrderedPairs(integers []int) {
        for _, first := range integers {        
                for _, second := range integers {
                        fmt.Println(first, second)
                }
        }
}
```

Here we're nesting two loops. If our list has nnn items, our outer loop runs nnn times and our inner loop runs nnn times for each iteration of the outer loop, giving us n^2 total prints. **Thus this function runs in O(n<sup>2</sup>) time (or "quadratic time")**. If the list has 10 items, we have to print 100 times. If it has 1,000 items, we have to print 1,000,000 times.

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
The cool thing about Big O's are they care more about what happens as n **gets arbitrarily large**. As n gets really big, adding 100 or dividing by 2 has a decreasingly significant effect. (Yeah!)

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
This is O(2n), but we just call O(n).

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

This is O(1+n/2+100), but we just call O(n).

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
Here our runtime is O(n+n<sup>2</sup>), which we just call O(n<sup>2</sup>). Even if it was O(n<sup>2</sup>/2+100n), it would still be O(n<sup>2</sup>).

Similarly:
+ O(n<sup>3</sup> + 50n<sup>2</sup> + 10000) is O(n<sup>3</sup>)
+ O((n+30) * (n+5)) is O(n<sup>2</sup>)



#### It's about the "worse case"
When we say this algorithm takes O(n) time, we're actually talking about a worst-case scenerio.

```go
func printHello(words []string) {
        for word := range words {
                if strings.Contains(word, "hello") {
                        fmt.Println(word)
                }
        }
}
```

This function could have taken O(1) time if a string containing "hello" is actually the first (or tenth, for what it's worth. See [Drop the constants](#drop-the-constants)) member in the slice. **Worst case is, the culprit is the last member, making this case an O(n)**.

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

This function takes O(n) space (the size of `hi_list` scales with the size of the input):

```go
func listOfHiNTimes(n int) []string {
        hiList := []string
        for time := range n {
                hiList = append(hiList, "hi")
        }
        return hiList
}
```
