# General Linux Hardening EPP 

## AppArmor 
   1. Secrutiy Kernel module that you can use to define the allowed behavior of applications. For example, you can block an Apache server from accessing any files outside of its root directory. 
   2. [Basic info and setup video](https://www.youtube.com/watch?v=zPkrcTidwQI)
   3. Comes pre-installed in Ubuntu 

### Install 
| Command | Description |
| - | - |
| `sudo apt install apparmor-profiles`  | Install. Don't worry about errors, continue to next command.
| `sudo apt install apparmor-utils` | Second install command. |
| `sed -i -e 's/GRUB_CMDLINE_LINUX_DEFAULT="/&security=apparmor /' /etc/default/grub` then `update-grub` then `reboot` | Only use after a fresh install of apparmor | 

### Profiles 
Apparmor works on each application by its (the application's) profile. You should probably create a profile for every networked application, **especially any and all web servers.** The video I linked above will show you how to do this, but here is a summary of the commands. *Please do watch the video before attempting to set this up.* [Here is a link to the timestamp after installing.](https://youtu.be/zPkrcTidwQI?t=585)

- `aa-genprof` 
  - You'll then be asked for the service to setup a profile for
  - When ready, begin the scan
  - Follow on-screen instructions to allow/deny/etc each item
- `aa-complain <service>` 
  - Put the service in complain-mode 
  - `tail -f /var/log/syslog` 
    - Use the service normally and make sure everything is accepting from apparmor that should be
- `aa-enforce <service>`
  - Put the service into enforce mode

## SELinux
This is a kernel module that allows you to completely define behavior for users/applications. 

It seems fairly complicated to setup and manage, so lots of practice with this tools is required before production use. 

- [Basics video](https://www.youtube.com/watch?v=gma-IRr5mnk)
- /etc/selinux/config
  - Main config file to set SELinux's operating mode
- `sestatus`
  - Get the current status of SELinux 

## grsecurity and PaX
This is collection of linux kernel modifications that greatly enchance the secruity of the OS, even before any configuration. I would look into this if you're curious, but I imagine it would be very hard to implement in a competition. 

- [Blog about installing it in Debian](https://micahflee.com/2016/01/debian-grsecurity/) 

## sudo
`sudo` is a very powerful command and needs to be locked down. Be sure to remove any and all users from the sudoer group that do not *need* to be in it. 

## [Artillery](https://github.com/BinaryDefense/artillery)
This is a good honeypot monitoring tool that can be installed on Linux or Windows (make sure python is installed). 

To install, clone the repo, move into its root directory, and run `sudo python3 setup.py`. There will be on-screen questions for setup. The config file is located in /var/artillery and you can also find other Artillery utility python scripts there as well. 

## [OSSEC](https://www.ossec.net/)
This is a host-based IDS that is open-source and widely used. It has many features, of which inlcudes FIM (file integrity monitoring), log analysis, Windows registry monitoring, rootkit detection, alerting, and **active response**. 

OSSEC requires and agent and a server to work. As such, I'm not sure if this is something that we can implement in competitions, but having the server and agent on the same machine might work. 

## [Tripwire](https://github.com/Tripwire/tripwire-open-source)
This is an open-source FIM. I haven't been able to get this to work locally, but it's apparently quite good. 

## [Sysdig](https://github.com/draios/sysdig/wiki/How-to-Install-Sysdig-for-Linux)
This is another open-source tool that is able to monitor system events. 
  - `sysdig` 
  - `sysdig proc.name=ssh`
    - Only get information about a certain process - in this case, ssh. 
  - `sysdig -c topprocs_net`
    - Display the top networking processes 
    - You could potentially find a C2 process here
    - You can replace "net" with "cpu" to get top CPU processes. This method is **better than `ps` or `top`** as it would be harder to hide processes from this command due to how it works 
  - `sysdig -p"%evt.arg.path" "evt.type=chdir and user.name=root"`
    - See which directories the user root has been accessing 
  - `sysdig fd.name contains /some/file/or/dir`
    - Monitor actions taken on a file or directory 
    - add `evt.type=open` joined with an `and` to only care about open files within a directory 
  - `sysdig -qw filename.scap` 
    - Write output to a file for later analysis
    - These files can also be opened and viewed with the `Csysdig` GUI 
  - `sudo csysdig`[^1]
    - This is a terminal-based GUI for Sysdig. This is a great option for actively trying to figure out what is happening on a given system in live time. By pressing F2, you can get other information about your system, such as currently open files, network connections, etc.
    - Some of the more interesting views:
      - Spy Users
      - Spy syslog
      - Connections
      - Directories
      - Files

## unhide
This is a simple but neat linux tool that tries to find "hidden" processes. 
- `sudo apt install -y unhide`
- `unhide brute` 



[^1]: If you get an error with something about xterm-256color, try running `export TERM=xterm`, make sure xterm is installed, and try again. 