from ctypes import *

logger = cdll.LoadLibrary("liblogger.so")

logger.Log("Some text", 0x04)