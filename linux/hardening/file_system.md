# Linux File System Hardening

# Info
## Logs 
Active log monitoring can be an extremely effective way to determine if a system is compromised and how. However, log monitoring is only as good as the logs that are kept. 
- Ensure good logging from `iptables` 
  - Have a LOG rule as the last one in each chain

| Command | Description | 
| --------| --------| 
| `tar -cvf ${HOSTNAME}_logs.tar.gz /var/log/` | Collect all logs into a tar file |

### Log locations 
| Location | Description |
| ---- | ----| 
| /var/log/apache2/ | Access and error logs for Apache web server | 
| /var/log/auth.log | Logs on user logins, privileged access, and remote authentication | 
| /var/log/kern.log | Kernel logs. This is where `iptables` LOG rules go | 
| /var/log/messages | General system info | 
| /var/log/syslog | General logs | 

## System info
You may not need most of this, but just in-case these might be helpful. 
| Command | Description |
| --- | - |
|`uname -a `| Information about the OS version | 
| `ifconfig` | NIC information | 
| `route` | Routing table | 
| `netstat -a` | Display network connections | 
| `mount` | Show filesystems | 
| `ps -e` | Show running processes | 
| `top` | Live feed of running processes | 

## Searching filesystem 
Here are some useful commands for finding files based on certain criteria. 
| Command | Description | 
| - | - | 
| `find DIR -name '*word*'` | Find all files in directory DIR (and subdirectories) whose name contains 'word'. | 
| `find DIR -name '.*'` | Hidden files | 
| `find DIR -size +1G ` | Larger than 1GB | 
| See command below | Find the largest files in the system | 
| `find DIR -mmin -5` | All files modified within the last 5 minutes | 
| `grep -i -r DIR -e 'word'` | Search directory DIR for files with 'word' in them |  

Command: 
`find / -type f -exec ls -s '{}' \; | sort -n -r 
| head -5`

---
---

# Tools
## rkhunter
This is a CLI application that does rookit checks along with some various other security checks. 

- `sudo apt install -y rkhunter`
- `sudo rkhunter --propudp`
  - Establish a baseline for the FS
- `sudo rkhunter -c --enable all --disable none`
  - Do a complete scan
  - append `--rwo` to only show warnings 
  - The config file is in /etc/rkhunter.conf. Edit this so that it won't show the same false warnings then reset your baseline. 

## chkrootkit
Similar to rkhunter except I think it's been updated more recently. 

- `sudo apt install -y chkrootkit`
- `sudo chkrootkit` 

## [ClamAV](https://www.clamav.net/)
Linux Anti-Virus program that is free. I've had trouble in the past trying to get this to work properly, but you should look into it and figure out how to get it up and running on linux machines. 