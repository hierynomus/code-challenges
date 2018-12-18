input = 846021
# input = 2018

elfs = [0, 1]
recipes = [3, 7]

for t in range(input + 10):
    r1 = recipes[elfs[0]]
    r2 = recipes[elfs[1]]
    new_recipe = r1 + r2
    if new_recipe >= 10:
        recipes.append(new_recipe // 10)
        recipes.append(new_recipe % 10)
    else:
        recipes.append(new_recipe)

    elfs[0] = (elfs[0] + r1 + 1) % len(recipes)
    elfs[1] = (elfs[1] + r2 + 1) % len(recipes)
    if len(recipes) >= input + 10:
        break

print(''.join(map(str, recipes[input:input + 10])))