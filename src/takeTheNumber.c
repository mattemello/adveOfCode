#include <string.h>
#include "library.h"

static int mergeNumber(char firstNum, char secondNum){

    int firstN = firstNum - '0', secondN = secondNum - '0';

    return ((firstN * 10) + secondN);
}

int takeNumber(char *string){
    char firstNumber = -1, secondNumber = -1;
    int dimS = strlen(string);
    int i = 0, j = dimS;

    while(i <= j){
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


