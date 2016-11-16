grammar ExprLR;

formula : expr EOF ;

expr : <assoc=right> expr '^' expr #EXP
     | expr '*' expr  #MULT // match subexpressions joined with '*' operator
     | expr '+' expr  #SUM  // match subexpressions joined with '+' operator
     | INT            #INT  // matches simple integer atom
     ;

INT : '0'..'9'+ ;
WS  : [ \n]+ -> skip ;
PLUS : '+';
STAR : '*';