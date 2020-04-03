from ctypes import *
from sys import *

logger = cdll.LoadLibrary('liblogger.so')

logger.Log(create_string_buffer(bytes(' '.join(argv[1:]), 'utf-8')), 0x03)
