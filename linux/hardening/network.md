# Linux Networking Hardening 

## General Guidelines 
- Make sure network daemons are run by a non-root system user
- Only allow the outside network (internet) to have access to the very least it needs
  - Configure SSH and the other services to only be accessible to the private network, and not to the internet
  - Things like web server DBs certainly do not need to be accessible from outside the private network
- Stop any network (or general) services that don't need to be running 
  - You can use the `systemctl` command for this (see the persistence.md file) 
  - See `netstat` and `lsof` info below 
  - `nmap localhost -p-` 
  - Check /etc/xinetd.d/ for xinetd services 
    - To disable, in the service's config file, add "disable = yes" 
- ### Secure SSH 

## **SSH**
SSH is the **S**ecure **SH**ell protocol and it's used primarily to control a system remotely. As such, SSH is often targetted by attackers. Here are some good ideas for securing the SSH server on your machine: 
- Consider using SSH keys instead of passwords to authenticate. You can generate the keys easily and then disable password authentication, which will prevent brute force attacks 
- /etc/ssh/sshd_config
    - PubkeyAuthentication yes 
- `ssh-keygen` 
  - It will ask for a password 
  - Creates ~/.ssh/id_rsa and ~/.ssh/id_rsa.pub 
  - `ssh-copy-id [user@]host` to copy the keys to a remote host
    - Adds pub key to ~/.ssh/authorized_keys in remote host
  - PasswordAuthentication no 
    - Put this in the config file if you want to only authenticate via keys 
    - If this is in place, even if an attacker acquire's the user's password, authentication will not be possible without the correct key
- ### Disable root login 
  - PermitRootLogin no
  - If you only want to allow root to login with a key:
    - PermitRootLogin without-password 
    - I recommend just disabling root login all together 
- ### Only allow certain users to SSH
  - AllowUsers user1 user2 
    - Put this in the config, with user1, user2, etc. being the users that will be permitted to ssh in
  - AllowGroups group1 gropu2 
    - Same as before except with groups 
  - DenyUsers user1 user2
  - DenyGroups group1 group2 
- AllowTcpForwarding no
- GatewayPorts no
- Protocol 2
  - This is a more secure protcol for SSH
- ListenAddress address1
  - Specify on which address SSH should be listening on
  - You can have multiple ListenAddress directives 
  - **Use this to have SSH only listening on the private network and not the internet**
- Port 2222 
  - Specify the port to use
- Banner none
  - This makes sure that no data about how the SSH server is running is sent before connections are made 

### **After any changes, make sure to reload the SSH service.** 
`systemctl reload sshd` 

---

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