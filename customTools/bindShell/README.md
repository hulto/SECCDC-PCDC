# Bind Shell

This program simply starts up a TCP server on the host. When new connections arrive, it povided that connection a direct shell into the host machine (the machine the server is running on). This is similar to the functionality of ssh. 

## How to connect remotely 
Simply use the tcpClient program in this same Git repo or use netcat. 

### Note
My tcpClient code has *slightly* better output that netcat will provide. However, netcat works a bit better than mine (curious why? Ask me. Hint: buf). So take your pick. If my client starts breaking or not working properly, try using netcat instead. 