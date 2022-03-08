# Linux Persistence 

This page details some information about how an attacker will aquire persistence and where to look on the system for persistent malware. 

## Packages 
Important files:
- /var/lib/dpkg/status
- /var/log/dpkg.log 

## cron
Cron is a program scheduler in linux. You can use it to schedule certain tasks to be run periodically. An attacker may setup a cronjob to automatically run some sort of malware. 

- `ls -la /var/spool/cron` 
- `crontab -l`
  - List cronjobs for the user that runs this command
- `nano /etc/crontab`
  - View the file with crontabs in it
- You might want to look in /etc/cron.hourly/ to see if anything strange is in it 
  - There is also /etc/cron.monthly/ and /etc/cron.weekly/, but those are unlikely to be used by an attacker at a CDC competition given the length of the competition 

## systemctl
- `systemctl list-units -t service`
  - List services 
  - append `--all` to include inactive ones
- `systemctl status <service>`
- `systemctl stop <service>`
  - or `start`
- `systemctl disable <service>` 
  - You will need to manually stop this service if it was running when you disabled it
  - You can use `enable` here as well
- `systemctl mask <service>`
  - Link the service to /dev/null, redering it unusable 
  - You can also use `unmask` to restore it 
- `systemctl cat <service>` 
- `systemctl list-unit-files --type=service`
  - Show what is enabled/disabled at system start
  
## Important files
An attacker may utilize the following files for persistence: 
- /etc/profile.d
- /etc/profile
- /etc/bash.bashrc
- The equivalent files in each user's home directory 