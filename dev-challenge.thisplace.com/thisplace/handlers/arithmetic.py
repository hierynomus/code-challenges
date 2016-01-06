import re
from thisplace.handlers import BodyHandler


class ArithmeticQuestion(BodyHandler):
    operators = {
        'plus': lambda x, y: x+y,
        'times': lambda x, y: x*y,
        'minus': lambda x, y: x-y
    }

    @classmethod
    def can_handle(cls, text):
        return "Arithmetic question" in text

    def handle(self):
        match = re.search("(?P<lhs>[0-9]+) (?P<operator>[a-z]+) (?P<rhs>[0-9]+)", self.body)
        if not match:
            raise Exception("Could not parse arithmetic question: %s" % self.body)

        d = match.groupdict()
        answer = self.operators[d['operator']](int(d['lhs']), int(d['rhs']))
        return self.location.next.do_request({'answer': answer})
