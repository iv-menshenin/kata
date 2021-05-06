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

