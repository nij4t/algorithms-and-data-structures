from ctypes import *

logger = cdll.LoadLibrary("liblogger.so")

logger.Log(create_string_buffer(b'Some text'), 0x04)
