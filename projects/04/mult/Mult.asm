@i
M=0
@R2
M=0

(LOOP)
    @i
    D=M // D=i
    @R1
    D=D-M //Dにi-R1を代入
    @END
    D;JGE // if i>=R1 goto END
    @R0
    D=M // D=R0
    @R2
    M=M+D // R2=R2+R0
    @i
    M=M+1 // i=i+1
    @LOOP
    0;JMP // goto LOOP
(END)
    @END
    0;JMP // infinite loop
// R0 * R1 = R0をR1回加算した値
