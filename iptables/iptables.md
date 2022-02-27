# iptables info
- iptables will **need** to be enabled on each linux machine and rules **need** to be creates to block **AND** log traffic (in that reverse order). 

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

