#include <stdio.h>

main(int argc, char const *argv[])
{
    int a, b, c, e = 0;

    a = 1;
    // initialization (17-35)
    c = 2 * 2 * 19 * 11 + (3 * 22 + 3);
    if (a != 0)
    {
        c += (27 * 28 + 29) * 30 * 14 * 32;
        a = 0;
    }
    b = 1;
    while (b <= c / 2)
    {
        if (c % b == 0)
        {
            a += b;
        }
        b += 1;
    }
    a += c;
    printf("%d\n", a);
}
