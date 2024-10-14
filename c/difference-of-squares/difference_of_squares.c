#include "difference_of_squares.h"
#include <stdio.h>

unsigned int sum_of_squares(unsigned int number)
{
    int sum = 0;

    for (unsigned int i = 1; i <= number; i++)
    {
        sum += i * i;
    }

    return sum;
}

unsigned int square_of_sum(unsigned int number)
{
    int sum = 0;

    for (unsigned int i = 1; i <= number; i++)
    {
        sum += i;
    }

    return sum * sum;
}

unsigned int difference_of_squares(unsigned int number)
{
    int i = square_of_sum(number);
    int j = sum_of_squares(number);
    return i - j;
}