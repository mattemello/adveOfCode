#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "library.h"

int DIMENSION_MAX_WORD = 1024;

int main(int argc, char *argv[]){

    if(argc < 1){
        printf("ERROR - not enough arguments\n");
        exit(EXIT_FAILURE);
    }

    FILE* fileInput;
    fileInput = fopen(argv[1], "r");

    int add = 0;
    int count = 0;
    char buffer[DIMENSION_MAX_WORD];
    char *word;

    while(fgets(buffer, DIMENSION_MAX_WORD, fileInput) != NULL){
        word = malloc(strlen(buffer) * sizeof(char*));

        strcpy(word, buffer);

        add = add + takeNumber(word);

        count++;
        free(word);
    }

    fclose(fileInput);

    printf("count %d value of the add: \n %d \n", count, add);

    return 0;
}
