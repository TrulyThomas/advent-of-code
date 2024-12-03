with open("./input.txt", "r") as fp:
    reports = fp.read().splitlines()
    safe_reports = 0

    for report in reports:
        report = report.split()
        max_length = len(report) - 1
        descending = None

        for index in range(max_length):
            diff = int(report[index]) - int(report[index + 1])
            
            if diff == 0 or abs(diff) > 3:
                break
            
            if descending is None:
                descending = diff > 0

            if (descending and diff < 0) or (not descending and diff > 0):
                break

            if max_length - 1 == index:
                safe_reports += 1
                break
    
    print(safe_reports)
