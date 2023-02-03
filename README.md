# KATA
A code kata is an exercise in programming which helps programmers hone their skills through practice and repetition.

## FizzBuzz
FizzBuzz is a very simple programming task, used in software developer job interviews, to determine whether the job candidate can actually write code. It was invented by Imran Ghory, and popularized by Jeff Atwood. Here is a description of the task:
```ignorelang
    Write a program that prints the numbers from 1 to 100.
    But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz”.
    For numbers which are multiples of both three and five print “FizzBuzz”.
```
It’s very well known in software development circles. There are multiple implementations in every language, joke implementations, and plenty of articles discussing its usefulness during the hiring process.
  
I present here my implementation of this task.  

## Ref or Val?
Sometimes I see a developer using pointers in the resulting parameters of a function to indicate no result. It is possible that in some cases this option is justified, but most often it adds overhead to calling this function. The limit of a developer's dirty fantasies is to wrap a slice in a separate data structure, and a reference to it is then used as the resulting parameter.
  
I saw this example in a real project and immediately wrote a code that proves the unreasonableness of this logic.
Pay attention to the benchmarks of this project.

## Viewers counter
I was offered this task at an interview in a well-known company, as a task for live programming in real time.
Of course, at that moment I was confused and did not find a solution right away, but nevertheless I did it and got an offer.
I didn’t get a job with this company, I chose another one that doesn’t scare me with live programming during an interview :).
However, I liked the task for its complexity and simplicity at the same time.

I present two solutions to this problem. Both solutions are correct but one of them leads to data copying, and the other saves memory but changes the original data so that it can no longer be used.

## Binary search
Binary search algorithm written in different ways.
### find1
It seemed to me that this is the simplest way of implementation that immediately comes to mind.
I split the array into two parts in each iteration until the desired number appears in the middle of the array.

A certain problem when reading the code is created by several consecutive conditional operators with other conditional operators nested inside.
### find2
This is a sophisticated version of the first algorithm.
The division of the array is performed taking into account the values of the boundary elements of the array.
The algorithm can achieve very high performance if the values in the array are evenly distributed.

As in the first case, it is quite difficult to perceive the code here, and there is also more mathematics, which also does not facilitate the algorithm.
### find3
This code is much easier to read thanks to the recursion.
I am using tail recursion to avoid wasting resources.

Despite the fact that the logic remains the same, this code is much easier to read.
The only remaining difficulty is `mid + 1` math.
### find4
This function is resource intensive compared to the others, and in addition, it is even more difficult to read, since it contains iteration and recursion and complicated mathematics.
Dividing an array into blocks performs a Fibonacci series.
The progression is rather weak, so a full traversal of the array takes relatively longer than in other cases.
In fact, this algorithm is not competitive and cannot be used.

## Is Contains a Pair
Given: an array of integers from 0 to 1000, the length of the array is no more than 500 values.
You need to bypass the array and find a pair of numbers in which one number is twice as big as the other.

```
i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]
2 <= arr.length <= 500
0 <= arr[i] <= 10^3
```
I confess that in the interview I performed the worst implementation possible.
In a stressful situation, I tend to make wrong decisions and didn't take into account that we have small values and a small array size.
In this case, it is better to operate with `slice` rather than `map`. Here are the results of the benchmarks:

```
cpu: Intel(R) Core(TM) i7-9700F CPU @ 3.00GHz
Benchmark_byArraySlow
Benchmark_byArraySlow-8   	 2718829	       613.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_byMap
Benchmark_byMap-8         	   84037	     20058 ns/op	   10928 B/op	       3 allocs/op
Benchmark_byIntArr
Benchmark_byIntArr-8      	10659265	       113.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_byMap2
Benchmark_byMap2-8        	  535474	      7657 ns/op	   10904 B/op	       2 allocs/op
```
I will now try to figure out what happens to the variables in this code:  
`go tool compile -m main.go`

```
main.go:11:6: can inline byArraySlow
main.go:23:11: inlining call to sort.Ints
main.go:11:18: arr does not escape
main.go:22:12: leaking param: arr
main.go:23:11: sort.Interface(sort.IntSlice(sort.x)) escapes to heap
main.go:24:14: make(map[int]struct {}, len(arr)) does not escape
main.go:37:13: arr does not escape
main.go:38:15: make(map[int]struct {}, len(arr)) does not escape
main.go:53:15: arr does not escape
```
 
Do you see that on line 23, the data escape into a heap? That was my mistake.
The implementation with an array of 1001 elements showed better results, which is not surprising at all,
because its cyclomatic complexity is O(n) and there are no allocations and all data are on the stack.

## Optimization
Everyone probably already knows that if you know the boundaries of an array (slice) beforehand, it is very effective to specify the necessary amount of elements when declaring it.
However, not every time we apply this, for any other buffered data structures, in my example bytes.Buffer declared without
a pre-created memory buffer, the first time we try to write to it, causes an allocation in memory in the procedure bytes.(*Buffer).grow().

Note the effect on performance of creating a variable with at least a minimum initial buffer. This is probably because the bytes.NewBuffer function call was compiled as inline, so the buffer does not escape from its stack frame. This allows the runtime to allocate the necessary amount of memory on the stack in advance. And it does not cause allocation in the heap.
```
cpu: Intel(R) Core(TM) i7-9700F CPU @ 3.00GHz
BenchmarkLength1
BenchmarkLength1-8   	41940806	        61.58 ns/op	      64 B/op	       1 allocs/op
BenchmarkLength2
BenchmarkLength2-8   	115965088	        10.21 ns/op	       0 B/op	       0 allocs/op
```
