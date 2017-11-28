import re
from thisplace.handlers import BodyHandler


class NumberGuesser:
    def __init__(self, low, high):
        self.low = low
        self.high = high
        self.range = range(low, high + 1)

    def guess(self):
        return self.range[int(len(self.range) / 2)]

    def guess_higher(self):
        self.range = self.range[1 + int(len(self.range) / 2):]
        return self.guess()

    def guess_lower(self):
        self.range = self.range[0:int(len(self.range) / 2)]
        return self.guess()


class NumberGuessingQuestion(BodyHandler):

    @classmethod
    def can_handle(cls, text):
        return "Guess a number question" in text

    def handle(self):
        match = re.search("from (?P<low>[0-9]+) to (?P<high>[0-9]+)", self.body)
        if not match:
            raise Exception("Could not find the bounds")
        d = match.groupdict()
        guesser = NumberGuesser(int(d['low']), int(d['high']))
        return self.start_guessing(guesser)

    def start_guessing(self, guesser):
        answer = guesser.guess()
        text = ""
        answer_location = self.location.next
        while True:
            print("--> Guessing %s" % answer)
            result = answer_location.do_request({'answer': answer})
            text = answer_location.body_text
            if "greater than" in text or "higher than" in text:
                answer = guesser.guess_higher()
            if "less than" in text or "lower than" in text:
                answer = guesser.guess_lower()
            else:
                return result
