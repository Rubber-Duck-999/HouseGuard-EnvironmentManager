'''
Created on 13 Feb 2020

@author: Rubber-Duck-999
'''

#!/usr/bin/env python
import pika
import sys, json
import subprocess
import time
###
# Environment Manager Integrator
# This is to show how the EVM could manage on
# its necessary pub & sub topics with rabbitmq

### Setup of EVM Integrator connection
print("## Beginning EVM Integrator")
credentials = pika.PlainCredentials('guest', 'password')
connection = pika.BlockingConnection(pika.ConnectionParameters('localhost', 5672, '/', credentials))
channel = connection.channel()
channel.exchange_declare(exchange='topics', exchange_type='topic', durable=True)
key_motion = 'Motion.Response'
key_failure = 'Failure.Component'
key_event = 'Event.EVM'
key_detected = 'Motion.Detected'
#

# Publishing
result = channel.queue_declare('', exclusive=False, durable=True)
queue_name = result.method.queue
channel.queue_bind(exchange='topics', queue=queue_name, routing_key=key_event)
channel.queue_bind(exchange='topics', queue=queue_name, routing_key=key_failure)
channel.queue_bind(exchange='topics', queue=queue_name, routing_key=key_detected)
#
motion = '{ "severity": 5 }'
channel.basic_publish(exchange='topics', routing_key=key_motion, body=motion)
#
print("Waiting for Messages")
count = 0
queue_empty = False

def callback(ch, method, properties, body):
    print(" Received %r:%r" % (method.routing_key, body))
    print("Count is : ", count)
    time.sleep(0.3)


while not queue_empty:
    method, properties, body = channel.basic_get(queue=queue_name, auto_ack=False)
    if not body is None:
        callback(channel, method, properties, body)
        count = count + 1
       
