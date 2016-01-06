import itertools

class Ingredient:
    def __init__(self, name, capacity, durability, flavor, texture, calories):
        self.name = name
        self.capacity = capacity
        self.durability = durability
        self.flavor = flavor
        self.texture = texture
        self.calories = calories

    def __repr__(self):
        return self.name


class Recipe:
    def __init__(self):
        self.amounts = {}

    def add_ingredient(self, ingredient, amount):
        self.amounts[ingredient] = amount

    def calc(self):
        capacity = 0
        durability = 0
        flavor = 0
        texture = 0
        for k, v in self.amounts.iteritems():
            capacity += k.capacity * v
            durability += k.durability * v
            flavor += k.flavor * v
            texture += k.texture * v

        if min([capacity, durability, flavor, texture]) <= 0:
            return 0
        else:
            return capacity * durability * flavor * texture

    def calories(self):
        return sum([i.calories * a for i, a in self.amounts.iteritems()])

ingredients = []

with open('day15.in', 'r') as f:
    for l in f:
        parts = l.strip().split(': ')
        ia = list(itertools.chain.from_iterable([p.strip().split(' ') for p in parts[1].split(',')]))

        ingredients.append(Ingredient(parts[0], int(ia[1]), int(ia[3]), int(ia[5]), int(ia[7]), int(ia[9])))

def recipe_generator():
    for a in range(0, 101):
        for b in range(0, 101 - a):
            for c in range(0, 101 - a - b):
                r = Recipe()
                r.add_ingredient(ingredients[0], a)
                r.add_ingredient(ingredients[1], b)
                r.add_ingredient(ingredients[2], c)
                r.add_ingredient(ingredients[3], 101 - a - b - c)
                yield r

print("1: %s" % max([r.calc() for r in recipe_generator()]))
print("2: %s" % max([r.calc() for r in recipe_generator() if r.calories() == 500]))
