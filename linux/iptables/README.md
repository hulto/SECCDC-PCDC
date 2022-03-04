# iptables info
- iptables will **need** to be enabled on each linux machine and rules **need** to be created to block **AND** log traffic (in that reverse order). 

## Quick command examples
| Command | Description |
| ----- | ----- |
| `iptables -A INPUT -p tcp --dport 4000 -j DROP` | Block TCP port 4000 |
| `iptables -L -v --line-numbers` | List the current iptables rules with line numbers of the default table (filter) |
| `iptables -A INPUT -m state --state ESTABLISHED,RELATED -j ALLOW` | Allow only established/related packets |
| `iptables -A OUTPUT -m state --state NEW,ESTABLISHED,RELATED -j ALLOW` | Allow new/established/related packets out |
| `iptables -A INPUT -m state --state INVALID -j DROP` | Drop strange/invalid packets |
| `iptables -P INPUT DROP` | Set the default policy of the INPUT chain to drop packets | 
| `iptables -A INPUT -p tcp -s 2.2.2.2 -j DROP` | Block the IP address 2.2.2.2 from INPUT | 
| `iptables -A OUTPUT -p icmp --icmp-type echo-request -j ACCEPT` | Allow ICMP echo requests | 
| `iptables -A INPUT -j LOG --log-prefix "INPUT DROP: " --log-ip-options --log-tcp-options` | Log INPUT packets with TCP/IP header details | 
| `iptables -A INPUT -p tcp -m state --state ESTABLISHED -m string --string "/etc/shadow" --algo bm -j LOG --log-prefix "ETC_SHADOW "` | Log INPUT packets that have "/etc/shadow" in them. You can do the same for any other string, like "/etc/passwd." |
| `iptables -A INPUT -p tcp -m state --state ESTABLISHED -m string --string "AAAAAAAAAA" --algo bm -j LOG --log-prefix "OVERFLOW "` | Log strange packets that might be trying to find a buffer overflow. | 

These are just some good example commands. You can modify them to your specific needs. This is by no means an exhaustive list. 

## Saving rules
- iptables rules do not save. In other words, if you set iptables rules and then reboot, they will no longer be there. This is good if you need to test a rule as you can always reboot if you accidentally block yourself out of the computer... definitely not speaking from experience. 
- Here is how to save: 
- 1. iptables-save > /root/ipt.save         
- 2. cat /root/ipt.save | iptables-restore  
- There are other methods to do this, so if you don't like this one, a simple good search will reveal others. 

## Other Info

- It's a good idea to have the final rule in each chain be a LOG rule with TCP/IP options enabled and a good log prefix. These logs can be analyzed later and help determine if a system has been comprimosed and, if it has, how it happened. 
- Generally speaking, the default policy for each chain should be to DROP the packet. This means that for any packet to be accepted, you must have a rule to allow it first. 
- Keep in mind the commands above only have rules for the **filter** table of iptables. There are other important tables, like **nat**, with their own chains. To view the rules in the other tables, use the same -L flag as before, but also specfiy the table with -t. 
- You can always refer to the easy-to-read iptables man page for more information. The -m flag and the proceeding words of said man page is a good place to read to really understand the power of iptables. 
- iptables is *very* powerful. It can do so many things. There are even options for load-balancing and DoS protection. It's unlikely that you'll need either of these things in a cyber defense competition, though. Do a quick google search if you're interested. 

## Book
I learned most of my iptables knowledge from the No Starch Press book, [Linux Firewalls](https://nostarch.com/firewalls.htm) by Michael Rash. Even though the book is a bit out-dated, I highly recommend giving it a read. As of writing this, I've only read through the iptables portions and am just getting started with psad, but I've already learned so much. 

