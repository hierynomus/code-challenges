#include <stdio.h>
#include <stdlib.h>

int main(int argc, char const *argv[])
{
    int e, f = 0;
    char *seen = (char *)malloc(1024 * 1024 * 1024);
    int last = 0;

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
            f = (((f + (e & 255)) & 16777215) * 65899) & 16777215; // L8-12
            e >>= 8;
        } while (e > 0); // L27
        if (seen[f] == 1)
        {
            break;
        }
        seen[f] = 1;
        last = f;
    } while (1); // L28-30: JMP 31 | 6
    printf("%d\n", last);
}
