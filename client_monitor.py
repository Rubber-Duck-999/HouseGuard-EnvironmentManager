from machine import Pin
from time import sleep
import socket, network, sys

led = Pin(2, Pin.OUT)

port = '9000'

file = open("file.txt", "r")
print(file.readlines())
ssid = file.readline()
pswd = file.readline()
print(ssid)

def connect():
    station = network.WLAN(network.STA_IF)
    if station.isconnected() == True:
        print("Already connected")
        return
 
    station.active(True)
    station.connect(ssid, pswd)
 
    while station.isconnected() == False:
        pass
 
    print("Connection successful")
    print(station.ifconfig())

connect()
print("Connected")
sock = socket.socket()
addrinfos = socket.getaddrinfo('192.168.0.25', int(port))
# (host and port to connect to are in 5th element of the first tuple in the addrinfos list
sock.connect(addrinfos[0][-1])
count = 0
while True:
    led.value(not led.value())
    sleep(0.5)
    print("Loop: " + str(count))
    sock.send("From Client " + str(count))
    count = count + 1
sock.close()
   
