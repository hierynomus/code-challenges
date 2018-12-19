#include <stdio.h>

main(int argc, char const *argv[])
{
    int a, b, c, e, f = 0;

    a = 1;
    // initialization (17-35)
    f = 3 * 22 + 3;
    c = 2 * 2 * 19 * 11 + f;
    if (a != 0)
    {
        f = (27 * 28 + 29) * 30 * 14 * 32;
        c += f;
        a = 0;
    }

    b = 1;
    while (b <= c)
    {
        e = 1;
        while (e <= c)
        {
            f = b * e;
            if (f == c)
            {
                a += b;
            }
            e += 1;
        }
        b += 1;
        if (b > c)
        {
            break;
        }
    }
    printf("%d\n", a);
}
