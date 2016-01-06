from thisplace.handlers import BodyHandler


class WelcomePage(BodyHandler):

    @classmethod
    def can_handle(cls, text):
        return "Welcome," in text

    def handle(self):
        return self.location.next.do_request({'name': 'Jeroen'})
