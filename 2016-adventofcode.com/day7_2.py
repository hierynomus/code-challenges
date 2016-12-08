import re

abba = re.compile("([a-z])((?!\\1).)\\2\\1")
hyper_abba = re.compile("\\[[a-z]*([a-z])((?!\\1).)\\2\\1[a-z]*\\]")
r_aba = re.compile("(([a-z])(?!\\2)[a-z]\\2)")
tls = 0
ssl_2 = 0


def window(fseq, window_size=5):
    for i in range(len(fseq) - window_size + 1):
        yield fseq[i:i + window_size]


def has_ssl(ip):
    def find_aba_bab(parts, aba, bab):
        if not parts:
            return aba, bab
        head, rest = parts[0], parts[1:]
        aba.extend([a for a in window(head, 3) if a[0] == a[2] and a[0] != a[1]])
        return find_aba_bab(rest, bab, aba)

    def aba_to_bab(aba):
        return aba[1] + aba[0] + aba[1]

    aba, bab = find_aba_bab(re.split("\\[|\\]", ip), [], [])
    for a in aba:
        if aba_to_bab(a) in bab:
            return True

    return False


with open('day7.in', 'r') as f:
    for line in f:
        ip = line.strip()
        if abba.search(ip) and not hyper_abba.search(ip):
            tls += 1
        if has_ssl(ip):
            ssl_2 += 1


print("Day 7.1: %s" % tls)
print("Day 7.2: %s" % ssl_2)
