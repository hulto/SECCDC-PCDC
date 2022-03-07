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