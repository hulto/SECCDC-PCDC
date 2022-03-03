#!/bin/bash 

# !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
# Run this after changing your passwd  !
# and making a new sudoer user to use  !
# via ssh (if needed). This script     !
# makes it so that root cannot ssh in   !
# and locks the root account           !
# !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

##########################################
# Run in as the first script in          #
# a cyber defense competition.           #
# It has just a few initial              #
# settings to alter.                     #
#                                        #
# This is not complete. If you           #
# see something that can/should          #
# be added, make a pull request          #
# with it added!                         #
#                                        #
# Written by: Dan                        #
##########################################

echo "You have read this script's comments and are ready to run it [y/n]:"
read ANS 
if [[ $ANS != "y" ]] 
then 
    exit 1
fi 

echo "Showing apt sources...\n"
cat /etc/apt/sources.list 
echo 
echo "Look good? [y/n]"
read ANS 
if [[ $ANS != "y" ]] 
then 
    exit 1
fi

echo "APT update and upgrade..."
sudo apt clean 
sudo apt update && sudo apt upgrade -y 

echo "Locking root account..."
# Lock root account 
sudo passwd -l root 

echo "Disabling root login..."
# Set the root account default shell
# to /usr/sbin/nologin 
sudo usermod -s /usr/sbin/nologin root 

echo "Finding all users with blank passwords..."
sudo getent shadow | grep -Po '^[^:]*(?=:.?:)'

echo "\nFinding all sudoers..."
grep -Po '^sudo.+:\K.*$' /etc/group

echo "\nFinding all groups with sudo access..."
cat /etc/sudoers

echo "Finding secondary root accounts..."
grep 'x:0:' /etc/passwd

echo "Save this information, fix any problems, then reboot."
