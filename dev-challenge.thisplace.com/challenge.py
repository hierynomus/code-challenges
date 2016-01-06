from thisplace.location import Location


class Challenge:

    def __init__(self, host="http://dev-challenge.thisplace.com"):
        self.host = host

    def challenge_accepted(self):
        # First the first page
        next_loc = Location(self.host, "/", "GET")
        while next_loc:
            next_loc = next_loc.do_request()

Challenge().challenge_accepted()
