# Boot Sector

This program attempts to read the boot sector of a Linux drive. The boot sector is the first 512 bytes of the boot disk. Analyzing the boot sector can reveal boot-kits. I think it's unlikely that, in a cyber-defense competition, we will encounter such malware, but who knows.

I'd recommend running this program and printing the results to a file (there is a CLI flag to do this) and checking the boot sector against the file every so often to ensure nothing has changed. 

Make sure when you run this program that you apply it to the correct disk. The default is "/dev/sda" but the boot disk could be different on different systems. 