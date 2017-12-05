#!/bin/sh

ip netns exec temp iptables -A OUTPUT -p tcp --dport 5050 -j REJECT
