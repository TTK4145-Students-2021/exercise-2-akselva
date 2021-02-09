Exercise 2 - Theory questions
-----------------------------

### What is an atomic operation?
> An atomic operation is an operation that must be performed entirely, it cannot be split up.

### What is a critical section?
> Critical sections are sections of code sharing the same resources

### What is the difference between race conditions and data races?
> A race condition is a situation where the result of an operation is dependent on the order of different individual operations, while a data race happens when different threads try to access a shared variable at the same time

### What are the differences between semaphores, binary semaphores, and mutexes?
> Semaphores are positive integers that may only be incremented, to make a resource available, or decremented, to make a resource unavailable. Binary semaphores can only have the values 0 and 1. Mutexes are similar to semaphores, however only the threads that takes control and makes the mutex unavailable can make the mutex available again.

### What are the differences between channels (in Communicating Sequential Processes, or as used by Go, Rust), mailboxes (in the Actor model, or as used by Erlang, D, Akka), and queues (as used by Python)? 
> Mailboxes have only one dedicated reciever, while queues can have multiple senders and recievers. Channels allow for using multiple queues at the same time

### List some advantages of using message passing over lock-based synchronization primitives.
> With message passing there are no conflicts over shared resources, which eliminates many potential errors.

### List some advantages of using lock-based synchronization primitives over message passing.
> Using shared memory can be faster as it doesn't require communication between threads.