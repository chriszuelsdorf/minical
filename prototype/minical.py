# Prototype for a Golang project.

FILEBUILD = 12
BUILDDATE = "2022.05.26"


class Timeslot:
    def __init__(self, timecode: str, desc=None):
        if not type(timecode) == str:
            raise TypeError(f"Timecode should be str, was {type(timecode)}")
        if (
            not len(timecode) == 4
            or not int(timecode[2:]) in range(60)
            or not int(timecode[:2]) in range(24)
        ):
            raise ValueError(
                "Timecode must be a 4-digit time of type 24h; was {timecode}"
            )
        self.timecode = timecode
        self.desc = desc

    def __repr__(self):
        return (
            f"{self.timecode} : {str(self.desc)[:60] if self.desc else '<<< FREE >>>'}"
        )

    def set_desc(self, desc=None):
        self.desc = desc


def comp_slots(pscron):
    if type(pscron) != str:
        raise TypeError(f"pscron was {type(pscron)}, expected str")
    if pscron.count(" ") != 1:
        raise ValueError(
            f"pscron expected only one space, got `{pscron}` which did not have exactly one splace"
        )
    fixln = lambda x: x if len(x) == 2 else "0" + x
    subproc = lambda u: [
        y
        for z in [
            [fixln(x[0])]
            if len(x) == 1
            else [fixln(str(w)) for w in range(int(x[0]), int(x[1]) + 1)]
            for x in [v.split("-") for v in u.split(",")]
        ]
        for y in z
    ]
    mins, hrs = [subproc(x) for x in pscron.split()]
    return [w for x in [[y + z for z in mins] for y in hrs] for w in x]


VALID_SLOTS = ["0,30 8-15", "30 7"]


class minical:
    def __init__(self, valid_slots):
        sts = sorted(
            list(set([x for y in [comp_slots(z) for z in valid_slots] for x in y]))
        )
        self.sts = [Timeslot(x) for x in sts]
        self.idx = sts

    def __repr__(self):
        return (
            f'| {"TIME : TASK".ljust(67)[:67]} |\n| ---- : '
            + "-" * 60
            + " |\n"
            + "\n".join([f"| {repr(x).ljust(67)} |" for x in self.sts])
        )

    def chg(self, tstmp, desc):
        if tstmp not in self.idx:
            raise ValueError("tstmp must be in idx")
        self.sts[self.idx.index(tstmp)].set_desc(desc)


def main(vslots):
    showv = lambda: print(f"minical b.{FILEBUILD} d.{BUILDDATE}")
    print("\n" * 50)
    showv()
    print()
    m = minical(vslots)
    print(m)
    doGoOn = True
    while doGoOn:
        print()
        i = input(">>> ")
        if i in ("?", "help"):
            print()
            showv()
            print(
                "\n".join(
                    [
                        "",
                        "? or help - show this page",
                        "set x y - set slot x to description y",
                        "show - refresh the view",
                        "clear - clear all slots",
                        "exit or quit - exit the program",
                    ]
                )
            )
        elif i.startswith("set "):
            if not len(i.split()) >= 3:
                print("\nERROR - SET MUST FOLLOW PATTERN SET X Y")
            try:
                m.chg(i.split()[1], " ".join(i.split()[2:]))
                print("\n" * 50)
                print(m)
            except Exception as e:
                print("ERROR - " + str(e))
        elif i == "show":
            print("\n" * 50)
            print(m)
        elif i == "clear":
            i = input("Type `Y` in cap to confirm clearing (any other to cancel): ")
            if i == "Y":
                m = minical(vslots)
                print("\n" * 50)
                print(m)
        elif i in ("exit", "quit"):
            doGoOn = False
        else:
            print("\nUNKNOWN COMMAND, USE ? FOR HELP")


if __name__ == "__main__":
    main(VALID_SLOTS)
