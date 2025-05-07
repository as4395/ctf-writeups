# CTF@CIT 2025

## Read Only

> Here we go!  
> Author: ronnie  
> [`readonly`](readonly)

**Tags:** reverse_engineering

---

## Solution

First, I inspected the file to understand its format:

```bash
file readonly
```

Output:

```
readonly: ELF 64-bit LSB executable, x86-64, statically linked, stripped
```

Since the binary was stripped and statically linked, I checked for embedded strings using:

```bash
strings readonly | grep CIT
```

This returned:

```
CIT{87z1BjG1968G}
```

The flag was stored directly in the binary as plain text and there was no need for debugging or reverse engineering.

---

## Flag

```
CIT{87z1BjG1968G}
```
