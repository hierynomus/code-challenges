
boss = {
    'hp': 100,
    'damage': 8,
    'armor': 2
}

no_item = [{ 'cost': 0, 'damage': 0, 'armor': 0 }]

weapons = [
    { 'cost': 8, 'damage': 4, 'armor': 0 },
    { 'cost': 10, 'damage': 5, 'armor': 0 },
    { 'cost': 25, 'damage': 6, 'armor': 0 },
    { 'cost': 40, 'damage': 7, 'armor': 0 },
    { 'cost': 74, 'damage': 8, 'armor': 0 }
]

armors = [
    { 'cost': 13, 'damage': 0, 'armor': 1 },
    { 'cost': 31, 'damage': 0, 'armor': 2 },
    { 'cost': 53, 'damage': 0, 'armor': 3 },
    { 'cost': 75, 'damage': 0, 'armor': 4 },
    { 'cost': 102, 'damage': 0, 'armor': 5 }
]

rings = [
    { 'cost': 25, 'damage': 1, 'armor': 0 },
    { 'cost': 50, 'damage': 2, 'armor': 0 },
    { 'cost': 100, 'damage': 3, 'armor': 0 },
    { 'cost': 20, 'damage': 0, 'armor': 1 },
    { 'cost': 40, 'damage': 0, 'armor': 2 },
    { 'cost': 80, 'damage': 0, 'armor': 3 }
]

def equipment_generator():
    for weapon in weapons:
        for armor in no_item + armors:
            for ring1 in no_item + rings:
                for ring2 in no_item + [r for r in rings if r != ring1]:
                    yield [weapon, armor, ring1, ring2]

def stats(eq):
    my_stats = {'hp': 100}
    my_stats['damage'] = sum([e['damage'] for e in eq])
    my_stats['armor'] = sum([e['armor'] for e in eq])
    return my_stats

def winning(eq):
    my_stats = stats(eq)
    my_damage = max([my_stats['damage'] - boss['armor'], 1])
    boss_damage = max([boss['damage'] - my_stats['armor'], 1])
    win = my_damage >= boss_damage
    return win

def cost(eq):
    return sum([e['cost'] for e in eq])

min_eq = min([e for e in equipment_generator() if winning(e)], key=cost)
print("1: %s " % cost(min_eq))

max_losing_eq = max([e for e in equipment_generator() if not winning(e)], key=cost)
print("2: %s" % cost(max_losing_eq))
