from machine import Pin
from time import sleep
import socket, network, sys

led = Pin(2, Pin.OUT)
print("Starting Processor")


class Client:
    def __init__(self):
        self.ssid = ""
        self.paswd = ""
        self.port = 9000
        self.addr = '192.168.0.25'
        self.error = False

    def openFile(self):
        file = open("ssid.txt", "r")
        self.ssid = file.readline()
        file = open("pswd.txt", "r")
        self.pswd = file.readline()
        print("Ssid is:" + self.ssid + ".")
        print("Paswd is length:" + self.pswd + ".")
        if len(self.ssid) == 0:
            print("Variables in file are empty")
            self.error = True
            return

    def connect(self):
        if not self.error:
            print("Connecting")
            station = network.WLAN(network.STA_IF)
            if station.isconnected() == True:
                print("Already connected")
                return
    
            station.active(True)
            station.connect(self.ssid, self.pswd)
    
            while station.isconnected() == False:
                pass

            print("Connection successful")
            print(station.ifconfig())

    def run(self):
        try:
            sock = socket.socket()
            addrinfos = socket.getaddrinfo(self.addr, self.port)
            # (host and port to connect to are in 5th element of the first tuple in the addrinfos list
            sock.connect(addrinfos[0][-1])
            print("Reconnected")
            while not self.error:
                sleep(2)
                sock.send('{ "micro": true, "ultra": true, "motion": true }\n')
            sock.close()
        except OSError:
            print("Connection Refused")
            self.error = True
            self.flash()
            sleep(2)
        except:
            print("Error on socket")
            self.error = True
            self.flash()
            sleep(2)

    def flash(self):
        if self.error:
            led.value(1)
            sleep(0.5)
            led.value(0)
            self.error = False
   

myClient = Client()
myClient.openFile()
myClient.connect()
while True:
    myClient.run()


