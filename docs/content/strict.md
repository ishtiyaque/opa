---
title: Compiler Strict Mode
kind: misc
weight: 9
---

The Rego compiler supports `strict mode`, where additional constraints and safety checks are enforced during compilation.
Compiler rules that will be enforced by future versions of OPA, but will be a breaking change once introduced, are incubated in strict mode. 
This creates an opportunity for users to verify that their policies are compatible with the next version of OPA before upgrading. 

Compiler Strict mode is supported by the `check` command, and can be enabled through the `-S` flag.

```
-S, --strict enable compiler strict mode
```

## Strict Mode Constraints and Checks

Name | Description | Enforced by default in OPA version
--- | --- | ---
Duplicate imports | Duplicate [imports](../policy-language/#imports), where one import shadows another, are prohibited. | 1.0
Unused local assignments | Unused [assignments](../policy-reference/#assignment-and-equality) local to a rule, function or comprehension are prohibited | 1.0
`input` and `data` reserved keywords | `input` and `data` are reserved keywords, and may not be used as names for rules and variable assignment. | 1.0 
`any()` and `all()` removed | The `any()` and `all()` built-in functions have been deprecated, and will be removed in OPA 1.0. Use of these functions is prohibited. | 1.0