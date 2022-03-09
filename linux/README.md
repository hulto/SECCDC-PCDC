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
  - ssh 
- Check `lastlog` frequently
  - This shows the last time each account has logged in
  - This should be never for most accounts 
- Log monitoring 
- `w`
  - Show who is currently logged in and what they're doing

## Transfer data 
You can move files from a system to another securely with `scp`. As an example, if we want to move a file *logs.tar.gz* to another system at user@2.2.2.2, we could run this command: 

`scp logs.tar.gz user@2.2.2.2:/home/user/logs.tar.gz` 