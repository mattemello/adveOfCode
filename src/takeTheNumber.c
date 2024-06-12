#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "library.h"

static int mergeNumber(char firstNum, char secondNum){

    int firstN = firstNum - '0', secondN = secondNum - '0';

    return ((firstN * 10) + secondN);
}

static char controllWord(char *subString){

    if(strstr("one", subString) != NULL)
        return '1';
    else if(strstr("two", subString) != NULL)
        return '2';
    else if(strstr("three", subString) != NULL)
        return '3';
    else if(strstr("four", subString) != NULL)
        return '4';
    else if(strstr("five", subString) != NULL)
        return '5';
    else if(strstr("six", subString) != NULL)
        return '6';
    else if(strstr("seven", subString) != NULL)
        return '7';
    else if(strstr("eight", subString) != NULL)
        return '8';
    else if(strstr("nine", subString) != NULL)
        return '9';

    return -1;

}

int takeNumber(char *string){
    char firstNumber = -1, secondNumber = -1;
    int dimS = strlen(string);
    int i = 0, j = dimS;
    char *substring1 = (char*)malloc(dimS * sizeof(char));
    char *substring2 = (char*)malloc(dimS * sizeof(char));

    while(i < dimS && j > 0){

        substring1[i] = string[i];
        substring1[j] = string[j];

        if(firstNumber == -1){
            firstNumber = controllWord(substring1);
        }
        if(secondNumber == -1){
            secondNumber = controllWord(substring2);
        }

        if(firstNumber == -1 && string[i] >= 48 && string[i] <= 57){
            firstNumber = string[i];
        }
        if(secondNumber == -1 && string[j] >= 48 && string[j] <= 57){
            secondNumber = string[j];
        }

        if(firstNumber != -1 && secondNumber != -1){
            break;
        }

        i++;
        j--;
    }


    if(firstNumber == -1){
        firstNumber = secondNumber;
    }

    if(secondNumber == -1){
        secondNumber = firstNumber;
    }


    return mergeNumber(firstNumber, secondNumber);
}


