'''
Created on 13 Feb 2020

@author: Rubber-Duck-999
'''

#!/usr/bin/env python
import pika
import sys, time, json
import subprocess
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
key_weather = 'Weather'
key_motion = 'Motion.Response'
key_failure = 'Failure.Component'
key_event = 'Event.EVM'
key_detected = 'Motion.Detected'
#

# Publishing
result = channel.queue_declare('', exclusive=False, durable=True)
queue_name = result.method.queue
channel.queue_bind(exchange='topics', queue=queue_name, routing_key=key_weather)
channel.queue_bind(exchange='topics', queue=queue_name, routing_key=key_motion)
#
text = '{ "severity": 0, "component": "EVM", "action": null }'
failure = '{ "time":"14:56:00", "type": "sensor", "severity": 1 }'
channel.basic_publish(exchange='topics', routing_key=key_event, body=text)
channel.basic_publish(exchange='topics', routing_key=key_failure, body=failure)
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
       
