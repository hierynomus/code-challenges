#include <stdio.h>

int main(int argc, char const *argv[])
{
    // printf("Start\n");
    int a = 0;
    int b = 0;
    int c = 0;
    int d = 0;
    int e = 0;
    int f = 0;
    LINE0:
        // printf("0: a=%d, b=%d, c=%d, e=%d, f=%d\n", a, b, c, e, f);
        goto LINE17;
    LINE1:
        // printf("1: a=%d, b=%d, c=%d, e=%d, f=%d\n", a, b, c, e, f);
        b = 1;
    LINE2:
        // printf("2: a=%d, b=%d, c=%d, e=%d, f=%d\n", a, b, c, e, f);
        e = 1;
    LINE3:
        // printf("3: a=%d, b=%d, c=%d, e=%d, f=%d\n", a, b, c, e, f);
        f = b * e;
    LINE4:
        // printf("4: a=%d, b=%d, c=%d, e=%d, f=%d\n", a, b, c, e, f);
        f = (f == c) ? 1 : 0;
    LINE5:
        // printf("5: a=%d, b=%d, c=%d, e=%d, f=%d\n", a, b, c, e, f);
        if (f == 1)
            goto LINE7;
    LINE6:
        // printf("6: a=%d, b=%d, c=%d, e=%d, f=%d\n", a, b, c, e, f);
        goto LINE8;
    LINE7:
        // printf("7: a=%d, b=%d, c=%d, e=%d, f=%d\n", a, b, c, e, f);
        a += b;
    LINE8:
        e += 1;
    LINE9:
        // printf("9: a=%d, b=%d, c=%d, e=%d, f=%d\n", a, b, c, e, f);
        f = (e > c) ? 1 : 0;
    LINE10:
        if (f == 1)
            goto LINE12;
    LINE11:
        goto LINE3;
    LINE12:
        b += 1;
    LINE13:
        f = (b > c) ? 1 : 0;
    LINE14:
        if (f == 1)
            goto LINE16;
    LINE15:
        goto LINE2;
    LINE16:
        printf("value of a: %d\n", a);
        return 0;
    LINE17:
        // printf("Line 17\n");
        f = ((f + 3) * 22) + 3;
        c = (c + 2) * (c + 2) * 19 * 11 + f;
        printf("17: a=%d, b=%d, c=%d, e=%d, f=%d\n", a, b, c, e, f);
        if (a == 0)
        {
            goto LINE1;
        }
        else {
            goto LINE27;
        }
    LINE27:
        // printf("Line 27\n");
        f = (27 * 28 + 29) * 30 * 14 * 32;
        c = c + f;
        a = 0;
        printf("27: a=%d, b=%d, c=%d, e=%d, f=%d\n", a, b, c, e, f);
        return 0;
        goto LINE1;
    }
