# SECCDC-PCDC

## Info
This repository is made with both learning and practical goals in mind, with the intention that it be used for preparation of CDCs, but also for following along with things to do during the CDCs. 

There is a lot of information here. As such, each team member needs to make sure that he/she understands everything needed and is ready to apply the knowledge during the competition with no hiccups. 

I recommend that team members specialize in either Windows or Linux tasks/knowledge. That way there can be a Windows group and a Linux group within the team that help eachother. 

## Team Comp 
In general, CDC team sizes range from 4 to 8 members, with one needing to be selected for the Team Captain role. The Team Captain serves as the liaison between the White Team/Competition managers and the rest of the team members. This person also gets information on the Injects and is responsible for submitting them. 

It is my recommendation to establish an unofficial Team Lead that coordinates each member of the team with tasks. This person can be the official Team Captain or someone else. This would help with the overall organization of the team. In each of my experiences competing in a CDC, our organization has been lackluster to say the least. (See the below section on organization.) You may even consider assigning a Windows group lead, and a Linux group lead. That may be a bit extra, but it might work well. Some of this can change depending on the members of the group, their personalities, skill levels, etc. 

Team work is of utmost importance in CDCs. Although individual work is common also, team efforts can often lead to quicker and easier solutions to problems. This usually doesn't mean the entier team working on one problem, but rather a couple members working together and bouncing ideas off each other. If you don't get along with someone for whatever reason, make sure to identify those issues and avoid them.  

## Organization 
Even if every team member has the skills required for these competitions, if organization is failing within the team, often times the entire team will fail too. Your team should prepare to be as organized as possible. You only have a certain number of hours to compete, so make sure you make the most of it by staying organized. Here are a few notes I have on staying organized:

- The team lead should be aware of what everyone is doing at a given time 
- Passwords should be known by all team members, but not displayed in the open (if competing in-person). See my notes on passwords below. 
- Proper preparation is **required**. (See the Skills Needed section below.) If you get to the competition unprepared, you're probably going to be useless during most of it. There are a ton of resources in this repo already, so look through them and start preparing/learning. Remember, we all start at 0, but it's the first steps that matter so much. If you're having problems finding resources for a particular skill or having trouble understanding something, don't fear asking for help. 
- Keep all the details well-placed. If you have access to a large whiteboard, for example, draw a diagram of your network and how things are connected with IPs listed for each node. 

## Passwords 
Passwords are very important. So important that changing passwords is one of the very first steps you do to secure a system. It's my recommendation that the team for any given CDC competion come up with a password-generating algorithm that is to be used by each member. 
As an example, let's say we make the algorithm as follows:

1. Linux? Begin the password with "pwLinux". Windows? Begin the password with "pwWindows" 
2. Append the year "1296" to the password
3. Append "@@" to the password

This will mean that everyone can generate the password for each computer and not have to memorize all the passwords - just the algorithm. I would say make the password length **at least** 12 characters. This will make brute-force attempts meaningless. 

Your team should figure out a way to also keep track of password changes. Perhaps a file on an **out-of-scope** system? Perhaps a on a piece of paper that one person in the group controls and gives access to? Figure it out **before the competition**. 

## Skills Needed
I obviously cannot enumerate every possible skill you may need during a CDC, but here are some highlights. Also keep in mind that if you're team is diverse in it's knowledge, team members can specialize, and probably should. I haven't listed them in a particular order, but, if you should focus on anything, I would say Active Defense and Server Administration. Of course, neither of those things matter if you can't use the CLI (in most cases). 
- CLI
  - This is broad, but in general, know how to operate systems at the command-line level. The top two CLIs that I can think of being Bash and Powershell. 
  - File navigation, file creation, file manipulation, finding files
  - Process control 
  - Controlling permissions, users, groups
  - Controlling services/applications
  - System administration 
  - ssh 
  - Controlling access to privileged permissions
  - etc.
- Server Administration 
  - In CDCs, you're given a set of servers that you are responsible for protecting **and managing**. 
  - Practice setting up servers from scratch that you know you'll have to work with in the CDC. For example, practice setting up and managing an Apache web server in whichever OS needed. VirtualBox is great for this.
  - Make sure to figure out and understand (practice) how to configure the servers. Do some research (some info might even be in this repo) on how to best configure the given server for security. 
- Active Defense
  - This means being able to determine if a system is compromised and what to do next. 
  - This can include log monitoring, creating and setting firewall rules, scanning for malware, removing malware, removing a hacker, security patching, SIEM management and use, security hardening of the OS and services, changing passwords, and so on. 
- Networking 
  - Networking devices and what they do
  - DHCP, ARP, IP, TCP, UDP, ports, port scanning 
  - Subnetting 
  - Top firewalls like Palo Alto, pfsense, etc 
  - See the networking folder for more
- Common attacks / Red Team Skills
  - Being aware of how hackers actually *hack* is vital information when trying to figure out how to defend. 
  - Research common attacks and system misconfigurations for the system you're given for the CDC and, **most importantly**, how to stop the attacks and fix the misconfigurations. 
  - You should know how Red Teams operate and you should be familiar with basic hacking skills 
  - Learn about and practice using (in a safe environment) popular hacking tools