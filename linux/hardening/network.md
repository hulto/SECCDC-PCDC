# Linux Networking Hardening 

## netstat
This is a common linux tool, but it is also supported on Windows. It's best to run `netstat` as root or with `sudo`. 
- `netstat -aon | more` 
  - Remove the `-n` for `netstat` to resolve port numbers (e.g. convert port 25 to show SMTP) 

Important flags:
| Flag | Description | 
| - | - |
| `p` | Show the process ID | 
| `t` or `u` | Filter by TCP or UDP| 
| `l`| Show only listening connections | 
| `c` | Continuously refresh | 
| `r` | Show the routing table| 

- `watch -d -n0 "netstat -atp | grep ESTA"`
  - Better method to continuously watch established connections than to use the `-c` flag in `netstat`. 
- `netstat -tn 2> /dev/null | awk '{print $5}' | sort | uniq -c | sort -nr` 
  - Sort and display the TCP connections based on the remote IP

## lsof 
"List open files" - the command description in its man page. This is accurate, and it can definitely be used for normal files, but let's focus on it's capablities that involve networking here. 

- `lsof`
  - `-P` to not resolve port numbers 
  - `-i` to list IP sockets 
    - append "4" to get only IPv4, or "6" for IPv6
    - append "TCP" or "UDP" for those filters
    - append "@\<ip:port>" for that filter 
    - append ":\<port>" for that filter 
  - `-n` to not resolve IP addresses 
- `watch -d -n 1 lsof -Pni`
  - Monitor for changes with live updates every 1 seconds 
  - Changes are highlighted 