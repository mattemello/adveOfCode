#include <stdlib.h>
#include <stdio.h>
#include "library.h"

void testWithuno(){
    char *string = "a1b2c3d4e5f";
    int want = 15;

    int ret = takeNumber(string);

    if(want != ret){
        printf("ERROR testWithuno - wanted %d ris %d\n", want, ret);
        exit(EXIT_FAILURE);
    }

    printf("testWithuno ok\n");
}

void testWithtwo(){
    char *string = "two1nine";
    int want = 29;

    int ret = takeNumber(string);

    if(want != ret){
        printf("ERROR testWithtwo - wanted %d ris %d\n", want, ret);
        exit(EXIT_FAILURE);
    }

    printf("testWithtwo ok\n");
}

void testWithfour(){
    char *string = "treb7uchet";
    int want = 77;

    int ret = takeNumber(string);

    if(want != ret){
        printf("ERROR testWithtwo - wanted %d ris %d\n", want, ret);
        exit(EXIT_FAILURE);
    }

    printf("testWithtwo ok\n");
}

void testWiththree(){
    char *string = "1abc2";

    int want = 12;

    int ret = takeNumber(string);

    if(want != ret){
        printf("ERROR testWiththree - wanted %d ris %d\n", want, ret);
        exit(EXIT_FAILURE);
    }

    printf("testWithree ok\n");
}

int main(){

    testWiththree();
    testWithuno();
    testWithtwo();
    testWithfour();

    return 0;
}
