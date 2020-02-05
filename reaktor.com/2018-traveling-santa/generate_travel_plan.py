# Help Santa reduce his carbon footprint by optimizing his logistics!
# At your disposal, you have Santa’s most highly guarded trade secret, the Nice List.
# The coveted List details which present shall be given to which child.
# Here’s the kicker: Santa’s sleigh can only carry 10 metric tons (10,000kg) at a time,
# and for each trip Santa makes, you'll need to tell him which items to pack.
#
# You can download the list here. In the file, you'll find the wish of each child.
# Each row contains one wish. For privacy purposes, we've left out the names of the children: instead,
# their files include a numerical ID, their coordinates on Earth, and the weight of their present in grams.
# Your job is to find the most optimal routes to deliver all the presents.
# Santa starts at Korvatunturi, Finland (68.073611N 29.315278E) and, based on your list, flies directly
# from one coordinate to another until all presents are delivered, and then returns to Korvatunturi.
# The shorter the overall length of the trip, the less emissions there will be.
# For the purposes of this task, we assume Earth to be a sphere with radius of 6,378km.

# To submit your solution, upload a .csv file below. You are welcome to upload multiple solutions to improve your score.
# Each row should contain a single gift run starting from Korvatunturi, listing all the children who Santa will visit on
# the run, their IDs separated with a semicolon. See example file here. You cannot exceed the capacity of the sleigh on a
# single run. Once you've sent over your solution, we will tally up the total distance covered for all rows.
# Whoever delivers all the presents while covering the least distance, wins!

# You might also be eligible to win more than just fame and glory – see rules and eligibility here!

import fileinput

Korvatunturi = (68.073611, 29.315278)  # (68.073611N 29.315278E)
earth_r = 6378  # KM
sleigh_capacity = 10000 * 1000  # gram

children = []
for l in fileinput.input():
    id, n, e, weight = l.rstrip('\n').split(';')
    children.append((id, n, e, int(weight)))

