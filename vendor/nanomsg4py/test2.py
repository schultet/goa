#!/usr/bin/env python2

import nanomsg2

print(dir(nanomsg2))
recv = nanomsg2.Socket(nanomsg2.AF_SP, nanomsg2.NN_PULL)
recv.bind('inproc://recv')

send = nanomsg2.Socket(nanomsg2.AF_SP, nanomsg2.NN_PUSH)
send.connect('inproc://recv')

send.send(b'adf')
send.send(b'')
print(recv.recv())
print(recv.recv())

send.shutdown()
send.close()

recv.shutdown()
recv.close()

