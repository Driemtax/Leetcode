#include <stdio.h>

int romanToInt(char* s) {
    int result = 0;

    int counter = 0;
    while (s[counter] != '\0')
    {
        switch (s[counter])
        {
        case 'I':
            if (s[counter + 1] == 'V')
            {
                result += 4;
                counter += 2;
            }
            else if (s[counter + 1] == 'X')
            {
                result += 9;
                counter += 2;
            }
            else
            {
                result += 1;
                counter ++;   
            }
            break;
        case 'V':
            result += 5;
            counter++;
            break;
        case 'X':
            if (s[counter + 1] == 'L')
            {
                result += 40;
                counter += 2;
            }
            else if (s[counter + 1] == 'C')
            {
                result += 90;
                counter += 2;
            }
            else
            {
                result += 10;
                counter ++;   
            }
            break;
        case 'L':
            result += 50;
            counter++;
            break;
        case 'C':
            if (s[counter + 1] == 'D')
            {
                result += 400;
                counter += 2;
            }
            else if (s[counter + 1] == 'M')
            {
                result += 900;
                counter += 2;
            }
            else
            {
                result += 100;
                counter ++;   
            }
            break;
        case 'D':
            result += 500;
            counter++;
            break;
        case 'M':
            result += 1000;
            counter++;
            break;
        }
    }
    return result;
}

int main(){
    char test[] = "III";

    int testResult = romanToInt(test);

    printf("Test Result: %d", testResult);
}