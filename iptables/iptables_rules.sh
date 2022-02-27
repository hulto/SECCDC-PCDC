#!/bin/bash

###############################################
# This is a sample iptables bash file.        #
# You can run this (or something similar)     #
# to easily fill out iptables rules quickly.  #
#                                             #
# Dan Rochester                               #
###############################################

###########################################
# To save iptables rules:                 #
# iptables-save > /root/ipt.save          #
# cat /root/ipt.save | iptables-restore   #
###########################################

IPTABLES=/sbin/iptables
MODPROBE=/sbin/modprobe

echo "[+] Flushing existing iptables rules..." 
$IPTABLES -F
$IPTABLES -F -t nat
$IPTABLES -X # delete every non-built-in chain 
$IPTABLES -P INPUT DROP
$IPTABLES -P OUTPUT DROP 
$IPTABLES -P FORWARD DROP

# Load connection-tracking modules 
$MODPROBE ip_conntrack 
$MODPROBE iptable_nat
$MODPROBE ip_conntrack_ftp
$MODPROBE ip_nat_ftp 

##### INPUT CHAIN #####
echo "[+] Setting up INPUT chain..."
# state tracking rules 
$IPTABLES -A INPUT -m state --state INVALID -j LOG --log-prefix "DROP INVALID " --log-ip-options --log-tcp-options 
$IPTABLES -A INPUT -m state --state INVALID -j DROP
$IPTABLES -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT

# ACCEPT rules
$IPTABLES -A INPUT -i lo -j ACCEPT 
$IPTABLES -A INPUT -p icmp --icmp-type echo-request -j ACCEPT 

### Bad stuff INPUT ###
$IPTABLES -A INPUT -p tcp -m state --state ESTABLISHED -m string --string "/etc/shadow" --algo bm -j LOG --log-prefix "ETC_SHADOW " 
$IPTABLES -A INPUT -p tcp -m state --state ESTABLISHED -m string --string "/etc/passwd" --algo bm -j LOG --log-prefix "ETC_PASSWD " 
# check for strange "AAAAAAA..." (buffer overflow attempts?) 
$IPTABLES -A INPUT -p tcp -m state --state ESTABLISHED -m string --string "AAAAAAAAAA" --algo bm -j LOG --log-prefix "OVERFLOW "  

# Default log rule 
$IPTABLES -A INPUT ! -i lo -j LOG --log-prefix "DROP " --log-ip-options --log-tcp-options

##### OUTPUT CHAIN #####
echo "[+] Setting up OUTPUT chain..."
# state tracking 
$IPTABLES -A OUTPUT -m state --state INVALID -j LOG --log-prefix "DROP INVALID " --log-ip-options --log-tcp-options
$IPTABLES -A OUTPUT -m state --state INVALID -j DROP 
$IPTABLES -A OUTPUT -m state --state ESTABLISHED,RELATED -j ACCEPT 

# Accept rules for allowing outbound connections 
$IPTABLES -A OUTPUT -p tcp --dport 21 --syn -m state --state NEW -j ACCEPT 
$IPTABLES -A OUTPUT -p tcp --dport 22 --syn -m state --state NEW -j ACCEPT 
$IPTABLES -A OUTPUT -p tcp --dport 25 --syn -m state --state NEW -j ACCEPT 
$IPTABLES -A OUTPUT -p tcp --dport 80 --syn -m state --state NEW -j ACCEPT 
$IPTABLES -A OUTPUT -p tcp --dport 443 --syn -m state --state NEW -j ACCEPT 
$IPTABLES -A OUTPUT -p udp --dport 53 -m state --state NEW -j ACCEPT 
$IPTABLES -A OUTPUT -p tcp --dport 53 -m state --state NEW -j ACCEPT 
$IPTABLES -A OUTPUT -p icmp --icmp-type echo-request -j ACCEPT 

# default OUTPUT LOG rule 
$IPTABLES -A OUTPUT ! -o lo -j LOG --log-prefix "DROP " --log-ip-options --log-tcp-options

##### FORWARD CHAIN ##### 
echo "[+] Setting up FORWARD chain..." 

# default log rule
$IPTABLES -A FORWARD ! -i lo -j LOG --log-prefix "DROP " --log-ip-options --log-tcp-options

##### forwarding (for nat) #####
# echo 1 > /proc/sys/net/ipv4/ip_forward
