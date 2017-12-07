#!/bin/sh

ip netns exec temp iptables -A OUTPUT -p $1 --dport $2 -j REJECT
