# Linux

## Tasks
- Change your user's password 
- Update/Upgrade
  - `sudo apt clean`
  - `sudo apt update`
  - `sudo apt upgrade -y`
- Create a new user (add to sudoer group) to use instead of root
  - `sudo adduser USERNAME`
  - `sudo usermod -aG sudo USERNAME` 
- Run initial.sh 
  - `bash initial.sh`
  - This was made for debian-based linux distros
- Use the output provided by initial.sh and fix and potential security threats 
  - `sudo userdel -r USERNAME`
- Setup `iptables` policies and rules 
  - See the iptables folder
- Make backups
  - Notably of any web servers 
- ### **Security Hardening**
  - See hardening folder
- Check `lastlog` frequently
  - This shows the last time each account has logged in
  - This should be never for most accounts 
- Log monitoring 
- `w`
  - Show who is currently logged in and what they're doing

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
| `find DIR -size +1GB ` | Larger than 1GB | 
| See command below | Find the largest files in the system | 
| `find DIR -mmin -5` | All files modified within the last 5 minutes | 
| `grep -i -r DIR -e 'word'` | Search directory DIR for files with 'word' in them |  

Command: 
`find / -type f -exec ls -s '{}' \; | sort -n -r 
| head -5`

## Transfer data 
You can move files from a system to another securely with `scp`. As an example, if we want to move a file *logs.tar.gz* to another system at user@2.2.2.2, we could run this command: 

`scp logs.tar.gz user@2.2.2.2:/home/user/logs.tar.gz` 