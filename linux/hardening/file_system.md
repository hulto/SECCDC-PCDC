# Linux File System Hardening

## rkhunter
This is a CLI application that does rookit checks along with some other various security checks. 

- `sudo apt install -y rkhunter`
- `sudo rkhunter --propudp`
  - Get a baseline for the FS
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