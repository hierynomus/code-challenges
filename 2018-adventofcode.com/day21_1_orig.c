#include <stdio.h>

int main(int argc, char const *argv[])
{
    int a, b, c, e, f = 0;
    a = 11513432;

    // Instructions 0-4
    do
    {
        f = 123 & 456;
    } while (f != 72);

    // 5...
    f = 0;
    do
    {
        e = f | 65536; // L6
        f = 8858047; // L7
        do
        {
            c = e & 255; // L8
            f += c; // L9
            f &= 16777215; // L10
            f *= 65899; // L11
            f &= 16777215; // L12
            if (e < 256) // L13-16
            {
                break; // JMP 28
            }
            c = 0; // L17
            do
            {
                b = c + 1; // L18
                b *= 256; // L19
                if (b > e) // L20-22
                {
                    break; // L23: JMP 26
                }
                c += 1; // L24
            } while (1); // L25: JMP 18
            e = c; // L26
        } while (1); // L27
    } while (f != a); // L28-30: JMP 31 | 6
}
