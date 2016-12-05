import re

code_chars = 0
string_chars = 0
encoded_chars = 0

with open('day8.in', 'r') as f:
    for l in f:
        code_chars += len(l.strip())
        string_representation = l.strip()[1:-1]
        string_representation = string_representation.replace('\\\\', '\\')
        string_representation = string_representation.replace('\\"', '"')
        string_representation = re.sub("\\\\x[0-9a-fA-F]{2}", 'X', string_representation)
        string_chars += len(string_representation)

        encoded_representation = l.strip()
        encoded_representation = encoded_representation.replace('\\', '\\\\')
        encoded_representation = encoded_representation.replace('"', '\\"')
        encoded_representation = '"' + encoded_representation + '"'
        encoded_chars += len(encoded_representation)

print("Day 8.1: %s" % str(code_chars - string_chars))
print("Day 8.2: %s" % str(encoded_chars - code_chars))

