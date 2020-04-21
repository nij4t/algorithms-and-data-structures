# The following example asserts
# the exsistance of attribute 'foo' in library 
# 
from library import Base

assert hasattr(Base, 'foo'), "method does not exist"

class User(Base):
    def bar(self):
        return self.foo()