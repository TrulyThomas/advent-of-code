from copy import deepcopy


with open("./input.txt", "r") as fp:
    reports = fp.read().splitlines()
    safe_reports = 0
    single_fail = False

    for report in reports:
        report = report.split()
        perm_reports: list[list[str]] = []
        for i in range(len(report)):
            report_copy = deepcopy(report)
            report_copy.pop(i)
            perm_reports.append(report_copy)

        try:
            for report_perm in perm_reports:
                max_length = len(report_perm) - 1
                descending = None

                for index in range(max_length):
                    diff = int(report_perm[index]) - int(report_perm[index + 1])
                    
                    if diff == 0 or abs(diff) > 3:
                        if single_fail:
                            break
                        else:
                            single_fail = True
                            continue
                    
                    if descending is None:
                        descending = diff > 0

                    if (descending and diff < 0) or (not descending and diff > 0):
                        if single_fail:
                            break
                        else:
                            single_fail = True
                            continue

                    if max_length - 1 == index:
                        safe_reports += 1
                        raise Exception()
        except Exception:
            pass
    
    print(safe_reports)
