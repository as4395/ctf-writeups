# CTF@CIT 2025

## Baby Keygen

> I didn't strip this one — you're welcome  
> Author: ronnie  
> [`babykeygen`](babykeygen)

---

## 🧾 Overview

This challenge involved a 64-bit ELF Linux binary that wasn't stripped, making it ideal for reverse engineering with Binary Ninja. The goal was to find a valid 16-character key that would make the program output a flag.

---

## Reversing the Binary

### Step 1: Check `main()`

Here's the code for `main()`:

```asm
00407ff5    int64_t main()
0040802d        std::cout << "Enter key: "
00408043        std::getline(std::cin, input)
00408062        check_key(input)
```

The program prompts the user and passes the input to `check_key()`. 

---

### Step 2: Look Inside `check_key()`

```asm
00407ead        if (input.length() != 0x10) return
00407ed3        std::string::substr(&var_48, input)
00407ee9        if (substr != "KEY_") return
00407f33        for (i = 4; i < 16; i++)
                  if (!isalnum(input[i])) return
00407f67        validate(input)
```

From this, it was clear the program expected:

- An input that is exactly 16 characters long
- Starts with `"KEY_"`
- Has 12 alphanumeric characters following it

If the input passes those checks, it's sent to `validate()`.

---

### Step 3: Check `validate()`

Here’s what Binary Ninja showed for the `validate()` function:

```asm
00407b65    int64_t validate(std::string arg1)
00407b65        int64_t arg7
00407b65        int64_t arg4
00407b65        int64_t arg3
00407b65        int64_t arg8
00407b65        int64_t arg2
00407b65        int64_t arg5
00407b65        int64_t arg6
00407b65        int64_t arg9
00407b65        int64_t arg10
00407b65        int64_t arg11
00407b65        int64_t arg12
00407b65        int64_t arg13
00407b65        return j_sub_a6f594(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10,
                arg11, arg12, arg13) __tailcall
```

The `validate()` function is just a tail call to another address, so I skipped reversing that and tested valid keys directly.

---

## Testing It

I entered a key that matched the structure from `check_key()`:

```bash
$ ./babykeygen
Enter key: KEY_0123456789XX
CIT{41jN8BKzz388}
```

---

## Flag

```
CIT{41jN8BKzz388}
```
