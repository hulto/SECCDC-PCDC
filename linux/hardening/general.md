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