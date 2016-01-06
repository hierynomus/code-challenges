import re
import requests
from thisplace.handlers.arithmetic import ArithmeticQuestion
from thisplace.handlers.guess import NumberGuessingQuestion
from thisplace.handlers.welcome import WelcomePage
from thisplace.handlers.word import WordQuestion

all_handlers = [ArithmeticQuestion, NumberGuessingQuestion, WelcomePage, WordQuestion]


class Location:
    def __init__(self, host, path, method):
        self.host = host
        self.path = path
        self.method = method
        self.url = host + path
        self.next = None

    def find_next(self, body_text):
        canonicalized_text = body_text.replace('\n', ' ')
        match = re.search("(?P<method>GET|POST).*?(?P<url>/[^ ,]+)", canonicalized_text)
        if not match:
            return None
        d = match.groupdict()
        return Location(self.host, d['url'], d['method'])

    def __do_get(self):
        resp = requests.get(self.url)
        if resp.status_code != 200:
            raise Exception("Error firing GET at %s, status %s\nResponse = %s" % (self.url, resp.status_code, resp))
        return resp.text

    def __do_post(self, fields):
        resp = requests.post(self.url, data=fields)
        if resp.status_code != 200:
            raise Exception("Error firing POST at %s, status %s\nResponse = %s" % (self.url, resp.status_code, resp))
        return resp.text

    def do_request(self, fields={}):
        print(self)
        body_text = None
        if self.method == 'GET':
            body_text = self.__do_get()
        elif self.method == 'POST':
            if fields:
                body_text = self.__do_post(fields)
            else:
                raise Exception("Should post with fields")
        else:
            raise Exception("Unexpected method: %s" % self.method)

        print(body_text)
        self.next = self.find_next(body_text)
        self.body_text = body_text
        return self.handle_body()

    def handle_body(self):
        for handler in all_handlers:
            if handler.can_handle(self.body_text):
                return handler(self).handle()
        # If nobody could handle, just return next
        return self.next

    def __str__(self):
        return "--> %s on %s" % (self.method, self.path)
