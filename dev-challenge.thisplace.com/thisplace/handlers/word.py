import re
from thisplace.handlers import BodyHandler


class WordQuestion(BodyHandler):
    functions = {
        'first': lambda x, y: x[:y],
        'last': lambda x, y: x[-y:]
    }

    @classmethod
    def can_handle(cls, text):
        return "Word question" in text

    def handle(self):
        match = re.search("(?P<which>first|last) (?P<many>[0-9]+).*?\"(?P<of>[a-z]+)\"", self.body)
        if not match:
            raise Exception("Could not parse word question %s" % self.body)

        d = match.groupdict()
        answer = self.functions[d['which']](d['of'], int(d['many']))
        return self.location.next.do_request({'answer': answer})
