class BodyHandler:
    def __init__(self, location):
        self.location = location
        self.body = location.body_text

    def __str__(self):
        return self.body
