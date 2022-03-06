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
`sudo` is a very powerful command and needs to be locked down. Besure to remove any and all users from the sudoer group that do not *need* to be in it. 

## [Artillery](https://github.com/BinaryDefense/artillery)
This is a good honeypot monitoring tool that can be installed on Linux or Windows (make sure python is installed). 

To install, clone the repo, move into its root directory, and run `sudo python3 setup.py`. There will be on-screen questions for setup. The config file is located in /var/artillery and you can also find other Artillery utility python scripts there as well. 