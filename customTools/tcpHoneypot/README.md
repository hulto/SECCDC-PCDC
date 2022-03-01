# TCP Honeypot

This is a very simple TCP honeypot. While running, this looks like a normal server to the outside world. If anyone tries to connect, the server will print to the screen some details of the connection. I would block the IP trying to access it using the local firewall but also on the network firewall. 

You can specify which port the server will listen on. You should change which port your SSH server is listening on and put this one on port 22. With a quick nmap scan on my local machine, it says that this is an SSH server, even though it's not, if it's running on port 22. This will make is more enticing to an attacker. 

### Note
This will not catch all types of connections or scanning. For example, if one tries to scan this with a SYN scan (`nmap -sS ...`), no output will be displayed and you will be none the wiser. 