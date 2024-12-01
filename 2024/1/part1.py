

with open("./input.txt", "r") as fp:
    left_ids, right_ids = zip(*[line.split() for line in fp.readlines()])
    distance_sum = 0
    
    for left, right in zip(sorted(left_ids), sorted(right_ids)):
        distance_sum += abs(int(left) - int(right))
    print(distance_sum)
