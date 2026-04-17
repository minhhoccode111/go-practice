# DSL to AST

```txt
# input: [rd_name1] == true && [rd_name1] == [rd_another_name1] || ([d_num2] != 5 && [t_name3] == "hello")

# output AST:

BinaryExpr (||)
  BinaryExpr (&&)
    BinaryExpr (==)
      Identifier: [rd_name1]
      Literal: true (bool)
    BinaryExpr (==)
      Identifier: [rd_name1]
      Identifier: [rd_another_name1]
  BinaryExpr (&&)
    BinaryExpr (!=)
      Identifier: [d_num2]
      Literal: 5 (float64)
    BinaryExpr (==)
      Identifier: [t_name3]
      Literal: hello (string)

# output JSON:

{
  "kind": "binary",
  "left": {
    "kind": "binary",
    "left": {
      "kind": "binary",
      "left": {
        "kind": "identifier",
        "name": "rd_name1"
      },
      "op": "==",
      "right": {
        "kind": "literal",
        "value": true
      }
    },
    "op": "&&",
    "right": {
      "kind": "binary",
      "left": {
        "kind": "identifier",
        "name": "rd_name1"
      },
      "op": "==",
      "right": {
        "kind": "identifier",
        "name": "rd_another_name1"
      }
    }
  },
  "op": "||",
  "right": {
    "kind": "binary",
    "left": {
      "kind": "binary",
      "left": {
        "kind": "identifier",
        "name": "d_num2"
      },
      "op": "!=",
      "right": {
        "kind": "literal",
        "value": 5
      }
    },
    "op": "&&",
    "right": {
      "kind": "binary",
      "left": {
        "kind": "identifier",
        "name": "t_name3"
      },
      "op": "==",
      "right": {
        "kind": "literal",
        "value": "hello"
      }
    }
  }
}
```
