from copy import deepcopy
from collections import namedtuple

Player = namedtuple('Player', ['hp', 'armor', 'mana', 'Shield', 'Poison', 'Recharge'])
Boss = namedtuple('Boss', ['hp', 'damage'])

def memoize(f):
    """ Memoization decorator for functions taking one or more arguments. """
    class memodict(dict):
        def __init__(self, f):
            self.f = f
        def __call__(self, *args):
            if args in self:
                pass
                # print("Cache hit!")
            return self[args]
        def __missing__(self, key):
            ret = self[key] = self.f(*key)
            return ret
    return memodict(f)

DEBUG = False

def debug(msg):
    if DEBUG:
        print(msg)

def damage(player, amount):
    return player._replace(hp=player.hp - amount)

def magic_missile(player, opponent):
    return (player, damage(opponent, 4))

def drain(player, opponent):
    return (damage(player, -2), damage(opponent, 2))

def shield(player, opponent, timer):
    debug("Shield's timer is %s" % timer)
    if player.armor == 0:
        return (player._replace(armor=7), opponent)
    if timer == 0:
        debug("Shield wears off.")
        return (player._replace(armor=0), opponent)
    return (player, opponent)

def poison(player, opponent, timer):
    debug("Poison deals 3 damage, its timer is now %s" % timer)
    if timer == 0:
        debug("Poison wears off.")
    return (player, damage(opponent, 3))

def recharge(player, opponent, timer):
    debug("Recharge provides 101 mana, its timer is now %s" % timer)
    if timer == 0:
        debug("Recharge wears off.")
    return (player._replace(mana=player.mana + 101), opponent)

spells = {
    'Magic Missile': { 'mana': 53, 'cast': magic_missile },
    'Drain': { 'mana': 73, 'cast': drain },
    'Shield': { 'mana': 113, 'timer': 6, 'effect': shield, 'set_timer': lambda p, t: p._replace(Shield=t) },
    'Poison': { 'mana': 173, 'timer': 6, 'effect': poison, 'set_timer': lambda p, t: p._replace(Poison=t) },
    'Recharge': { 'mana': 229, 'timer': 5, 'effect': recharge, 'set_timer': lambda p, t: p._replace(Recharge=t) }
}

def print_stats(player, opponent):
    debug("Player: %s - Boss: %s" % (player.hp, opponent.hp))

def boss_turn(player, opponent):
    debug("-- Boss turn --")
    print_stats(player, opponent)
    boss_damage = boss.damage - player.armor
    debug("Boss deals %s damage" % boss_damage)
    return (damage(player, boss_damage), opponent)

def player_turn(player, opponent, spell_name):
    """ Take Player's turn, return 'True' if the spell was successfully cast, 'False' otherwise. """
    debug("-- Player turn --")
    print_stats(player, opponent)
    spell = spells[spell_name]

    if player.mana < spell['mana']:
        debug("Player cannot cast %s" % spell_name)
        return (player, opponent, False)
    elif spell_name in player and getattr(player, spell_name) > 0:
        debug("Effect %s already active, cannot cast" % spell_name)
        return (player, opponent, False)
    else:
        player = player._replace(mana=player.mana - spell['mana'])

    if 'cast' in spell:
        player, opponent = spell['cast'](player, opponent)
    else:
        player = spell['set_timer'](player, spell['timer'])

    return (player, opponent, True)

def full_turn(player, opponent, spell_name, hard=False):
    if hard:
        player = damage(player, 1)
        if is_dead(player):
            return (player, opponent, None)
    player, opponent = apply_effects(player, opponent)
    if is_dead(opponent):
        return (player, opponent, None)
    player, opponent, was_cast = player_turn(player, opponent, spell_name)
    if was_cast:
        if is_dead(opponent):
            debug("Won!")
            return (player, opponent, spell_name)
        player, opponent = apply_effects(player, opponent)
        if is_dead(opponent):
            debug("Won!")
            return (player, opponent, spell_name)
        player, opponent = boss_turn(player, opponent)
        if is_dead(player):
            return (player, opponent, None)

        return (player, opponent, spell_name)
    else:
        return (player, opponent, None)

def apply_effects(player, opponent):
    for e in ['Shield', 'Poison', 'Recharge']:
        timer = getattr(player, e)
        if timer > 0:
            timer -= 1
            player, opponent = spells[e]['effect'](player, opponent, timer)
            player = spells[e]['set_timer'](player, timer)
    return (player, opponent)

def is_dead(p):
    return p.hp <= 0

def mana_cost(spells_cast):
    return sum([spells[s]['mana'] for s in spells_cast])

boss = Boss(hp=71, damage=10)
me = Player(hp=50, mana=500, armor=0, Shield=0, Recharge=0, Poison=0)

@memoize
def minimum_win_spell_sequence(player, opponent, hard=False):
    spell_sequences = []
    for spell_name in spells:
        player_copy, opponent_copy, spell_cast = full_turn(player, opponent, spell_name, hard)
        if not spell_cast:
            continue
        if is_dead(player_copy):
            continue
        elif is_dead(opponent_copy):
            spell_sequences.append([spell_name])
        else:
            min_win = minimum_win_spell_sequence(player_copy, opponent_copy, hard)
            if min_win:
                spell_sequences.append([spell_name] + min_win)

    if not spell_sequences:
        return None
    else:
        return min(spell_sequences, key=mana_cost)

sequence = minimum_win_spell_sequence(me, boss)
print("1: %s" % mana_cost(sequence))

sequence = minimum_win_spell_sequence(me, boss, True)
print("2: %s" % mana_cost(sequence))
