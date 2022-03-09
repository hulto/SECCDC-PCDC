# TCP Client

This is a simple TCP client written in Go. Not much else to say. But if you're curious, I'll talk a bit about my implementation here, which is kind of strange. 

---

If you look through my code, you'll likely spot some rather strange uses of concurrent code via goroutines. When I was creating this program, I ran into many issues in my testing of it. You wouldn't think creating a simple TCP client would be that difficult. I'm sure there are far better (and simpler) implementations that do the job just as well, but I tried my best. Anyway, here's the deal. After many trials and errors, I decided to play around with channels and goroutines to handle all output and input. Both input (from stdin) and output (from the connection) are recived and printed from within goroutines. Using the blocking features of channels, this actually works pretty well, since I made the channels be only of size 1. 

You might be wondering why I decided to throw the conn reading code into its own anonymous function which is made into its own goroutine. I did this because my client would stop functioning properlly whenever I sent a command to the server that had no output. This keeps things running smoothly. So, if you think about it, if you run commands with no output, you'll have multiple goroutines trying to read from the connection. According to the documentation, calling multiple functions on the same conn is perfectly fine, so this isn't a problem. 