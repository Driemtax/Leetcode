#include <stdio.h>

int romanToInt(char* s) {
    int result = 0;
    char lastChar = '\0';

    int counter = 0;
    while (s != '\0')
    {
        switch (s[counter])
        {
        case 'I':

            lastChar = 'I';
            result += 1;
            break;
        case 'V':
            if (lastChar = 'I')
            {
                result += 4;
            }
            else{
                result += 5;
            }
            lastChar = 'V';
            break;
        case 'X':
            lastChar = 'X';
            break;
        case 'L':
            lastChar = 'L';
            break;
        case 'C':
            lastChar = 'C';
            break;
        case 'D':
            lastChar = 'D';
            break;
        case 'M':
            lastChar = 'M';
            break;
        default:
            break;
        }
        counter++;
    }
    
    
}