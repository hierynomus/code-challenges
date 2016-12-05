import re

nice_1 = 0
nice_2 = 0

re_1 = re.compile("[aeuio]")
re_2 = re.compile("(\\w)\\1")
re_3 = re.compile("ab|cd|pq|xy")

re_4 = re.compile("(\\w{2}).*\\1")
re_5 = re.compile("(\\w).\\1")


def cond_1(s):
    return len(re_1.findall(s)) >= 3


def cond_2(s):
    return len(re_2.findall(s)) > 0


def cond_3(s):
    return len(re_3.findall(s)) == 0


def cond_4(s):
    return len(re_4.findall(s)) > 0


def cond_5(s):
    return len(re_5.findall(s)) > 0


with open('day5.in', 'r') as f:
    for s in f:
        if cond_1(s) and cond_2(s) and cond_3(s):
            nice_1 += 1
        if cond_4(s) and cond_5(s):
            nice_2 += 1

print("Day 5.1: %s" % str(nice_1))
print("Day 5.2: %s" % str(nice_2))
