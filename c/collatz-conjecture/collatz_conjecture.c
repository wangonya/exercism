#include "collatz_conjecture.h"

int steps(int start)
{
    if (start < 1)
    {
        return ERROR_VALUE;
    }

    int steps = 0;

    while (start != 1)
    {
        if (start % 2 == 0)
        {
            start /= 2;
        }
        else
        {
            start *= 3;
            start++;
        }
        steps++;
    }

    return steps;
}

/*
int steps(int start){
    if(start<1) return ERROR_VALUE;
    int n = 0;
    for(; start>1; n++) start = start%2? start*3+1 : start/2;
    return n;
}
*/