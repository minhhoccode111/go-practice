// TypeScript of the JSON return
type Operator = "==" | "!=" | ">" | "<" | ">=" | "<=" | "&&" | "||";

interface BinaryExpr {
  kind: "binary";
  left: Expr;
  op: Operator;
  right: Expr;
}

interface Identifier {
  kind: "identifier";
  name: string;
}

interface Literal {
  kind: "literal";
  value: string | number | boolean;
}

type Expr = BinaryExpr | Identifier | Literal;
