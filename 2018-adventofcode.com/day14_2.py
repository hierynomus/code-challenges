input = 846021
# input = 2018

elfs = [0, 1]
recipes = [3, 7]
to_match = [int(c) for c in str(input)]
m_idx = 0
while m_idx < len(to_match):
    r1 = recipes[elfs[0]]
    r2 = recipes[elfs[1]]
    new_recipe = r1 + r2
    if new_recipe >= 10:
        nr1 = new_recipe // 10
        nr2 = new_recipe % 10
        recipes.append(nr1)
        recipes.append(nr2)
        for i in [nr1, nr2]:
            if m_idx == len(to_match):
                break
            elif to_match[m_idx] == i:
                m_idx += 1
            else:
                m_idx = 0
    else:
        recipes.append(new_recipe)
        if new_recipe == to_match[m_idx]:
            m_idx += 1
        else:
            m_idx = 0
            

    elfs[0] = (elfs[0] + r1 + 1) % len(recipes)
    elfs[1] = (elfs[1] + r2 + 1) % len(recipes)
    # print(recipes)
    
print(recipes[-10:])
print(len(recipes) - len(to_match) - 1 if recipes[-1] != to_match[-1] else 0)