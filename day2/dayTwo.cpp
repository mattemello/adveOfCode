#include <cstdlib>
#include <iostream>
#include <charconv>
#include <cstdio>
#include <fstream>
#include <iterator>
#include <ostream>
#include <string>

bool isGood(int* number, int dimension) {

    /* for(int i = 1; i < dimension + 1; i++) {
        std::cout << number[i-1] << " ";
    }

    std::cout << std::endl; */

    bool decrease = true;

    for(int i = 1; i < dimension; i++) {
        if(number[i - 1] <= number[i]) {
            decrease = false;        
            break;
        }

        if(number[i - 1] - number[i] < 1 || number[i - 1] - number[i] > 3){
            // std::cout << number[i] - number[i - 1] << " ";
            decrease = false;
            break;
        } 
    }


    bool increase = true;

    for(int i = 1; i < dimension; i++) {
        if(number[i - 1] >= number[i]) {
            increase = false;        
            break;
        }

        if(number[i] - number[i - 1] < 1 || number[i] - number[i - 1] > 3){
            increase = false;
            break;
        }
    }

    return decrease || increase;
}

int absoluteValue(int num){
    if(num < 0){
        return num * -1;
    }

    return num;
}

int* eliminateElemen(int* number, int dimension, int position) {
    for(int i = position; i < dimension - 1; i++) {
        number[i] = number[i+1];
    }
    
    return number;
}

bool isGood2(int* number, int dimension) {

    for (int i = 0; i < dimension; i++) {
        std::cout << number[i] << " "; 
    }
    std::cout << std::endl;
    bool decrease = true;
    bool entered = false;
    int* dec = number; int dimensionDec = dimension;

    for(int i = 1; i < dimensionDec; i++) {
        int num = dec[i - 1] - dec[i];
        if(dec[i - 1] <= dec[i]) {
            if(!entered){
                dec = eliminateElemen(dec, dimensionDec, i);
                dimensionDec--;
                entered = true;
                i = 1;
            }else{
                decrease = false;
            }
        }else if(absoluteValue(num) < 0 || absoluteValue(num) > 3){
            if(!entered){
                dec = eliminateElemen(dec, dimensionDec, i);
                dimensionDec--;
                i = 1;
                entered = true;
            }else{
                decrease = false;
            }
        }
    }

    entered = false;
    bool increase = true;
    int* inc = number;
    int dimensionInc = dimension;

    for(int i = 1; i < dimensionInc; i++) {
        int num = inc[i - 1] - inc[i];
        if(inc[i - 1] >= inc[i]) {
            if(!entered){
                inc = eliminateElemen(inc, dimensionInc, i);
                dimensionInc--;
                entered = true;
                i = 1;
            }else{
                decrease = false;
            }
        }else if(absoluteValue(num) < 0 || absoluteValue(num) > 3){
            if(!entered){
                inc = eliminateElemen(inc, dimensionInc, i);
                dimensionInc--;
                i = 1;
                entered = true;
            }else{
                decrease = false;
            }
        }
    }

    for (int i = 0; i < dimension; i++) {
        std::cout << number[i] << " "; 
    }
    std::cout << std::endl;

    std::cout << decrease << " " << increase << std::endl;

    return decrease || increase;
}

bool parseLine(std::string text) {

    std::string* number = new std::string[text.length()];
    std::string num = "";
    int i = 0, dimensionNumber = 0;

    while(text[i] != '\0') {

        if(text[i] == ' ' || text[i] == '\n'){
            if(num != "" && num != " "){
                number[dimensionNumber] = num;
                dimensionNumber++;
            }
            num = "";
        }else{
            num += text[i];
        }

        i++;
    }
    if(num != "" && num != " "){
        number[dimensionNumber] = num;
        dimensionNumber++;
    }
    num = "";

    int* toReturn  = new int[dimensionNumber];

    for (int i = 0; i < dimensionNumber; ++i) {
        toReturn[i] = std::stoi(number[i]); 
    }

    return isGood2(toReturn, dimensionNumber);
}

int firstSolution(std::ifstream &readfile) {

    std::string text;

    int number = 0;

    while (getline(readfile, text)) {
        if(parseLine(text)){
//            std::cout << text << std::endl;
            number++;
        }
    }

    return number;
}

int main() {

    std::ifstream readfile("try2");

    std::cout << firstSolution(readfile) << std::endl;



    return 0;
}
