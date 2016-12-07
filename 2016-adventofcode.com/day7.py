import re

abba = re.compile("([a-z])((?!\\1).)\\2\\1")
hyper_abba = re.compile("\\[[a-z]*([a-z])((?!\\1).)\\2\\1[a-z]*\\]")
aba_bab = re.compile("([a-z])((?!\\1).)\\1[a-z]*\\[[a-z]*\\2\\1\\2[a-z]*\\]")
bab_aba = re.compile("\\[[a-z]*([a-z])((?!\\1).)\\1[a-z]*\\][a-z]*\\2\\1\\2")
tls = 0
ssl = 0
ssl_2 = 0


def has_ssl(ip):
    def split_ip(ip, supernets, hypernets, is_hyper):
        if not ip:
            return supernets, hypernets
        if not is_hyper:
            supernet, _, rest = ip.partition('[')
            supernets.append(supernet)
            return split_ip(rest, supernets, hypernets, True)
        elif is_hyper:
            hypernet, _, rest = ip.partition(']')
            hypernets.append(hypernet)
            return split_ip(rest, supernets, hypernets, False)

    def aba_to_bab(aba):
        return aba[1] + aba[0] + aba[1]

    supernets, hypernets = split_ip(ip, [], [], False)
    bab_needed = []
    for supernet in supernets:
        for i in range(len(supernet) - 2):
            if supernet[i] == supernet[i + 2] and supernet[i] != supernet[i + 1]:
                bab_needed.append(aba_to_bab(supernet[i:i + 3]))

    for hypernet in hypernets:
        for bab in bab_needed:
            if bab in hypernet:
                return True

    return False


with open('day7.in', 'r') as f:
    for line in f:
        ip = line.strip()
        if abba.search(ip) and not hyper_abba.search(ip):
            tls += 1
        if aba_bab.search(ip) or bab_aba.search(ip):
            ssl += 1
        if has_ssl(ip):
            ssl_2 += 1


print("Day 7.1: %s" % tls)
print("Day 7.2: %s" % ssl)
print("Day 7.2: %s" % ssl_2)
