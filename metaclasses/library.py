# BaseMeta is a metaclass that enforces the
# implementation of bar method to a user class 
#
class BaseMeta(type):
    def __new__(cls, name, bases, body):
        if name != 'Base' and 'bar' not in body:
            raise TypeError('bad user class')
        return super().__new__(cls, name, bases, body)

class Base(metaclass=BaseMeta):
    def foo(self):
        return self.bar()
