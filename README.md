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