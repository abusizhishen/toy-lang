grammar Toy;

IDEN:[a-zA-Z_]+;
NUMBER:[0-9]+;
STRING:'"' [A-Za-z0-9_] '"';

WS: [\r\n ]+ -> skip;

item: IDEN|NUMBER|STRING;
express:IDEN;
assignStatement: IDEN '=' item;
returnStatement: 'return' express;
argument: IDEN IDEN;
arguments: argument','(argument)*;

statement
    :assignStatement
    |returnStatement
    ;

fun: 'fun' IDEN '(' arguments ')' '{'
    statement* ';'
'}';

