#include <stdio.h>
#include <stdint.h>

#define MODULUS 0x7FFFFFFF

#define START_A 679
#define START_B 771
#define FACT_A 16807
#define FACT_B 48271

uint64_t rng_next(uint64_t x, uint64_t a, uint64_t divisor)
{
    while (1)
    {
        uint64_t t = a * x;
        uint64_t p = t >> 31;
        uint64_t q = t & 0x7FFFFFFF;
        x = p + q;
        if (x > MODULUS)
            x -= MODULUS;
        if ((x & divisor) == 0)
            return x;
    }
}

uint64_t rounds(uint64_t r, uint64_t div_a, uint64_t div_b)
{
    uint64_t a = START_A;
    uint64_t b = START_B;
    uint64_t count = 0;

    for (uint64_t i = 0; i < r; i++)
    {
        a = rng_next(a, FACT_A, div_a);
        b = rng_next(b, FACT_B, div_b);
        // printf("a = %llu, b = %llu\n", a, b);
        // if (i > 5)
        // {
        //     break;
        // }
        if (((a ^ b) & 0xFFFF) == 0) {
            count++;
        }
    }
    return count;
}

int main(int argc, char *argv[])
{
    printf("Day 15.1: %llu\n", rounds(40000000, 0, 0));
    printf("Day 15.2: %llu\n", rounds(5000000, 3, 7));
}
