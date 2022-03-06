# Networking
## Things you should know/learn
- TCP/IP
- UDP 
- Ports 
- Networking devices and how they work
  - Router
  - Access Point
  - Switch
  - Firewall 
- vLAN
- Subnets 
- NAT
- DMZ
- Load Balancing 
- SIEM 
- Wireshark
  - TShark 
- tcpdump 
- CIDR Notation 
- nmap 
- DNS 
- Packet anlysis 
- Security Monitoring 

## Packet Capture and Display
### [Wireshark](https://www.wireshark.org/)  
Wireshark is probably the best GUI application for viewing and analyzing packet captures. You can either set capture filters, or display filters for when you need to analyze packets. You can easily view every layer of each packet individually. Everyone in security needs to know at least the basics Wireshark.

### TShark
This is the CLI version of Wireshark. It has many of the same features as Wireshark, just in a command-line form.

### tcpdump
This is another CLI packet capturing tool. It's pretty common for this to be installed by default in Linux. It is not supported on Windows OS's. 

## Commands 
Here are some useful commands for TShark. I'm only showing TShark commands as it's supported on both Linux and Windows, but most of these can also be done using tcpdump. 

### General 
| Command | Desciption | 
| -- | -- | 
| `tshark` | Start a packet capture on the default NIC | 
| `tshark -D` | Show the available devices to listen on | 
| `tshark -i 1` | Use the first device from previous command output to listen on | 
| `tshark -w capture.pcap` | Write to the provided pcap file | 
| `tshark -r capture.pcap` | Read from the pcap file | 
| `tshark -r pcap.pcap -c20` | Read only the first 20 packets | 

### Output
| Command | Desciption |
| --- | --- |
| `tshark -V` | Verbose output - show more details of each packet | 
| `tshark -x` | View the hex and ASCII versions of the packets |

### Name resolution 
| Command | Desciption | 
| -- | -- |
| `tshark -n` | Disable name resolution | 

### Applying Filters 
| Command | Desciption |
| --- | ---| 
| `tshark -f "tcp port 80"` | Use BPF syntax to apply a capture filter | 
| `tshark -Y "tcp.dstport == 80"` | Use the Wireshark filter syntax to apply a display filter | 

### Time formatting 
| Command | Desciption |
| --- | ----|
| `tshark -t ad` | Use absolute time and date | 

### Advanced statistics information 
| Command | Desciption | 
| --- | --- |
| `tshark -z help` | Show a full listing of available statistic filters | 
| `tshark -z conv,ip` | Endpoint and conversation stats | 
| http,tree | HTTP requests/responses in a table | 
| `tshark -z follow,tcp,ascii,1` | Follow the first tcp stream and output as ASCII format | 
| ip_hosts,tree | Display every IP in a given packet capture as well as the rate/percentage of traffic corresponding to each | 
| io,phs | Protocol hierarchy | 
| http_req,tree | HTTP request stats | 
| smb,srt | SMB command stats | 


## Info
You can use TShark and the others to record and monitor traffic. Try to establish a baseline (i.e. what packet captures *should* look like in your network or on a system). Once you have that, monitoring traffic for anomalies can help you identify malicious network traffic. The statistical data provided by TShark can be especially useful here. 

## [Zeek](https://zeek.org/) 
This is an open-source and really amazing pcap analysis tool. I have yet to be able to install it on my laptop, but learning and being able to use this tool can help your network monitoring out a lot. 
