

from collections import defaultdict


with open("./input.txt", "r") as fp:
    left_ids, right_ids = zip(*[line.split() for line in fp.readlines()])
    
    similarity_sum = 0
    count_dict: dict[str, int] = defaultdict(int)

    for right in right_ids:
        if right in left_ids:
            count_dict[right] += 1


    print(sum((int(key) * count for key, count in count_dict.items())))
