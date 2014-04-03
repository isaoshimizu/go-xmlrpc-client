# coding=utf-8
import xmlrpclib
import sys

proxy = xmlrpclib.ServerProxy("http://localhost:8000/")
msg = proxy.MessageService.Send(sys.argv[1])

print "Response: %s" % msg
