# CTF@CIT 2025

## Rotten

> PVG{LxxdJwAXJGcsDoncKfRctddA}  
> Author: ronnie  
> [`rotten`](rotten)

**Tags:** crypto

---

## 🧾 Overview

The challenge just gave this:

```
PVG{LxxdJwAXJGcsDoncKfRctddA}
```

I noticed the flag started with PVG{}, but all the flags for this CTF use CIT{}, so I figured something was up. The challenge name was “rotten,” which immediately made me think of ROT13. I checked if PVG turns into CIT with a ROT13 shift — it did, so that confirmed it.

---

## ✅ Solution

I pasted the whole string into CyberChef and ran the “Magic” operation. It decoded cleanly to:

```
CIT{YkkqWjNKWTpfQbapXsepqqN}
```

---

## Flag

```
CIT{YkkqWjNKWTpfQbapXsepqqN}
```
