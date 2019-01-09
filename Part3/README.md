# Reasons for concurrency and parallelism


To complete this exercise you will have to use git. Create one or several commits that adds answers to the following questions and push it to your groups repository to complete the task.

When answering the questions, remember to use all the resources at your disposal. Asking the internet isn't a form of "cheating", it's a way of learning.

 ### What is concurrency? What is parallelism? What's the difference?
 > Concurrency and parallelism are similar concepts which can both be explained as "running multiple things _at the same time_". The key difference between the two are that concurrent execution need not be executed _physically_ at the same time, i.e in different processor cores. Parallel execution on the other hand is exactly that, and occurs for instance when running calculations on two or more processor cores, or two or more workers machines in a distributed system.
 
 ### Why have machines become increasingly multicore in the past decade?
 > Historically, CPUs could be speed up by decreasing the size of the chip to minimize distances within the chip. But in the past decade, the performance gain from this approach has become less and less substantial. As a result, processor manufacturers have seeked other ways to improve performance. One of which is the introduction of multiple cores, which allowes execution to run in parallel, thereby speeding up calculations even though the actual processing units might not be faster. 
 
 ### What kinds of problems motivates the need for concurrent execution?
 (Or phrased differently: What problems do concurrency help in solving?)
 > - Any problem that requires doing multiple things simultaneously.
 >   - In this case concurrent implementation would likely be more readable than non-concurrent counterparts.
 > - Calculations that can benefit from being spread over multiple processor cores, if available
 
 ### Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
 (Come back to this after you have worked on part 4 of this exercise)
 > Concurrency makes implementing things that should run at the same time easier, but they introduce potential for bugs that are very hard to spot and usually impossible to notice at compile time. Problems like these are for instance deadlocks and race conditions. Deadlocks occur when two threads are waiting for each other to continue execution. Race conditions are situations where the result of a concurrent computation depends on the order of execution for the individual threads. 
 
 ### What are the differences between processes, threads, green threads, and coroutines?
 > Processes are the routines running as part of a program. A computer system typically has multiple processes running at the same time, and these are time shared on the available processors. Each process can in turn spawn multiple threads. Threads are routines running concurrently as part of a process. These can share memory through messaging or resource sharing. The execution of threads are handled by a _scheduler_, which is a part of the OS. Green threads are not dependent on the OS providing a scheduler, and are instead scheaduled by a runtime library such or a virtual machine such asthe Java VM. Coroutines on the other hand are somewhat different. In the most basic case, these introduce no parallel execution. Coroutines are just multiple subroutines which switch between which one gets to execute cooperatively. Cooperatively in this sense refers to the fact that the coroutines themselves signal when they can be switched from (This is known as cooperative multitasking). This occurs typically when executing a blocking operation, such as waiting for a network result. Coroutines can be used to simplify concurrency. Go's implementation for instance, moves other coroutines to a different thread when one is blocking, to allow the others to keep executing. By letting the coroutines themselves signal when they're blocking, the need for a scheduler disappears. By removing this overhead, we can run hundreds of thousands of coroutines compared to only tens of threads. 
 
 ### Which one of these do `pthread_create()` (C/POSIX), `threading.Thread()` (Python), `go` (Go) create?
 > - `pthread_create()` creates threads
 > - `threading.Thread()` creates threads
 > - `go` created coroutines
 
 ### How does pythons Global Interpreter Lock (GIL) influence the way a python Thread behaves?
 > The GIL is a mutex preventing multiple python threads from accessing the same python object at the same time. This means the threads need to wait for the object to become available before it can be accessed. The need for the GIL is a result of python's memory management not being thread-safe. The GIL effectively limits python bytecode execution to a single core, regardless of how many python threads it is spread out on. 
 
 ### With this in mind: What is the workaround for the GIL (Hint: it's another module)?
 > To circumvent the GIL one can use the `multiprocessing` library instead of threads. This library lets you create processes instead of threads to run your program, "effectively side-stepping the GIL"[\[1\]](https://docs.python.org/3/library/multiprocessing.html )
 ### What does `func GOMAXPROCS(n int) int` change? 
 > The number of processors that can execute the program code simultaneously
