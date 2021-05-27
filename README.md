# cp-bot (INCOMPLETE)

A tool using golang for simplifying the process of participating in a competitive programming contest.

# working

The tool will generate a directory structure as follows:

```
contest
  |- problem a
  |    |- solution.cpp
  |    |- i_[1-9].txt
  |    |- o_[1-9].txt
  |    |- e_[1-9].txt
  |
  |- problem b...
  | ...
  | ...
  |- run.sh
```

1. `contest` is the main directory named after some unique identifier used by the cp platform.
2. `problem a` will be named after the unique identifier used by the cp platform.
3. `solution.cpp` is the main file containing the solution containing your template.
4. `i_[1-9].text` are the files with the input for various test cases.
5. `o_[1-9].cpp` will contain the output of running the respective `i` files with your program.
6. `e_[1-9].cpp` is the expected output for the respective `i` file.
7. `run.sh` will run the specified program against all the `i` files and diff the `o` with the `e`.

# goals

1. Support Codeforces and AtCoder.
2. Support CPlusPlus and Python.
3. `run.batch` for windows support.
