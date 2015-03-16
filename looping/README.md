# Looping

Another simple example that I want to touch upon is looping. The test here is that I want to print out `"Loop number %d"` one hundred thousand times where `%d` is the current iteration of the loop.

I'm running this test 11 times for both Go and Hack.

## Go

| Test No | Real     | User     | System   |
|---------|----------|----------|----------|
| 0       | 0m2.215s | 0m0.344s | 0m0.555s |
| 1       | 0m1.295s | 0m0.194s | 0m0.300s |
| 2       | 0m1.623s | 0m0.234s | 0m0.412s |
| 3       | 0m1.309s | 0m0.000s | 0m0.509s |
| 4       | 0m1.489s | 0m0.222s | 0m0.354s |
| 5       | 0m1.867s | 0m0.296s | 0m0.452s |
| 6       | 0m1.832s | 0m0.246s | 0m0.470s |
| 7       | 0m1.915s | 0m0.270s | 0m0.443s |
| 8       | 0m2.382s | 0m0.378s | 0m0.641s |
| 9       | 0m1.529s | 0m0.191s | 0m0.394s |
| 10      | 0m1.410s | 0m0.201s | 0m0.355s |


|             | Real     | User     | System   |
|-------------|----------|----------|----------|
| **Best**    | 0m1.295s | 0m0.000s | 0m0.300s |
| **Average** | 0m1.715s | 0m0.234s | 0m0.444s |
| **Worst**   | 0m2.382s | 0m0.378s | 0m0.641s |

## Hack

| Test No | Real     | User     | System   |
|---------|----------|----------|----------|
| 0       | 0m1.340s | 0m0.162s | 0m0.496s |
| 1       | 0m1.169s | 0m0.207s | 0m0.327s |
| 2       | 0m1.192s | 0m0.185s | 0m0.375s |
| 3       | 0m1.236s | 0m0.173s | 0m0.394s |
| 4       | 0m1.453s | 0m0.219s | 0m0.399s |
| 5       | 0m1.520s | 0m0.510s | 0m0.205s |
| 6       | 0m1.676s | 0m0.392s | 0m0.450s |
| 7       | 0m1.147s | 0m0.189s | 0m0.313s |
| 8       | 0m1.378s | 0m0.306s | 0m0.390s |
| 9       | 0m1.206s | 0m0.339s | 0m0.188s |
| 10      | 0m1.171s | 0m0.164s | 0m0.357s |

|             | Real     | User     | System   |
|-------------|----------|----------|----------|
| **Best**    | 0m1.147s | 0m0.164s | 0m0.188s |
| **Average** | 0m1.317s | 0m0.258s | 0m0.354s |
| **Worst**   | 0m1.676s | 0m0.510s | 0m0.496s |

## Results

It's interesting if you take a look at the two tables I've created and put them together because Hack wins for Real and System while Go wins for User only.

|             | Real | User | System |
|-------------|------|------|--------|
| **Best**    | Hack | Go   | Hack   |
| **Average** | Hack | Go   | Hack   |
| **Worst**   | Hack | Go   | Hack   |
